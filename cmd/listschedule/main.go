// 今日から一週間分のスケジュールをロード
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/marrbor/go-groupsession/webapi/schedule"

	"github.com/marrbor/go-groupsession/cmd"
	"github.com/marrbor/go-groupsession/webapi/user"
)

func main() {
	c, err := cmd.SetupAndLogin()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	ui, err := user.LoadLoginUser(c)

	from := time.Now()
	to := from.Add(7 * 24 * time.Hour)
	sl, err := schedule.LoadUserSchedule(c, ui.Usrsid, from, to)
	for _, s := range sl {
		s.Dump()
	}
	fmt.Printf("total:%d件\n", len(sl))
}
