package main

import (
	"github.com/chainreactors/logs"
	"github.com/chainreactors/malice-network/client/command"
	"github.com/chainreactors/malice-network/client/console"
	"github.com/chainreactors/malice-network/helper/constant"
	"github.com/pterm/pterm"
)

func init() {
	logs.Log.SetFormatter(map[logs.Level]string{
		logs.Debug:     constant.Cloud + pterm.BgLightYellow.Sprint("[debug]") + " %s ",
		logs.Warn:      constant.Zap + pterm.BgYellow.Sprint("[warn]") + " %s ",
		logs.Info:      constant.Rocket + pterm.BgCyan.Sprint("[+]") + " %s ",
		logs.Error:     constant.Monster + pterm.BgRed.Sprint("[-]") + " %s ",
		logs.Important: constant.Fire + pterm.BgMagenta.Sprint("[*]") + " %s ",
	})
}

func main() {
	err := console.Start(command.BindCommands)
	if err != nil {
		return
	}
}
