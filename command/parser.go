package command

import (
	"flag"
	"github.com/tottokotkd/dockervlm/common"
	"io"
)

// dockervlm options

const (
	helpOptionDescription = "show help"
)

func parseDockervlmOptions(args []string, errorStream io.Writer) (common.DockervlmOptions, []string, error) {

	flags := flag.NewFlagSet("dockervlm", flag.ContinueOnError)
	flags.SetOutput(errorStream)
	result := common.DockervlmOptions{}

	// help
	flags.BoolVar(&result.Help, "h", false, "")
	flags.BoolVar(&result.Help, "help", false, helpOptionDescription)

	// config file
	ymlFile := "./docker-compose.yml"
	flags.StringVar(&result.YmlFile, "f", ymlFile, "")
	flags.StringVar(&result.YmlFile, "file", ymlFile, helpOptionDescription)

	err := flags.Parse(args)
	return result, args, err
}

// export options

const (
	archiveDirectoryOptionDescription = "show help"
)

func parseExportOptions(args []string, errorStream io.Writer) (common.ExportOptions, error) {

	flags := flag.NewFlagSet("dockervlm", flag.ContinueOnError)
	flags.SetOutput(errorStream)
	result := common.ExportOptions{}

	err := flags.Parse(args)
	return result, err
}

//import options

func parseImportOptions(args []string, errorStream io.Writer) (common.ImportOptions, error) {

	flags := flag.NewFlagSet("dockervlm", flag.ContinueOnError)
	flags.SetOutput(errorStream)
	result := common.ImportOptions{}

	err := flags.Parse(args)
	return result, err
}
