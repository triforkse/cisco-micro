package terraform
import (
        "testing"
        "cisco/micro/provider/gce"
)

func TestMergeVars(t *testing.T) {
        config := gce.Config{Project: "test-project"}
        expectedValue := "asked-value"

        result := gatherVars(&config, func(a string, b string) (string, error) {
              return "asked-value", nil
        })

        if val := result["datacenter"]; val != expectedValue {
                t.Errorf("Expected %s, got %s", expectedValue, val)
        }

        if val := result["ssh_username"]; val != expectedValue {
                t.Errorf("Expected %s, got %s", expectedValue, val)
        }

        expectedValue = "test-project"
        if val := result["project"]; val != expectedValue {
                t.Errorf("Expected %s, got %s", expectedValue, val)
        }
}
