package handler
import (
	"golang.org/x/net/context"
	"{{.PrefixFromGoSrcPath}}/{{.DirName}}/pb"
)

type RpcHandler struct {
}

func NewRpcHandler() *RpcHandler {
	return &RpcHandler{}
}

// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_HANDLER_IMPL_BEGIN__

func (this *RpcHandler) Echo (ctx context.Context, in *pb.EchoRequest) (out *pb.EchoResponse, err error) {
	out = &pb.EchoResponse{Msg:in.Msg}
    return 
}

// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_HANDLER_IMPL_END__

