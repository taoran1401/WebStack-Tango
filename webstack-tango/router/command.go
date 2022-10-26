package router

import (
	"taogin/app/command"
	"taogin/config/global"
)

func InitRouteCommand(cmd string) {
	switch cmd {
	case "test:cmd":
		command.NewTestCmd(global.LOG).Handle()
	case "colly:cmd":
		command.NewCollyCmd(global.LOG).Handle()
	}
}
