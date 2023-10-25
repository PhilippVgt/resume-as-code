package model

import (
	log "github.com/sirupsen/logrus"
)

type Resume struct {
	Work      []Work      `yaml:"work" json:"work" required:"true"`
	Education []Education `yaml:"education" json:"education" required:"true"`
	Projects  []Project   `yaml:"projects" json:"projects"`

	Publications []Publication   `yaml:"publications" json:"publications"`
	Awards       []Certification `yaml:"awards" json:"awards"`
	Certificates []Certification `yaml:"certificates" json:"certificates"`

	SkillSets []SkillSet `yaml:"skillSets" json:"skillSets"`
	Languages []Language `yaml:"languages" json:"languages"`
	Interests []Interest `yaml:"interests" json:"interests"`
}

type Work struct {
	Name       string   `yaml:"name" json:"name" required:"true"`
	Address    string   `yaml:"address" json:"address"`
	Website    string   `yaml:"website" json:"website"`
	Role       string   `yaml:"role" json:"role" required:"true"`
	From       *Date    `yaml:"from" json:"from" required:"true"`
	Until      *Date    `yaml:"until" json:"until"`
	Summary    string   `yaml:"summary" json:"summary"`
	Highlights []string `yaml:"highlights" json:"highlights"`
}

type Education struct {
	Institution string   `yaml:"institution" json:"institution" required:"true"`
	Address     string   `yaml:"address" json:"address"`
	Website     string   `yaml:"website" json:"website"`
	Degree      string   `yaml:"degree" json:"degree"`
	Field       string   `yaml:"field" json:"field" required:"true"`
	From        *Date    `yaml:"from" json:"from" required:"true"`
	Until       *Date    `yaml:"until" json:"until"`
	Grade       string   `yaml:"grade" json:"grade"`
	Summary     string   `yaml:"summary" json:"summary"`
	Highlights  []string `yaml:"highlights" json:"highlights"`
}

type Project struct {
	Name       string   `yaml:"name" json:"name" required:"true"`
	Website    string   `yaml:"website" json:"website"`
	Role       string   `yaml:"role" json:"role" required:"true"`
	From       *Date    `yaml:"from" json:"from"`
	Until      *Date    `yaml:"until" json:"until"`
	Summary    string   `yaml:"summary" json:"summary"`
	Highlights []string `yaml:"highlights" json:"highlights"`
	Keywords   []string `yaml:"keywords" json:"keywords"`
}

type Certification struct {
	Certificate string `yaml:"certificate" json:"certificate" required:"true"`
	Institution string `yaml:"institution" json:"institution" required:"true"` // Certifying or awarding institution
	Date        *Date  `yaml:"date" json:"date" required:"true"`
	Website     string `yaml:"website" json:"website"`
	Description string `yaml:"description" json:"description"`
}

type Publication struct {
	Publication string `yaml:"publication" json:"publication" required:"true"`
	Institution string `yaml:"institution" json:"institution" required:"true"` // Publishing institution
	Author      string `yaml:"author" json:"author" required:"true"`
	Date        *Date  `yaml:"date" json:"date" required:"true"`
	Website     string `yaml:"website" json:"website"`
	Description string `yaml:"description" json:"description"`
}

type SkillSet struct {
	Title  string  `yaml:"title" json:"title" required:"true"`
	Skills []Skill `yaml:"skills" json:"skills" required:"true"`
}
type Skill struct {
	Skill string  `yaml:"skill" json:"skill" required:"true"`
	Level float64 `yaml:"level" json:"level"` // skill level in percent within [0, 1]
}

func (l *Skill) Levels(num int) []bool {
	return skillLevels(l.Skill, l.Level, num)
}

type Language struct {
	Language    string  `yaml:"language" json:"language" required:"true"`
	Proficiency string  `yaml:"proficiency" json:"proficiency" required:"true"` // proficiency level in text (e.g. 'Intermediate')
	Level       float64 `yaml:"level" json:"level"`                             // proficiency level in percent within [0, 1]
}

func (l *Language) Levels(num int) []bool {
	return skillLevels(l.Language, l.Level, num)
}

type Interest struct {
	Interest string   `yaml:"interest" json:"interest" required:"true"`
	Keywords []string `yaml:"keywords" json:"keywords"`
}

func skillLevels(skill string, level float64, num int) []bool {
	levels := make([]bool, num)

	increment := 1.0 / float64(num)
	for i := 0; i < num; i++ {
		levels[i] = level > float64(i)*increment
	}

	log.Debugf("Computed skill levels for %s: %v", skill, levels)
	return levels
}
