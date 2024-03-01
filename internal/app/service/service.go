package service

import (
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

const apiKey = "2fb43ddacc174d0d96f0e39048699a1b"

type Rates struct {
	EUR string `json:"EUR"`
	USD string `json:"USD"`
	BYN string `json:"BYN"`
}

type Curr struct {
	Date  string `json:"date"`
	Base  string `json:"base"`
	Rates Rates  `json:"rates"`
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetCurr() (Curr, error) {
	resp, err := http.Get("https://api.currencyfreaks.com/v2.0/rates/latest?apikey=" + apiKey + "&symbols=EUR,USD,BYN")
	if err != nil {
		return Curr{}, err
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)

	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		return Curr{}, err
	}

	var c Curr
	err = json.Unmarshal(buf, &c)
	if err != nil {
		return Curr{}, err
	}

	return c, nil
}
