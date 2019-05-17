package db

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
)

var GlobalSession *mgo.Session

type RatingS struct {
	Average      string `json:"average" bson:"average"`
	RatingPeople string `json:"rating_people" bson:"rating_people"`
}

type People struct {
	Name string `json:"name" bson:"name"`
	ID   string `json:"id" bson:"id"`
}

type Film struct {
	ID        string    `json:"_id" bson:"_id"`
	Directors []People  `json:"directors" bson:"directors"`
	Genres    []string  `json:"genres" bson:"genres"`
	Rating    RatingS 	`json:"rating" bson:"rating"`
	Pubdate   []string  `json:"pubdate" bson:"pubdate"`
	Countries []string  `json:"countries" bson:"countries"`
	Casts     []People  `json:"casts" bson:"casts"`
	Title     string    `json:"title" bson:"title"`
	Poster    string    `json:"poster" bson:"poster"`
	Summary   string    `json:"summary" bson:"summary"`
	Languages []string  `json:"languages" bson:"languages"`
	Writers   []People  `json:"writers" bson:"writers"`
	Year      string    `json:"year" bson:"year"`
	Duration  string    `json:"duration" bson:"duration"`
}

func InitializeGlobalDB(url string) error {
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}

	GlobalSession = session
	return nil
}

func InitData() error {
	db := GlobalSession.Clone()
	defer db.Close()
	database := db.DB("Films")

	data, err := ioutil.ReadFile("G:/TJ/Lessons/Web/Assignment3/backend/db/films_all.json")
    if err != nil {
        return err
	}
	
	var films []Film
	err = json.Unmarshal(data, &films)
	if err != nil {
		return err
	}

	for _, value := range films {
		database.C("films").Insert(value)
	}

	return nil
}

func GetAllFilms(c echo.Context) (err error) {
	db := GlobalSession.Clone()
	defer db.Close()
	database := db.DB("Films")

	var films []Film
	err = database.C("films").Find(nil).All(&films)
	if err != nil {
		return err
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, films)
}

func GetFilm(c echo.Context) (err error) {
	id := c.Param("id")
	db := GlobalSession.Clone()
	defer db.Close()
	database := db.DB("Films")

	film := Film{}
	err = database.C("films").Find(bson.M{"_id": id}).One(&film)
	if err != nil {
		return err
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, film)
}
