package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/shariquehaider/ecom-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Picture  string             `json:"picture"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"password"`
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
	trimmedObjectId := strings.TrimPrefix(id, "ObjectID(")
	trimmedObjectId = strings.TrimSuffix(trimmedObjectId, ")")
	trimmedObjectId = strings.Trim(trimmedObjectId, `"`)
	objID, err := primitive.ObjectIDFromHex(trimmedObjectId)

	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByUsername(key string) (*User, error) {
	var user User

	isEmail := utils.IsValidEmail(key)
	if isEmail {
		err := collection.FindOne(context.TODO(), bson.M{"email": key}).Decode(&user)
		if err != nil {
			return nil, err
		}
	} else {
		err := collection.FindOne(context.TODO(), bson.M{"username": key}).Decode(&user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func ChangePasswordByID(id string, oldPassword, newPassword string) (*mongo.UpdateResult, error) {
	objID, err := utils.VerifyObjectId(id)
	if err != nil {
		return nil, err
	}

	var user User

	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	isVerified := utils.CompareHashPassword(oldPassword, user.Password)
	if !isVerified {
		return nil, errors.New("Invalid Current Password")
	}

	hashPassword, err := utils.GenerateHashPassword(newPassword)
	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"password": hashPassword}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateProfileByID(id string, user User) (*mongo.UpdateResult, error) {
	objId, err := utils.VerifyObjectId(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{"$set": bson.M{
		"profile": user.Picture,
		"name":    user.Name,
	}}
	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objId}, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}
