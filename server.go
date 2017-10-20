package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/m1nam1/porn-thumbnail-api/scraping"
	"log"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/pornhub", Pornhub),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Result struct {
	Status        string   `json:"status"`
	SiteName      string   `json:"site_name"`
	Url           string   `json:"url"`
	ThumbnailUrls []string `json:"thumbnail_urls"`
}

func Pornhub(w rest.ResponseWriter, r *rest.Request) {
	url := r.URL.Query().Get("url")
	if url != "" {
		fmt.Println(url)

		thumbnail_urls := scraping.Pornhub(url)
		if thumbnail_urls == nil {
			rest.NotFound(w, r)
			return
		}

		w.WriteJson(
			Result{
				Status:        "ok",
				SiteName:      "PornHub",
				Url:           url,
				ThumbnailUrls: thumbnail_urls,
			},
		)
	}
}
