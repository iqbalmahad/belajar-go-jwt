package main

import (
	"fmt"
	"net/http"

	"github.com/iqbalmahad/belajar-go-jwt/internal/jwt"
)

func main() {
	// Setup HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// Generate JWT token
	token, err := jwt.GenerateJWTToken("username", []byte("your_secret_key"))
	if err != nil {
		fmt.Println("Error generating JWT token:", err)
		return
	}

	fmt.Println("JWT Token:", token)

	// Verify JWT token
	claims, err := jwt.VerifyJWTToken(token, []byte("your_secret_key"))
	if err != nil {
		fmt.Println("Error verifying JWT token:", err)
		return
	}

	fmt.Println("Username from token:", claims.Username)

	// Start HTTP server
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
