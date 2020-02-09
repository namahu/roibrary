package main

import (
	"net/http"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"

	"fmt"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func getBook(c *gin.Context) {
	c.Writer.Write([]byte("Hello\n"))
}

func initializeFirebase() *firebase.App {
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_DEV"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	return app
}

func authMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {

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
			c.Abort()
			return
		}
		fmt.Printf("Verify ID Token: %v\n", token)
	}
}

func firestoreMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := app.Firestore(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		defer client.Close()

		iter := client.Collection("books").Documents(context.Background())
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Printf("Faild to iterate: %v", err)
			}
			fmt.Println(doc.Data())
		}
	}
}

func main() {
	router := gin.Default()

	app := initializeFirebase()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}

	router.Use(cors.New(config))
	// router.Use(authMiddleware(app))

	router.GET("/getbook", firestoreMiddleware(app), getBook)

	router.Run(":8081")

}
