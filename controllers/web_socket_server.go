package controllers

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"go-git-webhook-client/cache"
	"fmt"
)


var upgrader = websocket.Upgrader{} // use default options




func WebSocketServer(w http.ResponseWriter, r *http.Request)  {

	if valid := verification(r);!valid {
		fmt.Fprint(w,"Permission denied.")
		return
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	
	defer func() {
		c.Close()
		destroy(r)
	}()
	log.Println("Connectioned:",c.RemoteAddr().String())

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

func verification(r *http.Request) bool {

	token := r.Header.Get("x-smarthook-token")

	return cache.TokenCache.Contains(token);
}

func destroy(r *http.Request)  {
	token := r.Header.Get("x-smarthook-token")
	cache.TokenCache.Delete(token)
}