package file

import (
	"testing"
	"os"
	"fmt"
)

func createEmptyFile (dirName string, name string) {
	_, err := os.Create(dirName + "/" + name)
	if err != nil {
		fmt.Println("could not create an Empty File", err)
	}
}

func TestGetFilePath(t *testing.T) {
	dirName := "TestFolder"
	recDirName := "RecurrTestFolder"
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Cannot create a directory!", err)
	}
	err = os.Mkdir(dirName + "/" + recDirName, 0755)
	if err != nil {
		fmt.Println("Cannot create a directory!", err)
	}
	defer os.RemoveAll(dirName)
	defer os.RemoveAll(recDirName)

	testFileNames := []string{"testFile1", "testFile2", "testFile3", "testFile4", "RecurrTestFolder/RtestFile1", "RecurrTestFolder/RtestFile2"}

	for _, fileName := range testFileNames {
		createEmptyFile(dirName, fileName)
	}
	inCorrectfiles, err := GetFilePath([]string{dirName}, false)
	if err != nil {
		t.Errorf("Error while getting the files: %v", err)
	}
	if len(inCorrectfiles) != 4 {
		t.Errorf("Expected %v, got %v", testFileNames, inCorrectfiles)
	}


	files, err := GetFilePath([]string{dirName}, true)
	if err != nil {
		t.Errorf("Error while getting the files: %v", err)
	}

	if len(files) != (len(testFileNames)) {
		t.Errorf("Expected %v, got %v", testFileNames, files)
	}

}
