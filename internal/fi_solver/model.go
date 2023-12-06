package fisolver

import "strconv"

type Data struct {
	Quote Quote
	url   string
	html  string
}

type Quote struct {
	Current float64
	Min     float64
	Max     float64
}

func (data *Data) ToMap() map[string]float64 {
	return map[string]float64{
		"minQuote":     data.Quote.Min,
		"maxQuote":     data.Quote.Max,
		"currentQuote": data.Quote.Current,
	}
}

func FromMap(values map[string]string) Data {
	data := Data{Quote: Quote{}}
	data.Quote.Current, _ = strconv.ParseFloat(values["currentQuote"], 64)
	data.Quote.Min, _ = strconv.ParseFloat(values["minQuote"], 64)
	data.Quote.Max, _ = strconv.ParseFloat(values["maxQuote"], 64)
	return data
}
