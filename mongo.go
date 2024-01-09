package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func init() {
	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, _ = mongo.Connect(context.Background(), clientOptions)

	collection = client.Database("your-db-name").Collection("people")
}

// InsertPerson menyimpan data diri ke MongoDB
func InsertPerson(person Person) (interface{}, error) {
	// Set timestamp
	person.CreatedAt = time.Now()

	// Simpan data diri ke MongoDB
	result, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result.InsertedID, nil
}

// FindPersonByQRCode mencari data diri berdasarkan QR Code
func FindPersonByQRCode(qrCode string) (*Person, error) {
	var person Person
	err := collection.FindOne(context.Background(), bson.M{"qrcode": qrCode}).Decode(&person)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

// UpdateMaritalStatus memperbarui status pernikahan
func UpdateMaritalStatus(personID string) error {
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": personID}, bson.M{"$set": bson.M{"maritalStatus": "Sudah Menikah"}})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
