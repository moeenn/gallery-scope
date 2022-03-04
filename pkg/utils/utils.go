package utils

import (
	"crypto/rand"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

/**
 *  zero pad count numbers
 */
func ZeroPad(n int, digits int) string {
	n_str := strconv.Itoa(n)
	n_zeros := digits - len(n_str)
	zeros := ""

	for i := 0; i < n_zeros; i++ {
		zeros += "0"
	}

	return zeros + n_str
}

/**
 *  generate a crypto-random string
 */
func RandomStr(length int) (string, error) {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes), nil
}

/**
 *  in the base (URL) the count placeholder is '{n}'
 *  this function replaces the placeholder with the actual count
 */
func InsertCount(url, count string) string {
	return strings.Replace(url, "{n}", count, -1)
}

/**
 *  conver raw string into URL-safe string
 */
func URLEncode(s string) string {
	return url.QueryEscape(s)
}

/**
 *  extract filename from URL
 */
func ExtractURLFilename(url string) string {
	pieces := strings.Split(url, "/")
	return pieces[len(pieces)-1]
}

/**
 *  download and save a file to disk
 */
func Download(url, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
