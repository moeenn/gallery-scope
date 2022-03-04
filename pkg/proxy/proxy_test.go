package proxy

import (
	"testing"
)

func TestNewProxyURL(t *testing.T) {
	input := "http://www.google.com"
	got := NewProxyURL(input)

	if got.Full == "" || got.Thumb == "" {
		t.Errorf("[failed] Produced ProxyURL object is invalid %+v", got)
	}
}
