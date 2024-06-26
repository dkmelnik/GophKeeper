package utils

import (
	"encoding/json"
	"log"
	"os"
	"os/user"

	"github.com/dkmelnik/GophKeeper/configs"
)

var filename = "/config.json"

func init() {
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}
}

func loadConfig() error {
	path, err := configFile()
	if err != nil {
		return err
	}
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = createDefaultConfigFile(path); err != nil {
			return err
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&configs.Client); err != nil {
		return err
	}

	return nil
}

func createDefaultConfigFile(fl string) error {
	configs.Client.ADDR = "localhost:1234"
	ph, err := configFilePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(ph, 0755); err != nil {
		return err
	}

	file, err := os.Create(fl)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(configs.Client); err != nil {
		return err
	}

	return nil
}

func configFilePath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.HomeDir + "/gophkeeper", nil
}

func configFile() (string, error) {
	path, err := configFilePath()
	if err != nil {
		return "", err
	}
	return path + filename, nil
}

func UpdateConfigFile(cnf configs.ClientConf) error {
	path, err := configFile()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err = encoder.Encode(&cnf); err != nil {
		return err
	}

	return nil
}
