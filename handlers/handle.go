package handlers
import (
	"encoding/json"
	"github.com/ignis/galvatron/protocol"
)

type Data struct {
	Call string				`json:"call"`
	Data json.RawMessage 	`json:"data"`
}

func Handle(raw []byte) {
	v := Data{}

	err := json.Unmarshal([]byte(raw), &v)
	if err != nil {
		panic(err)
	}
	switch v.Call {

		case "connecting": {	
			data := protocol.Connecting{}

			err := json.Unmarshal(v.Data, &data)
			if err != nil {
				panic(err)
			}
			Connecting(data)
		}

		case "disconnecting": {
			data := protocol.Disconnecting(v.Data)
			Disconnecting(data)
		}

		case "login": {
			data := protocol.Login{}

			err := json.Unmarshal(v.Data, &data)
			if err != nil {
				panic(err)
			}
			Login(data)
		}

		case "checkAssetsVersion": {
			data := protocol.CheckAssetsVersion{}

			err := json.Unmarshal(v.Data, &data)
			if err != nil {
				panic(err)
			}
			CheckAssetsVersion(data)
		}
	}
}