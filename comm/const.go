package comm

const (
	CMD_CREATEPROJ         = "create"
	CMD_ADDINTERFACE       = "addinterface"
	TAG_SEGMENT_BEGIN_TMPL = "// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_%s_%s_BEGIN__" // file_func
	TAG_SEGMENT_END_TMPL   = "// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_%s_%s_END__"   // file_func

	// NOT USE
	TAG_LINE_BEGIN_TMPL = "// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_LINE_%s_%s_BEGIN__" // file_func
	TAG_LINE_END_TMPL   = "// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_LINE_%s_%s_END__"   // file_func

	// PB tmpl
	CONTENT_TMPL_PB_MSG     = "message %s {\n\t//TODO\n}"
	CONTENT_TMPL_PB_SERVICE = "rpc %s (%s) returns (%s) {}"

	// handler impl
	CONTENT_TMPL_HANDLER_HEADER = "package handler\nimport (\n\t\"golang.org/x/net/context\"\n\t\"%s\"\n)\n\ntype RpcHandler struct {\n}\n\nfunc NewRpcHandler() *RpcHandler {\n\treturn &RpcHandler{}\n}\n" // import pb
	CONTENT_TMPL_HANDLER_IMPL   = "func (this *RpcHandler) %s (ctx context.Context, in *pb.%s) (out *pb.%s, err error) {\n\t//TODO\n\treturn\n}"
)
