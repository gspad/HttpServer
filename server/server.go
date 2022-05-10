package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (server *PlayerServer) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	playerScore := server.store.GetPlayerScore(player)
	fmt.Fprint(writer, playerScore)
}
