package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func getBook(c *gin.Context) {
	c.String(200, "Hello")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		opt := option.WithCredentialFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Prinf("error: %v\n", err)
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		authHeader := router.Header.Get("Authorization")
		idToken := String.Replace(authHeader, "Bearer ", "", 1)
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Printf("error Verifying ID Token: %v\n", err)
		}
	}
}

func main() {
	router := gin.Default()
	router.Use(authMiddleware)

	router.GET("/getbook", getBook)

	router.Run()

}