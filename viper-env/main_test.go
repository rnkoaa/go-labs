package main

import "testing"

type testCase struct {
	Description    string
	Config         Config
	ExpectedResult DeploymentRequest
}

var testCases = []Config{
	{
		Vela:            "true",
		Application:     "batchconsumer",
		VelaDescription: "dry=true",
		VelaBuildRef:    "heads/tags/641",
		VelaBuildEvent:  "deployment",
		VelaTarget:      "stage",
	},
	{
		Vela:            "true",
		Application:     "batchconsumer",
		VelaDescription: "dry=true",
		VelaBuildRef:    "heads/branches/main",
		VelaBuildEvent:  "deployment",
		VelaTarget:      "stage",
	},
	{
		Vela:            "true",
		Application:     "batchconsumer",
		VelaDescription: "dry=true;cluster=batchconsumer-stage",
		VelaBuildRef:    "heads/tags/641",
		VelaBuildEvent:  "deployment",
		VelaTarget:      "stage",
	},
	{
		Vela:            "true",
		Application:     "batchconsumer",
		VelaDescription: "dry=true;cluster=batchconsumer-stage",
		VelaBuildRef:    "heads/branches/main",
		VelaBuildEvent:  "deployment",
		VelaTarget:      "stage",
	},
	{
		Vela:            "true",
		Application:     "batchconsumer",
		VelaDescription: "cluster=batchconsumer-stage",
		VelaBuildRef:    "heads/tags/641",
		VelaBuildEvent:  "deployment",
		VelaTarget:      "stage",
	},
	{
		Vela:            "true",
		Application:     "batchconsumer",
		VelaDescription: "cluster=batchconsumer-stage",
		VelaBuildRef:    "heads/branches/main",
		VelaBuildEvent:  "deployment",
		VelaTarget:      "stage",
	},
	{
		Vela:           "true",
		Application:    "batchconsumer",
		VelaBuildRef:   "heads/tags/641",
		VelaBuildEvent: "deployment",
		VelaTarget:     "stage",
	},
	{
		Vela:           "true",
		Application:    "batchconsumer",
		VelaBuildRef:   "heads/branches/main",
		VelaBuildEvent: "deployment",
		VelaTarget:     "stage",
	},
}

func Test_deployment_requests(t *testing.T) {
	if len(testCases) != 8 {
		t.Fatalf("error expected 8 cases, got %d", len(testCases))
	}
}
