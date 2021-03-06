package user

import (
	"fmt"

	"github.com/marrbor/go-groupsession/webapi"
)

type (
	// Info ユーザ情報
	Info struct {
		Usrsid        int64  //	ユーザSID
		Usisei        string //	姓
		Usimei        string //	名
		Usiseikn      string //	姓カナ
		Usimeikn      string //	名カナ
		UsrUkoFlg     int64  //	ユーザ無効フラグ	0:有効	1:無効
		SyainNo       int64  //	社員/職員番号
		Syozoku       string //	所属
		YakusyokuSid  int64  //	役職SID
		YakusyokuName string //	役職名称
		Birthday      string //	生年月日(西暦)	yyyy/MM/dd
		BirthdayKf    int64  //	生年月日          (西暦)公開フラグ	0:公開, 1:非公開
		Mail1         string //	メールアドレス1
		Mail1Comment  string //	メールアドレスコメント1
		Mail1Kf       int64  //	メールアドレス1公開フラグ	0:公開, 1:非公開
		Mail2         string //	メールアドレス2
		Mail2Comment  string //	メールアドレスコメント2
		Mail2Kf       int64  //	メールアドレス2公開フラグ	0:公開, 1:非公開
		Mail3         string //	メールアドレス3
		Mail3Comment  string //	メールアドレスコメント3
		Mail3Kf       int64  //	メールアドレス3公開フラグ	0:公開, 1:非公開
		Zip1          string //	郵便番号1
		Zip2          string //	郵便番号2
		ZipKf         int64  //	郵便番号公開フラグ	0:公開, 1:非公開
		TodofukenSid  int64  //	都道府県SID
		TodofukenName string //	都道府県名称
		TodofukenKf   int64  //	都道府県公開フラグ	0:公開, 1:非公開
		ImageKubun    int64  //	プロフィール画像区分 0:なし/1:あり(公開)/2:あり(非公開) Belong のレスポンスにはない。
		Address1      string //	住所1
		Address1Kf    int64  //	住所1公開フラグ	0:公開, 1:非公開
		Address2      string //	住所2
		Address2Kf    int64  //	住所2公開フラグ	0:公開, 1:非公開
		Tel1          string //	電話番号1
		Tel1Naisen    string //	電話番号1内線
		Tel1Comment   string //	電話番号1コメント
		Tel1Kf        int64  //	電話番号1公開フラグ	0:公開, 1:非公開
		Tel2          string //	電話番号2
		Tel2Naisen    string //	電話番号2内線
		Tel2Comment   string //	電話番号2コメント
		Tel2Kf        int64  //	電話番号2公開フラグ	0:公開, 1:非公開
		Tel3          string //	電話番号3
		Tel3Naisen    string //	電話番号3内線
		Tel3Comment   string //	電話番号3コメント
		Tel3Kf        int64  //	電話番号3公開フラグ	0:公開, 1:非公開
		Fax1          string //	FAX番号1
		Fax1Comment   string //	FAX番号1コメント
		Fax1Kf        int64  //	FAX番号1公開フラグ	0:公開, 1:非公開
		Fax2          string //	FAX番号2
		Fax2Comment   string //	FAX番号2コメント
		Fax2Kf        int64  //	FAX番号2公開フラグ	0:公開, 1:非公開
		Fax3          string //	FAX番号3
		Fax3Comment   string //	FAX番号3コメント
		Fax3Kf        int64  //	FAX番号3公開フラグ	0:公開, 1:非公開
		Bikou         string //	備考
		AddDateTime   string //	登録日 yyyy/MM/dd	hh:mm:ss
		EditDateTime  string //	変更日	yyyy/MM/dd	hh:mm:ss	変更していない場合は登録日と同じ日付

		///// 以下は whoami の問い合わせの時だけに返される属性。
		Usid      int64  // ユーザSID
		LoginId   string // ログインID
		NameSei   string // 姓
		NameMei   string // 名
		NameSeiKn string // 姓カナ
		NameMeiKn string // 名カナ
	}

	// WhoamiResultSet ログインユーザの情報取得レスポンス
	WhoamiResultSet struct {
		Url    string `xml:",attr"`
		Result Info
	}

	// LoginResponse 一次認証レスポンスフィールド
	LoginResponse struct {
		URL             string `xml:"url,attr"`
		Type            int64  // トークン種別 0 :一次トークン、1:トークン
		Token           string // トークン
		SendMailFrom    string // 送信者名 ワンタイムパスワード認証時のみ(※トークン種別に一次トークンが返る場合、ユーザのワンタイムパスワード通知先アドレスにメール送信が行われる)
		SendMailTo      string // 送信先アドレス 開始3文字以降は*でエスケープ
		SendMailSubject string // 送信タイトル ワンタイムパスワード認証時のみ
		SendMailDate    string // メールの送信日時 ワンタイムパスワード認証時のみ
	}

	// SearchRequest ユーザ検索リクエスト
	SearchRequest struct {
		GroupSid       int64   `query:"gsid"`       // グループSID 個人情報を取得するグループSID。
		SearchKana     string  `query:"searchKana"` // 頭文字（カナ）
		UserID         string  `query:"userId"`     // ユーザID
		EmployNo       string  `query:"shainNo"`    // 社員番号
		FamilyName     string  `query:"sei"`        // ユーザ名 姓
		FirstName      string  `query:"mei"`        // ユーザ名 名
		FamilyNameKana string  `query:"seiKn"`      // ユーザ名 姓カナ
		FirstNameKana  string  `query:"meiKn"`      // ユーザ名 名カナ
		AgeFrom        int64   `query:"ageFrom"`    // 年齢From
		AgeTo          int64   `query:"ageTo"`      // 年齢To
		PositionSid    int64   `query:"yakushoku"`  // 役職Sid
		Mail           string  `query:"mail"`       // メール
		PrefSid        int64   `query:"tdfkcd"`     // 都道府県SID
		LabelSid       []int64 `query:"labelSid"`   // ラベルSID
		Page           int64   `query:"page"`       // ページ default: 1
		Results        int64   `query:"results"`    // 結果を取得する件数 default: 50 max 100
		SortKey1       int64   `query:"sortKey"`    // ソート1キー 1:名前, 2:社員番号, 3:役職, 4:年月日
		SortOrder      int64   `query:"sortOrder"`  // ソート1昇順降順 0:昇順, 1降順 default:0
		Sort2Key       int64   `query:"sortKey2"`   // ソート2キー 1:名前, 2:社員番号, 3:役職, 4:年月日
		Sort2Order     int64   `query:"sortOrder2"` // ソート2昇順降順 0:昇順, 1降順 default:0
	}

	// SearchResultSet ユーザ検索レスポンスフィールド
	SearchResultSet struct {
		Url        string `xml:",attr"` // リクエストを送ったURL
		Count      int    `xml:",attr"` //要素数
		TotalCount int    `xml:",attr"` //検索にマッチした件数
		Page       int    `xml:",attr"` //ページ番号
		MaxPage    int    `xml:",attr"` //ページ数
		Result     []Info
		LabelSet   webapi.LabelSet
		GroupSet   webapi.GroupSet
	}

	// InfoResultSet は、whoami/inf で取得されるユーザ情報の取得型です。
	InfoResultSet struct {
		URL    string `xml:",attr"`
		Result Info
	}
)

// Dump 表示
func (i *Info) Dump() {
	fmt.Printf(
		`
社員番号(GS ID) <有効>: %d(%d) <%+v>
名前: %s %s(%s %s)
メールアドレス:%s
所属: %s(%s)
`,
		i.SyainNo, i.Usrsid, i.UsrUkoFlg,
		i.Usisei, i.Usimei, i.Usiseikn, i.Usimeikn,
		i.Mail1,
		i.Syozoku, i.YakusyokuName)
}

///////// クラスメソッド

// Login は、ログインして接続情報を返します。
func Login(id, pw, domain string) (*webapi.Context, error) {
	var l LoginResponse
	cp := webapi.CreateContext(domain, id, pw)
	if err := webapi.Request(cp, webapi.LoginURL, nil, &l); err != nil {
		return nil, err
	}
	cp.UpdateToken(l.Token) // Basic認証文字列から取得したトークンに更新
	return cp, nil
}

// LoadAllUser は全ユーザを返します。
func LoadAllUser(c *webapi.Context) ([]Info, error) {
	ret := make([]Info, 0)
	page := 1
	for {
		var r SearchResultSet
		if err := webapi.Request(c, webapi.UserSearchURL, webapi.PageSpecifyQuery{Page: page}, &r); err != nil {
			return nil, err
		}

		for i := range r.Result {
			ret = append(ret, r.Result[i])
		}

		if page == r.MaxPage {
			break
		}
		page += 1
	}
	return ret, nil
}

// LoadLoginUser はログインユーザの情報を取得します
func LoadLoginUser(c *webapi.Context) (*Info, error) {
	var r InfoResultSet
	if err := webapi.Request(c, "/api/user/whoami.do", nil, &r); err != nil {
		return nil, err
	}

	// whoami 独特のパラメータを、Info のパラメータにコピーする
	r.Result.Usrsid = r.Result.Usid
	r.Result.Usisei = r.Result.NameSei
	r.Result.Usimei = r.Result.NameMei
	r.Result.Usiseikn = r.Result.NameSeiKn
	r.Result.Usimeikn = r.Result.NameMeiKn
	return &r.Result, nil
}

// LoadUser は指定されたユーザの情報を取得します
func LoadUser(c *webapi.Context, sid int) (*Info, error) {
	p := struct {
		UsrSid int `query:"usrSid"`
	}{UsrSid: sid}
	var r InfoResultSet
	if err := webapi.Request(c, "/api/user/inf.do?", p, &r); err != nil {
		return nil, err
	}
	return &r.Result, nil
}
