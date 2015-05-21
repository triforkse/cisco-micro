package provider

import (
  "testing"
  "cisco/micro/provider/aws"
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
  config.(aws.Provider)
}

