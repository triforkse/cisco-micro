package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "bytes"
)

func initCmd(providerId string, filePath string) {
  config := newProvider(providerId)
  config.populate();

  debugf("Generating Config: %+v", config)

  data, err := json.Marshal(config)

  if err != nil {
    log.Fatal("Could not write configuration." + err.Error())
  }

  var out bytes.Buffer
  json.Indent(&out, data, "", "  ")

  err = ioutil.WriteFile(filePath, out.Bytes(), 0644)

  if err != nil {
    log.Fatal("Could not write configuration. " + err.Error())
  }
}
