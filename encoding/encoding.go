package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/hive-bootcamp/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	data, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf(err.Error())
		return err

	}
	if err := json.Unmarshal(data, &j.DockerCompose); err != nil {
		fmt.Printf(err.Error())
		return err
	}
	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	if err := os.WriteFile(j.FileOutput, yamlData, 0644); err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	data, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf(err.Error())
		return err

	}
	if err := yaml.Unmarshal(data, &y.DockerCompose); err != nil {
		fmt.Printf(err.Error())
		return err
	}
	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Printf(err.Error())
		return err

	}
	if err := os.WriteFile(y.FileOutput, jsonData, 0644); err != nil {
		fmt.Printf(err.Error())
		return err

	}
	return nil
}
