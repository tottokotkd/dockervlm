package command

import (
	"bytes"
	"fmt"
	"github.com/tottokotkd/dockervlm/common"

	"strings"
	"testing"
)

// test of 'dockervlm version'
func Test_OfCli_VersionCommand(t *testing.T) {
	command := "dockervlm version"
	cli := commandTest(command, common.ExitCodeOK, t)

	expected := fmt.Sprintf("dockervlm version %s", common.Version)
	if cli.errStream.String() != "" {
		t.Errorf("cli gets '%s', it should write '' but actually '%s'", command, cli.errStream.String())
	}
	if !strings.Contains(cli.outStream.String(), expected) {
		t.Errorf("cli gets '%s', it should write '%s' but actually %s", command, expected, cli.outStream.String())
	}
}

// test of 'dockervlm --help'
func Test_OfCli_DockervlmHelpOption(t *testing.T) {
	commands := []string{"dockervlm -h", "dockervlm -help", "dockervlm --help", "dockervlm --help version"}
	for _, command := range commands {
		cli := commandTest(command, common.ExitCodeOK, t)
		expected := common.DocOfDockervlm
		if !strings.Contains(cli.outStream.String(), expected) {
			t.Errorf("'%s' should return '%s' but actually '%s'", command, expected, cli.outStream.String())
		}
	}
}

//test of no args calling
func Test_OfCli_NoArgs(t *testing.T) {
	command := "dockervlm"
	cli := commandTest(command, common.ExitCodeNoArgsError, t)
	if !strings.Contains(cli.errStream.String(), common.DocOfDockervlm) {
		t.Error("naked 'dockervlm' command should return help doc but it does not.")
	}
}

// test of invalid sub commands
func Test_OfCli_InvalidSubCommands(t *testing.T) {
	command := "dockervlm hoge"
	cli := commandTest(command, common.ExitCodeOK, t)
	if !strings.Contains(cli.errStream.String(), "unknown sub command: hoge") {
		t.Errorf("error message for invalid sub command '%s' is incorrect: '%s'", command, cli.errStream.String())
	}
}

// util
func commandTest(command string, exitCode common.ExitCode, t *testing.T) DebugCLI {
	cli := makeCli()
	status := cli.run(strings.Split(command, " "))
	if status != exitCode {
		t.Errorf("cli gets '%s', it should return %v but actually %d", command, exitCode, status)
	}
	return cli
}

type DebugCLI struct {
	outStream, errStream *bytes.Buffer
}

func makeCli() DebugCLI {
	stdoutStream, stderrStream := new(bytes.Buffer), new(bytes.Buffer)
	return DebugCLI{outStream: stdoutStream, errStream: stderrStream}
}

func (dcli DebugCLI) run(args []string) common.ExitCode {
	cli := common.Env{OutStream: dcli.outStream, ErrStream: dcli.errStream, Args: args}
	return Run(cli)
}
