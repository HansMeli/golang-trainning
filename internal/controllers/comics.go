package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	comic "github.com/golang-trainning/internal/models"
	encrypter "github.com/golang-trainning/internal/pkg/middlewares/encrypter"
	comicsRepository "github.com/golang-trainning/internal/repositories/comicsRepository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// GetAllComics return all the available comics from marvel page
func GetAllComics() comic.Data {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MarvelBaseURL := os.Getenv("MARVEL_BASEURL")
	MarvelPublicKey := os.Getenv("MARVEL_PUBLIC_KEY")
	MarvelPrivateKey := os.Getenv("MARVEL_PRIVATE_KEY")

	HASHKEY := encrypter.DigestString("1", MarvelPublicKey, MarvelPrivateKey)

	response, err := http.Get(MarvelBaseURL + "comics?ts=1&apikey=" + MarvelPublicKey + "&hash=" + HASHKEY)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject comic.Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject.Data
}

// GetLibraryComics return all the available comics in the library
func GetLibraryComics(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	comics := comicsRepository.GetRepLibraryComics

	c.JSON(200, comics)
}

// UpdateComicsLibrary uopdate each period of time the comics library database
func UpdateComicsLibrary() string {

	listComics := GetAllComics()

	response := comicsRepository.UpdateDatabaseFromMarvel(listComics)

	return response
}
