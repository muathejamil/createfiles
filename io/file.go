package io

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type File struct {
	FileName string
	FilePath string
	FileSize int32
}

func NewFile(filename string, filepath string, filesize int32) File {
	newFile := File{
		FileName: filename,
		FilePath: filepath,
		FileSize: filesize,
	}
	return newFile
}

func CreateFile(file File) {
	newFile, err := os.Create(filepath.Join(file.FilePath, filepath.Base(file.FileName)))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("Error in creating the file: %s", file.FileName)
		os.Exit(1)
	}
	defer newFile.Close()
}

func PopulateTheFile(filePath string, filesize int32) {
	readPath := "inFiles/1024bytesFile.txt"
	readFile, err := os.OpenFile(readPath, os.O_RDONLY, 0644)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("Error in opening the file: %s", readPath)
		os.Exit(1)
	}
	defer readFile.Close()
	fileData, dataErr := os.ReadFile(readPath)
	if dataErr != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("Error in reading the file: %s", readPath)
		os.Exit(1)
	}
	fmt.Println(string(fileData))
}
