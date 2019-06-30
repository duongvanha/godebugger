package configuration

import (
	"fmt"
	"strings"
)

type moduleConfig struct {
	Name string `xml:"name,attr"`
}

type parameter struct {
	Value string `xml:"value,attr"`
}

type method struct {
	Value string `xml:"v,attr"`
}

type Configuration struct {
	Default          bool         `xml:"default,attr"`
	Name             string       `xml:"name,attr"`
	Type             string       `xml:"type,attr"`
	FactoryName      string       `xml:"factoryName,attr"`
	Module           moduleConfig `xml:"module"`
	WorkingDirectory parameter    `xml:"working_directory"`
	GoParameters     parameter    `xml:"go_parameters"`
	Parameters       parameter    `xml:"parameters"`
	Kind             parameter    `xml:"kind"`
	FilePath         parameter    `xml:"filePath"`
	Package          parameter    `xml:"package"`
	Directory        parameter    `xml:"directory"`
	Method           method       `xml:"method"`
}

type component struct {
	Name          string        `xml:"name,attr"`
	Configuration Configuration `xml:"configuration"`
}

type IdeConfig struct {
	Component component `xml:"component"`
}

func New(cicd Cicd) component {

	AdditionArgs := []string{""}

	if cicd.AppType == "api" {
		AdditionArgs = append(AdditionArgs, fmt.Sprintf("-http-port=%v", cicd.Deploy.DevExposePort))
	} else {
		AdditionArgs = append(AdditionArgs, fmt.Sprintf("-grpc-port=%v", cicd.Deploy.DevExposePort))
	}

	AdditionArgs = append(AdditionArgs, "-config-remote-address=127.0.0.1:8500")


	if len(cicd.Deploy.ConfigRemoteKeys) > 0 {
		fullConfigsArgs := fmt.Sprintf("-config-remote-keys=%v", strings.Join(cicd.Deploy.ConfigRemoteKeys, ","))
		AdditionArgs = append(AdditionArgs, fullConfigsArgs)
		AdditionArgs = append(AdditionArgs, "-config-type=remote")
	}

	return component{
		Name: "ProjectRunConfigurationManager",
		Configuration: Configuration{
			Default:     false,
			Name:        cicd.Deploy.ServiceName,
			Type:        "GoApplicationRunConfiguration",
			FactoryName: "Go Application",
			Module: moduleConfig{
				Name: "sbgo",
			},
			WorkingDirectory: parameter{
				Value: "$PROJECT_DIR$/",
			},
			GoParameters: parameter{
				Value: "-i",
			},
			Parameters: parameter{
				Value: strings.Join(AdditionArgs, " "),
			},
			Kind: parameter{
				Value: "FILE",
			},
			FilePath: parameter{
				Value: fmt.Sprintf("$PROJECT_DIR$/cmd/%v/main.go", cicd.Build.CmdBinDir),
			},
			Package: parameter{
				Value: "sbgo",
			},
			Directory: parameter{
				Value: "$PROJECT_DIR$/",
			},
			Method: method{
				Value: "2",
			},
		},
	}
}
