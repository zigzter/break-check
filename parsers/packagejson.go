package parsers

import (
	"encoding/json"
	"os"
)

type PackageJSONData struct {
	Name            string
	Dependencies    map[string]string
	DevDependencies map[string]string
}

func ParsePackageJSON() (PackageJSONData, error) {
	content, err := os.ReadFile("package.json")
	if err != nil {
		return PackageJSONData{}, err
	}
	var parsedJson PackageJSONData
	err = json.Unmarshal(content, &parsedJson)
	if err != nil {
		return PackageJSONData{}, err
	}
	return parsedJson, nil
}
