package main

import (
	"github.com/PhilippVgt/resume-as-code/app/model"
	"github.com/PhilippVgt/resume-as-code/app/service"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	log.SetLevel(log.InfoLevel)
	if os.Getenv("LOG_LEVEL") != "" {
		if level, err := log.ParseLevel(os.Getenv("LOG_LEVEL")); err == nil {
			log.SetLevel(level)
		} else {
			log.Warnf("Failed to parse log level from environment variable LOG_LEVEL, default to %s", log.GetLevel().String())
		}
	}

	log.SetOutput(os.Stdout)
}

func main() {
	log.Warn("==========================================")
	log.Warn("Resume-as-Code")
	log.Warn("==========================================")

	if len(os.Args) < 2 {
		log.Fatal("usage: resume-as-code resume-file [template-name]")
	}
	inputFile := os.Args[1]
	inputDir := filepath.Dir(inputFile)
	inputFileName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))

	var templatePath string
	if len(os.Args) > 2 {
		templatePath = os.Args[2]
	} else {
		templatePath = service.BasicTemplatePath
	}

	definition, err := service.ReadDefinitionFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, coverLetter := range definition.CoverLetters {
		template, err := service.FillTemplate(templatePath, service.TemplateCoverLetter, &model.CoverLetterDefinition{
			Definition:  definition,
			CoverLetter: coverLetter,
		})
		if err != nil {
			log.Fatal(err)
		}

		err = service.GeneratePdf(template, inputDir, inputFileName+"_cover_letter_"+strings.ToLower(coverLetter.Name))
		if err != nil {
			log.Fatal(err)
		}
	}

	template, err := service.FillTemplate(templatePath, service.TemplateResume, &definition)
	if err != nil {
		log.Fatal(err)
	}

	err = service.GeneratePdf(template, inputDir, inputFileName+"_resume")
	if err != nil {
		log.Fatal(err)
	}

	log.Warn("==========================================")
	log.Warn("Done")
	log.Warn("==========================================\n\n")
}
