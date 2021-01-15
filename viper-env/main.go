package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type DeploymentRequest struct {
	Application string `json:"application" yaml:"application"`
	Cluster     string `json:"cluster" yaml:"cluster"`
	Dry         bool   `json:"dry" yaml:"dry"`
	Environment string `json:"environment" yaml:"environment"`
	Suffix      string `json:"suffix" yaml:"suffix"`
	Version     string `json:"version" yaml:"version"`
}

type Config struct {
	CI              string `yaml:"ci" mapstructure:"CI"`
	Vela            string `yaml:"vela" mapstructure:"vela"`
	Application     string `yaml:"application" mapstructure:"application"`
	VelaDeployment  string `yaml:"vela_deployment" mapstructure:"vela_deployment"`
	VelaTarget      string `yaml:"vela_build_target" mapstructure:"vela_build_target"`
	VelaDescription string `yaml:"vela_description" mapstructure:"vela_description"`
	VelaBuildRef    string `yaml:"vela_build_ref" mapstructure:"vela_build_ref"`
	VelaBuildEvent  string `yaml:"vela_build_event" mapstructure:"vela_build_event"`
	Suffix          string `yaml:"suffix" mapstructure:"suffix"`
	Environment     string `yaml:"env" mapstructure:"env"`
}

func parseDeploymentVersion(req *DeploymentRequest, velaBuildRef string) {
	// only set
	if strings.Contains(velaBuildRef, "tags") {
		refParts := strings.Split(velaBuildRef, "/")
		if len(refParts) == 3 {
			req.Version = refParts[2]
		}
	}
}

func parseDeploymentDescription(req *DeploymentRequest, velaDescription string) {
	descParts := strings.Split(velaDescription, ";")
	if len(descParts) > 0 {
		for _, p := range descParts {
			dp := strings.Split(p, "=")
			if len(dp) == 2 {
				key := dp[0]
				value := dp[1]
				switch key {
				case "cluster":
					req.Cluster = value
				case "dry":
					b, e := strconv.ParseBool(value)
					if e != nil {
						fmt.Printf("error parsing bool value: %v\n", e)
						b = false
					}
					req.Dry = b
				}
			}
		}
	}
}

func (c Config) GetDeploymentInfo() DeploymentRequest {
	req := DeploymentRequest{}
	if c.Suffix != "" {
		req.Suffix = c.Suffix
	}

	if c.Environment != "" {
		req.Environment = c.Environment
	}
	if c.Application != "" {
		req.Application = c.Application
	}

	if c.VelaTarget != "" {
		req.Environment = c.VelaTarget
	}

	if c.VelaBuildEvent == "deployment" {
		if c.VelaDescription != "" {
			parseDeploymentDescription(&req, c.VelaDescription)
		}

		if c.VelaBuildRef != "" {
			parseDeploymentVersion(&req, c.VelaBuildRef)
		}
	}
	return req
}

var envsToBind = []string{
	"application", "ci", "vela", "vela_deployment", "vela_build_target", "vela_description", "vela_build_ref",
}

var config Config

func init() {
	v := viper.New()

	// bind all environment variables manually
	for _, e := range envsToBind {
		_ = v.BindEnv(e)
	}

	// enable viper to read environment variables
	v.AutomaticEnv()

	err := v.Unmarshal(&config)
	if err != nil {
		fmt.Printf("error unmarshalling viper into config [%v]", err)
	}
}

func main() {
	config := Config{
		Application:  "batchconsumer",
		VelaBuildRef: "heads/branches/main",
		// VelaBuildRef:    "heads/tags/641",
		VelaBuildEvent:  "deployment",
		VelaDescription: "cluster=batchconsumer-poison",
		// VelaDescription: "dry=true;cluster=batchconsumer-poison",
		Suffix: "poison",
		// Environment:     "lab2",
		VelaTarget: "dev",
	}

	deploymentInfo := config.GetDeploymentInfo()
	b, _ := json.MarshalIndent(deploymentInfo, "", "  ")
	fmt.Printf("%s\n", b)
	// fmt.Println(config)
}
