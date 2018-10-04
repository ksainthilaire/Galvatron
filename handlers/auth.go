package handlers
import (
	"fmt"
	"github.com/ignis/galvatron/protocol"
)

func 	Login(data protocol.Login){

}

func 	Connecting(data protocol.Connecting){
	fmt.Printf("Version de l'application: %d", data.AppVersion)
}

func 	Disconnecting(data protocol.Disconnecting){

}

func 	CheckAssetsVersion(data protocol.CheckAssetsVersion){

}