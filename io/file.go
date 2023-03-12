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

func CreateBatch(path string, count int, sizeInKb int) {
	prefix := "file"
	for i := 0; i < count; i++ {
		filename := fmt.Sprintf("%s%d.txt", prefix, i)
		filePath := filepath.Join(path, filepath.Base(filename))
		_, err := os.Create(filePath)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Errorf("Error in creating the file: %s", filename)
			os.Exit(1)
		}
		PopulateTheFile(filePath, sizeInKb)
	}
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

func PopulateTheFile(filePath string, filesize int) {
	readPath := "C:\\Users\\Lenovo\\GolandProjects\\createfiles\\inFiles\\1024bytesFile.txt"
	//TODO: Integrate with 3rd party to generate random 1024 bytes
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
			"error": dataErr,
		}).Errorf("Error in reading the file: %s", readPath)
		os.Exit(1)
	}
	data := string(fileData)

	toBePopulatedFile, PopulateErr := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if PopulateErr != nil {
		log.WithFields(log.Fields{
			"error": PopulateErr,
		}).Errorf("Error in opning the file: %s", toBePopulatedFile.Name())
		os.Exit(1)
	}
	for i := 0; i < filesize; i++ {
		if _, err = toBePopulatedFile.WriteString(data); err != nil {
			panic(err)
		}
	}
	defer toBePopulatedFile.Close()
}
