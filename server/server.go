package server

import (
	"fmt"
	"net/http"
)

func run_server() {

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		// Ручка для получения токена
		guid := r.FormValue("guid")
		ip := r.RemoteAddr

		fmt.Fprintf(w, "guid: %s ip: %s", guid, ip)

	})

	http.HandleFunc("/refresh_token", func(w http.ResponseWriter, r *http.Request) {
		// Ручка для обновления токена
		fmt.Fprint(w, "Refresh token")
	})

	fmt.Println("Server is listening http://127.0.0.1:8181")
	http.ListenAndServe(":8181", nil)
}
