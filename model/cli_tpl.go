package model

import (
	"time"
)

type CliTplModel struct {
	ModuleName       string    `json:"moduleName"`
	Date             time.Time `json:"time"`
	SvrIp            string    `json:"svrIp"`
	SvrPort          uint32    `Json:"svrPort"`
	PrefixFromGoPath string    `Json:"prefixFromGoPath"`
}
