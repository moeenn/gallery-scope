package record

import (
	"errors"
	"gallery-scope/pkg/utils"
)

type Record struct {
	Base  string `json:"base" binding:"required"`
	Start int    `json:"start"`
	End   int    `json:"end" binding:"required"`
	Zero  int    `json:"zero"`
}

/**
 *  validate and set record default values
 */
func (r *Record) normalize() error {
	if r.Base == "" || r.End == 0 {
		return errors.New("Mandatory fields missing")
	}

	if r.Start == 0 {
		r.Start = 1
	}

	return nil
}

/**
 *  process the input record
 */
func (r *Record) Process() ([]string, error) {
	urls := []string{}
	r.normalize()

	for i := r.Start; i <= r.End; i++ {
		count := utils.ZeroPad(i, r.Zero)
		url := utils.InsertCount(r.Base, count)
		urls = append(urls, url)
	}

	return urls, nil
}
