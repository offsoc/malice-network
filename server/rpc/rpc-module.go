package rpc

import (
	"context"
	"github.com/chainreactors/malice-network/helper/types"
	"github.com/chainreactors/malice-network/proto/client/clientpb"
	"github.com/chainreactors/malice-network/proto/implant/implantpb"
)

func (rpc *Server) ListModules(ctx context.Context, _ *implantpb.Empty) (*clientpb.Task, error) {
	greq, err := newGenericRequest(ctx, &implantpb.Request{Name: types.MsgModules.String()})
	if err != nil {
		return nil, err
	}
	ch, err := rpc.asyncGenericHandler(ctx, greq)
	if err != nil {
		return nil, err
	}

	go greq.HandlerAsyncResponse(ch, types.MsgModules)
	return greq.Task.ToProtobuf(), nil
}

func (rpc *Server) LoadModule(ctx context.Context, req *implantpb.LoadModule) (*clientpb.Task, error) {
	greq, err := newGenericRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	ch, err := rpc.asyncGenericHandler(ctx, greq)
	if err != nil {
		return nil, err
	}

	go greq.HandlerAsyncResponse(ch, types.MsgNil)
	return greq.Task.ToProtobuf(), nil
}