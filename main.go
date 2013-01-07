package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
 "regexp"
 //"strings"
)

func main() {
  resp, err := http.Get("http://xxlgamers.gameme.com/tf")
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  serverUrlRegex, err := regexp.Compile(`http://xxlgamers.gameme.com/overview/\d.`)
  res := serverUrlRegex.FindAllString(string(body), -1)
  severUrls := []string{}
  var seen bool
  for i := range res {
    seen = false
    for j := range severUrls {
      if res[i] == severUrls[j] {
        seen = true
      }
    }
    if seen == false {
      severUrls = append(severUrls, res[i])
    }
  }
  fmt.Printf("%v", severUrls)
  if err != nil {
    fmt.Println(err)
  }
}
