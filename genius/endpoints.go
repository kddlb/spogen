package genius

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/k3a/html2text"
	"github.com/masatana/go-textdistance"
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
			strings.ToLower(re.ReplaceAllString(s.Result.FullTitle, "")))
		searchMap = append(searchMap, s.Result)
	}

	sort.Slice(searchMap, func(i, j int) bool {
		return searchMap[i].SpogenDistance < searchMap[j].SpogenDistance
	})

	if searchMap == nil {
		searchMap = []Result{}
	}

	ctx.Header("X-SpoGen-Search", query)
	ctx.Header("X-SpoGen-Compare-To", queryDistance)

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

	sx := strings.Split(html2text.HTML2Text(html), "\r\n")
	sx2 := make([]string, len(sx))

	for i, v := range sx {
		sx2[i] = strings.TrimSpace(v)
	}

	ctx.JSON(200, GetResult{
		HTML: strings.TrimSpace(html),
		Text: strings.Join(sx2, "\r\n"),
	})

}

type ParseReq struct {
	Text string `form:"text" json:"text" xml:"text" binding:"required"`
}

func Parse(ctx *gin.Context) {

	var data ParseReq

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(data.Text))
	doc.Find("div.lyrics").Find("a").Each(func(i int, selection *goquery.Selection) {
		h, _ := selection.Html()
		selection.ReplaceWithHtml(h)
	})

	type GetResult struct {
		HTML string `json:"html"`
		Text string `json:"text"`
	}

	lyrics := doc.Find("div.lyrics")

	if lyrics.Nodes == nil {
		doc.Find(".jgQsqn").Find("a").Each(func(i int, selection *goquery.Selection) {
			h, _ := selection.Html()
			selection.ReplaceWithHtml(h)
		})
		lyrics = doc.Find(".jgQsqn")
	}

	var lxx []string

	for _, vx := range lyrics.Nodes {
		htmlx, _ := doc.FindNodes(vx).Html()
		lxx = append(lxx, htmlx)
	}

	html := strings.Join(lxx, "<br/><br/>")

	sx := strings.Split(html2text.HTML2Text(html), "\r\n")
	sx2 := make([]string, len(sx))

	for i, v := range sx {
		sx2[i] = strings.TrimSpace(v)
	}

	ctx.JSON(200, GetResult{
		HTML: strings.TrimSpace(html),
		Text: strings.Join(sx2, "\r\n"),
	})

}
