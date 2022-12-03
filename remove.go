package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"path/filepath"
	"strings"
)

const version = "1.0"

const usage = `Usage: remove [OPTION]... [FILE]...
  -h, --help
  -i, --input file     line separated list of files
  -v, --verbose        output what is being done
  [FILE]...            file or directory to remove

Example: remove --input list.txt
`

func main() {
	println("remove by Philipp Speck [Version " + version + "]")
	println("Copyright (C) 2022 Typomedia Foundation.")
	println()

	var list string

	flag.StringVarP(&list,
		"input",
		"i",
		"",
		"Line separated file list")
	flag.BoolP("help", "h", false, "")
	verbose := flag.BoolP("verbose", "v", false, "")
	flag.Usage = func() {
		//flag.PrintDefaults()
		fmt.Print(usage) // override default usage
	}
	flag.Parse()

	args := flag.Args()

	if len(list) == 0 && len(args) == 0 {
		flag.Usage()
	}

	if len(args) > 0 {
		for _, path := range args {
			remove(path, *verbose)
		}
	}

	if list != "" {
		file, err := os.Open(list)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			remove(scanner.Text(), *verbose)
		}

		file.Close()
	}
}

func remove(path string, verbose bool) {
	path = strings.Replace(path, "/", string(filepath.Separator), -1)

	files, err := filepath.Glob(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if verbose {
			fmt.Println("remove " + file)
		}
		if err := os.RemoveAll(file); err != nil {
			fmt.Println(err)
		}
	}
}
