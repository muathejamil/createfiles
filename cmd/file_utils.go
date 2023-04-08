package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
)

const PREFIX = "file"

var wg sync.WaitGroup

type FileWriter struct {
	fileName string
	filePath string
	fileSize int
}

func NewFileWriter(filename string, filepath string, filesize int) *FileWriter {
	newFile := FileWriter{
		fileName: filename,
		filePath: filepath,
		fileSize: filesize,
	}
	return &newFile
}

func CreateBatch(path string, count int, sizeInKb int) {
	dummyData := ReadDummyContent()

	for i := 0; i < count; i++ {
		fileName := fmt.Sprintf("%s%d.txt", PREFIX, i)
		filePath := filepath.Join(path, filepath.Base(fileName))
		fmt.Printf("The path of the file is %s\n", filePath)
		_, err := os.Create(filePath)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Errorf("Error in creating the file: %s", fileName)
			os.Exit(1)
		}
		file := NewFileWriter(fileName, filePath, sizeInKb)
		wg.Add(1)
		go PopulateFileContent(file, dummyData)
	}
	wg.Wait()
}

func ReadDummyContent() []byte {
	readPath := viper.GetString("createfiles.path")
	fileData, err := os.ReadFile(readPath)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("Error in reading the file: %s", readPath)
		os.Exit(1)
	}
	return fileData
}

func CreateFile(file FileWriter) {
	newFile, err := os.Create(filepath.Join(file.filePath, filepath.Base(file.fileName)))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("Error in creating the file: %s", file.fileName)
		os.Exit(1)
	}
	defer newFile.Close()
}

func PopulateFileContent(file *FileWriter, data []byte) {
	defer wg.Done()

	log.Infof("write to file len %d", len(data))
	toBePopulatedFile, err := os.OpenFile(file.filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer toBePopulatedFile.Close()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("Error in opning the file: %s", toBePopulatedFile.Name())
		os.Exit(1)
	}
	chunkSize := len(data) * 4
	for i := 0; i < file.fileSize/chunkSize; i++ {
		if _, err = toBePopulatedFile.Write(data); err != nil {
			panic(err)
		}
	}

	if remainingBytes := file.fileSize % chunkSize; remainingBytes > 0 {
		if _, err = toBePopulatedFile.Write(data[:remainingBytes]); err != nil {
			panic(err)
		}
	}
}
