package search

import (
	"fmt"
	"sync"
	"bufio"
	"regexp"
	"os"
	"strings"
)

var (
	compiledPattern *regexp.Regexp
	mcount int
)

type Results struct {
        Matched bool
        Line string
        File_path string
}

func printResults(result_slice []Results) {
	for i := range result_slice{
		if result_slice[i].Matched {
			fmt.Println(result_slice[i].File_path, strings.TrimRight(result_slice[i].Line, "\n"))
		}
	}
}

func searchSingleFile(filename string) []Results {
	var reader *bufio.Reader
	result := make([]Results, 1)
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
                        result = append(result, Results{match, readbuffer, filename})

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

func searchAndPrint(wg *sync.WaitGroup, filepath string) {
	defer wg.Done()
	result := searchSingleFile(filepath)
	printResults(result)
}

func Search(pattern string, filepaths []string) {
        var wg sync.WaitGroup
        compiledPattern = regexp.MustCompile(pattern)
	for _, path := range filepaths {
		wg.Add(1)
		go searchAndPrint(&wg, path)
	}
	wg.Wait()
}
