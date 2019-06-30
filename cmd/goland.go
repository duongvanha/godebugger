package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"sbgo/tools/bedebug/configuration"
	"sbgo/tools/bedebug/log"
)

var golandConfig = &cobra.Command{
	Use:   "goland",
	Short: "Export debug configuration goland",
	Run:   ExportGolandConfig,
}

func init() {
	rootCmd.AddCommand(golandConfig)
}

func ExportGolandConfig(cmd *cobra.Command, args []string) {
	var currentDir string
	if len(args) > 0 {
		currentDir = args[0]
	} else {
		currentDir, _ = os.Getwd()
	}

	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if info.Name() == "cicd.json" {
			RunCommand(filepath.Dir(path))
		}
		return nil
	})

	if err != nil {
		log.Warn(fmt.Sprintf("Error %v", err))
	}
}

func RunCommand(dir string) {
	println(dir)

	file, err := ioutil.ReadFile(dir + "/cicd.json")

	cicd := configuration.Cicd{}

	err = json.Unmarshal([]byte(file), &cicd)

	nameFile := fmt.Sprintf("%v/src/sbgo/.idea/runConfigurations/%v.xml", os.Getenv("GOPATH"), cicd.Deploy.ServiceName)

	writer, err := os.OpenFile(nameFile, os.O_RDWR|os.O_CREATE, 0777)

	output, err := xml.MarshalIndent(configuration.New(cicd), "  ", "    ")
	if err != nil {
		println(fmt.Sprintf("Error 1 %v", err))
	}

	_, err = writer.Write(output)

	if err != nil {
		println(fmt.Sprintf("Error 2 %v", err))
	}

}
