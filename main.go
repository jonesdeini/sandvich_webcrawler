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
  serverUrlRegex, err := regexp.Compile(`http://xxlgamers.gameme.com/overview/\d`)
  res := serverUrlRegex.FindAllString(string(body), -1)
  fmt.Printf("%v", res)
  if err != nil {
    fmt.Println(err)
  }
}
