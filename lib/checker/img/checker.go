package img

import (
	"encoding/json"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/hachi-n/page_checker/lib/page"
	"github.com/hachi-n/page_checker/lib/status"
	"github.com/hachi-n/page_checker/lib/util"
	"io/ioutil"
)

func Apply(jsonPath string, outputPath string) error {
	b, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	var urls []string
	if err := json.Unmarshal(b, &urls); err != nil {
		return err
	}
	urls = util.UniqSlice(urls)

	jsonBytes, err := pagesCheck(page.NewPages(urls))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outputPath, jsonBytes, 0666)
	if err != nil {
		return err
	}

	fmt.Printf("Please check file: %s\n", outputPath)

	return nil
}

const pbMaxWidth = 100

func pagesCheck(pages []*page.Page) ([]byte, error) {
	var err error
	var statuses []*status.Status

	count := len(pages)

	bar := pb.Simple.Start(count)
	bar.SetMaxWidth(pbMaxWidth)

	for _, page := range pages {
		statuses = append(statuses, page.ImageUrlCheck()...)
		bar.Increment()
	}
	bar.Finish()

	jsonBytes, err := json.MarshalIndent(statuses, "", "    ")
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}
