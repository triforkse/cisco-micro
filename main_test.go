package main

import (
  "testing"
)

func TestReadConfig(t *testing.T) {
  config := ReadConfig("testdata/aws_test.json")

  if config.Provider != "aws" {
    t.Error("expected provider 'aws'")
  }

  if config.Id != "test-aws-123" {
    t.Error("expected another 'id' attribute")
  }
}
