package schedule

import (
	"fmt"
	"time"

	"github.com/marrbor/go-groupsession/webapi"
)

const (
	scheduleDateFormat = "2006/01/02"
)

type (
	// Info スケジュール情報
	Info struct {
		Schsid              int64  // スケジュールSID
		SchEf               string // 編集区分
		SchKf               string // 公開区分
		Title               string // タイトル
		Naiyo               string // 内容
		StartDateTime       string // 開始日時 yyyy/MM/dd hh:mm:ss
		EndDateTime         string // 終了日時 yyyy/MM/dd hh:mm:ss
		TimeKbn             int64  // 時間指定区分 0:有り,1:無し
		UserKbn             int64  // 0:ユーザ,1:グループ
		UserSid             int64  // スケジュール登録対象のSID
		UserName            string // スケジュール登録対象の名前
		UserUkoFlg          int64  // スケジュール登録対象のユーザ無効フラグ 0:有効 1:無効
		ColorKbn            string // 文字色 1:青 2:赤 3:緑 4:黄 5:黒
		AbleEdit            string // 編集権限 0:編集不可 1:編集可能
		Biko                string // 備考
		AddUserName         string // 登録者名
		AddUserJkbn         string // 登録者 ユーザ状態区分 0:通常1:削除
		AddUserUkoFlg       string // 登録者 ユーザ無効フラグ 0:有効 1:無効
		AdateTime           string // 登録日時
		EDateTime           string // 編集日時
		AttendKbn           string // 出欠確認区分 0:出欠確認しない 1:出欠確認する
		AttendAns           string // 出欠確認応答 0:未登録 1:出席 2:欠席
		AttendAuKbn         string // 出欠確認登録者区分 0:登録者 1:登録者以外
		SameScheduleUserSet []SameScheduleUserSet
		ReserveSet          []ReserveSet
		CompanySet          []CompanySet // 会社拠点情報配列
		AdressSet           []AddressSet // 連絡先情報配列
	}

	// SearchRequest スケジュール検索リクエスト
	SearchRequest struct {
		Target       string `query:"target" default:"person"`  // person or group
		GSid         int64  `query:"gsid"`                     // グループSID
		USid         int64  `query:"usid"`                     // ユーザSID
		StartFrom    string `query:"startFrom"`                // 開始日 yyyy/MM/dd hh:mm
		StartTo      string `query:"startTo"`                  // 終了日 yyyy/MM/dd hh:mm
		EndFrom      string `query:"endFrom"`                  // 開始日 yyyy/MM/dd hh:mm
		EndTo        string `query:"endTo"`                    // 終了日 yyyy/MM/dd hh:mm
		KeyWord      string `query:"keyWord"`                  // キーワード
		KeyWordKbn   int64  `query:"keyWordKbn" default:"0"`   // キーワードand/or	0:and,1:or
		KeyTitile    int64  `query:"keytitle" default:"0"`     // キーワード対象 利用目的 0:on, 1:off
		KeyBody      int64  `query:"keybody" default:"0"`      // キーワード対象 内容・備考 0:on, 1:off
		Sort1        int64  `query:"sort1" default:"2"`        // ソート1キー 1:名前,2:開始日時,3:終了日時,4:タイトル/内容
		Order1       int64  `query:"order1" default:"0"`       // ソート1昇順降順 0:昇順, 1降順
		Sort2        int64  `query:"sort2" default:"3"`        // ソート2キー 1:名前,2:開始日時,3:終了日時,4:タイトル/内容
		Order2       int64  `query:"order2" default:"0"`       // ソート2昇順降順 0:昇順, 1降順
		Results      int64  `query:"results" default:"50"`     // 結果を取得する件数 MAX100
		Start        int64  `query:"start" default:"0"`        // 取得開始位置
		GroupShowKbn int64  `query:"grpShowKbn" default:"0"`   //
		SameInputFlg int64  `query:"sameInputFlg" default:"1"` // 同時登録、施設予約情報を取得する
		EscapeFlg    int64  `query:"escapeFlg" default:"1"`    // 内容、備考をhtmlエスケープを行う
	}

	// UserDurationSearchRequest 期間を指定してユーザのスケジュールを検索する
	UserDurationSearchRequest struct {
		USid      int64  `query:"usid"`      // ユーザSID
		StartFrom string `query:"startFrom"` // 開始日FROM yyyy/MM/dd
		StartTo   string `query:"startTo"`   // 開始日TO yyyy/MM/dd
		EndFrom   string `query:"endFrom"`   // 終了日FROM yyyy/MM/dd hh:mm
		EndTo     string `query:"endTo"`     // 終了日TO yyyy/MM/dd hh:mm
	}

	// SearchResultSet スケジュール検索レスポンスフィールド
	SearchResultSet struct {
		Url        string `xml:",attr"`
		Start      int64  `xml:"start,attr"`
		TotalCount int64  `xml:",attr"` // 検索結果にマッチした件数
		Result     []Info
	}

	// SameScheduleUserSet 同時登録ユーザ情報配列
	SameScheduleUserSet struct {
		Count             int64  `xml:",attr"` // 検索結果にマッチした件数
		User              string // ユーザ情報
		Name              string // 連絡先名
		UsrSid            int64  // ユーザSID
		UsrUkoFlg         int64  // ユーザ無効フラグ 0:有効 1:無効
		isEdit            int64  // 編集権限 0:権限なし 1:権限あり
		AttendAns         int64  // 出欠確認応答 0:未登録 1:出席 2:欠席
		AttendAuKbn       int64  // 出欠確認登録者区分 0:登録者 1:登録者以外
		AttendAnsDateTime string // 出欠回答日時 yyyy/MM/dd hh:mm 未登録時は「-」
	}

	// ReserveSet 連動施設予約情報配列
	ReserveSet struct {
		Reserve string // 施設情報
		RsdSid  int64  // 施設SID
		Name    string // 施設名
	}

	// CompanySet 会社拠点情報配列
	CompanySet struct {
		Count     int64      `xml:",attr"` // 検索結果にマッチした件数
		Company   string     // 会社拠点情報
		AcoSid    int64      // 会社SID
		AbaSid    int64      // 会社拠点SID
		Name      string     // 会社拠点名
		AdressSet AddressSet // 連絡先情報配列
	}

	// AddressSet 連絡先情報配列
	AddressSet struct {
		Count  int64  `xml:",attr"` // 検索結果にマッチした件数
		Adress string // 連絡先情報
		AdrSid int64  // 連絡先SID
		Name   string // 連絡先名
	}
)

func (i *Info) Dump() {
	fmt.Printf(
		`
「%s」(%s-%s)(%sが登録)
`,
		i.Title, i.StartDateTime, i.EndDateTime, i.AddUserName)
}

///////// クラスメソッド

// time2Str time.Time を日付検索パラメータ文字列にコンバート
func time2Str(d time.Time) string {
	return d.Format(scheduleDateFormat)
}

// LoadUserSchedule 指定ユーザ・指定期間のスケジュール読み出し
func LoadUserSchedule(c *webapi.Context, sid int64, from, to time.Time) ([]Info, error) {
	s := time2Str(from)
	e := time2Str(to)
	p := UserDurationSearchRequest{
		USid:      sid,
		StartFrom: s,
		StartTo:   e,
		EndFrom:   s,
		EndTo:     e,
	}
	var r SearchResultSet
	if err := webapi.Request(c, "/api/schedule/search.do", p, &r); err != nil {
		return nil, err
	}
	return r.Result, nil
}
