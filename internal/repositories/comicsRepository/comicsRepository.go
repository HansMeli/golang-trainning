package comicsrepository

import (
	"context"
	"fmt"
	"log"
	"os"

	database "github.com/golang-trainning/internal/database"
	comicModel "github.com/golang-trainning/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GET SECTION

// GetRepLibraryComics get the comics data from database
func GetRepLibraryComics() []comicModel.Comic {
	// we created Comics array
	comic := comicModel.New()

	//Connection mongoDB with database class
	collection := database.ConnectDB()

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var item comicModel.Comic
		// & character returns the memory address of the following variable.
		err := cur.Decode(&item) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array

		comic.Add(item)
	}

	return comic.GetAll()
}

// SET SECTION

// UpdateDatabaseFromMarvel the comics data from database
func UpdateDatabaseFromMarvel(j comicModel.Data) int {

	//Connection mongoDB with database class
	collection := database.ConnectDB()

	var count int = 0

	for k, v := range j.Results {

		res, errSearch := collection.CountDocuments(context.TODO(), bson.D{{"idMarvel", v.ID}})

		if errSearch != nil {
			fmt.Printf("find fail #%d %v\n", k, errSearch)
			os.Exit(1)
		} else {
			if res == 0 {
				var comic comicModel.Comic

				comic.IDMarvel = v.ID
				comic.Title = v.Title
				comic.Description = v.Description
				comic.SetStock()

				// bson.M{},  we passed empty filter. So we want to get all data.
				_, errInsert := collection.InsertOne(context.Background(), comic)
				if errInsert != nil {
					fmt.Printf("insert fail #%d %v\n", k, errInsert)
					os.Exit(1)
				}
				count++
			}
		}
	}

	return count
}
