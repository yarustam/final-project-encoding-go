package encoding

import (
	"os"
	"testing"

	"github.com/Yandex-Practicum/final-project-encoding-go/utils"
	"github.com/stretchr/testify/suite"
)

func TestEncodingSuite(t *testing.T) {
	suite.Run(t, new(EncodingSuite))
}

type EncodingSuite struct {
	suite.Suite
}

func (s *EncodingSuite) TestEncoding() {
	utils.CreateJSONFile()
	utils.CreateYAMLFile()
	s.Run("yaml equal", func() {
		jsonData := JSONData{FileInput: "jsonInput.json", FileOutput: "yamlOutput.yml"}

		err := jsonData.Encoding()
		s.Assert().NoError(err)

		expected, err := os.ReadFile("yamlInput.yml")
		s.Assert().NoError(err)

		res, err := os.ReadFile(jsonData.FileOutput)
		s.Assert().NoError(err)

		s.Assert().YAMLEq(string(expected), string(res), "Результат выполнения функции (*JSONData)Encoding() не совпадает с ожидаемым.")
	})

	s.Run("json equal", func() {
		yamlData := YAMLData{FileInput: "yamlInput.yml", FileOutput: "jsonOutput.json"}

		err := yamlData.Encoding()
		s.Assert().NoError(err)

		expected, err := os.ReadFile("jsonInput.json")
		s.Assert().NoError(err)

		res, err := os.ReadFile(yamlData.FileOutput)
		s.Assert().NoError(err)

		s.Assert().JSONEq(string(expected), string(res), "Результат выполнения функции (*YAMLData)Encoding() не совпадает с ожидаемым.")
	})
}
