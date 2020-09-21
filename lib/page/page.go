package page

import (
	"fmt"
	"github.com/hachi-n/page_checker/lib/scraper"
	"github.com/hachi-n/page_checker/lib/status"
	"github.com/hachi-n/page_checker/lib/util"
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

func (p *Page) ImageUrlCheck() []*status.Status {
	imageUrls, documentStatus := scraper.GetSelectorAttributes(p.Url.String(), "img", "src")
	var statuses []*status.Status
	if documentStatus != nil {
		statuses = append(statuses, documentStatus)
	}

	imageUrls = util.UniqSlice(imageUrls)

	for _, imageUrl := range imageUrls {
		normalizedUrl, err := p.normalizeUrl(imageUrl)
		if err != nil {
			s := status.NewStatus(p.Url.String(), normalizedUrl, false, err)
			statuses = append(statuses, s)
			continue
		}
		resp, err := http.Get(normalizedUrl)
		if err != nil {
			s := status.NewStatus(p.Url.String(), normalizedUrl, false, err)
			statuses = append(statuses, s)
			continue
		}
		if !p.httpStatusCheck(resp.StatusCode) {
			err = fmt.Errorf("Image Access Error.")
			s := status.NewStatus(p.Url.String(), normalizedUrl, false, err)
			statuses = append(statuses, s)
			continue
		}

		s := status.NewStatus(p.Url.String(), normalizedUrl, true, nil)
		statuses = append(statuses, s)
	}

	return statuses
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
