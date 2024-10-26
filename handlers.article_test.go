package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := io.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0
		return statusOK && pageOK
	})
}

func TestArticleListJSON(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := io.ReadAll(w.Body)
		var articles []article
		err = json.Unmarshal(p, &articles)
		jsonOK := err == nil && len(articles) > 0
		return statusOK && jsonOK
	})
}

func TestArticleXML(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)
	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Set("Accept", "application/xml")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := io.ReadAll(w.Body)
		var article article
		err = xml.Unmarshal(p, &article)
		xmlOK := err == nil && article.ID == 1
		return statusOK && xmlOK
	})
}
