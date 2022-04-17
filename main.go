package main

import (
	"github.com/PhilippVgt/resume-as-code/app/service"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	// set up logging
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	log.Warnf("\n\n")
	log.Warnf("Resume-as-Code")
	log.Warnf("==========================================")
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("usage: resume-as-code resume-file [template-name]")
	}
	inputFile := os.Args[1]

	var templatePath string
	if len(os.Args) > 2 {
		templatePath = os.Args[2]
	} else {
		templatePath = service.BasicTemplatePath
	}

	resume, err := service.ReadResumeFromFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	template, err := service.FillTemplate(templatePath, resume)
	if err != nil {
		log.Fatal(err)
	}

	inputFileName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
	err = service.GeneratePdf(template, filepath.Dir(inputFile), inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Warnf("==========================================")
	log.Warnf("Done")
	log.Warnf("\n\n")
}
