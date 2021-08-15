// 登録ユーザのリストを出力する
package main

import (
	"fmt"
	"os"

	"github.com/marrbor/go-groupsession/webapi/user"

	"github.com/marrbor/go-groupsession/cmd"
)

func main() {
	cp, err := cmd.SetupAndLogin()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
	ul, err := user.LoadAllUser(cp)
	for _, u := range ul {
		u.Dump()
	}
	fmt.Printf("total:%d人\n", len(ul))
}
