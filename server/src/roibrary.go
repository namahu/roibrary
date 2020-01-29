package main

import (
	"net/http"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"

	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"os"
	"strings"
)

func getBook(c *gin.Context) {
	c.Writer.Write([]byte("Hello\n"))
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_DEV"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Printf("error Verifying ID Token: %v\n", err)
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("error verifying ID token\n"))
			return
		}
		fmt.Printf("Verify ID Token: %v\n", token)
	}
}

func main() {
	router := gin.Default()
	router.Use(authMiddleware())

	router.GET("/getbook", getBook)

	router.Run()

}
