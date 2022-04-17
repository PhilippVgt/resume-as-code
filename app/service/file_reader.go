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

func ReadResumeFromFile(file string) (*model.Resume, error) {
	var input []byte
	var err error
	if input, err = ioutil.ReadFile(file); err != nil {
		return nil, fmt.Errorf("unable to read file %s: %v", file, err)
	}

	extension := filepath.Ext(file)

	resume := prepareDefaultResume()
	if extension == ".yaml" || extension == ".yml" {
		if err = yaml.Unmarshal(input, &resume); err != nil {
			return nil, fmt.Errorf("unable to parse YAML file %s: %v", file, err)
		}
		log.Infof("successfully read resume from YAML file %s", file)
	} else if extension == ".json" {
		if err = json.Unmarshal(input, &resume); err != nil {
			return nil, fmt.Errorf("unable to parse JSON file %s: %v", file, err)
		}
		log.Infof("successfully read resume from JSON file %s", file)
	}

	if err = validate(&resume); err != nil {
		return nil, fmt.Errorf("resume validation failed: %v", err)
	}

	return &resume, nil
}

func validate(resume *model.Resume) error {
	// Major version updates contain breaking changes and should not be compatible
	resumeVersions := strings.Split(resume.Version, ".")
	codeVersions := strings.Split(model.ResumeVersion, ".")
	if len(resumeVersions) != 3 || resumeVersions[0] != codeVersions[0] {
		return fmt.Errorf("version of resume definition is not compatible, version %s.x.x expected", codeVersions[0])
	}

	if resume.Theme.Background != "" && !strings.HasPrefix(resume.Theme.Background, "#") {
		return fmt.Errorf("background color must be hex code with leading hash, e.g. '#305285', but '%s' was found", resume.Theme.Background)
	}
	if resume.Theme.Accent != "" && !strings.HasPrefix(resume.Theme.Accent, "#") {
		return fmt.Errorf("accent color must be hex code with leading hash, e.g. '#305285', but '%s' was found", resume.Theme.Accent)
	}

	// Either full name or first and last name must be provided
	if resume.Personal.FullName == "" && (resume.Personal.FirstName == "" || resume.Personal.LastName == "") {
		return errors.New("either 'personal.name' or both 'personal.firstName' and 'personal.lastName' must be set")
	}

	// Only A4 oder Letter as supported paper sizes
	if strings.ToUpper(resume.Page.Size) == "A4" {
		resume.Page.Size = "A4"
	} else if strings.ToLower(resume.Page.Size) == "letter" {
		resume.Page.Size = "letter"
	} else {
		return fmt.Errorf("'settings.pageSize' can only be 'A4' or 'Letter' but '%s' was found", resume.Page.Size)
	}

	return nil
}

func prepareDefaultResume() model.Resume {
	return model.Resume{
		Version: model.ResumeVersion,
		Page: model.Page{
			Size:    "A4",
			Margins: model.PageMargins{Top: 15, Bottom: 20, Left: 5, Right: 10},
		},
	}
}
