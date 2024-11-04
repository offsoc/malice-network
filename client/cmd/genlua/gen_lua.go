package main

import (
	"fmt"
	"github.com/chainreactors/malice-network/client/command"
	"github.com/chainreactors/malice-network/client/core/intermediate"
	"github.com/chainreactors/malice-network/client/core/plugin"
	"github.com/chainreactors/malice-network/client/repl"
	"github.com/chainreactors/malice-network/helper/proto/services/clientrpc"
	"github.com/spf13/cobra"
)

func main() {
	con, err := repl.NewConsole()
	if err != nil {
		fmt.Println(err)
		return
	}
	var cmd = &cobra.Command{
		Use:   "client",
		Short: "",
		Long:  ``,
	}
	cmd.TraverseChildren = true
	command.BindBuiltinCommands(con, cmd)
	command.RegisterClientFunc(con)
	command.RegisterImplantFunc(con)
	rpc := clientrpc.NewMaliceRPCClient(nil)
	intermediate.RegisterBuiltin(rpc)
	vm := plugin.NewLuaVM()
	plugin.GenerateLuaDefinitionFile(vm, "define.lua")
	plugin.GenerateMarkdownDefinitionFile(vm, "lua.md")
}
