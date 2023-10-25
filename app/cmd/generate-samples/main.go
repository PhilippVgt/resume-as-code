package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PhilippVgt/resume-as-code/app/model"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v3"
)

func main() {
	resume := model.Definition{
		Version: model.ResumeVersion,
		Personal: model.Personal{
			Profiles: make([]model.Profile, 1),
		},
		Resume: model.Resume{
			Work: []model.Work{
				{From: &model.Date{Time: time.Now()}},
			},
			Education: []model.Education{
				{From: &model.Date{Time: time.Now()}},
			},
			Projects: []model.Project{
				{From: &model.Date{Time: time.Now()}},
			},
			Publications: []model.Publication{
				{Date: &model.Date{Time: time.Now()}},
			},
			Awards: []model.Certification{
				{Date: &model.Date{Time: time.Now()}},
			},
			Certificates: []model.Certification{
				{Date: &model.Date{Time: time.Now()}},
			},
			SkillSets: []model.SkillSet{
				{
					Skills: make([]model.Skill, 1),
				},
			},
			Languages: make([]model.Language, 1),
			Interests: make([]model.Interest, 1),
		},
		CoverLetters: []model.CoverLetter{
			{Date: &model.Date{Time: time.Now()}},
		},
	}

	var err error
	var yamlData bytes.Buffer

	yamlEncoder := yaml.NewEncoder(&yamlData)
	yamlEncoder.SetIndent(2)
	if err = yamlEncoder.Encode(&resume); err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("schema/resume.yaml", yamlData.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}

	var jsonData []byte
	if jsonData, err = json.MarshalIndent(&resume, "", "  "); err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("schema/resume.json", jsonData, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("samples written")
}
