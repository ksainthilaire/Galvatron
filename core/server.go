package core
import (
	"net/http"
	"log"
	"github.com/rs/cors"
	"github.com/googollee/go-engine.io"
	"github.com/jinzhu/gorm"
)

type Server struct {
	conf 	Conf
	db 		*gorm.DB
	clients []*Client
}

func 	NewServer(conf Conf, db *gorm.DB) *Server {
	ctx := Server{}
	ctx.conf = conf;
	ctx.db = db;
	server, err := engineio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}
	go func(){
		for {
			conn, err := server.Accept()

			if err != nil {
				log.Fatal(err)
			}
			client := NewClient(conn)
			ctx.clients = append(ctx.clients, client)
		}
	}()
	c := cors.New(conf.Cors)
	http.Handle(conf.Path, server)
	http.ListenAndServe(conf.Port, c.Handler(server))
	return &ctx
}


func (ctx *Server) bar(){

}