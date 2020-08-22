package page

import (
	"fmt"
	"github.com/hachi-n/page_checker/lib/scraper"
	"net/http"
	u "net/url"
	"os"
)

type Page struct {
	Url *u.URL
}

func NewPage(url string) *Page {
	parsedUrl, err := u.Parse(url)
	if err != nil {
		fmt.Printf("Invalid Url. Please Check. %s \n", url)
		os.Exit(1)
	}
	return &Page{
		Url: parsedUrl,
	}
}

func NewPages(urls []string) []*Page {
	var pages []*Page

	for _, url := range urls {
		pages = append(pages, NewPage(url))
	}
	return pages
}

func (p *Page) ImageUrlCheck() []error {
	imageUrls, errors := scraper.GetSelectorAttributes(p.Url.String(), "img", "src")

	for _, imageUrl := range imageUrls {
		normalizedUrl, err := p.normalizeUrl(imageUrl)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		resp, err := http.Get(normalizedUrl)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		if !p.httpStatusCheck(resp.StatusCode) {
			errors = append(errors, fmt.Errorf(
				"Image Access Error. Url: %s, ImageUrl: %s",
				p.Url.String(), normalizedUrl),
			)
		}
	}

	return errors
}

func (p *Page) normalizeUrl(url string) (string, error) {
	parsedUrl, err := u.Parse(url)
	if err != nil {
		return "", err
	}
	// Relative Url Check
	if parsedUrl.Host == "" {
		parsedUrl.Host = p.Url.Host
		parsedUrl.Scheme = p.Url.Scheme
	}
	return parsedUrl.String(), nil
}

func (p *Page) httpStatusCheck(status int) bool {
	return status == 200
}
