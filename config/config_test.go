package config

import (
  "testing"
  "fmt"
  "os"
)

const (
  TestConfigFile string = "test_config_file.json"
  TestConfigJSON string = `
  {
    "UntappdAPIKey": "SomeAPIKey123"
  }
  `
)

func CreateTestConfigFile() error {
  f, err := os.Create(TestConfigFile)
  defer f.Close()
  if err != nil {
    return err
  }
  f.Write([]byte(TestConfigJSON))
  return nil
}

func DeleteTestConfigFile() error {
  return os.Remove(TestConfigFile)
}

func TestMain(t *testing.T) {
  err := CreateTestConfigFile()
  if err != nil {
    t.Error("Test aborted. Cannot create config file: ", err)
    return
  }

  TestingNewConfig(t)
  TestingReadJSON(t)
  TestingReadConfig(t)

  err = DeleteTestConfigFile()
  if err != nil {
    t.Error("Cannot delete config file: ", err)
    return
  }

}

func TestingNewConfig(t *testing.T) {
  testConfig := NewConfig(TestConfigFile)
  if testConfig.File != TestConfigFile {
    t.Error(
      fmt.Sprintf("Config file should be %s but got %s",
        TestConfigFile,  testConfig.File))
  }
}

func TestingReadJSON(t *testing.T) {
  testConfig := NewConfig(TestConfigFile)
  err := testConfig.ReadJSON([]byte(TestConfigJSON))
  if err != nil {
    t.Error("Error reading file: ", err)
  }
}

func TestingReadConfig(t *testing.T) {
  testConfig := NewConfig(TestConfigFile)
  err := testConfig.Read()
  if err != nil {
    t.Error("Error reading file: ", err)
  }
}
