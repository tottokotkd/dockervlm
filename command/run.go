package command

import (
	"fmt"
	"github.com/tottokotkd/dockervlm/common"
	"github.com/tottokotkd/dockervlm/compose"
	"github.com/tottokotkd/dockervlm/vlmexport"
	"github.com/tottokotkd/dockervlm/vlmimport"
	"path/filepath"
)

// 引数処理を含めた具体的な処理
func Run(env common.Env) common.ExitCode {

	dockervlmArgs := env.Args[1:]

	// exit when args invalid
	if len(dockervlmArgs) == 0 {
		fmt.Fprintln(env.ErrStream, common.DocOfDockervlm)
		return common.ExitCodeNoArgsError
	}

	// handle dockervlm options
	options, parsedArgs, err := parseDockervlmOptions(dockervlmArgs, env.ErrStream)
	if err != nil {
		fmt.Fprintf(env.ErrStream, err.Error())
		return common.ExitCodeParseFlagError
	}

	// help
	if options.Help {
		fmt.Fprintln(env.OutStream, common.DocOfDockervlm)
		return common.ExitCodeOK
	}

	// config file
	if !filepath.IsAbs(options.YmlFile) {
		options.YmlFile = filepath.Join(env.CurrentDirectory, options.YmlFile)
	}

	// sub commands
	subCommand := parsedArgs[0]
	subCommandArgs := parsedArgs[1:]

	switch subCommand {
	case "export":
		containers, err := compose.GetComposeConfig(options.YmlFile)
		if err != nil {
			fmt.Fprintf(env.ErrStream, err.Error())
			return common.ExitCodeComposeConfigFileNotFoundError
		}
		exportOptions, err := parseExportOptions(subCommandArgs, env.ErrStream)
		if err != nil {
			return common.ExitCodeParseExportFlagError
		}
		return vlmexport.VlmExport(containers, exportOptions, env)
	case "import":
		containers, err := compose.GetComposeConfig(options.YmlFile)
		if err != nil {
			fmt.Fprintf(env.ErrStream, err.Error())
			return common.ExitCodeComposeConfigFileNotFoundError
		}
		importOptions, err := parseImportOptions(subCommandArgs, env.ErrStream)
		if err != nil {
			return common.ExitCodeParseExportFlagError
		}
		return vlmimport.VlmImport(containers, importOptions, env)
	case "version":
		fmt.Fprintf(env.OutStream, "dockervlm version %s\n", common.Version)
		return common.ExitCodeOK
	default:
		fmt.Fprintf(env.ErrStream, "unknown sub command: %s\n", subCommand)
	}

	fmt.Fprintln(env.ErrStream, common.DocOfDockervlm)
	return common.ExitCodeOK
}
