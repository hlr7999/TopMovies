package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"flag"
	"log"

	"Films/db"
)

func main() {
	var addr string
	var initData bool
	flag.StringVar(&addr, "addr", ":10080", "server listens at this addr")
	flag.BoolVar(&initData, "init-data", false, "initialize development data")
	flag.Parse()

	err := db.InitializeGlobalDB("127.0.0.1")
	if err != nil {
		log.Panic(err)
	}

	if initData {
		err = db.InitData()
		if err != nil {
			log.Panic(err)
		}
		return
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/films/:page/:sort", db.GetPageFilms)
	e.GET("/films/:id", db.GetFilm)
	e.Logger.Fatal(e.Start(addr))
}
