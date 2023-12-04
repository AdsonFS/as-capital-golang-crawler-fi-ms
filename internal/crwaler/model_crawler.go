package crwaler

import (
	"io"
	"net/http"
	"regexp"
	"sort"
)

type Crwaler struct {
	fiisUrl []string
}

func NewCrawler() (Crwaler, error) {
	crawler := Crwaler{[]string{}}
	url := "https://fiis.com.br/lista-de-fundos-imobiliarios/"
	resp, err := http.Get(url)
	if err != nil {
		return crawler, err
	}
	defer resp.Body.Close()
	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return crawler, err
	}

	re, _ := regexp.Compile(`https:\/\/fiis.com.br\/\w{4}11\/`)
	matches := re.FindAllString(string(html), -1)
	sort.Strings(matches)

	var lastMatch = ""
	for _, match := range matches {
		if match != lastMatch {
			crawler.fiisUrl = append(crawler.fiisUrl, match)
		}
		lastMatch = match
	}
	return crawler, nil
}
