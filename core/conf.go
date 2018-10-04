package core
import (
	"os"
	"encoding/json"
	"log"
	"github.com/rs/cors"
)

type Conf struct {
	Port	string
	Path	string
	Cors	cors.Options
	DBUrl	string
}

func GetConf() Conf {
	var conf Conf

	file, err := os.Open("settings.json")
	if err != nil {
		log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, fi.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}