package status

import "fmt"

type Status struct {
	RefererUrl   string
	ImageUrl     string
	Judge        bool
	ErrorMessage string
}

func NewStatus(referer, imageUrl string, judge bool, err error) *Status {
	message := ""
	if err != nil {
		message = fmt.Sprintf("%v", err)
	}

	return &Status{
		RefererUrl:   referer,
		ImageUrl:     imageUrl,
		Judge:        judge,
		ErrorMessage: message,
	}
}

