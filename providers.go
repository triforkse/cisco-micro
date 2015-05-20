package main

import (
  "log"
)

func newProvider(providerId string) Provider {
  providers := map[string]Provider {
    "aws": new(AWSProvider),
    "gcc": new(GCCProvider),
  }

  provider, known := providers[providerId]

  if !known {
    log.Fatal("Unknown provider: '" + providerId + "'")
  }

  return provider
}
