package main

import (
	"fmt"
	"os"

	devfileParser "github.com/devfile/library/v2/pkg/devfile"
)

// const devfilePath = "/home/tpetkos/github/thepetk/devfile-samples/devfile-sample-python-basic/devfile.yaml"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("exiting: no path provided\n\ncorrect usage: <binary> <path>\n")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Printf("exiting: too many arguments provided\n")
		os.Exit(1)
	}

	devfilePath := os.Args[1]

	fmt.Printf("** parsing devfile: %s **\n", devfilePath)
	devfileObject, err := devfileParser.ParseAndValidate(devfilePath)
	if err != nil {
		fmt.Printf("error upon validating: %v\n", err)
		return
	}
	fmt.Println("** Devfile Parsed and Validated Successfully **")
	fmt.Printf("devfile:\t%s\n", devfileObject.GetMetadataName())
	fmt.Printf("schemaVersion:\t%s\n", devfileObject.Data.GetSchemaVersion())
}
