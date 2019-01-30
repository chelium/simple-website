package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chelium/simple-website/mongo"
	"github.com/chelium/simple-website/server"
	todo "github.com/chelium/simple-website/todo/service"
	user "github.com/chelium/simple-website/user/service"

	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	jose "gopkg.in/square/go-jose.v2"
)

const (
	defaultPort       = ":8080"
	defaultMongoDBURL = "127.0.0.1:27017"
	defaultDBName     = "simplewebsite"
)

var (
	audience string
	domain   string
)

func main() {
	var (
		addr   = envString("PORT", defaultPort)
		dburl  = envString("MONGODB_URL", defaultMongoDBURL)
		dbname = envString("DB_NAME", defaultDBName)
	)

	session, err := mgo.Dial(dburl)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var (
		todos, _ = mongo.NewTodoRepository(dbname, session)
		users, _ = mongo.NewUserRepository(dbname, session)
	)

	var us user.Service
	us = user.NewService(todos, users)

	var ts todo.Service
	ts = todo.NewService(todos, users)

	setAuth0Variables()
	//	r := gin.Default()

	// authorized := r.Group("/")
	// authorized.Use(authRequired())
	srv := server.New(ts, us)

	srv.Router.Run(addr)
}

func setAuth0Variables() {
	audience = os.Getenv("AUTH0_API_IDENTIFIER")
	domain = os.Getenv("AUTH0_DOMAIN")
}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var auth0Domain = "https://" + domain + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{audience}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(c.Request)

		if err != nil {
			log.Println(err)
			terminateWithError(http.StatusUnauthorized, "token is not valid", c)
			return
		}
		c.Next()
	}
}

func terminateWithError(statusCode int, message string, c *gin.Context) {
	c.JSON(statusCode, gin.H{"error": message})
	c.Abort()
}

func envString(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	return val
}
