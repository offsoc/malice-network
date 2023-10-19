package implant

import (
	"fmt"
	"github.com/chainreactors/malice-network/proto/implant/commonpb"
	"github.com/chainreactors/malice-network/tests/common"
	"testing"
)

func TestRegister(t *testing.T) {
	client := common.NewClient(common.DefaultListenerAddr)
	spite := &commonpb.Spite{
		TaskId: 1,
	}
	body := &commonpb.Register{
		Os: &commonpb.Os{
			Name: "windows",
		},
		Process: &commonpb.Process{
			Name: "test",
			Pid:  123,
			Uid:  "admin",
			Gid:  "root",
		},
		Timer: &commonpb.Timer{
			Interval: 10,
		},
	}
	client.BuildSpite(spite, body)
	resp := client.RequestSpite(spite)
	fmt.Println(resp)
}
