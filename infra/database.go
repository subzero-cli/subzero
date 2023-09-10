package infra

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"

	"github.com/subzero-cli/subzero/domain"
	"github.com/subzero-cli/subzero/utils"
	"gopkg.in/yaml.v2"
)

var (
	instance       *Database
	configInstance *Configuration
	once           sync.Once
)

type Database struct {
	FilePath string
	Lock     sync.Mutex
}

var defaultFilePath = path.Join(utils.GetHomeDir(), ".subzero")
var databasePath = path.Join(defaultFilePath, "database.yaml")

func NewDatabaseInstance() *Database {
	once.Do(func() {
		instance = &Database{
			FilePath: databasePath,
		}

		if err := os.MkdirAll(path.Join(utils.GetHomeDir(), ".subzero"), 0755); err != nil {
			panic(err)
		}

		if _, err := os.Stat(databasePath); os.IsNotExist(err) {
			err := instance.writeYAML([]domain.FileInfo{})
			if err != nil {
				log.Fatalf("Erro ao criar o arquivo YAML: %v", err)
			}
		}
	})

	return instance
}

func GetDatabaseInstance() *Database {
	if instance == nil {
		instance = &Database{
			FilePath: databasePath,
		}
	}

	return instance
}

func (db *Database) Create(info domain.FileInfo) error {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	data, err := db.readYAML()
	if err != nil {
		return err
	}

	for i := range data {
		if data[i].ID == info.ID {
			data[i] = info
			return db.writeYAML(data)
		}
	}

	data = append(data, info)

	return db.writeYAML(data)
}

func (db *Database) ReadAll() ([]domain.FileInfo, error) {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	return db.readYAML()
}

func (db *Database) GetByFileName(sanitizedName string) (*domain.FileInfo, error) {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	data, err := db.readYAML()
	if err != nil {
		return nil, err
	}

	for i := range data {
		if data[i].SanitizedName == sanitizedName {
			return &data[i], nil
		}
	}

	return nil, errors.New("yaml database read error")
}

func (db *Database) Update(id string, info domain.FileInfo) error {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	data, err := db.readYAML()
	if err != nil {
		return err
	}

	updated := false
	for i := range data {
		if data[i].ID == id {
			data[i] = info
			updated = true
			break
		}
	}

	if !updated {
		return errors.New("FileInfo não encontrado para atualização")
	}

	return db.writeYAML(data)
}

func (db *Database) Delete(id string) error {
	db.Lock.Lock()
	defer db.Lock.Unlock()

	data, err := db.readYAML()
	if err != nil {
		return err
	}

	deleted := false
	for i := range data {
		if data[i].ID == id {
			data = append(data[:i], data[i+1:]...)
			deleted = true
			break
		}
	}

	if !deleted {
		return errors.New("FileInfo não encontrado para deleção")
	}

	return db.writeYAML(data)
}

func (db *Database) readYAML() ([]domain.FileInfo, error) {
	yamlFile, err := ioutil.ReadFile(db.FilePath)
	if err != nil {
		return nil, err
	}

	var data []domain.FileInfo
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *Database) writeYAML(data []domain.FileInfo) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(db.FilePath, yamlData, 0644)
	if err != nil {
		return err
	}

	return nil
}
