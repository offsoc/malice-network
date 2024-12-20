package intermediate

import (
	"context"
	"fmt"
	"github.com/chainreactors/files"
	"github.com/chainreactors/logs"
	"github.com/chainreactors/malice-network/client/assets"
	"github.com/chainreactors/malice-network/helper/consts"
	"github.com/chainreactors/malice-network/helper/proto/client/clientpb"
	"github.com/chainreactors/malice-network/helper/proto/implant/implantpb"
	"github.com/chainreactors/malice-network/helper/proto/services/clientrpc"
	"github.com/chainreactors/malice-network/helper/utils/fileutils"
	"github.com/chainreactors/malice-network/helper/utils/handler"
	"math"
	"os"
	"path/filepath"
)

func GetResourceFile(pluginName, filename string) (string, error) {
	resourceFile := filepath.Join(assets.GetMalsDir(), pluginName, "resources", filename)
	if files.IsExist(resourceFile) {
		return resourceFile, nil
	}
	return "", fmt.Errorf("file not found")
}

func GetGlobalResourceFile(filename string) (string, error) {
	resourceFile := filepath.Join(assets.GetResourceDir(), filename)
	if files.IsExist(resourceFile) {
		return resourceFile, nil
	}
	return "", fmt.Errorf("file not found")
}

func NewSacrificeProcessMessage(ppid int64, hidden, block_dll, disable_etw bool, argue string) (*implantpb.SacrificeProcess, error) {
	return &implantpb.SacrificeProcess{
		Ppid:     uint32(ppid),
		Hidden:   hidden,
		BlockDll: block_dll,
		Argue:    argue,
		Etw:      !disable_etw,
	}, nil
}

func NewBinary(module string, path string, args []string, output bool, timeout uint32, arch string, process string, sac *implantpb.SacrificeProcess) (*implantpb.ExecuteBinary, error) {
	bin, err := os.ReadFile(fileutils.FormatWindowPath(path))
	if err != nil {
		return nil, err
	}

	return &implantpb.ExecuteBinary{
		Name:        filepath.Base(path),
		Bin:         bin,
		Type:        module,
		Args:        args,
		Output:      output,
		Timeout:     timeout,
		Arch:        consts.MapArch(arch),
		ProcessName: process,
		Sacrifice:   sac,
	}, nil
}

func NewExecutable(module string, path string, args []string, arch string, sac *implantpb.SacrificeProcess) (*implantpb.ExecuteBinary, error) {
	bin, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	process := filepath.Base(path)
	return &implantpb.ExecuteBinary{
		Name:        filepath.Base(path),
		Bin:         bin,
		Type:        module,
		Args:        args,
		Output:      true,
		Timeout:     math.MaxUint32,
		Arch:        consts.MapArch(arch),
		ProcessName: process,
		Sacrifice:   sac,
	}, nil
}

func WaitResult(rpc clientrpc.MaliceRPCClient, task *clientpb.Task) (*clientpb.TaskContext, error) {
	task.Need = -1
	content, err := rpc.WaitTaskFinish(context.Background(), task)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func GetResult(rpc clientrpc.MaliceRPCClient, task *clientpb.Task, index int32) (*clientpb.TaskContext, error) {
	task.Need = index
	content, err := rpc.GetTaskContent(context.Background(), task)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func PrintTask(task *clientpb.TaskContext) (*implantpb.Spite, error) {
	logs.Log.Consolef("Session: %s, Task: %d, Index:%d \n", task.Task.SessionId, task.Task.TaskId, task.Task.Need)
	err := handler.HandleMaleficError(task.Spite)
	if err != nil {
		return nil, err
	}
	logs.Log.Consolef("%v", task.Spite.GetBody())
	return task.Spite, nil
}

func ParseAssembly(spite *implantpb.Spite) (string, error) {
	err := handler.HandleMaleficError(spite)
	if err != nil {
		return "", err
	}
	response := spite.GetBinaryResponse()
	if response == nil {
		return "", fmt.Errorf("assembly response is nil")
	}
	if response.GetErr() != "" {
		return fmt.Sprintf("exit status: %d, %s", response.Status, response.Err), nil
	}
	return string(response.GetData()), nil
}

func ParseStatus(spite *implantpb.Spite) (bool, error) {
	if spite.Error == 6 {
		return false, nil
	} else if spite.Error == 0 {
		return true, nil
	} else {
		return false, handler.HandleMaleficError(spite)
	}
}
