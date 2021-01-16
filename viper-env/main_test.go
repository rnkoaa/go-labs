package main

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v3"
)

type TestCase struct {
	Description    string            `yaml:"description" json:"description"`
	Config         Config            `yaml:"config" json:"config"`
	ExpectedResult DeploymentRequest `yaml:"expected_result" json:"expected_result"`
}

func Test_deployment_requests(t *testing.T) {
	velaTestCases := readGoldenTestFile(t, "./env-test-cases.yml")
	if len(velaTestCases) != 8 {
		t.Fatalf("error expected 8 cases, got %d", len(velaTestCases))
	}

	for _, c := range velaTestCases {
		t.Run(c.Description, func(t *testing.T) {
			gotDeploymentInfo := c.Config.GetDeploymentInfo()
			if diff := cmp.Diff(c.ExpectedResult, gotDeploymentInfo); diff != "" {
				t.Fatalf("want mismatch (%s) (-want +got):\n%s", c.Description, diff)
			}
		})
	}
}

func readGoldenTestFile(t *testing.T, path string) []TestCase {
	t.Helper()
	var res []TestCase
	b, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("error while reading test cases file, %v", err)
	}

	err = yaml.Unmarshal(b, &res)
	if err != nil {
		t.Fatalf("error while unmarshalling yaml test cases file, %v", err)
	}

	return res
}
