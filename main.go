package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"github.com/Yandex-Practicum/final-project-encoding-go/utils"
	"gopkg.in/yaml.v3"
)

type JSONData struct {
	User       *models.User
	fileInput  string
	fileOutput string
}

type YAMLData struct {
	User       *models.User
	fileInput  string
	fileOutput string
}

type MyEncoder interface {
	Encoding() error
}

func (j *JSONData) Encoding() error {
	jsonFile, err := os.ReadFile(j.fileInput)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonFile, &j.User)
	if err != nil {
		return err
	}

	yamlData, err := yaml.Marshal(&j.User)
	if err != nil {
		return err
	}

	yamlFile, err := os.Create(j.fileOutput)
	if err != nil {
		return err
	}

	defer yamlFile.Close()

	_, err = yamlFile.Write(yamlData)
	if err != nil {
		return err
	}

	return nil
}

func (y *YAMLData) Encoding() error {
	yamlFile, err := os.ReadFile(y.fileInput)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &y.User)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(&y.User)
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(y.fileOutput)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func Encode(data MyEncoder) error {
	return data.Encoding()
}

func main() {
	utils.CreateJSONFile()
	utils.CreateYAMLFile()

	jsonData := JSONData{fileInput: "jsonInput.json", fileOutput: "yamlOutput.yml"}
	err := Encode(&jsonData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из JSON в YAML: %s", err.Error())
	}

	yamlData := YAMLData{fileInput: "yamlInput.yml", fileOutput: "jsonOutput.json"}
	err = Encode(&yamlData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из YAML в JSON: %s", err.Error())
	}
}
