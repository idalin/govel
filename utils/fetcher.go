package utils

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	// "log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
)

const UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36"

func GetPage(url, ua string) (io.Reader, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.DebugF("GetPage error:%s\n", err.Error())
		return nil, err

	}

	if ua == "" {
		ua = UserAgent
	}
	req.Header.Set("User-Agent", ua)

	resp, err := client.Do(req)
	if err != nil {
		log.DebugF("GetPage error:%s\n", err.Error())
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error code:%d is not 200.", resp.StatusCode))
	}

	// defer resp.Body.Close()
	// return resp.Body, nil
	return DecodeHTMLBody(resp.Body)
}

func PostPage(url, key string) (io.Reader, error) {
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(key))
	if err != nil {
		log.DebugF("PostPage error:%s\n", err.Error())
		return nil, err
	}
	return DecodeHTMLBody(res.Body)
}

func detectContentCharset(body io.Reader) (string, io.Reader) {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		_, name, _ := charset.DetermineEncoding(data, "")
		return name, r
	}
	return "utf-8", r
}

// DecodeHTMLBody returns an decoding reader of the html Body for the specified `charset`
// If `charset` is empty, DecodeHTMLBody tries to guess the encoding from the content
func DecodeHTMLBody(body io.Reader) (io.Reader, error) {
	// if charset == "" {
	charset, r := detectContentCharset(body)
	// }
	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}
	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(r)
	} else {
		body = r
	}
	return body, nil
}
