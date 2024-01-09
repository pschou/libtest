package main

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

var version string

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("LIBTEST - Tool to call arbitrary dynamic library function.\n"+
			"  https://github.com/pschou/libtest  Version: %s\n\n"+
			"Usage: %s library type function [optional inputs]\n"+
			"Types: int_fn_string, fn_string, string_fn, int_fn\n", version, os.Args[0])
		os.Exit(1)
	}
	lib, err := purego.Dlopen(os.Args[1], purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if len(os.Args[2]) > 1 && os.Args[2][len(os.Args[2])-1] == '_' {
		defer func() { fmt.Println("") }()
		os.Args[2] = os.Args[2][:len(os.Args[2])-1]
	}
	if len(os.Args[2]) > 3 && os.Args[2][len(os.Args[2])-3:] != "_fn" && len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Missing input")
		os.Exit(1)
	}
	switch os.Args[2] {
	case "int_fn_string":
		var f func(string) int
		purego.RegisterLibFunc(&f, lib, os.Args[3])
		fmt.Printf("%v", f(os.Args[4]))
	case "string_fn_string":
		var f func(string) string
		purego.RegisterLibFunc(&f, lib, os.Args[3])
		fmt.Printf("%v", f(os.Args[4]))
	case "char_fn_char":
		var f func(byte) byte
		purego.RegisterLibFunc(&f, lib, os.Args[3])
		fmt.Printf("%c", f(os.Args[4][0]))
	case "fn_string":
		var f func(string)
		purego.RegisterLibFunc(&f, lib, os.Args[3])
		f(os.Args[4])
	case "string_fn":
		var f func() string
		purego.RegisterLibFunc(&f, lib, os.Args[3])
		fmt.Printf("%v", f())
	case "int_fn":
		var f func() int
		purego.RegisterLibFunc(&f, lib, os.Args[3])
		fmt.Printf("%v", f())
	default:
		fmt.Fprintf(os.Stderr, "Unsupported type\n")
		os.Exit(1)

	}
}

func RegisterLibFunc(fptr interface{}, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load function %q\n", name)
		os.Exit(1)
	}
	purego.RegisterFunc(fptr, sym)
}
