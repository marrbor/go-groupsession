package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/marrbor/go-groupsession/util"

	"github.com/marrbor/go-groupsession/control"
)

// SetupAndLogin ログイン処理
func SetupAndLogin() (*control.CommunicateParam, error) {
	var uid string
	var pw string
	var domain string
	var log bool
	flag.StringVar(&uid, "u", "", "ユーザID")
	flag.StringVar(&pw, "p", "", "ログインパスワード")
	flag.StringVar(&domain, "d", "", "接続先。例 http(s)://gs.test.com")
	flag.BoolVar(&log, "v", false, "デバッグ出力オン")
	flag.Parse()

	// コマンドラインで指定されていなければ、環境変数から読み出す
	if len(uid) <= 0 {
		uid = os.Getenv("GO-GROUPSESSION-UID")
	}

	if len(pw) <= 0 {
		pw = os.Getenv("GO-GROUPSESSION-PW")
	}

	if len(domain) <= 0 {
		domain = os.Getenv("GO-GROUPSESSION-DOMAIN")
	}

	// それでも設定されていなければ、エラー
	if len(uid) <= 0 {
		return nil, fmt.Errorf("ユーザIDを指定してください")
	}

	if len(pw) <= 0 {
		return nil, fmt.Errorf("パスワードを指定してください")
	}

	if len(domain) <= 0 {
		return nil, fmt.Errorf("ドメインを指定してください")
	}

	if !strings.HasPrefix(domain, "http") {
		return nil, fmt.Errorf("ドメインはスキーマ(http/https)から始めてください")
	}

	// ログ出し
	if log {
		if err := util.LoggingSwitchOn(); err != nil {
			return nil, err
		}
	}

	cl, token, err := control.Login(uid, pw, domain)
	if err != nil {
		return nil, err
	}
	return &control.CommunicateParam{
		Client: cl,
		Domain: domain,
		Token:  token,
	}, nil
}
