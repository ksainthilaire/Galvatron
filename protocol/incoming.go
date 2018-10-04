package protocol

type Connecting struct {
	Language		string		`json:"language"`
	Client 			string		`json:"client"`
	AppVersion		float64		`json:"appVersion"`
	BuildVersion	float64		`json:"buildVersion"`
}

type Disconnecting string

type CheckAssetsVersion struct {
	Version 		float64 `json:"version"`
	AssetsVersion 	float64 `json:"assetsVersion"`
}

type Login struct {
	Username	string	`json:"username"`
	Token		string	`json:"password"`
	Salt		string	`json:"salt"`
	Key			string 	`json:"key"`
}