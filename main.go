package main

import (
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/encoding"
	"github.com/Yandex-Practicum/final-project-encoding-go/utils"
)

func Encode(data encoding.MyEncoder) error {
	return data.Encoding()
}

func main() {

	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		fmt.Printf("ошибка при создании папки data: %s", err.Error())
	}

	err = os.MkdirAll("data/input", os.ModePerm)
	if err != nil {
		fmt.Printf("ошибка при создании папки input: %s", err.Error())
	}

	err = os.MkdirAll("data/output", os.ModePerm)
	if err != nil {
		fmt.Printf("ошибка при создании папки output: %s", err.Error())
	}

	utils.CreateJSONFile()
	utils.CreateYAMLFile()

	jsonData := encoding.JSONData{FileInput: "data/input/jsonInput.json", FileOutput: "data/output/yamlOutput.yml"}
	err = Encode(&jsonData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из JSON в YAML: %s", err.Error())
	}

	yamlData := encoding.YAMLData{FileInput: "data/input/yamlInput.yml", FileOutput: "data/output/jsonOutput.json"}
	err = Encode(&yamlData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из YAML в JSON: %s", err.Error())
	}
}
