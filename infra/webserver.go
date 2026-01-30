package infra

import (
	"fmt"
	"net/http"

	"diegorosa.com/poker-game/core"
)

var (
	// TODO: map with route to method ex. {"/api/deck": core.ApiDeck}
	routes = []string{
		"/api",
		"/api/cards",
	}
)

func ServerStart() {
	fmt.Println("Starting server: localhost:8080")

	for _, route := range routes {
		fmt.Printf("route: %s", route)
	}

	http.HandleFunc("/api/poker", func(resp http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(resp, "Counter")
	})

	http.HandleFunc("/api/poker/deck/dealer", core.DealerCards)

	er := http.ListenAndServe("localhost:8080", nil)
	if er != nil {
		fmt.Printf("Error %s", er.Error())
	}
}
