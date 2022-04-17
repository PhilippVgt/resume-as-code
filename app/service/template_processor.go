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
	templateIndexFile      = "index.html"
	templateResourceFolder = "res"

	BasicTemplatePath = "templates/basic"
)

func FillTemplate(templatePath string, resume *model.Resume) (*http.ServeMux, error) {
	log.Infof("using template from: %s", templatePath)

	htmlTemplate, err := template.ParseFiles(path.Join(templatePath, templateIndexFile))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template %s: %v", templatePath, err)
	}

	localPhoto := false
	localPhotoDir := ""
	// If the user photo is a local file, we serve its directory and reference the local server instead
	// If the photo is a URL, we leave it unchanged (chromedp will resolve it)
	if resume.Personal.Photo != "" {
		if isUrl(resume.Personal.Photo) {
			log.Infof("using photo from URL: %s", resume.Personal.Photo)
		} else {
			log.Infof("using photo from local file: %s", resume.Personal.Photo)
			localPhotoDir = filepath.Dir(resume.Personal.Photo)
			resume.Personal.Photo = "/photo/" + filepath.Base(resume.Personal.Photo)
			localPhoto = true
		}
	}

	buf := new(bytes.Buffer)
	err = htmlTemplate.Execute(buf, resume)
	if err != nil {
		return nil, fmt.Errorf("failed to fill template with resume: %v", err)
	}
	log.Infof("successfully prepared template: %s", templatePath)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		if _, err := io.WriteString(res, strings.TrimSpace(buf.String())); err != nil {
			log.Fatal("Failed to serve template from local test server: %v", err)
		}
		log.Debug("local test server served template html")
	})

	templateFiles := http.FileServer(http.Dir(path.Join(templatePath, templateResourceFolder)))
	mux.Handle("/res/", http.StripPrefix("/res", templateFiles))

	if localPhoto {
		photoFiles := http.FileServer(http.Dir(localPhotoDir))
		mux.Handle("/photo/", http.StripPrefix("/photo", photoFiles))
	}

	return mux, nil
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
