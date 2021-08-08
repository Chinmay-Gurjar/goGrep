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
)

type results struct {
        matched bool
        line string
        file_path string
}

func printResults(result results) {
        if result.matched {
		fmt.Println(result.file_path, strings.TrimRight(result.line, "\n"))
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

func search(wg *sync.WaitGroup, filename string) {
        defer wg.Done()
	var reader *bufio.Reader
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
                        result := results{match, readbuffer, filename}

			if !*count && !*write{
				printResults(result)
			} else if !*count && *write{
				writeResults(result)
			}
                } else {
                        break
                }
        }
}

func main() {
        flag.Parse()
        args := flag.Args()
        pattern := args[0]
        filenames := args[1:]
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
		go search(&wg, "")
	}

        for _, path := range filenames {
                filepath.Walk(path, func(file_path string, file os.FileInfo, err error) error {
                        if err != nil {
                                return nil
                        }
			if file.IsDir() && !*recursive {
                                return filepath.SkipDir
                        }

                        wg.Add(1)
                        go search(&wg, file_path)
                        return nil
                })
        }
        wg.Wait()
	if *count {
		fmt.Println(mcount)
	}
}
