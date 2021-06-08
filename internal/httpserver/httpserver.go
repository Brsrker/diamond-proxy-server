package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"brsrker.com/diamond/proxyserver/internal/logger"
	"brsrker.com/diamond/proxyserver/internal/websocket"
)

const TAG = "httpserver"

//Server port
const port = "2020"

func Start() error {

	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/health", healthCheckHandler).Methods("GET").Schemes("http")
	router.HandleFunc("/ws", socketConnection)

	logger.Info(TAG, fmt.Sprintf("Server Started at port: %v", port))

	return http.ListenAndServe(":"+port, handlers.CORS(headers, methods, origins)(router))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"alive": true})
}

func socketConnection(w http.ResponseWriter, r *http.Request) {
	var err error
	var upgrader = websocket.Upgrader
	client, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	websocket.Client = client

	websocket.ReceiveMessage()
}
