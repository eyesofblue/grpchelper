# grpchelper
## 简介
### 这个工具是干什么的？<br>
这个工具针对go环境下，使用grpc框架开发的场景，可以自动生成相关项目的目录结构和grpc的框架代码<br>
### 为什么需要使用这个工具？<br>
当然protoc帮我们做了大部分的事情，但是没有做彻底，你还是需要<br>
&emsp;&emsp;1、手动创建一个svrmain.go的文件，写一些grpc生成服务器的框架代码<br>
&emsp;&emsp;2、手动创建一个client.go的文件，写一些客户端的框架代码，并且实现一些客户端的rpc调用来调试你的服务器逻辑<br>
而这些重复性的代码，本可以自动化的生成，这也是这个工具的意思<br>

## 使用<br>
### 依赖<br>

> module_name<br>
>   |-pb<br>
>     |-service.proto<br>
    |-svr<br>
        |-svr_main.go<br>
        |-handler<br>
            |-handler.go<br>
    |-cli_tool<br>
        |-cli_tool_main.go<br>
        |-stub<br>
            |-stub.go<br>

