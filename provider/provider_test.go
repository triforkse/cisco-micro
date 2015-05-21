package provider

import (
  "testing"
  //"cisco/micro/provider/aws"
  "reflect"
  "cisco/micro/provider/aws"
  //"fmt"
)

func TestReadConfig(t *testing.T) {
  config := FromFile("testdata/aws_test.json")

  if config.ProviderId() != "aws" {
    t.Error("expected provider 'aws'")
  }

  if config.ConfigId() != "test-aws-123" {
    t.Error("expected another 'id' attribute")
  }

  // Make sure we produce the correct type.

  var expectedType aws.Config
  if reflect.TypeOf(config) != reflect.TypeOf(&expectedType) {
    t.Error("the wrong config type was produced")
  }
}

