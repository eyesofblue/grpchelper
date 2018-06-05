/*
 * author: jinwei
 * date  : {{Time2Date .Date}}
 */

package main

import (
	// "golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"os"
    "strconv"
    log "github.com/cihub/seelog"
    "github.com/eyesofblue/grpchelper/public"
	"{{.PrefixFromGoSrcPath}}/{{.DirName}}/pb"
)

const (
    SVR_IP = "{{.SvrIp}}"
    SVR_PORT = {{.SvrPort}}
)

func init() {
	logger, err := log.LoggerFromConfigAsString(public.FileLogConfig)
	if err != nil {
		log.Critical("Failed to parse config params")
		return
	}
	log.ReplaceLogger(logger)
}

type server struct{}

/*
func (s *server) RawSend(ctx context.Context, in *pb.RawSendReq) (*pb.RawSendRsp, error) {
	// TODO
	return logic.DoSend(in)
}
*/

func main() {
	defer log.Flush()

	// 创建监听端口
    address := SVR_IP + ":" + strconv.FormatUint(uint64(SVR_PORT), 10)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Errorf("Failed to Listen. msg:%s", err)
		os.Exit(-1)
	}
	// 创建grpc服务器实例
	grpcSvr := grpc.NewServer()
	// 将自己的服务器实现注册到grpc服务器上
	pb.Register{{.ProjName}}Server(grpcSvr, &server{})
	// 开启服务
    log.Tracef("Sever Start. Address:%s", address)
	grpcSvr.Serve(lis)
}
