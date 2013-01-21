package main

import (
 "time"
 "fmt"
 "io/ioutil"
 "net/http"
 "regexp"
 "strings"
)

func backpackRetriever(steamUrl string) {
  regex, err := regexp.Compile(`\d+`)
  errorHandler(err)
  res := regex.FindString(steamUrl)
  fmt.Println(res)
}

func crawler(url string, c chan string) {
  /* fmt.Println(url) */
  regex, err := regexDeterminer(url)
  errorHandler(err)
  if regex != nil {
    res := regex.FindAllString(urlFetcher(url), -1)
    urls := uniq(res)
    for i := range urls {
      go crawler(urls[i], c)
    }
  } else {
    backpackRetriever(url)
  }
  c <- url
}

func errorHandler(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func regexDeterminer(url string) (*regexp.Regexp, error) {
  if strings.Contains(url, "tf") {
    return regexp.Compile(`http://\w+.gameme.com/overview/\d+`)
  } else if strings.Contains(url, "overview") {
    return regexp.Compile(`http://\w+.gameme.com/playerinfo/\d+`)
  } else if strings.Contains(url, "playerinfo") {
    return regexp.Compile(`http://steamcommunity.com/profiles/\d+`)
  }
  return nil, nil
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

func urlFetcher(url string) string {
  resp, err := http.Get(url)
  errorHandler(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  errorHandler(err)
  return string(body)
}

func main() {
  clans := []string {"xxlgamers", "db"}
  myCh := make(chan string)
  for i := range clans {
    clanUrl := "http://" + clans[i] + ".gameme.com/tf"
    go crawler(clanUrl,myCh)
    //fmt.Println("finished!: " + crawler(clanUrl))
  }
  for {
    select {
    case out := <- myCh:
      fmt.Println("finished: " + out)
    case <- time.After(7 * 1e9):
      return
    }
  }
}
