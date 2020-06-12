package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func TokenRefresher() func(context *gin.Context) {
	return func(context *gin.Context) {
		reToken, err := context.Cookie("refreshToken")
		if err != nil {
			context.Next()
		} else {
			dateExpireT, _ := context.Cookie("dateTokenExpires")
			dateExpire, _ := time.Parse(time.RFC3339, dateExpireT)
			if time.Now().After(dateExpire) {
				type OAuthRefreshToken struct {
					AccessToken string `json:"access_token"`
					TokenType   string `json:"token_type"`
					Scope       string `json:"scope"`
					ExpiresIn   int    `json:"expires_in"`
				}

				tokenResp, _ := Resty.R().
					SetFormData(map[string]string{
						"grant_type":    "refresh_token",
						"refresh_token": reToken,
					}).
					SetHeader("Accept", "application/json").
					SetBasicAuth(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET")).
					SetResult(&OAuthRefreshToken{}).
					Post("https://accounts.spotify.com/api/token")
				token := tokenResp.Result().(*OAuthRefreshToken)

				dateTokenExpires := time.Now().Add(3600 * time.Second).Format(time.RFC3339)

				context.SetCookie("accessToken", token.AccessToken, 2592000, "", "", false, true)
				context.SetCookie("dateTokenExpires", dateTokenExpires, 2592000, "", "", false, true)
				context.Redirect(http.StatusFound, context.Request.URL.String())

			} else {
				context.Next()
			}
		}
	}
}
