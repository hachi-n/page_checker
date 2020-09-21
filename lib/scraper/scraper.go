package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hachi-n/page_checker/lib/status"
)

func GetSelectorAttributes(url, selector, attributeKey string) ([]string, *status.Status) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, status.NewStatus(url, "Undifined", false, err)
	}

	var attributes []string
	selection := doc.Find(selector)
	selection.Each(func(index int, s *goquery.Selection) {
		attr, exists := s.Attr(attributeKey)
		_ = exists

		// src attribute does not exist.

		//if !exists {
		//	message := fmt.Errorf(
		//		"Could not find attribute. url:%s, selector:%s,  attributeKey: %s.",
		//		url,
		//		selector,
		//		attributeKey,
		//	)
		//	errors = append(errors, message)
		//}
		attributes = append(attributes, attr)
	})
	return attributes, nil
}
