package main

import (
 "net/http"
 "io/ioutil"
 "fmt"
)

func main() {
  resp, err := http.Get("http://xxlgamers.gameme.com/tf")
  if err != nil {
    // handle error
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  // lol wut
  fmt.Println(body)
}
