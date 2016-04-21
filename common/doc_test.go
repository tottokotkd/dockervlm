package common

import (
	"strings"
	"testing"
)

// test format
// function name: 'Test_Target_Purpose'
func Test_OfTestUtils_ForTestFormat(t *testing.T) {
	var msg = "cmd subcmd -opt1 a -opt2 b x y z"
	args := parseString(msg)
	if len(args) != 9 {
		// error message: more description, less misunderstanding
		// bad one
		t.Errorf("args=%s, to be %s", args, []string{"cmd", "subcmd", "-opt1", "a", "-opt2", "b", "x", "y", "z"})

		// not values but YOU are (able) to tell what's happened!
		t.Errorf("test utilities broken! makeArgs() should parse %s into 9 pieces, but it returns %d.", args, len(args))
	}
}

func parseString(msg string) []string {
	return strings.Split(msg, " ")
}
