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

func NewShortener() Shortener {
	shortener := Shortener{}
	shortener.urlStore = make(map[string]string)
	return shortener
}

func (s *Shortener) Encode(url string) string {
	hash := sha1.New()
	hash.Write([]byte(url))
	hexHash := hex.EncodeToString(hash.Sum(nil))
	shortUrl := hexHash[:10]
	s.urlStore[shortUrl] = url
	return "https://short-url.com/" + shortUrl
}

func (s *Shortener) Decode(url string) (string, error) {
	index := strings.LastIndex(url, "/")
	param := url[index+1:]
	if _, ok := s.urlStore[param]; !ok {
		return "", errors.New("url not found")
	}
	return s.urlStore[param], nil
}
