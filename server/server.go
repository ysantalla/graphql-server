package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	gql "github.com/ysantalla/graphql-server/gql"
	"github.com/ysantalla/graphql-server/resolver"
)

const addr = "0.0.0.0:5000"

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the gql/generated.go file
	// Resolver is in the resolver/resolver.go file
	h := handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolver.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	/*ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("graphql_db")

	db.Collection("user")

	fmt.Println("Connected to MongoDB!")*/

	r := gin.Default()
	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(addr)

	log.Printf("Connect to http://%s/ for GraphQL playground", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
