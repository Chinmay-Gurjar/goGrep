package file

import (
        "os"
        "path/filepath"
        "sample.com/search"
	"strings"
	"fmt"
        )

func GetFilePath(filenames []string, isRecursive bool) ([]string, error){
	filepaths := make([]string, 0)
        for _, path := range filenames {
		filepath.Walk(path, func(file_path string, file os.FileInfo, err error) error {
			if path == file_path {
				return nil
			}
                        if err != nil {
                                return nil
                        }
			if !file.IsDir() {
				filepaths = append(filepaths, file_path)
				return nil
			}
			if file.IsDir() && !isRecursive {
				return filepath.SkipDir
                        }

                        return nil
                })
	}
	return filepaths, nil
}

func WriteResults(data string, outfile string){
	fp, err := os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer fp.Close()
	_, ferr := fp.WriteString(data)
	if ferr != nil {
		fmt.Println("Error while writing the file:", ferr)
	}
}

func WriteGrepResults(result search.Results, outfile string) {
	data := result.File_path + "\t" + strings.TrimRight(result.Line, "\n") + "\n"
	WriteResults(data, outfile)
}
