package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(c *mongo.Client) Models {
	client = c
	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (l *LogEntry) Insert(entry LogEntry) error {
	collection := client.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		fmt.Println("Error inserting into logs: ", err)
		return err
	}

	return nil
}

func (l *LogEntry) All() ([]*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	collection := client.Database("logs").Collection("logs")

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		fmt.Println("Error getting logs from MongoDB: ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []*LogEntry

	for cursor.Next(ctx) {
		var log LogEntry
		err := cursor.Decode(&log)
		if err != nil {
			fmt.Println("Error decoding data from MongoDB: ", err)
			return nil, err
		}

		logs = append(logs, &log)
	}

	return logs, nil
}

func (l *LogEntry) GetLogEntryByID(id string) (*LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	collection := client.Database("logs").Collection("logs")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error getting log id from MongoDB: ", err)
		return nil, err
	}

	var log LogEntry
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&log)
	if err != nil {
		fmt.Println("Error getting log from MongoDB: ", err)
		return nil, err
	}

	return &log, nil
}

func (l *LogEntry) Drop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	collection := client.Database("logs").Collection("logs")

	err := collection.Drop(ctx)
	if err != nil {
		fmt.Println("Error dropping logs from MongoDB: ", err)
		return err
	}

	return nil
}

func (l *LogEntry) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	collection := client.Database("logs").Collection("logs")

	docID, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		fmt.Println("Error getting log id from MongoDB: ", err)
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: l.Name},
				{Key: "data", Value: l.Data},
				{Key: "updated_at", Value: time.Now()},
			}},
		},
	)
	if err != nil {
		fmt.Println("Error updating log in MongoDB: ", err)
		return nil, err
	}

	return result, nil
}
