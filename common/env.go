package common

import "io"

type Env struct {
	OutStream, ErrStream io.Writer
	CurrentDirectory     string
	Args                 []string
}
