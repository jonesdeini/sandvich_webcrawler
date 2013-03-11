package main

import (
  "time"
  "fmt"
  "encoding/json"
  "regexp"
  "strings"
)

type jsonObject struct {
  Result ResultType
}

type ResultType struct {
  Status        int
  BackpackSlots int `json:"num_backpack_slots"`
  Items         []ItemsType
}

type ItemsType struct {
  Defindex    int
  Level       int
  Quality     int
  Attributes  []AttributesType
}

type AttributesType struct {
  FloatValue float32 `json:"float_value"`
}

func backpackRetriever(steamUrl string) {
  // I tried to avoid this extra regexp. Idealy just the steam id would be passed into this function.
  // Due to the use of FindAllString() in crawler we're stuck with entire steam profile url
  regex, err := regexp.Compile(`\d+`)
  errorHandler(err)
  steamId := regex.FindString(steamUrl)
  apiCall := "http://api.steampowered.com/IEconItems_440/GetPlayerItems/v0001/?key=" + apiKey()
  apiCall = apiCall + "&steamid=" + steamId
  backpack := apiFetcher(apiCall)
  var jsonType jsonObject
  err = json.Unmarshal(backpack, &jsonType)
  errorHandler(err)
  fmt.Printf("Results: %+v\n", jsonType)
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
