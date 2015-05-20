package main

import (
  "crypto/rand"
)

func randStr(strSize int) string {

  dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

  var bytes = make([]byte, strSize)
  rand.Read(bytes)
  for k, v := range bytes {
    bytes[k] = dictionary[v%byte(len(dictionary))]
  }

  return string(bytes)
}
