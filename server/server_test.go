package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		responseRecorder := httptest.NewRecorder()
		PlayerServer(responseRecorder, request)

		got := responseRecorder.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("returns Pablo's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pablo", nil)
		responseRecorder := httptest.NewRecorder()
		PlayerServer(responseRecorder, request)

		got := responseRecorder.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}
