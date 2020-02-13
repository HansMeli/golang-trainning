package comicsrepository

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	database "github.com/golang-trainning/internal/database"
	comic "github.com/golang-trainning/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GET SECTION

// GetRepLibraryComics get the comics data from database
func GetRepLibraryComics(w http.ResponseWriter) []comic.Comic {
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

// SET SECTION

// UpdateDatabaseFromMarvel the comics data from database
func UpdateDatabaseFromMarvel(j comic.Data) string {

	//Connection mongoDB with database class
	collection := database.ConnectDB()

	var msj string = ""

	for k, v := range j.Results {

		res, errSearch := collection.CountDocuments(context.TODO(), bson.D{{"id", v.ID}})

		if errSearch != nil {
			fmt.Printf("insert fail #%d %v\n", k, errSearch)
			os.Exit(1)
		} else {
			if res == 0 {
				// bson.M{},  we passed empty filter. So we want to get all data.
				_, errInsert := collection.InsertOne(context.Background(), v)
				if errInsert != nil {
					fmt.Printf("insert fail #%d %v\n", k, errInsert)
					os.Exit(1)
				}
				msj = "La base de datos fue actualizada"
			} else {
				msj = "Todo al dia"
				break
			}

		}

	}

	return msj
}
