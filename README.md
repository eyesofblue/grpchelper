#grpchelper<br>
##简介<br>
###这个工具是干什么的？<br>
这个工具就是go环境下，使用grpc框架开发的时候，自动生成相关项目的目录结构和grpc的框架代码<br>
###为什么需要使用这个工具？<br>
当然protoc帮我们做了很多事情，且生成了<br>

>module_name<br>
    |-pb<br>
        |-service.proto<br>
    |-svr<br>
        |-svr_main.go<br>
        |-handler<br>
            |-handler.go<br>
    |-cli_tool<br>
        |-cli_tool_main.go<br>
        |-stub<br>
            |-stub.go<br>

