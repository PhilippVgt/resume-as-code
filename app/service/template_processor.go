package service

import (
	"bytes"
	"fmt"
	"github.com/PhilippVgt/resume-as-code/app/model"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strings"

	"html/template"
)

const (
	TemplateResume         = "resume.html"
	TemplateCoverLetter    = "cover_letter.html"
	templateResourceFolder = "res"

	BasicTemplatePath = "templates/basic"
)

func FillTemplate(templatePath string, templateFile string, definition any) (*http.ServeMux, error) {
	log.Infof("using template from: %s", templatePath)

	htmlTemplate, err := template.New(filepath.Base(templateFile)).
		Funcs(template.FuncMap{"mul": Mul, "lines": Lines}).
		ParseFiles(path.Join(templatePath, templateFile))

	if err != nil {
		return nil, fmt.Errorf("failed to parse template %s: %v", templatePath, err)
	}

	var photo *string
	var signature *string
	switch definition.(type) {
	case *model.Definition:
		photo = &definition.(*model.Definition).Personal.Photo
		signature = &definition.(*model.Definition).Personal.Signature
	case *model.CoverLetterDefinition:
		photo = &definition.(*model.CoverLetterDefinition).Personal.Photo
		signature = &definition.(*model.CoverLetterDefinition).Personal.Signature
	}

	mux := http.NewServeMux()

	servePicture(photo, mux, "photo")
	servePicture(signature, mux, "signature")

	buf := new(bytes.Buffer)
	err = htmlTemplate.Execute(buf, definition)
	if err != nil {
		return nil, fmt.Errorf("failed to fill template with resume: %v", err)
	}
	log.Infof("successfully prepared template: %s", templatePath)

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		if _, err := io.WriteString(res, strings.TrimSpace(buf.String())); err != nil {
			log.Fatal("Failed to serve template from local test server: %v", err)
		}
		log.Debug("local test server served template html")
	})

	templateFiles := http.FileServer(http.Dir(path.Join(templatePath, templateResourceFolder)))
	mux.Handle("/res/", http.StripPrefix("/res", templateFiles))

	return mux, nil
}

func servePicture(picture *string, mux *http.ServeMux, endpoint string) {
	localPictureDir := ""
	// If the user photo is a local file, we serve its directory and reference the local server instead
	// If the photo is a URL, we leave it unchanged (chromedp will resolve it)
	if *picture != "" {
		if isUrl(*picture) {
			log.Infof("using picture from URL: %s", *picture)
		} else {
			log.Infof("using picture from local file: %s", *picture)
			localPictureDir = filepath.Dir(*picture)
			*picture = "/" + endpoint + "/" + filepath.Base(*picture)

			pictureFileServer := http.FileServer(http.Dir(localPictureDir))
			mux.Handle("/"+endpoint+"/", http.StripPrefix("/"+endpoint, pictureFileServer))
		}
	}
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func Mul(param1 float64, param2 float64) float64 {
	return param1 * param2
}

func Lines(text string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n\n", "</p><p>", -1))
}
