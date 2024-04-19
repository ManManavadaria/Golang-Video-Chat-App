package server

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func NewRoomMap() *RoomMap {
	return &RoomMap{
		Map: make(map[string][]Participant),
	}
}

func (r *RoomMap) CreatRoom() string {

	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 15)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	roomId := string(b)

	r.Map[roomId] = []Participant{}

	return roomId
}

func (r *RoomMap) InsertIntoRoom(roomId string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	p := Participant{Host: host, Conn: conn}

	r.Map[roomId] = append(r.Map[roomId], p)
}

func (r *RoomMap) DeleteRoom(roomId string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomId)
}

func (r *RoomMap) Get(roomid string) []Participant {
	return r.Map[roomid]
}
