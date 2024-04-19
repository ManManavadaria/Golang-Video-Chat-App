package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms = NewRoomMap()

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.CreatRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}

	log.Println(AllRooms.Map)
	json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type BroadcastMsg struct {
	msg    map[string]interface{}
	roomId string
	Client *websocket.Conn
}

var Broadcast = make(chan BroadcastMsg)

func Broadcaster() {
	for {
		msg := <-Broadcast

		for _, Client := range AllRooms.Map[msg.roomId] {
			if Client.Conn != msg.Client {
				err := Client.Conn.WriteJSON(msg.msg)

				if err != nil {
					log.Fatal(err)
					Client.Conn.Close()
				}
			}
		}
	}
}

func JoinRoomReqHandler(w http.ResponseWriter, r *http.Request) {
	roomid := r.URL.Query().Get("roomID")

	if roomid == "" {
		log.Println("roomid missing")
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("web socket upgrad error")
	}

	AllRooms.InsertIntoRoom(roomid, false, ws)

	go Broadcaster()

	for {
		var msg BroadcastMsg

		err := ws.ReadJSON(&msg.msg)

		if err != nil {
			log.Fatal("read error")
		}

		msg.Client = ws

		msg.roomId = roomid

		Broadcast <- msg
	}
}
