package db

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"sort"
	"strconv"
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

type Films []Film
func (f Films) Len() int { return len(f) }
func (f Films) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

type SortByRating struct{ Films }
func (f SortByRating) Less(i, j int) bool {
	return f.Films[i].Rating.Average > f.Films[j].Rating.Average
}

type SortByRatingPeople struct{ Films }
func (f SortByRatingPeople) Less(i, j int) bool {
	str1 := f.Films[i].Rating.RatingPeople
	str2 := f.Films[j].Rating.RatingPeople
	if (len(str1) != len(str2)) {
		return len(str1) > len(str2)
	} else {
		return str1 > str2
	}
}

type SortByPubdate struct{ Films }
func (f SortByPubdate) Less(i, j int) bool {
	return f.Films[i].Pubdate[0] > f.Films[j].Pubdate[0]
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

	data, err := ioutil.ReadFile("./films_all.json")
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

func GetPageFilms(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.Param("page"))
	sortWay := c.Param("sort")
	db := GlobalSession.Clone()
	defer db.Close()
	database := db.DB("Films")

	var films Films
	err = database.C("films").Find(nil).All(&films)
	if err != nil {
		return err
	}

	switch sortWay {
	case "0":
		sort.Sort(SortByRating{films})
	case "1":
		sort.Sort(SortByRatingPeople{films})
	case "2":
		sort.Sort(SortByPubdate{films})
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.JSON(http.StatusOK, films[(page-1)*10:page*10])
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
