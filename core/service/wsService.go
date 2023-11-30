package service

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
)

var clients []*websocket.Conn

func WsConnect() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		clients = append(clients, c)
		// c.Locals are added to the *websocket.Conn
		//log.Println(c.Locals("allowed"))  // true
		//log.Println(c.Params("id"))       // 123
		//log.Println(c.Query("v"))         // 1.0
		//log.Println(c.Cookies("session")) // ""
		// 当连接关闭时，从广播列表中移除
		defer func() {
			for i, client := range clients {
				if client == c {
					clients = append(clients[:i], clients[i+1:]...)
					break
				}
			}
		}()
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			// 广播消息给所有连接的客户端
			for _, client := range clients {
				if err := client.WriteMessage(mt, msg); err != nil {
					log.Println("write:", err)
					break
				}
			}
		}
	})
}

func SendMsg(msg string) error {
	log.Println("Number of clients:", len(clients))
	for _, client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Println("write:", err)
			return err
		}
	}
	return nil
}
