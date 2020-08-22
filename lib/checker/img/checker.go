package img

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/page_checker/lib/page"
	"io/ioutil"
)

func Apply(jsonPath string) error {
	b, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	var urls []string
	if err := json.Unmarshal(b, &urls); err != nil {
		return err
	}

	return pagesCheck(page.NewPages(urls))
}

func pagesCheck(pages []*page.Page) error {
	var err error
	for _, page := range pages {
		errors := page.ImageUrlCheck()
		if errors != nil {
			err = fmt.Errorf("%v errors: %v\n", err, errors)
			fmt.Printf("url: %s, status: ng\n", page.Url.String())
			continue
		}
		fmt.Printf("url: %s, status: ok\n", page.Url.String())
	}

	return err
}
