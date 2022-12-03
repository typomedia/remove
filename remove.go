package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"path/filepath"
	"strings"
)

const info = `Usage: remove [OPTION]... [FILE]...
  -h, --help
  -i, --input file     Line separated list of files
  [FILE]...            File or directory to remove

Example: remove --input list.txt
`

func main() {
	println("Remove by Philipp Speck [Version 1.0]")
	println("Copyright (C) 2022 Typomedia Foundation.")
	println()

	var list string

	flag.StringVarP(&list,
		"input",
		"i",
		"",
		"Line separated file list")
	flag.BoolP("help", "h", false, "")
	flag.Usage = func() {
		//flag.PrintDefaults()
		fmt.Print(info) // override default usage
	}
	flag.Parse()

	args := flag.Args()

	if len(list) == 0 && len(args) == 0 {
		flag.Usage()
	}

	if len(args) > 0 {
		for _, path := range args {
			remove(path)
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
			remove(scanner.Text())
		}

		file.Close()
	}
}

func remove(path string) {
	path = strings.Replace(path, "/", string(filepath.Separator), -1)

	files, err := filepath.Glob(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println("Remove: " + file)
		if err := os.RemoveAll(file); err != nil {
			fmt.Println(err)
		}
	}
}
