package comm

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func Time2Date(t time.Time) string {
	return t.Format("2006-01-02")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetGoPath() string {
	return GetEnv("GOPATH")
}

func GetGoSrcPath() string {
	goPath := GetGoPath()
	if len(goPath) == 0 {
		panic("Not Found GOPATH")
	}
	return goPath + "/src"
}

// 获取模版路径
func GetTplPath() string {
	return GetGoSrcPath() + "/github.com/eyesofblue/grpchelper/tpl"
}

func GetTplPath4Pb() string {
	return GetTplPath() + "/proto.tpl"
}

func GetTplPath4Svr() string {
	return GetTplPath() + "/svr.tpl"
}

func GetTplPath4Cli() string {
	return GetTplPath() + "/cli.tpl"
}

func GetTplPath4Handler() string {
	return GetTplPath() + "/handler.tpl"
}

func GetTplPath4Stub() string {
	return GetTplPath() + "/stub.tpl"
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetPrefixFromGoSrcPath() (string, error) {
	goSrcPath := GetGoSrcPath()
	currentDir := GetCurrentDirectory()
	index := strings.Index(currentDir, goSrcPath)

	if index != 0 {
		return "", errors.New("Current Dir Not Under GOPATH")
	}

	ret := currentDir[len(goSrcPath):]
	if len(ret) > 0 && ret[0] == '/' {
		ret = ret[1:]
	}

	return string(ret), nil
}

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		if os.IsExist(err) {
			return true
		}
		return false
	}
}

func MakeDir(path string) error {
	if PathExist(path) {
		return nil
	}

	return os.MkdirAll(path, 0755)
}

func IsValidRawName(rawName string) bool {
	pattern := `[^a-zA-Z_]+`
	return !regexp.MustCompile(pattern).MatchString(rawName)
}

//abc_def_g ->AbcDefG
func CapitalizeStr(str string) string {
	splitList := strings.Split(str, "_")
	var upperStr string
	for i := 0; i != len(splitList); i++ {
		item := []byte(splitList[i])
		if len(item) > 0 {
			if item[0] >= 'a' {
				upperStr += string(item[0] - 32)
			} else {
				upperStr += string(item[0])
			}
			if len(item) > 1 {
				upperStr += string(item[1:])
			}
		}
	}
	return upperStr
}

func RawName2ProjName(moduleName string) (string, error) {
	if IsValidRawName(moduleName) == false {
		return "", errors.New("Invalid moduleName")
	}

	ret := CapitalizeStr(moduleName)
	if len(ret) == 0 {
		return "", errors.New("Invalid moduleName")
	}

	return ret, nil
}

func RawName2DirName(moduleName string) (string, error) {
	if IsValidRawName(moduleName) == false {
		return "", errors.New("Invalid moduleName")
	}

	ret := strings.Replace(moduleName, "_", "", -1)
	if len(ret) == 0 {
		return "", errors.New("Invalid moduleName")
	}

	return ret, nil
}

// 获取项目目录
func GetMainDir(rawName string) string {
	dirName, err := RawName2DirName(rawName)
	if err != nil {
		panic(err)
	}
	currentDir := GetCurrentDirectory()
	mainDir := currentDir + "/" + dirName
	return mainDir
}

func GetPbDir(mainDir string) string {
	return mainDir + "/pb"
}

func GetSvrDir(mainDir string) string {
	return mainDir + "/svr"
}

func GetCliToolDir(mainDir string) string {
	return mainDir + "/cli_tool"
}

func GetHandlerDir(mainDir string) string {
	return mainDir + "/svr/handler"
}

func GetStubDir(mainDir string) string {
	return mainDir + "/cli_tool/stub"
}

// 获取PB相关文件
func GetPbFilePath(pbDir string) string {
	return pbDir + "/service.proto"
}

// 获取Svr相关文件
func GetSvrMainFilePath(svrDir string) string {
	return svrDir + "/svr_main.go"
}

func GetHandlerFilePath(handlerDir string) string {
	return handlerDir + "/handler.go"
}

// 获取CliTool相关文件
func GetCliToolMainFilePath(cliToolDir string) string {
	return cliToolDir + "/cli_tool_main.go"
}

func GetStubFilePath(stubDir string) string {
	return stubDir + "/stub.go"
}

// 获取rpc入参出参名称
func GetRpcReqName(rpcName string) string {
	return rpcName + "Request"
}

func GetRpcRspName(rpcName string) string {
	return rpcName + "Response"
}

// 获取各种锚定Tag
func GetTagSegBegin4PbMsg() string {
	return fmt.Sprintf(TAG_SEGMENT_BEGIN_TMPL, "PB", "MESSAGE")
}

func GetTagSegEnd4PbMsg() string {
	return fmt.Sprintf(TAG_SEGMENT_END_TMPL, "PB", "MESSAGE")
}

func GetTagSegBegin4PbService() string {
	return fmt.Sprintf(TAG_SEGMENT_BEGIN_TMPL, "PB", "SERVICE")
}

func GetTagSegEnd4PbService() string {
	return fmt.Sprintf(TAG_SEGMENT_END_TMPL, "PB", "SERVICE")
}

func GetTagSegBegin4HandlerImpl() string {
	return fmt.Sprintf(TAG_SEGMENT_BEGIN_TMPL, "HANDLER", "IMPL")
}

func GetTagSegEnd4HandlerImpl() string {
	return fmt.Sprintf(TAG_SEGMENT_END_TMPL, "HANDLER", "IMPL")
}

// 获取PB内容模版
func GetContentTmpl4PbMsg(rpcName string) string {
	ret := fmt.Sprintf(CONTENT_TMPL_PB_MSG, GetRpcReqName(rpcName)) + "\n" + fmt.Sprintf(CONTENT_TMPL_PB_MSG, GetRpcRspName(rpcName))
	return ret
}

func GetContentTmpl4PbService(rpcName string) string {
	return fmt.Sprintf(CONTENT_TMPL_PB_SERVICE, rpcName, GetRpcReqName(rpcName), GetRpcRspName(rpcName))
}
