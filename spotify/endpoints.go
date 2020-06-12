package spotify

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

var Resty *resty.Client

func PlaybackInfo(ctx *gin.Context) {

	token, err := ctx.Cookie("accessToken")

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
	} else {

		type PIResult struct {
			IsPlaying bool  `json:"isPlaying"`
			IsPrivate bool  `json:"isPrivate"`
			Item      *Item `json:"item"`
		}

		piResp, _ := Resty.R().
			SetAuthToken(token).
			SetResult(&CurrentPlayback{}).
			Get("https://api.spotify.com/v1/me/player")

		currPlayback := piResp.Result().(*CurrentPlayback)

		piRx := PIResult{
			IsPlaying: piResp.StatusCode() != http.StatusNoContent,
			IsPrivate: currPlayback.Item == nil,
			Item:      currPlayback.Item,
		}

		ctx.JSON(200, piRx)

	}

}

func Action(ctx *gin.Context) {

	token, err := ctx.Cookie("accessToken")

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
	} else {
		_, _ = Resty.R().
			SetAuthToken(token).
			Post("https://api.spotify.com/v1/me/player/" + ctx.Param("action"))

		ctx.String(http.StatusNoContent, "")

	}

}
