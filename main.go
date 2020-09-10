package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/kddlb/spogen/genius"
	"github.com/kddlb/spogen/session"
	"github.com/kddlb/spogen/spotify"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	Resty              *resty.Client = resty.New()
	availableAddresses               = [4]*net.TCPAddr{}
)

func main() {

	availableIPs := [4]string{"173.234.25.154:", "173.234.25.156:", "173.234.25.157:", "173.234.25.158:"}

	for ix, ip := range availableIPs {
		var err error
		availableAddresses[ix], err = net.ResolveTCPAddr("tcp4", ip)
		if err != nil {
			log.Fatal(err)
		}
	}

	session.Resty = Resty
	spotify.Resty = Resty
	genius.Resty = Resty
	genius.GenerateTransportFromRandomIPAddress = generateTransportFromRandomIPAddress

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.Use(TokenRefresher())

	r.Use(static.Serve("/", static.LocalFile("www", false)))

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
		geniusEp.GET("/gl/:id", genius.GetId)
		geniusEp.POST("/parse", genius.Parse)
	}

	r.Run()
}

func generateTransportFromRandomIPAddress() *http.Transport {
	var pick *net.TCPAddr
	if os.Getenv("IS_DEV") == "yes" {
		pick, _ = net.ResolveTCPAddr("tcp4", "192.168.1.85:")
	} else {
		randomIndex := rand.Intn(len(availableAddresses))
		pick = availableAddresses[randomIndex]
	}

	log.Printf("IP picked: %s", pick)
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			LocalAddr: pick}).Dial,
		TLSHandshakeTimeout: 10 * time.Second}
}
