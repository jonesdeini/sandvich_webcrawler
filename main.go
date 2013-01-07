package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
 "regexp"
 //"strings"
)

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
  resp, err := http.Get("http://xxlgamers.gameme.com/tf")
  if err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  serverUrlRegex, err := regexp.Compile(`http://xxlgamers.gameme.com/overview/\d.`)
  res := serverUrlRegex.FindAllString(string(body), -1)
  severUrls := uniq(res)
  fmt.Printf("%v", severUrls)
  if err != nil {
    fmt.Println(err)
  }
}
