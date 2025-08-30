package main

import (
	"github.com/Zaba-web/master-of-tea/cli"
	"github.com/Zaba-web/master-of-tea/core"
)

func main() {
	app := core.InitCore()
	cli.InitCli(app)
}
