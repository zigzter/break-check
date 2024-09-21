package parsers

import (
	"encoding/json"
	"os"
)

type Data struct {
	Name            string
	Dependencies    map[string]string
	DevDependencies map[string]string
}

func ParsePackageJSON() (Data, error) {
	content, err := os.ReadFile("package.json")
	if err != nil {
		return Data{}, err
	}
	var parsedJson Data
	err = json.Unmarshal(content, &parsedJson)
	if err != nil {
		return Data{}, err
	}
	return parsedJson, nil
}
