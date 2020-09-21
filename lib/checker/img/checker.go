package img

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/page_checker/lib/page"
	"github.com/hachi-n/page_checker/lib/status"
	"github.com/hachi-n/page_checker/lib/util"
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
	urls = util.UniqSlice(urls)

	return pagesCheck(page.NewPages(urls))
}

func pagesCheck(pages []*page.Page) error {
	var err error
	var statuses []*status.Status

	for _, page := range pages {
		statuses = append(statuses, page.ImageUrlCheck()...)
	}

	jsonBytes, err := json.MarshalIndent(statuses, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonBytes))

	return nil
}
