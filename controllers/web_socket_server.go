package controllers

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"github.com/lifei6671/go-git-webhook-client/cache"
	"github.com/lifei6671/go-git-webhook-client/models"
	"github.com/lifei6671/go-git-webhook-client/commands"
	"fmt"
)


var upgrader = websocket.Upgrader{CheckOrigin : verification} // use default options


func WebSocketServer(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover();err != nil {
			log.Printf("%+v",err)
		}
	}()
	isWebSocketClosed := false

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	c.SetCloseHandler(func(code int, text string) error {
		isWebSocketClosed = true
		log.Println("code:", code, ";text:", text)
		destroy(r)
		return nil
	})
	defer func() {
		c.Close()

	}()
	log.Println("Connectioned:", c.RemoteAddr().String())

	for {
		if isWebSocketClosed {
			return
		}

		var body models.JsonResult

		err := c.ReadJSON(&body)
		if err != nil {
			if websocket.IsCloseError(err,1005) {
				log.Printf("WebSocket client closed:%s",err)
				break
			}

			res := models.JsonResult{
				ErrorCode: 5002,
				Message: "Parameter format error.",
			}
			err = c.WriteJSON(res)
			if err != nil {
				log.Printf("write error 4001:%s", err)
			}

			continue
		}

		if body.Command == "shell" {
			command := body.Data.(string)

			channel := make(chan []byte, 10)

			go commands.Command(command, channel)

			isChannelClosed := false
			for {
				if isChannelClosed {
					res := models.JsonResult{
						ErrorCode:0,
						Message:"ok",
						MsgId:body.MsgId,
						Command: "end",
					}

					err = c.WriteJSON(res)

					if err != nil {
						log.Println("write:", err)
						break
					}
					break
				}
				select {
				case out, ok := <-channel:
					{
						if !ok {
							log.Println("chan closed")
							isChannelClosed = true
							break
						}
						if len(out) > 0 {
							res := models.JsonResult{
								ErrorCode:0,
								Message:"ok",
								MsgId:body.MsgId,
								Command: body.Command,
								Data: string(out),
							}

							err = c.WriteJSON(res)

							if err != nil {
								log.Println("write:", err)
								break
							}
						}
					}
				}
			}

			log.Println("Command execute result")
		}else{
			res := models.JsonResult{
				ErrorCode: 5002,
				Message: "Command not support.",
			}
			err = c.WriteJSON(res)
			if err != nil {
				log.Println("write:", err.Error())
			}

			continue
		}

	}
}

//校验是否是授权客户端
func verification(r *http.Request) bool {

	token := r.Header.Get("x-smarthook-token")

	return cache.TokenCache.Contains(token);

}
//销毁已授权的Token
func destroy(r *http.Request)  {
	token := r.Header.Get("x-smarthook-token")
	cache.TokenCache.Delete(token)
	fmt.Println("websocket closed.token:",token)
}

