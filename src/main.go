package main

import loadConfig "main.go/connections/load_configs"

func init() {
	loadConfig.LoadConfigs("env")
}

func main() {

}
