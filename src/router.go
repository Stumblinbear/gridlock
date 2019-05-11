package gridlock

import (
	"flag"
	"log"
	"net/http"
	"encoding/json"
	"errors"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:1329", "gridlock service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var INVALID_METHOD error = errors.New("Method not allowed.")

func RegisterRoutes(g *Gridlock) {
	// Register the websocket server
	http.HandleFunc("/ws", handleWebSocket)

	g.AddEndpoint("v1/system", func(r *http.Request) (int, interface{}) {
		if r.Method != "GET" {
			return 405, INVALID_METHOD
		}

		return 200, g.system
	})

	g.AddEndpoint("v1/library", func(r *http.Request) (int, interface{}) {
		return 200, "hi bb"
	})

	g.AddEndpoint("v1/test", func(r *http.Request) (int, interface{}) {
		var param interface{}
		if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
			return 400, err
		}

		return 200, "hi bb"
	})

	log.Println("Started RPC/WebSocket HTTP server on:", *addr)
	log.Println("    WebSocket available at /ws")

	log.Println("")
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
