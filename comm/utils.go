package comm

import (
	"errors"
	"os"
	"path/filepath"
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
	return goPath + "/src"
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
