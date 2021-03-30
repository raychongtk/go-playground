package shortenurl

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"strings"
)

// demo for working with url shortener
type Shortener struct {
	urlStore map[string]string
}

func newShortener() Shortener {
	shortener := Shortener{}
	shortener.urlStore = make(map[string]string)
	return shortener
}

func (s *Shortener) encode(url string) string {
	hash := sha1.New()
	hash.Write([]byte(url))
	hexHash := hex.EncodeToString(hash.Sum(nil))
	shortUrl := hexHash[:10]
	s.urlStore[shortUrl] = url
	return "https://short-url.com/" + shortUrl
}

func (s *Shortener) decode(url string) (string, error) {
	index := strings.LastIndex(url, "/")
	param := url[index+1:]
	if _, ok := s.urlStore[param]; !ok {
		return "", errors.New("url not found")
	}
	return s.urlStore[param], nil
}

var shortener = newShortener()

func Encode(url string) string {
	return shortener.encode(url)
}

func Decode(url string) string {
	shortUrl, err := shortener.decode(url)

	if err != nil {
		panic("url not found")
	}

	return shortUrl
}
