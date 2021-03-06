package comm

const (
	CMD_CREATEPROJ         = "new"
	CMD_ADDINTERFACE       = "addrpc"
	TAG_SEGMENT_BEGIN_TMPL = "// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_%s_%s_BEGIN__" // file_func
	TAG_SEGMENT_END_TMPL   = "// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_%s_%s_END__"   // file_func

	// NOT USE
	TAG_LINE_BEGIN_TMPL = "// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_LINE_%s_%s_BEGIN__" // file_func
	TAG_LINE_END_TMPL   = "// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_LINE_%s_%s_END__"   // file_func

	// PB tmpl
	CONTENT_TMPL_PB_MSG     = "message %s {\n\t//TODO\n}"
	CONTENT_TMPL_PB_SERVICE = "\trpc %s (%s) returns (%s) {}"
	SERVICE_EXIST_FLAG      = "[\\s]*rpc[\\s]+%s[\\s]+\\(%s\\)[\\s]+returns[\\s]+\\(%s\\)[\\s]+"

	// handler impl
	// CONTENT_TMPL_HANDLER_HEADER = "package handler\nimport (\n\t\"golang.org/x/net/context\"\n\t\"%s\"\n)\n\ntype RpcHandler struct {\n}\n\nfunc NewRpcHandler() *RpcHandler {\n\treturn &RpcHandler{}\n}\n" // import pb
	CONTENT_TMPL_HANDLER_IMPL = "func (this *RpcHandler) %s (ctx context.Context, in *pb.%s) (out *pb.%s, err error) {\n\t//TODO\n\terr = errors.New(\"TODO\")\n\treturn\n}"

	// stub newreq
	CONTENT_TMPL_STUB_NEWREQ = "func New%s() interface{} {\n\treturn &pb.%s{}\n}"
	// stub register
	CONTENT_TMPL_STUB_REGISTER = "\tlogic.RegisterReqNew(\"%s\", New%s)"
)
