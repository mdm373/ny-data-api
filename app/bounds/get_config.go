package bounds

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type BoundConfig struct {
	TableName   string `json:"tableName"`
	Route       string `json:"route"`
	DisplayName string `json:displayName`
}

type BoundsConfig struct {
	Bounds []BoundConfig `json:"bounds"`
}

var GetConfig = func() ([]BoundConfig, error) {
	jsonFile, err := os.Open("./static/bounds-config.json")
	defer func() {
		err := jsonFile.Close()
		if err != nil {
			log.Printf("error closing config: %+v", err)
		}
	}()
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var config BoundsConfig
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}
	return config.Bounds, nil
}
