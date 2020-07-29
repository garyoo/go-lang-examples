package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Product struct {
	_id      primitive.ObjectID
	Brand    string             `json:"brand" bson:"brand"`
	Category string             `json:"category" bson:"category"`
	Image    string             `json:"image" bson:"image"`
	IsNew    bool               `json:"isNew" bson:"isNew"`
	ItemInfo bson.A             `json:"itemInfo" bson:"itemInfo"`
	LastDt   primitive.DateTime `json:"lastDt" bson:"lastDt"`
	Link     string             `json:"link" bson:"link"`
	RegDate  primitive.DateTime `json:"regDate" bson:"regDate"`
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("yoox").Collection("product")

	ctx, cancel = context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var results []Product

	if err := cur.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(results))

	//
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	id := c.Query("id")
	//	index, _ := c.Cookie("index")
	//
	//	var data Data
	//	json.Unmarshal([]byte(index), &data)
	//	print(data.disabledNotice)
	//
	//	c.JSON(200, gin.H{
	//		"message": fmt.Sprintf("%s%s", id,index),
	//	})
	//})
	//r.Run()
}
