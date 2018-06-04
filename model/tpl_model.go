package model

import (
	"time"
)

type TplModel struct {
	ModuleName          string    `json:"moduleName"` // create时入参的rawName
	DirName             string    `json:"dirName"`    // rawName去掉_, 全部转为小写
	ProjName            string    `json:"moduleName"` // rawName取消下划线且转为驼峰
	Date                time.Time `json:"time"`
	SvrIp               string    `json:"svrIp"`
	SvrPort             uint32    `Json:"svrPort"`
	PrefixFromGoSrcPath string    `Json:"prefixFromGoSrcPath"`
}

func NewTplModel() *TplModel {
	return &TplModel{}
}
