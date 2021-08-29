package main

import (
        "fmt"
        "os"
        "bufio"
        "regexp"
        "flag"
        "sync"
        "path/filepath"
	"strings"
        )

var (
        recursive       = flag.Bool("r", false, "recursive")
        count           = flag.Bool("c", false, "Just show counts")
        caseinsensitive = flag.Bool("i", false, "case-insensitive matching")
        write           = flag.Bool("o", false, "Write to file")
        compiledPattern *regexp.Regexp
        wg sync.WaitGroup
	mcount int
	outfile string
	grepResult	= make([][]results, 0)
)

type results struct {
        matched bool
        line string
        file_path string
}

func printResults(result_slice []results) {
	for i := range result_slice{
		if result_slice[i].matched {
			fmt.Println(result_slice[i].file_path, strings.TrimRight(result_slice[i].line, "\n"))
		}
	}
}

func writeResults(result results){
        if result.matched {
		fp, err := os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			panic(err)
		}
		defer fp.Close()

		_, ferr := fp.WriteString(result.file_path + "\t" + strings.TrimRight(result.line, "\n") + "\n")
		if ferr != nil {
			fmt.Println("Error while writing the file:", ferr)
		}
	}
}

func search(filename string) []results {
	var reader *bufio.Reader
	result := make([]results, 1)
	if filename == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		fp, _ := os.Open(filename)
		defer fp.Close()
		reader = bufio.NewReader(fp)
	}
        for {
                if readbuffer, err := reader.ReadString('\n'); err == nil{
                        match := compiledPattern.Match([]byte(readbuffer))
			if match {
				mcount++
			}
                        result = append(result, results{match, readbuffer, filename})

		//	if !*count && !*write{
		//		printResults(result)
		//	} else if !*count && *write{
		//		writeResults(result)
		//	}
                } else {
                        break
                }
        }
	return result
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

func searchAndPrint(wg *sync.WaitGroup, filepath string) {
	defer wg.Done()
	result := search(filepath)
	printResults(result)
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

        compiledPattern = regexp.MustCompile(pattern)
	if len(args) < 2 {
		wg.Add(1)
		go search("")
	}

	paths := getFilePath(filenames)


	for _, path := range paths {
		wg.Add(1)
		go searchAndPrint(&wg, path)
	}
        wg.Wait()
	if *count {
		fmt.Println(mcount)
	}
}
