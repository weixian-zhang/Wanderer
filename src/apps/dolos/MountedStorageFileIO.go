package main

import (
	"io/ioutil"
	"os"
	log "github.com/sirupsen/logrus"
	"time"
)

const fileEnvVarKey string = "filepath"
const content string = "file content from webapp Dolos"

type AzDiskFileIO struct {
	
}

func getFileName() (string) {

	var filePath string = ""

	val, err := utils.getEnvarVal(fileEnvVarKey)
	log.Error(err)

	filePath = val

	return filePath
}

func (diskFileIo AzDiskFileIO) RunUseCase(done chan bool) { //implement interface UserCaseRunner

	timeticker := time.NewTicker(5 * time.Second)

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
	filePath := getFileName()

			err := writeLocalFile(filePath)
			if hasErr(err) {
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