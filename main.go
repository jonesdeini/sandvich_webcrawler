package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
 "strings"
)

func main() {
  resp, err := http.Get("http://xxlgamers.gameme.com/tf")
  if err != nil {
    // handle error
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if strings.Contains(string(body), "http://xxlgamers.gameme.com/overview/") {
     fmt.Println("word")
   }
}
