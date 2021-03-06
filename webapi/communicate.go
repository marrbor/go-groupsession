package webapi

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/marrbor/go-groupsession/util"
)

const (
	// URL (2箇所以上で使うものだけ定義）

	// LoginURL ログインのURLは認証ヘッダを判断するのに必要
	LoginURL = "/api/cmn/login.do"

	// UserSearchURL ユーザ検索
	UserSearchURL = "/api/user/search.do"
)

type (
	// Context 接続情報
	Context struct {
		Client *http.Client
		Domain string
		Token  string
	}

	// ErrorResp 共通エラーレスポンス
	ErrorResp struct {
		Url      string   `xml:"url,attr"`
		Ecode    string   // エラーコード
		Messages []string `xml:"Message"` // エラーメッセージ
	}

	ErrMessage struct {
		Message string // エラーメッセージ
	}
)

// UpdateToken は、トークン文字列を更新します。
func (c *Context) UpdateToken(token string) {
	c.Token = token
}

// CreateContext は、コンテクストを生成します。
func CreateContext(domain, id, pw string) *Context {
	return &Context{
		Client: &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
		Domain: domain,
		Token:  base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", id, pw))),
	}
}

// IsErrorResp 受信 XML はエラー XML か？
func IsErrorResp(bytes []byte) error {
	var resp ErrorResp
	if err := xml.Unmarshal(bytes, &resp); err != nil {
		return err
	}

	if 0 < len(resp.Messages) {
		str := fmt.Sprintf("ERRORS: %d\n", len(resp.Messages))
		for i := range resp.Messages {
			str += fmt.Sprintf("%s\n", resp.Messages[i])
		}
		return fmt.Errorf(str)
	}
	return nil
}

// QueryParam2Str クエリパラメータ構造体から文字列を構築する
func QueryParam2Str(params interface{}) (string, error) {
	v := reflect.Indirect(reflect.ValueOf(params))
	t := v.Type()
	queries := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		// フィールド名
		key := t.Field(i).Name
		val := v.Field(i)
		q := t.Field(i).Tag.Get("query")
		r := t.Field(i).Tag.Get("required")
		util.Log(fmt.Sprintf("Query %s: %+v\n", t.Field(i).Name, val))
		if val.String() == "" {
			// 値なしの場合
			if r == "true" {
				// 必須パラメータだったらエラー。
				return "", fmt.Errorf("'%s.%s(query: '%s')' is required", t.Name(), key, q)
			}
			continue // 必須でないパラメータだったらクエリ文字列に入れない
		}
		queries = append(queries, fmt.Sprintf("%s=%+v", q, val))
	}
	return strings.Join(queries, "&"), nil
}

// Request リクエスト処理
func Request(c *Context, url string, params, dst interface{}) (err error) {
	query := ""
	if params != nil {
		q, err := QueryParam2Str(params)
		if err != nil {
			return err
		}
		if 0 < len(q) {
			query = "?" + q
		}
	}

	target := c.Domain + url + query
	util.Log(target)
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return err
	}

	if url == LoginURL {
		req.Header.Set("Authorization", "Basic "+c.Token)

	} else {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	if 400 <= resp.StatusCode {
		return fmt.Errorf("%d %+v", resp.StatusCode, resp.Status)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = IsErrorResp(bytes); err != nil {
		return err
	}

	util.Log(fmt.Sprintf("%s", string(bytes)))

	err = xml.Unmarshal(bytes, dst)
	util.Log(fmt.Sprintf("%+v", dst))
	return err
}
