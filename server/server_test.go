package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	store map[string]int
}

func (ps *StubPlayerStore) GetPlayerScore(name string) int {
	return ps.store[name]
}

func TestGETPlayers(t *testing.T) {

	playerStore := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Pablo":  10,
		},
	}

	t.Run("returns Pepper's score", func(t *testing.T) {
		server := &PlayerServer{playerStore}
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		responseRecorder := httptest.NewRecorder()
		server.ServeHTTP(responseRecorder, request)

		got := responseRecorder.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("returns Pablo's score", func(t *testing.T) {
		server := &PlayerServer{playerStore}
		request, _ := http.NewRequest(http.MethodGet, "/players/Pablo", nil)
		responseRecorder := httptest.NewRecorder()
		server.ServeHTTP(responseRecorder, request)

		got := responseRecorder.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}
