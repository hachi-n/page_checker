package scraper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func GetSelectorAttributes(url, selector, attributeKey string) ([]string, []error) {
	var errors []error
	doc, err := goquery.NewDocument(url)
	if err != nil {
		errors = append(errors, err)
		return nil, errors
	}

	var attributes []string
	selection := doc.Find(selector)
	selection.Each(func(index int, s *goquery.Selection) {
		attr, exists := s.Attr(attributeKey)
		if !exists {
			message := fmt.Errorf(
				"Could not find attribute. url:%s, selector:%s,  attributeKey: %s.",
				url,
				selector,
				attributeKey,
			)

			errors = append(errors, message)
		}
		attributes = append(attributes, attr)
	})
	return attributes, errors
}
