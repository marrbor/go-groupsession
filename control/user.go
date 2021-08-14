// Package control 共通・ユーザカテゴリの機能実現
package control

import (
	"crypto/tls"
	"net/http"

	"github.com/marrbor/go-groupsession/model"
)

// Login は、ログインしてhttpクライアントとトークンを返します。
func Login(id, pw, domain string) (*http.Client, string, error) {
	var l model.LoginResponse
	cl := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	if err := Request(cl, domain, LoginURL, GenBasicAuthString(id, pw), nil, &l); err != nil {
		return nil, "", err
	}
	return cl, l.Token, nil
}

// LoadAllUser は全ユーザを返します。
func LoadAllUser(cp *CommunicateParam) ([]model.UserInfo, error) {
	ret := make([]model.UserInfo, 0)
	page := 1
	for {
		var r model.UserSearchResponse
		if err := Request(cp.Client, cp.Domain, UserSearchURL, cp.Token,
			model.PageSpecifyQuery{Page: page}, &r); err != nil {
			return nil, err
		}

		for i := range r.UserInfos {
			ret = append(ret, r.UserInfos[i])
		}

		if page == r.MaxPage {
			break
		}
		page += 1
	}
	return ret, nil
}
