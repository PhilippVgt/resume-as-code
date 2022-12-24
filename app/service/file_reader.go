package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PhilippVgt/resume-as-code/app/model"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ReadDefinitionFromFile(file string) (model.Definition, error) {
	var input []byte
	var err error
	if input, err = ioutil.ReadFile(file); err != nil {
		return model.Definition{}, fmt.Errorf("unable to read file %s: %v", file, err)
	}

	extension := filepath.Ext(file)

	definition := prepareDefaultDefinition()
	if extension == ".yaml" || extension == ".yml" {
		if err = yaml.Unmarshal(input, &definition); err != nil {
			return model.Definition{}, fmt.Errorf("unable to parse YAML file %s: %v", file, err)
		}
		log.Infof("successfully read resume from YAML file %s", file)
	} else if extension == ".json" {
		if err = json.Unmarshal(input, &definition); err != nil {
			return model.Definition{}, fmt.Errorf("unable to parse JSON file %s: %v", file, err)
		}
		log.Infof("successfully read resume from JSON file %s", file)
	}

	if err = validate(&definition); err != nil {
		return model.Definition{}, fmt.Errorf("resume validation failed: %v", err)
	}

	return definition, nil
}

func validate(definition *model.Definition) error {
	// Major version updates contain breaking changes and should not be compatible
	resumeVersions := strings.Split(definition.Version, ".")
	codeVersions := strings.Split(model.ResumeVersion, ".")
	if len(resumeVersions) != 3 || resumeVersions[0] != codeVersions[0] {
		return fmt.Errorf("version of resume definition is not compatible, version %s.x.x expected", codeVersions[0])
	}

	if definition.Theme.Background != "" && !strings.HasPrefix(definition.Theme.Background, "#") {
		return fmt.Errorf("background color must be hex code with leading hash, e.g. '#305285', but '%s' was found", definition.Theme.Background)
	}
	if definition.Theme.Accent != "" && !strings.HasPrefix(definition.Theme.Accent, "#") {
		return fmt.Errorf("accent color must be hex code with leading hash, e.g. '#305285', but '%s' was found", definition.Theme.Accent)
	}

	// Either full name or first and last name must be provided
	if definition.Personal.FullName == "" && (definition.Personal.FirstName == "" || definition.Personal.LastName == "") {
		return errors.New("either 'personal.name' or both 'personal.firstName' and 'personal.lastName' must be set")
	}

	// Supported paper sizes are only A4 and Letter
	if strings.ToUpper(definition.Page.Size) == "A4" {
		definition.Page.Size = "A4"
	} else if strings.ToLower(definition.Page.Size) == "letter" {
		definition.Page.Size = "letter"
	} else {
		return fmt.Errorf("'settings.pageSize' can only be 'A4' or 'Letter' but '%s' was found", definition.Page.Size)
	}

	return nil
}

func prepareDefaultDefinition() model.Definition {
	return model.Definition{
		Version: model.ResumeVersion,
		Page: model.Page{
			Size:    "A4",
			Margins: model.PageMargins{Top: 15, Bottom: 20, Left: 5, Right: 10},
		},
	}
}
