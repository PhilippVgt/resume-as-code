package service

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"time"
)

func GeneratePdf(template *http.ServeMux, outputPath string, fileName string) error {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(template)
	defer ts.Close()

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(ts.URL, &buf)); err != nil {
		return fmt.Errorf("error printing from headless Chrome: %v", err)
	}

	outputFile := path.Join(outputPath, fileName+".pdf")
	if err := ioutil.WriteFile(outputFile, buf, 0644); err != nil {
		return fmt.Errorf("failed to write to file %s: %v", outputFile, err)
	}

	log.Infof("successfully exported PDF to %s", outputFile)
	return nil
}

// print a specific pdf page.
func printToPDF(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			time.Sleep(time.Second)
			buf, _, err := page.
				PrintToPDF().
				WithPreferCSSPageSize(true).
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
