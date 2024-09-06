package sys

import (
	"github.com/chainreactors/malice-network/client/console"
	"github.com/chainreactors/malice-network/helper/consts"
	"github.com/chainreactors/malice-network/proto/client/clientpb"
	"github.com/chainreactors/malice-network/proto/implant/implantpb"
	"github.com/chainreactors/malice-network/proto/services/clientrpc"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

func WhoamiCmd(cmd *cobra.Command, con *console.Console) {
	session := con.GetInteractive()
	if session == nil {
		return
	}
	sid := con.GetInteractive().SessionId
	task, err := Whoami(con.Rpc, session)
	if err != nil {
		console.Log.Errorf("Whoami error: %v", err)
		return
	}
	con.AddCallback(task.TaskId, func(msg proto.Message) {
		resp := msg.(*implantpb.Spite).GetResponse()
		con.SessionLog(sid).Consolef("%v\n", resp.GetOutput())
	})
}

func Whoami(rpc clientrpc.MaliceRPCClient, session *clientpb.Session) (*clientpb.Task, error) {
	task, err := rpc.Whoami(console.Context(session), &implantpb.Request{
		Name: consts.ModuleWhoami,
	})
	if err != nil {
		return nil, err
	}
	return task, err
}
