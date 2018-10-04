package core
import (
	"io/ioutil"
	"log"
	"github.com/googollee/go-engine.io"
	"github.com/ignis/galvatron/handlers"
)

type Client struct {
	conn *engineio.Conn
}

func NewClient(conn engineio.Conn) *Client {
	c := Client{&conn}

	go func(){
		defer conn.Close()
		for {
			t, r, err := conn.NextReader()
			if err != nil {
				log.Fatal(err)
			}
			b, _ := ioutil.ReadAll(r)
			log.Println(t, string(b))
			handlers.Handle(b)
		}
	}()

	return &c
}

func (c *Client) sendMessage(buffer []byte){
	w, _ := c.conn.NextWriter(engineio.MessageText)

	w.Write(buffer)
	w.Close()
}