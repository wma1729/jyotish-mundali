package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func (config *Config) LoadFromYaml(configFile string) error {
	f, err := os.Open(configFile)
	if err != nil {
		log.Println(err)
		return err
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		log.Printf("failed to unmarshal %s: %s", configFile, err)
		return err
	}

	return nil
}
