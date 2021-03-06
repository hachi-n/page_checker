package img

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/hachi-n/page_checker/lib/page"
	"github.com/hachi-n/page_checker/lib/status"
	"github.com/hachi-n/page_checker/lib/util"
	"golang.org/x/sync/semaphore"
	"io/ioutil"
	"strings"
	"sync"
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
		outputPath = "./result.json"
		fmt.Printf("write file err: %v\n Write here instead: %s\n", err, outputPath)
		ioutil.WriteFile(outputPath, jsonBytes, 0666)
	}

	fmt.Printf("Please check file: %s\n", outputPath)

	return nil
}

const (
	pbMaxWidth = 100

	threadLimit  = 30
	threadWeight = 1
)

type ImageCheckResult struct {
	statuses []*status.Status
	bar      *pb.ProgressBar
	mu       sync.Mutex
}

func NewImageCheckResult(count int) *ImageCheckResult {
	bar := pb.Simple.Start(count)
	bar.SetMaxWidth(pbMaxWidth)

	return &ImageCheckResult{
		bar: bar,
	}
}

func (i *ImageCheckResult) Store(s []*status.Status) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.statuses = append(i.statuses, s...)
	i.bar.Increment()
}

func pagesCheck(pages []*page.Page) ([]byte, error) {
	var w sync.WaitGroup
	smph := semaphore.NewWeighted(threadLimit)
	results := NewImageCheckResult(len(pages))

	for _, pg := range pages {
		w.Add(1)
		smph.Acquire(context.Background(), threadWeight)

		go func(p *page.Page) {
			sts := p.ImageUrlCheck()

			results.Store(sts)

			smph.Release(threadWeight)
			w.Done()
		}(pg)
	}

	w.Wait()
	results.bar.Finish()

	return json.MarshalIndent(
		results.statuses,
		"",
		strings.Repeat(" ", 4),
	)
}
