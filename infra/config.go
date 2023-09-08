package infra

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"

	"github.com/ceelsoin/subzero/domain"
	"github.com/ceelsoin/subzero/utils"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	FilePath string
	Lock     sync.Mutex
}

var configurationPath = path.Join(defaultFilePath, "config.yaml")

func NewConfigurationInstance() *Configuration {
	configInstance = &Configuration{
		FilePath: configurationPath,
	}

	if err := os.MkdirAll(path.Join(utils.GetHomeDir(), ".subzero"), 0755); err != nil {
		panic(err)
	}

	if _, err := os.Stat(configurationPath); os.IsNotExist(err) {
		err := instance.writeYAML([]domain.FileInfo{})
		if err != nil {
			log.Fatalf("Erro ao criar o arquivo YAML: %v", err)
		}
	}

	return configInstance
}

func GetConfigurationInstance() *Configuration {
	if configInstance == nil {
		configInstance = &Configuration{
			FilePath: configurationPath,
		}
	}
	return configInstance
}

func (db *Configuration) SaveConfig(config domain.Configuration) error {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	yamlData, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(db.FilePath, yamlData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (db *Configuration) GetConfig() (domain.Configuration, error) {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	yamlFile, err := ioutil.ReadFile(db.FilePath)
	if err != nil {
		return domain.Configuration{}, err
	}

	var data domain.Configuration
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return domain.Configuration{}, err
	}

	return data, nil
}
