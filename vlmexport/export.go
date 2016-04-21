package vlmexport

import (
	"github.com/tottokotkd/dockervlm/common"
	"github.com/tottokotkd/dockervlm/compose"
)

func VlmExport(containers []common.Container, options common.ExportOptions, env common.Env) common.ExitCode {
	compose.GetComposeConfig(env.CurrentDirectory)
	return common.ExitCodeOK
}
