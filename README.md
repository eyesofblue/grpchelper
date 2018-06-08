# grpchelper
## 简介
### 这个工具是干什么的？<br>
这个工具针对go环境下，使用grpc框架开发的场景，可以自动生成相关项目的目录结构和grpc的框架代码<br>
### 为什么需要使用这个工具？<br>
当然protoc帮我们做了大部分的事情，但是没有做彻底，你还是需要：<br>
&emsp;&emsp;1、手动创建一个svrmain.go的文件，写一些grpc生成服务器的框架代码<br>
&emsp;&emsp;2、手动创建一个client.go的文件，写一些客户端的框架代码，并且实现一些客户端的rpc调用来调试你的服务器逻辑<br>
<br>
但这些重复性的代码，本可以自动化的生成，这也是这个工具的意义<br>

## 使用<br>
### 依赖<br>
&emsp;&emsp;1、已经配置好Go环境，设置好GOROOT、GOPATH等环境变量<br>
&emsp;&emsp;2、按照标准路径安装好grpc-go ($GOPATH/src/google.golang.org/grpc)<br>
&emsp;&emsp;3、安装好protoc工具，且配置好相应的PATH环境变量<br>
&emsp;&emsp;4、安装好protoc的go环境插件<br>
### 安装<br>
```go 
go get -u github.com/eyesofblue/grpchelper
```
### 使用<br>
#### 创建一个项目
```go
grpchelper -c new -n [your_proj_name] -i [ip] -p [port]
```
工具会自动生成如下代码结构
> yourprojname<br>
&emsp;&emsp;|-pb<br>
&emsp;&emsp;&emsp;&emsp;|-service.proto<br>
&emsp;&emsp;|-svr<br>
&emsp;&emsp;&emsp;&emsp;|-svr_main.go<br>
&emsp;&emsp;&emsp;&emsp;|-handler<br>
&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;|-handler.go<br>
&emsp;&emsp;|-cli_tool<br>
&emsp;&emsp;&emsp;&emsp;|-cli_tool_main.go<br>
&emsp;&emsp;&emsp;&emsp;|-stub<br>
&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;|-stub.go<br>
&emsp;&emsp;build.sh
> 
这是一个空项目，包含一个Echo的rpc接口<br>
此时项目是可以编译通过的，执行build.sh会生成bin文件夹，其中含有一个svrmain的服务器bin程序和一个clitool的客户端调试程序<br>

### 新增一个接口
在项目的根目录下，执行
```go
grpchelper -c addrpc -n [rpc_name]
```
工具会做的事情：
&emsp;&emsp;1、自动在pb/service.proto文件中声明相关message和service
&emsp;&emsp;2、自动在svr/handle/handler.go文件中添加相关rpc函数的声明
&emsp;&emsp;3、自动在cli_tool/stub/stub.go文件中添加相关客户端桩代码，注册对应rpc函数
你只需要做：
&emsp;&emsp;1、在pb/service.proto文件中定义数据结构
&emsp;&emsp;2、在svr/handle/handler.go文件中对应rpc函数内实现业务逻辑

### 客户端调试程序
工具会在bin目录下自动生成名为clitool的客户端调试程序，可以方便对各个rpc接口进行调试
```go
clitool -f rpcname -d 'json_req'
```
