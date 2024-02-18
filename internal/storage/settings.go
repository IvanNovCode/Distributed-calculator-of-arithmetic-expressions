package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Settings struct {
	Addition       string `json:"addition"`
	Subtraction    string `json:"subtraction"`
	Multiplication string `json:"multiplication"`
	Division       string `json:"division"`
}

const settingsFile = "database/calcSettings.json"

func SetNewSetting(setting string) error {
	settings, err := getFromTimeoutsFile(settingsFile)
	if err != nil {
		logger.Printf("Cannot read getFromSettingsFile: %s", err)
		return err
	}
	value := ""
	var operation int
	for _, ch := range setting {
		char := string(ch)
		switch char {
		case "+":
			operation = 1
		case "-":
			operation = 2
		case "*":
			operation = 3
		case "/":
			operation = 4
		default:
			value += char
		}
	}

	switch operation {
	case 1:
		settings.Addition = value
	case 2:
		settings.Subtraction = value
	case 3:
		settings.Multiplication = value
	case 4:
		settings.Division = value
	}

	return saveToFile(settingsFile, settings)
}

func GetSettings() (*Settings, error) {
	return getFromTimeoutsFile(settingsFile)
}

func getFromTimeoutsFile(fileName string) (*Settings, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		logger.Printf("File does not exist: %s", err)
		return &Settings{}, nil
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		logger.Printf("Cannot read the file: %s", err)
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	if len(data) == 0 {
		logger.Println("File is empty, returning new Expressions object")
		return &Settings{}, nil
	}

	var timeout Settings
	if err := json.Unmarshal(data, &timeout); err != nil {
		logger.Printf("Cannot unmarshal data: %s", err)
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	logger.Printf("List retrieved from %s", expressionsFile)
	return &timeout, nil
}
