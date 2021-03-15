package infrastructure

//______________________________________________________________________________

import (
	"os"

	dbg "github.com/fatih/color"
	jsoniter "github.com/json-iterator/go"

	dom "github.com/cydside/cff/src/domain"
)

//______________________________________________________________________________

// AppConfig holds the configuration values from config.json file
var AppConfig dom.Configuration

//______________________________________________________________________________

// Initialize AppConfig
func initConfig() {
	loadAppConfig("config.json")
}

//______________________________________________________________________________

// Reads config.json and decode into AppConfig
func loadAppConfig(configFile string) {
	dbg.Blue("Lettura file di configurazioone...")
	file, err := os.Open(configFile)
	defer file.Close()
	if err != nil {
		dbg.Red("Errore nella lettura del file %s\n%s", configFile, err)
		return
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	decoder := json.NewDecoder(file)
	AppConfig = dom.Configuration{}

	err = decoder.Decode(&AppConfig)
	if err != nil {
		dbg.Red("Errore nella decodifica del file %s\n%s", configFile, err)
		return
	}

	for _, v := range AppConfig.Folders {
		dbg.Green("%s [%t]", v.AbsPath, v.CheckSubfolders)
	}
}
