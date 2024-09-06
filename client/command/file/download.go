package file

import (
	"github.com/chainreactors/malice-network/client/console"
	"github.com/chainreactors/malice-network/proto/client/clientpb"
	"github.com/chainreactors/malice-network/proto/implant/implantpb"
	"github.com/chainreactors/malice-network/proto/services/clientrpc"
	"github.com/spf13/cobra"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

func DownloadCmd(cmd *cobra.Command, con *console.Console) {
	path := cmd.Flags().Arg(1)
	session := con.GetInteractive()
	if session == nil {
		return
	}
	sid := con.GetInteractive().SessionId
	task, err := Download(con.Rpc, session, path)
	if err != nil {
		console.Log.Errorf("Download error: %v", err)
		return
	}
	con.AddCallback(task.TaskId, func(msg proto.Message) {
		con.SessionLog(sid).Importantf("Downloaded file %s ", path)
	})
}

func Download(rpc clientrpc.MaliceRPCClient, session *clientpb.Session, path string) (*clientpb.Task, error) {
	task, err := rpc.Download(console.Context(session), &implantpb.DownloadRequest{
		Name: filepath.Base(path),
		Path: path,
	})
	if err != nil {
		return nil, err
	}
	return task, err
}
