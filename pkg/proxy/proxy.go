package proxy

import (
	"gallery-scope/pkg/utils"
)

const BASE string = "https://external-content.duckduckgo.com/iu/?u="

type ProxyURL struct {
	Full  string `json:"full"`
	Thumb string `json:"thumb"`
}

func NewProxyURL(url string) ProxyURL {
	fullURL := BASE + utils.URLEncode(url)
	return ProxyURL{fullURL, fullURL}
}
