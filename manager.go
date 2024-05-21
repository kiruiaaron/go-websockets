package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:1024,
		WriteBufferSize:1024,
	}
)
type Manager struct{

}

func NewManager()*Manager{
	return &Manager{}
}

func (m *Manager) serveWS(w http.ResponseWriter, r*http.Request){
	log.Println("New Connection")
	//upgrade regular http connection into websocket
	conn, err := webSocketUpgrader.Upgrade(w, r, nil)

	if err != nil{
		log.Println(err)
		return
	}

	conn.Close()
}