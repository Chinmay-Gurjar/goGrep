package main

import (
        "fmt"
        "os"
        "regexp"
        "flag"
        "path/filepath"
	"strings"
	"sample.com/search"
        )

var (
        recursive       = flag.Bool("r", false, "recursive")
        count           = flag.Bool("c", false, "Just show counts")
        caseinsensitive = flag.Bool("i", false, "case-insensitive matching")
        write           = flag.Bool("o", false, "Write to file")
	mcount int
	outfile string
)



func writeResults(result search.Results){
        if result.Matched {
		fp, err := os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			panic(err)
		}
		defer fp.Close()

		_, ferr := fp.WriteString(result.File_path + "\t" + strings.TrimRight(result.Line, "\n") + "\n")
		if ferr != nil {
			fmt.Println("Error while writing the file:", ferr)
		}
	}
}



func parseInput() (string, []string,[]string) {
	flag.Parse()
        args := flag.Args()
        pattern := args[0]
        filenames := args[1:]
	return pattern, filenames, args
}

func getFilePath(filenames []string) []string{
	filepaths := make([]string, 0)
        for _, path := range filenames {
                filepath.Walk(path, func(file_path string, file os.FileInfo, err error) error {
                        if err != nil {
                                return nil
                        }
			if file.IsDir() && !*recursive {
                                return filepath.SkipDir
                        }
			filepaths = append(filepaths, file_path)
                        return nil
                })
	}
	return filepaths
}


func main() {
	pattern, filenames, args := parseInput()
	if *write {
		outfile = args[0]
		pattern = args[1]
		filenames = args[2:]
	}
        if *caseinsensitive {
		 pattern = "(?i)" + pattern
        }

	//if len(args) < 2 {
	//	wg.Add(1)
	//	go search("")
	//}

	paths := getFilePath(filenames)

	search.Search(patterg, paths)
	if *count {
		fmt.Println(mcount)
	}
}
