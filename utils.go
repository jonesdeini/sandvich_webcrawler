// Utility functions which shouldn't change much
package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func errorHandler(err error) {
  if err != nil {
    fmt.Println(err)
  }
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
