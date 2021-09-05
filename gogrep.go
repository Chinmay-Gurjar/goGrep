package main

import (
        "fmt"
        "os"
        "flag"
	"sample.com/search"
	"errors"
	"sample.com/file"
	"strconv"
        )

var (
        recursive       = flag.Bool("r", false, "recursive")
        count           = flag.Bool("c", false, "Just show counts")
        caseinsensitive = flag.Bool("i", false, "case-insensitive matching")
        write           = flag.Bool("o", false, "Write to file")
)

func parseInput() (string, []string, []string, error) {
	flag.Parse()
        args := flag.Args()
	if len(args) == 0 {
		err := errors.New("No arguments")
		return "", nil, nil, err
	}
        pattern := args[0]
	if len(pattern) == 0 {
		err := errors.New("Empty search pattern")
		return "", nil, nil, err
	}
        filenames := args[1:]
	return pattern, filenames, args, nil
}

func showErrorAndExit(err error) {
	fmt.Println("The code had follwing error: ", err)
	os.Exit(0)
}

func main() {
	var outfile string
	pattern, filenames, args, err := parseInput()
	if err != nil{
		showErrorAndExit(err)
	}
	if *write {
		outfile = args[0]
		pattern = args[1]
		filenames = args[2:]
	}
        if *caseinsensitive {
		 pattern = "(?i)" + pattern
        }

	if len(args) < 2 {
		search.Search(pattern, []string{})
	}

	paths, err := file.GetFilePaths(filenames, *recursive)
	if err != nil {
		os.Exit(0)
	}

	results, mcount, err := search.Search(pattern, paths)
	if err != nil {
		showErrorAndExit(err)
	}
	if *count {
		if *write {
			file.WriteResults(strconv.Itoa(mcount) + "\n", outfile)
		} else {
			fmt.Println(mcount)
		}
	} else if *write {
		for _, result := range results {
			for _, line := range result {
				file.WriteGrepResults(line, outfile)
			}
		}

	} else {
		for _, result := range results {
			search.PrintResults(result)
		}
	}


}
