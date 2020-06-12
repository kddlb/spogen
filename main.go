package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/kddlb/spogen/genius"
	"github.com/kddlb/spogen/session"
	"github.com/kddlb/spogen/spotify"
	"log"
)

var Resty *resty.Client = resty.New()

func main() {

	session.Resty = Resty
	spotify.Resty = Resty
	genius.Resty = Resty

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.Use(TokenRefresher())

	sessionEp := r.Group("/api/session")
	{
		sessionEp.GET("/new", session.New)
		sessionEp.GET("/callback", session.Callback)
		sessionEp.GET("/delete", session.Delete)
		sessionEp.GET("/info", session.Info)
	}

	spotifyEp := r.Group("/api/spotify")
	{
		spotifyEp.GET("/playbackInfo", spotify.PlaybackInfo)
		spotifyEp.GET("/action/:action", spotify.Action)
	}

	geniusEp := r.Group("/api/genius")
	{
		geniusEp.GET("/search", genius.Search)
		geniusEp.GET("/get/*path", genius.Get)
	}

	r.Run()
}
