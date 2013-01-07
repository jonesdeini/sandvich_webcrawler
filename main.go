package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
 "regexp"
 //"strings"
)

func urlFetcher(url string) string {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  // if err yada yada???
  return string(body)
}

func uniq(s []string) []string {
  var seen bool
  uniqSlice := []string{}

  for i := range s {
    seen = false
    for j := range uniqSlice {
      if s[i] == uniqSlice[j] {
        seen = true
      }
    }
    if seen == false {
      uniqSlice = append(uniqSlice, s[i])
    }
  }
  return uniqSlice
}

func main() {
  serverUrlRegex, _ := regexp.Compile(`http://xxlgamers.gameme.com/overview/\d.`)
  clanUrls := []string {"http://xxlgamers.gameme.com/tf"}
  severUrls := []string{}

  // NOTE this ain't gonna work
  // FIXME use a slice of structs contianing clan specific url and regex
  for i := range clanUrls {
    // fetch page and parse
    res := serverUrlRegex.FindAllString(urlFetcher(clanUrls[i]), -1)
    // unique results and save sever urls
    severUrls = append(severUrls, uniq(res)...)
  }
  fmt.Printf("%v", severUrls)
}
