package util

import (
	"os"
	"strings"
)

type MrtgType []string

const FileName = "mrtg.txt"

var fileHandle *os.File

func OpenFile(fileName string) {
	var err error
	fileHandle, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)

	CheckErr(err)
}

func CloseFile() {
	err := fileHandle.Close()
	CheckErr(err)
}

func WriteToFile(mrtgSummary MrtgType) {
	fileHandle.Write([]byte(mrtgSummary.toString()))
	fileHandle.Write([]byte("\n"))
}

func (t MrtgType) toString() string {
	return strings.Join([]string(t), ",")
}
