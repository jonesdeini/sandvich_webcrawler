package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
 "regexp"
)

func errorHandler(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func playerInfoUrlFetcher(serverUrl string) []string {
  playerInfoUrlsRegex, _ := regexp.Compile(`http://xxlgamers.gameme.com/playerinfo/\d+`)
  res := playerInfoUrlsRegex.FindAllString(urlFetcher(serverUrl), -1)
  return uniq(res)
}

func serverUrlFetecher() []string {
  serverUrlRegex, _ := regexp.Compile(`http://xxlgamers.gameme.com/overview/\d+`)
  clanUrl := "http://xxlgamers.gameme.com/tf"

  // fetch page and parse
  res := serverUrlRegex.FindAllString(urlFetcher(clanUrl), -1)
  // unique results and save sever urls
  return uniq(res)
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

//channels
func sendPlayerInfoUrls(playerInfoUrls []string, cs chan string) {
  for i := range playerInfoUrls {
    cs <- playerInfoUrls[i]
  }
}

func recievePlayerInfoUrl(playerInfoIdChannel chan string, steamIdChannel chan string) {
  playerInfoUrl := <-playerInfoIdChannel
  playerInfoPage := urlFetcher(playerInfoUrl)
  steamIdRegexp, _ := regexp.Compile(`http://steamcommunity.com/profiles/\d+`)
  res := steamIdRegexp.FindAllString(playerInfoPage, -1)
  fmt.Printf("%v", uniq(res))
}

func main() {
  //severUrls := serverUrlFetecher()
  playerInfoUrls := playerInfoUrlFetcher("http://xxlgamers.gameme.com/overview/18")
  playerInfoIdChannel := make(chan string)
  steamIdChannel := make(chan string)
  for i := range playerInfoUrls {
    go sendPlayerInfoUrls(playerInfoUrls, playerInfoIdChannel)
    recievePlayerInfoUrl(playerInfoIdChannel, steamIdChannel)
    fmt.Println("loop number: ", i)
  }
  fmt.Printf("%v", playerInfoUrls)
}
