package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Picture  string             `json:"picture"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"-"`
	Name     string             `json:"name,omitempty"`
	Username string             `json:"username,omitempty"`
}

const connectionString string = "mongodb+srv://sharique:heysharique123@cluster0.8o2va.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName string = "Ecom-express"
const colName string = "user"

var collection *mongo.Collection

func InitDB() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB is Connected", client)

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instant is ready")
}

func CreateUser(user User) error {
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func FindById(id string) (*User, error) {
	var user User
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByUsername(email string) (*User, error) {
	var user User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
