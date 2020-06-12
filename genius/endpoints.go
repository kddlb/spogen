package genius

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/k3a/html2text"
	"github.com/masatana/go-textdistance"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

var Resty *resty.Client

func Search(ctx *gin.Context) {
	re := regexp.MustCompile("\\s*(\\(.*\\))")
	simplifiedName := re.ReplaceAllString(strings.Split(ctx.Query("track"), " - ")[0], "")
	query := fmt.Sprintf("%s %s", ctx.Query("artist"), simplifiedName)
	queryDistance := fmt.Sprintf("%s by %s", simplifiedName, ctx.Query("artist"))
	sResp, _ := Resty.R().
		SetAuthToken(os.Getenv("GENIUS_ACCESS_TOKEN")).
		SetResult(&SearchResult{}).
		SetQueryParam("q", query).
		Get("https://api.genius.com/search")
	searchHits := sResp.Result().(*SearchResult).Response.Hits

	var searchMap []Result

	for _, s := range searchHits {
		s.Result.SpogenDistance = textdistance.DamerauLevenshteinDistance(strings.ToLower(queryDistance),
			strings.ToLower(s.Result.FullTitle))
		searchMap = append(searchMap, s.Result)
	}

	sort.Slice(searchMap, func(i, j int) bool {
		return searchMap[i].SpogenDistance < searchMap[j].SpogenDistance
	})

	ctx.JSON(http.StatusOK, searchMap)

}

func Get(ctx *gin.Context) {
	url := "https://genius.com" + ctx.Param("path")
	gRes, _ := Resty.R().Get(url)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(gRes.String()))
	doc.Find("div.lyrics").Find("a").Each(func(i int, selection *goquery.Selection) {
		h, _ := selection.Html()
		selection.ReplaceWithHtml(h)
	})

	type GetResult struct {
		HTML string `json:"html"`
		Text string `json:"text"`
	}

	html, _ := doc.Find("div.lyrics").Html()

	ctx.JSON(200, GetResult{
		HTML: strings.TrimSpace(html),
		Text: html2text.HTML2Text(html),
	})

}
