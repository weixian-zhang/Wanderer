package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const content string = "file content from webapp Dolos"

type AzDiskFileIO struct {
	
}

func getFilePath() (string, error) {

	if configs.AzDiskMountFilePath == "" {
		return "", errors.New("read write file path not found in EnvVar('filepath')")
	} else {
		return configs.AzDiskMountFilePath, nil
	}
}

func (diskFileIo AzDiskFileIO) RunUseCase(done chan bool) { //implement interface UserCaseRunner

	timeticker := time.NewTicker(15 * time.Second)

	go func() {

		for {
			
			select {
				case <- done:
					timeticker.Stop()
					return
				case <- timeticker.C:
					executeFileIOonAzDisk()
			}
		}
		
	}()
	
}

func executeFileIOonAzDisk() {
	filePath, err := getFilePath()

	if err != nil {
		log.Error(err)
	} else {
		log.Info(fmt.Sprintf("AzDisk path @ %v", filePath))
	}
	
	werr := writeLocalFile(filePath)
	if hasErr(werr) {
		return
	}

	time.Sleep(1 * time.Second)

	content, err := readLocalFile(filePath)
	if hasErr(err) {
		return
	}

	log.Info(content)
}

func writeLocalFile(filePath string) error {

	file, err := os.Create(filePath)

	if !hasErr(err) {
		_, err := file.WriteString(content)

		log.Error(err)

		return nil
	}

	return err
}

func readLocalFile(filePath string) (string, error) {
	byteCon, err := ioutil.ReadFile(filePath)

	if !hasErr(err) {
		return string(byteCon), nil
	} else {
		return "", err
	}
}

func hasErr(err error) bool {
	if err != nil {
		log.Error(err)
		return true
	} else {
		return false
	}
}