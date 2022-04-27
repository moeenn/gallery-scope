package actions

import (
	"fmt"
	"gallery-scope/pkg/proxy"
	"gallery-scope/pkg/record"
	"gallery-scope/pkg/utils"
)

func GeneratePreviewURLs(record record.Record) ([]proxy.ProxyURL, error) {
	var result []proxy.ProxyURL

	urls, err := record.Process()
	if err != nil {
		return result, err
	}

	for _, url := range urls {
		proxyUrl := proxy.NewProxyURL(url)
		result = append(result, proxyUrl)
	}

	return result, nil
}

func DownloadGallery(r record.Record) error {
	slug, err := utils.RandomStr(5)
	if err != nil {
		return err
	}

	urls, err := GeneratePreviewURLs(r)
	if err != nil {
		return err
	}

	for i, url := range urls {
		go func(i int, slug string, url proxy.ProxyURL, r record.Record) {
			path := fmt.Sprintf("./storage/%s_%s.jpg", slug, utils.ZeroPad(i+1, r.Zero))
			utils.Download(url.Full, path)
			fmt.Printf("[Complete] %v\n", url.Full)
		}(i, slug, url, r)
	}

	return nil
}
