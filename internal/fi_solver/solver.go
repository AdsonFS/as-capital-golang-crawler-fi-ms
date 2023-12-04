package fisolver

import (
	"as-capital-crawler-fi-ms/internal/redis_repository"
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func GetData(name string) (Data, error) {
	ctx := context.Background()
	dataRedis := redis_repository.Get(ctx, name)
	if len(dataRedis) == 0 {
		fmt.Println("Buscando no site")
		data, err := newData(fmt.Sprintf("https://fiis.com.br/%s/", name))
		if err == nil {
			redis_repository.Save(ctx, name, data.ToMap())
		}
		return data, err
	}
	fmt.Println("Pegando cache")
	return FromMap(dataRedis), nil
}

func newData(url string) (Data, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Data{}, err
	}
	defer resp.Body.Close()
	htmlString, err := io.ReadAll(resp.Body)

	if err != nil {
		return Data{}, err
	}

	quotes, err := extractPrice(string(htmlString))
	if err != nil {
		return Data{}, err
	}
	return Data{
		Quote: quotes,
	}, nil
}

func extractPrice(html string) (Quote, error) {
	re := regexp.MustCompile(`<div>\s*<span class="currency">R\$</span>\s*<span class="value">([0-9,]+)</span>`)
	matches := re.FindAllStringSubmatch(html, -1)

	if len(matches) < 3 {
		return Quote{}, fmt.Errorf("não foi possível encontrar o preço no HTML")
	}

	currentQuoteStr := strings.ReplaceAll(matches[0][1], ",", ".")
	minQuoteStr := strings.ReplaceAll(matches[1][1], ",", ".")
	maxQuoteStr := strings.ReplaceAll(matches[2][1], ",", ".")

	currentQuote, err1 := strconv.ParseFloat(currentQuoteStr, 64)
	minQuote, err2 := strconv.ParseFloat(minQuoteStr, 64)
	maxQuote, err3 := strconv.ParseFloat(maxQuoteStr, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return Quote{}, fmt.Errorf("erro ao converter preço para float64")
	}

	return Quote{Current: currentQuote, Min: minQuote, Max: maxQuote}, nil
}
