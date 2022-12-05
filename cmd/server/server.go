package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/harsh/project/internal/routes"
	"github.com/harsh/project/supertoken"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Server() {
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
