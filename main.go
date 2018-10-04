package main
import (
	"log"
	"github.com/ignis/galvatron/core"
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"

func main(){
	conf := core.GetConf();
	db, err := gorm.Open("mysql", conf.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

    core.NewServer(conf, db);
}