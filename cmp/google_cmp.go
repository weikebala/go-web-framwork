package cmp

import (
	"github.com/google/go-cmp/cmp"
)

func Google_cmp(a interface{}, b interface{}) bool {
	if diff := cmp.Diff(a, b); diff != "" {
		return false
	}
	return true
}