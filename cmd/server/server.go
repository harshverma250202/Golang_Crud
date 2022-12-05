package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/harsh/project/internal/routes"
	"github.com/harsh/project/supertoken"
	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Server() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	supertoken.SuperTokens()
	fmt.Println("Server started at port 8000")
	router := routes.Routes()
	// log.Fatal(http.ListenAndServe("localhost:3000", router))

	http.ListenAndServe("localhost:8000", handlers.CORS(
		handlers.AllowedHeaders(append([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"},
			supertokens.GetAllCORSHeaders()...)),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowCredentials(),
	)(supertokens.Middleware(router)))
}
