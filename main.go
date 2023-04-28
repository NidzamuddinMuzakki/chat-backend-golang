package main

import (
	"encoding/json"
	"log"

	"github.com/NidzamuddinMuzakki/chat-golang-backend/configs"
	"github.com/NidzamuddinMuzakki/chat-golang-backend/controllers"
	"github.com/NidzamuddinMuzakki/chat-golang-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

type MessageObject struct {
	Data string `json:"data"`
	From string `json:"from"`
}

func main() {
	// clients := make(map[string]string)
	app := fiber.New()

	configs.Connect()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/chats", controllers.GetChats)

	app.Get("/user/:id", controllers.GetUser)
	app.Use("/addUser", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	// ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
	// 	fmt.Println(fmt.Printf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("user_id")))
	// })

	// ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
	// 	fmt.Println(fmt.Printf("Connection event 2 - User: %s", ep.Kws.GetStringAttribute("user_id")))
	// })

	// // On message event
	// ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {

	// 	fmt.Println(fmt.Printf("Message event - User: %s - Message: %s", ep.Kws.GetStringAttribute("user_id"), string(ep.Data)))

	// 	message := MessageObject{}

	// 	// Unmarshal the json message
	// 	// {
	// 	//  "from": "<user-id>",
	// 	//  "to": "<recipient-user-id>",
	// 	//  "data": "hello"
	// 	//}
	// 	err := json.Unmarshal(ep.Data, &message)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	// Emit the message directly to specified user
	// 	for _, s := range clients {
	// 		log.Println(s, string(ep.Data))
	// 		err = ep.Kws.EmitTo(s, ep.Data)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}
	// })

	// // On disconnect event
	// ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
	// 	// Remove the user from the local clients
	// 	delete(clients, ep.Kws.GetStringAttribute("user_id"))
	// 	fmt.Println(fmt.Printf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	// })

	// // On close event
	// // This event is called when the server disconnects the user actively with .Close() method
	// ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
	// 	// Remove the user from the local clients
	// 	delete(clients, ep.Kws.GetStringAttribute("user_id"))
	// 	fmt.Println(fmt.Printf("Close event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	// })

	// // On error event
	// ikisocket.On(ikisocket.EventError, func(ep *ikisocket.EventPayload) {
	// 	fmt.Println(fmt.Printf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	// })

	app.Get("/addUser/:name", websocket.New(func(c *websocket.Conn) {
		// userId := kws.Params("name")

		// // Add the connection to the list of the connected clients
		// // The UUID is generated randomly and is the key that allow
		// // ikisocket to manage Emit/EmitTo/Broadcast
		// clients[userId] = kws.UUID

		// // Every websocket connection has an optional session key => value storage
		// kws.SetAttribute("user_id", userId)

		// //Broadcast to all the connected users the newcomer
		// kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true)
		// //Write welcome message
		// kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)))
		user := new(models.User)

		// data := responses.UserResponse{
		// 	Status: fiber.ErrBadRequest.Code,

		// 	Data: nil,
		// }
		name := c.Params("name")
		log.Println(name)

		var (
			mt  int
			msg []byte
			err error
		)
		for {
			datas := configs.Database.First(&user, "name=?", name)
			log.Println(datas.RowsAffected)
			if datas.RowsAffected == 0 {
				break

			}
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			if string(msg[:]) == "first" {

				log.Printf("recv: %s", msg)
				msg = []byte("welcome..." + name)
				if err = c.WriteMessage(mt, msg); err != nil {
					log.Println("write:", err)
					break
				}
			} else {
				var chat models.Chat
				var pods map[string]string
				json.Unmarshal(msg, &pods)
				chat.Name = pods["name"]
				chat.Message = pods["message"]
				configs.Database.Create(&chat)
				log.Printf("recv: %s", msg)
				// msg = []byte(pods["name"] + "|" + pods["message"])
				// err := c.WriteJSON(pods)
				// if err != nil {
				// 	log.Println(err)
				// 	break
				// }
				if err = c.WriteMessage(1, msg); err != nil {
					log.Println("write:", err)
					break
				}
			}
		}

	}))

	log.Fatal(app.Listen(":8080"))

}
