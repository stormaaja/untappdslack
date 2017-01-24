package main

import (
  "github.com/stormaaja/untappdslack/config"
)

const ConfigFile string = "config.json"

func main() {
  appConfig := config.NewConfig(ConfigFile)
  appConfig.Read()
}
