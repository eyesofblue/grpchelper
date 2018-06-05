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
)
