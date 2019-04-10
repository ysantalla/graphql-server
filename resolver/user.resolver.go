package resolver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/ysantalla/graphql-server/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://localhost:27017"

// DBNAME Database name
const DBNAME = "graphqlDB"

// COLLNAME Collection name
const COLLNAME = "user"

var db *mongo.Database

// Connect establish a connection to database
func init() {

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

func (r *mutationResolver) CreateUser(ctx context.Context, data model.UserCreateInput) (*model.User, error) {

	result, err := db.Collection(COLLNAME).InsertOne(context.Background(), data)

	if err != nil {
		return nil, err
	}

	findOneResult := db.Collection(COLLNAME).FindOne(context.Background(), bson.D{{"_id", result.InsertedID}})

	var user model.User

	if findOneResult != nil {
		findOneResult.Decode(&user)
		return &user, nil
	}

	return nil, errors.New("Error new user not inserted")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, data model.UserUpdateInput, where model.UserWhereUniqueInput) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, where model.UserWhereUniqueInput) (*model.User, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, where model.UserWhereUniqueInput) (*model.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context, where *model.UserWhereInput, orderBy *model.UserOrderByInput, skip *int, after *string, before *string, first *int, last *int) ([]model.User, error) {

	if where.Name == nil {
		return nil, errors.New("Field name not found")
	}

	fmt.Println("esto es una prueba")

	cur, err := db.Collection(COLLNAME).Find(context.Background(),
		bson.D{{"name", where.Name}}, nil)

	if err != nil {
		return nil, err
	}

	var elements []model.User
	var elem model.User
	// Get the next result from the cursor
	for cur.Next(context.Background()) {

		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements, nil
}
