package comicsrepository

import (
	"context"
	"log"
	"net/http"

	database "github.com/golang-trainning/internal/database"
	comic "github.com/golang-trainning/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetLibraryComics get the comics data from database
func GetLibraryComics(w http.ResponseWriter) []comic.Comic {
	// we created Comics array
	var res []comic.Comic

	//Connection mongoDB with database class
	collection := database.ConnectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		database.GetError(err, w)
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var comic comic.Comic
		// & character returns the memory address of the following variable.
		err := cur.Decode(&comic) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		res = append(res, comic)
	}

	return res
}
