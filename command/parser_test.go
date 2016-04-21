package command

import (
	"bytes"
	"strings"
	"testing"
)

// dockervlm options parser test
func Test_OfDockervlmOptionsParser(t *testing.T) {
	args := []string{"-h", "-help", "--help", "-h=true", "-help=true", "--help=true", "--help version", "-h export"}
	for _, a := range args {
		stream := new(bytes.Buffer)
		options, _, err := parseDockervlmOptions(strings.Split(a, " "), stream)
		if err != nil {
			t.Errorf("unexpected error '%s' with '%s'", err.Error(), a)
		}
		if !options.Help {
			t.Errorf("'%s' does not turn help option on", a)
		}
		if stream.String() != "" {
			t.Errorf("'%s' polluting stderr: '%s'", a, stream.String())
		}
	}
}

// export options parser test
func Test_OfExportOptionsParser(t *testing.T) {

}

// import options parser test
func Test_OfImportOptionsParser(t *testing.T) {

}
