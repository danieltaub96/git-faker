package yaml

import (
	"github.com/danieltaub96/git-faker/object"
	util "github.com/danieltaub96/git-faker/util/file"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
)

func ReadDataFile(name string) (*object.DataFile, error) {
	dataFile := GetDataFile()

	data, _ := ioutil.ReadFile(name)
	err := yaml.Unmarshal(data, &dataFile)

	if err != nil {
		log.Errorf("Error while parsing data file")
		return nil, err
	}

	log.Infof("Data loaded from file seccufully")
	return dataFile, nil
}

func ConfigDataFile(dataFilePath string) *object.DataFile {
	var dataFile = &object.DataFile{}
	var fileErr error

	if util.IsFileExists(dataFilePath) {
		log.Infof("File %s exists, starting reading...\n", util.AbsPath(dataFilePath))
		dataFile, fileErr = ReadDataFile(dataFilePath)

		if fileErr != nil {
			log.Infoln("Loading default data to git-faker")
			dataFile.LoadDefaults()
		} else {
			if len(dataFile.Messages) == 0 {
				log.Infoln("Loading default messages to git-faker")
				dataFile.SetDefaultMessages()
			}

			if len(dataFile.Emails) == 0 {
				log.Infoln("Loading default emails to git-faker")
				dataFile.SetDefaultEmails()
			}

			if len(dataFile.Names) == 0 {
				log.Infoln("Loading default names to git-faker")
				dataFile.SetDefaultNames()
			}
		}
	} else {
		log.Infoln("Loading default data to git-faker")
		dataFile.LoadDefaults()
	}

	log.Infof("git-faker data is: %s\n", dataFile)
	return dataFile
}

var (
	dataFileOnce sync.Once
	dataFile     *object.DataFile
)

func InitDataFile(dataFilePath string) {
	dataFileOnce.Do(func() {
		log.Infoln("Starting load data for git-faker")
		dataFile = ConfigDataFile(dataFilePath)
	})
}

func GetDataFile() *object.DataFile {
	return dataFile
}
