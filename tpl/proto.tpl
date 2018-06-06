syntax = "proto3";
package pb;

// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_PB_MESSAGE_BEGIN__

message EchoRequest {
    string msg = 1;
}
message EchoResponse {
    string msg = 1;
}

// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_PB_MESSAGE_END__

service {{.ProjName}} {
// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_PB_SERVICE_BEGIN__

    rpc Echo (EchoRequest) returns (EchoResponse) {}

// <DO NOT MODIFY THIS LINE> __GRPC_HELPER_SEGMENT_PB_SERVICE_END__
}

