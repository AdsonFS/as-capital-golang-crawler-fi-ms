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
	redisRepository := redis_repository.NewRedisRepository()
	ctx := context.Background()
	dataRedis := redisRepository.Get(ctx, name)
	if len(dataRedis) == 0 {
		fmt.Println("Buscando no site")
		data, err := newData(fmt.Sprintf("https://fiis.com.br/%s/", name))
		if err == nil {
			redisRepository.Save(ctx, name, data.ToMap())
		}
		return data, err
	}
	fmt.Println("Pegando cache")
	return FromMap(dataRedis), nil
}

func newData(url string) (Data, error) {
	data := Data{}
	resp, err := http.Get(url)
	if err != nil {
		return Data{}, err
	}
	defer resp.Body.Close()
	htmlString, err := io.ReadAll(resp.Body)

	if err != nil {
		return Data{}, err
	}
	data.html = string(htmlString)

	err = data.extractPrice()
	if err != nil {
		return Data{}, err
	}
	return data, nil
}

func (data *Data) extractPrice() error {
	re := regexp.MustCompile(`<div>\s*<span class="currency">R\$</span>\s*<span class="value">([0-9,]+)</span>`)
	matches := re.FindAllStringSubmatch(data.html, -1)

	if len(matches) < 3 {
		return fmt.Errorf("não foi possível encontrar o preço no HTML")
	}

	currentQuoteStr := strings.ReplaceAll(matches[0][1], ",", ".")
	minQuoteStr := strings.ReplaceAll(matches[1][1], ",", ".")
	maxQuoteStr := strings.ReplaceAll(matches[2][1], ",", ".")

	var err1, err2, err3 error
	data.Quote.Current, err1 = strconv.ParseFloat(currentQuoteStr, 64)
	data.Quote.Min, err2 = strconv.ParseFloat(minQuoteStr, 64)
	data.Quote.Max, err3 = strconv.ParseFloat(maxQuoteStr, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return fmt.Errorf("erro ao converter preço para float64")
	}
	return nil
}
