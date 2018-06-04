package comm

import (
	"errors"
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

func GetTplPath4Const() string {
	return GetTplPath() + "/const.tpl"
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
		return false
	}
}

func MakeDir(path string) error {
	if PathExist(path) {
		return nil
	}

	return os.MkdirAll(path, 0755)
}

func IsValidModuleName(moduleName string) bool {
	pattern := `[^a-zA-Z_]+`
	return !regexp.MustCompile(pattern).MatchString(moduleName)
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

func ModuleName2ProjName(moduleName string) (string, error) {
	if IsValidModuleName(moduleName) == false {
		return "", errors.New("Invalid moduleName")
	}

	ret := CapitalizeStr(moduleName)
	if len(ret) == 0 {
		return "", errors.New("Invalid moduleName")
	}

	return ret, nil
}

func ModuleName2DirName(moduleName string) (string, error) {
	if IsValidModuleName(moduleName) == false {
		return "", errors.New("Invalid moduleName")
	}

	ret := strings.Replace(moduleName, "_", "", -1)
	if len(ret) == 0 {
		return "", errors.New("Invalid moduleName")
	}

	return ret, nil
}
