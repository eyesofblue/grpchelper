package stub

import (
	"github.com/eyesofblue/grpchelper/logic"
	"{{.PrefixFromGoSrcPath}}/{{.DirName}}/pb"
)

// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_STUB_NEWREQ_BEGIN__

func NewEchoRequest() interface{} {
	return &pb.EchoRequest{}
}

// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_STUB_NEWREQ_END__

func init() {
	// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_STUB_REGISTER_BEGIN__

	logic.RegisterReqNew("Echo", NewEchoRequest)

	// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_STUB_REGISTER_END__
}
