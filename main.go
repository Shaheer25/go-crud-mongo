package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	user "awesomeProject/controllers"
)

// type Item struct {
// 	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
// 	Name  string             `json:"name" bson:"name"`
// 	Phone string             `json:"phone" bson:"phone"`
// }

// var client *mongo.Client

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading Environmental file")
// 	}
// 	mongoUrl := os.Getenv("MONGODB_URL")
// 	fmt.Printf("MongoURL %s",mongoUrl)
// 	clientOptions := options.Client().ApplyURI(mongoUrl)
// 	client, _ = mongo.Connect(context.Background(), clientOptions)
// 	if client != nil {
// 		fmt.Println("Connected to DB")
// 	} else {
// 		fmt.Println("Error Connecting to DB")
// 	}
// }

// func createItem(c *gin.Context) {

// 	var item Item
// 	if err := c.ShouldBindJSON(&item); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	collection := client.Database("mydb").Collection("items")
// 	itemss, err := collection.InsertOne(context.Background(), item)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating item"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"id": itemss.InsertedID,
// 	})
// }

// func getItems(c *gin.Context) {
// 	collection := client.Database("mydb").Collection("items")
// 	cursor, err := collection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving items"})
// 		return
// 	}

// 	var items []Item
// 	defer cursor.Close(context.Background())
// 	for cursor.Next(context.Background()) {
// 		var item Item
// 		cursor.Decode(&item)
// 		items = append(items, item)
// 	}

// 	c.JSON(http.StatusOK, items)
// }

// func updateItem(c *gin.Context) {
// 	id := c.Param("id")
// 	var updatedItem Item
// 	if err := c.ShouldBindJSON(&updatedItem); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	collection := client.Database("mydb").Collection("items")
// 	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": updatedItem})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating item"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
// }

// func deleteItem(c *gin.Context) {
// 	id := c.Param("id")

// 	// Convert the ID string to an ObjectID
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	collection := client.Database("mydb").Collection("items")
// 	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting item"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
// }
func TimeTaken(t time.Time) {
	elapsed := time.Since(t)
	log.Printf("Time Taken for Execution %v", elapsed)
}
func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.POST("/users", user.CreateUser)

	defer TimeTaken(time.Now())
	time.Sleep(time.Millisecond)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
