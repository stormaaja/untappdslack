package config

import (
  "io/ioutil"
  "encoding/json"
)

type Config struct {
  File string
  Values ConfigValues
}

type ConfigValues struct {
  UntappdAPIKey string
}

func NewConfig(file string) Config {
  return Config{file, ConfigValues{""}}
}

func (config *Config) Read() error {
  data, fileError := ioutil.ReadFile(config.File)
  if fileError != nil {
    return fileError
  }

  return config.ReadJSON(data)
}

func (config *Config) ReadJSON(data []byte) error {
  config.Values = ConfigValues{ "" }
  jsonError := json.Unmarshal(data, &config.Values)

  return jsonError
}
