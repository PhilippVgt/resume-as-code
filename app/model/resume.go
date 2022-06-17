package model

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	ResumeVersion = "1.0.0"
)

type Resume struct {
	Version string `yaml:"version" json:"version" required:"true"`
	Page    Page   `yaml:"page" json:"page"`
	Theme   Theme  `yaml:"theme" json:"theme"`

	Personal Personal `yaml:"personal" json:"personal" required:"true"`

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

type Page struct {
	Size    string      `yaml:"size" json:"size"`       // either 'A4' oder 'Letter', defaults to 'A4'
	Margins PageMargins `yaml:"margins" json:"margins"` // margins all in millimeters
}
type PageMargins struct {
	Top    float64 `yaml:"top" json:"top"`
	Bottom float64 `yaml:"bottom" json:"bottom"`
	Left   float64 `yaml:"left" json:"left"`
	Right  float64 `yaml:"right" json:"right"`
}

type Theme struct {
	Background string `yaml:"background" json:"background"` // theme background color in hex code (used if applicable)
	Accent     string `yaml:"accent" json:"accent"`         // theme accent color in hex code (used if applicable)
}

type Personal struct {
	FullName  string `yaml:"name" json:"name"`
	FirstName string `yaml:"firstName" json:"firstName"`
	LastName  string `yaml:"lastName" json:"lastName"`

	Role    string `yaml:"role" json:"role" required:"true"`
	Photo   string `yaml:"photo" json:"photo"`
	Summary string `yaml:"summary" json:"summary"`

	Contact  Contact   `yaml:"contact" json:"contact" required:"true"`
	Address  Address   `yaml:"address" json:"address" required:"true"`
	Profiles []Profile `yaml:"profiles" json:"profiles"`
}
type Contact struct {
	Phone   string `yaml:"phone" json:"phone" required:"true"`
	Email   string `yaml:"email" json:"email" required:"true"`
	Website string `yaml:"website" json:"website"`
}
type Address struct {
	Street     string `yaml:"street" json:"street"`
	City       string `yaml:"city" json:"city" required:"true"`
	PostalCode string `yaml:"postalCode" json:"postalCode"`
	Country    string `yaml:"country" json:"country"`
}

func (c *Contact) WebsiteUrl() string {
	return websiteUrl(c.Website)
}
func (c *Contact) WebsiteTitle() string {
	return websiteTitle(c.Website)
}

type Profile struct {
	Platform string `yaml:"platform" json:"platform" required:"true"`
	User     string `yaml:"user" json:"user" required:"true"`
	Website  string `yaml:"website" json:"website" required:"true"`
}

func (p *Profile) WebsiteUrl() string {
	return websiteUrl(p.Website)
}
func (p *Profile) WebsiteTitle() string {
	return websiteTitle(p.Website)
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
	Degree      string   `yaml:"degree" json:"degree" required:"true"`
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

func websiteUrl(website string) string {
	if !strings.HasPrefix(website, "https://") {
		website = "https://" + website
	}

	log.Debugf("Using website URL %s", website)
	return website
}
func websiteTitle(website string) string {
	website = strings.TrimPrefix(website, "https://")
	website = strings.TrimPrefix(website, "www.")
	for strings.HasSuffix(website, "/") {
		website = website[:len(website)-1]
	}

	log.Debugf("Using website title %s", website)
	return website
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
