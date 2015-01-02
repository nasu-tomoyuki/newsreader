/*
 * NHK からニュースを取得し、記事単位でファイルに出力します
 */
package main

import (
  "fmt"
  "log"
  "flag"
  "io/ioutil"
  "net/http"
  "encoding/xml"
)

type RssItem struct {
  Title string `xml:"title"`
  Description string `xml:"description"`
}

type RssChannel struct {
  Title string `xml:"title"`
  Description string `xml:"description"`
  Item []RssItem `xml:"item"`
}

type Rss struct {
  Channel RssChannel `xml:"channel"`
}

func main() {
  outfile := flag.String("out", "source/source", "output file")
  url := flag.String("url", "http://www3.nhk.or.jp/rss/news/cat0.xml", "nhk rss")
  count := flag.Int("count", 3, "read article count")
  flag.Parse()

  resp, _ := http.Get(*url)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)

  var rss Rss
  err := xml.Unmarshal([]byte(byteArray), &rss)

  if err != nil {
    log.Fatal(err)
    return
  }

  text := ""
  for k,v := range rss.Channel.Item {
    index := *count - k			// 古い記事を若い番号に
    switch {
    case index == 1:
      text = rss.Channel.Title + "です。　最初のニュースです。　"
    case k == 0:
      text = "　最後のニュースです。　"
    case index != 1:
      text = "　次のニュースです。　"
    }
    text += v.Title + "　" + v.Description
    if k == 0 {
      text += "　以上、" + rss.Channel.Title + "でした。"
    }
    output(*outfile, index, text)

    if index == 1 {
      break
    }
  }
}

func output(outfile string, index int, text string) {
  filename := fmt.Sprintf("%s%02d.txt",outfile, index)
  ioutil.WriteFile(filename, []byte(text), 0644)
}
