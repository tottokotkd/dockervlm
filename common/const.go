package common

const Version string = "0.0.1"

// 終了コード
type ExitCode int

const (
	ExitCodeOK ExitCode = iota
	ExitCodeNoArgsError
	ExitCodeGetwdError
	ExitCodeParseFlagError
	ExitCodeComposeConfigFileNotFoundError
	ExitCodeParseExportFlagError
	ExitCodeParseImportFlagError
)
