package search

import (
	"fmt"
	"sync"
	"bufio"
	"regexp"
	"os"
	"strings"
	"io"
)

var (
	compiledPattern *regexp.Regexp
	mcount int
	allResult = make([][]Results, 0)
)

type Results struct {
        Matched bool
        Line string
        File_path string
}

func PrintResults(result_slice []Results) {
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
		fp, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer fp.Close()
		reader = bufio.NewReader(fp)
	}
        for {
                if readbuffer, err := reader.ReadString('\n'); err == nil{
                        match := compiledPattern.Match([]byte(readbuffer))
			if match {
				mcount++
				result = append(result, Results{match, readbuffer, filename})
			}
			if filename == "" {
				PrintResults(result)
			}

                } else if(err==io.EOF) {
			break

                } else {
			fmt.Println("Error:", err)
                        break
		}
        }
	return result
}

func searchAndPrint(wg *sync.WaitGroup, filepath string) {
	defer wg.Done()
	result := searchSingleFile(filepath)
	allResult = append(allResult, result)
}

func Search(pattern string, filepaths []string) ([][]Results, int) {
        var wg sync.WaitGroup
        compiledPattern = regexp.MustCompile(pattern)
	if len(filepaths) == 0 {
		searchAndPrint(&wg, "")
	}
	for _, path := range filepaths {
		wg.Add(1)
		go searchAndPrint(&wg, path)
	}
	wg.Wait()
	return allResult, mcount
}
