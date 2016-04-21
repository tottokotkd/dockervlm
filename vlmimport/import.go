package vlmimport

import (
	"github.com/tottokotkd/dockervlm/common"
)

func VlmImport(containers []common.Container, options common.ImportOptions, env common.Env) common.ExitCode {
	return common.ExitCodeOK
}
