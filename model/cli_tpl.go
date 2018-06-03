package model

type CliTplModel struct {
	ModuleName string `json:"moduleName"`
	Date       string `json:"time"`
	SvrIp      string `json:"svrIp"`
	SvrPort    string `Json:"svrPort"`
	PathPrefix string `Json:"pathPrefix"`
}
