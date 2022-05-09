package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(writer http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprintf(writer, "20")
	}

	if player == "Pablo" {
		fmt.Fprintf(writer, "10")
	}
}
