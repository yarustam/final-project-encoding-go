package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

func CreateJSONFile() {
	jsonFile, err := os.Create("jsonInput.json")
	if err != nil {
		fmt.Printf("json file creation fail: %s", err.Error())
	}

	defer jsonFile.Close()

	user := models.User{
		ID:       "D45Gt72",
		Name:     "Вадим",
		Email:    "vadim_ka@email.com",
		Password: "IwBkn19hsF",
	}

	out, err := json.Marshal(&user)
	if err != nil {
		fmt.Printf("json encoding fail: %s", err.Error())
	}

	_, err = jsonFile.Write(out)
	if err != nil {
		fmt.Printf("writing data fail: %s", err.Error())
	}
}

func CreateYAMLFile() {
	yamlFile, err := os.Create("yamlInput.yml")
	if err != nil {
		fmt.Printf("yaml file creation fail: %s", err.Error())
	}

	defer yamlFile.Close()

	user := models.User{
		ID:       "NV343GF",
		Name:     "Виталик",
		Email:    "vitamin@emil.com",
		Password: "BadPassword",
	}

	out, err := yaml.Marshal(&user)
	if err != nil {
		fmt.Printf("yaml encoding fail: %s", err.Error())
	}

	_, err = yamlFile.Write(out)
	if err != nil {
		fmt.Printf("writing data fail: %s", err.Error())
	}
}
