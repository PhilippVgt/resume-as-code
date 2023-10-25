package model

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	ResumeVersion = "2.0.0"
)

type Definition struct {
	Version string `yaml:"version" json:"version" required:"true"`
	Page    Page   `yaml:"page" json:"page"`
	Theme   Theme  `yaml:"theme" json:"theme"`

	Personal Personal `yaml:"personal" json:"personal" required:"true"`

	Resume       Resume        `yaml:"resume" json:"resume" required:"true"`
	CoverLetters []CoverLetter `yaml:"coverLetters" json:"coverLetters"`
}

type CoverLetterDefinition struct {
	Definition
	CoverLetter CoverLetter `yaml:"coverLetters" json:"coverLetters"`
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

	Role    string  `yaml:"role" json:"role" required:"true"`
	Summary string  `yaml:"summary" json:"summary"`
	Contact Contact `yaml:"contact" json:"contact" required:"true"`

	Address  Address   `yaml:"address" json:"address" required:"true"`
	Profiles []Profile `yaml:"profiles" json:"profiles"`

	Photo     string `yaml:"photo" json:"photo"`
	Signature string `yaml:"signature" json:"signature"`
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

func (p *Personal) Name() string {
	if p.FirstName != "" && p.LastName != "" {
		return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
	} else {
		return p.FullName
	}
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
