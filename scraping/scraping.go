package scraping

import (
    "fmt"
    "log"
    "strings"
    "regexp"
    "github.com/PuerkitoBio/goquery"
)

// Return PornHub thumbnail urls
func Pornhub(url string) []string {
    var thumbnail_urls []string

    // HTML取得
    doc, err := goquery.NewDocument(url)
    if err != nil {
        log.Printf("url scraping failed")
        log.Fatal(err)
    }

    // 動画プレイヤーを取得
    player_script := doc.Find("#player").First().Children().First()

    // "image_url" 以降の文字列を取得
    str := strings.Split(player_script.Text(), "\"image_url\"")[1]

    // "\" を除去
    str = strings.Replace(str, "\\", "", -1)

    // 最初のURLを取得
    r := regexp.MustCompile(`https?://[-_.!~*\'()a-zA-Z0-9;/?:@&=+$,%#]+`)
    image_url := r.FindStringSubmatch(str)[0]

    // "/original/" より前のURLを取得
    str = strings.Split(image_url, "/original/")[0]

    // 1から16までの全てのサムネイルURLを格納
    for i := 1; i <= 16; i++ {
        thumbnail_url := fmt.Sprintf("%s/original/%d.jpg", str, i)
        thumbnail_urls = append(thumbnail_urls, thumbnail_url)
    }

    return thumbnail_urls
}
