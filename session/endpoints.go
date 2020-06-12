package session

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/kddlb/spogen/spotify"
	"net/http"
	"os"
	"strings"
	"time"
)

var Resty *resty.Client

func New(ctx *gin.Context) {
	authUrl := "https://accounts.spotify.com/authorize"
	type QueryParams struct {
		ClientId     string `url:"client_id"`
		ResponseType string `url:"response_type"`
		RedirectUri  string `url:"redirect_uri"`
		ShowDialog   bool   `url:"show_dialog"`
		Scope        string `url:"scope"`
	}
	scopes := []string{"user-read-playback-state", "user-read-currently-playing", "user-modify-playback-state"}
	opt := QueryParams{
		ClientId:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ResponseType: "code",
		RedirectUri:  os.Getenv("REDIRECT_URI"),
		ShowDialog:   true,
		Scope:        strings.Join(scopes[:], " "),
	}
	oq, _ := query.Values(opt)

	ctx.Redirect(http.StatusFound, authUrl+"?"+oq.Encode())
}

func Callback(ctx *gin.Context) {
	cUrl := ctx.Request.URL
	if len(cUrl.Query().Get("error")) != 0 {
		ctx.Redirect(http.StatusFound, "/")
	} else {

		type OAuthToken struct {
			AccessToken  string `json:"access_token"`
			TokenType    string `json:"token_type"`
			Scope        string `json:"scope"`
			ExpiresIn    int    `json:"expires_in"`
			RefreshToken string `json:"refresh_token"`
		}

		tokenResp, _ := Resty.R().
			SetFormData(map[string]string{
				"grant_type":   "authorization_code",
				"code":         cUrl.Query().Get("code"),
				"redirect_uri": os.Getenv("REDIRECT_URI"),
			}).
			SetHeader("Accept", "application/json").
			SetBasicAuth(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET")).
			SetResult(&OAuthToken{}).
			Post("https://accounts.spotify.com/api/token")
		token := tokenResp.Result().(*OAuthToken)

		dateTokenExpires := time.Now().Add(3600 * time.Second).Format(time.RFC3339)

		ctx.SetCookie("accessToken", token.AccessToken, 2592000, "", "", false, true)
		ctx.SetCookie("refreshToken", token.RefreshToken, 2592000, "", "", false, true)
		ctx.SetCookie("dateTokenExpires", dateTokenExpires, 2592000, "", "", false, true)
		ctx.Redirect(http.StatusFound, "/")
	}
}

func Delete(ctx *gin.Context) {
	ctx.SetCookie("accessToken", "DELETED", -1000, "", "", false, true)
	ctx.SetCookie("refreshToken", "DELETED", -1000, "", "", false, true)
	ctx.SetCookie("dateTokenExpires", "DELETED", -1000, "", "", false, true)
	ctx.Redirect(http.StatusFound, "/")
}

func Info(ctx *gin.Context) {

	type InfoResult struct {
		Authenticated bool                 `json:"authenticated"`
		User          *spotify.UserProfile `json:"user"`
	}

	token, err := ctx.Cookie("accessToken")
	if err != nil {
		ctx.JSON(200, InfoResult{
			Authenticated: false,
			User:          nil,
		})
	} else {
		userResp, _ := Resty.R().
			SetAuthToken(token).
			SetResult(&spotify.UserProfile{}).
			Get("https://api.spotify.com/v1/me")

		ctx.JSON(200, InfoResult{
			Authenticated: true,
			User:          userResp.Result().(*spotify.UserProfile),
		})

	}
}
