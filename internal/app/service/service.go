package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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

func (s *Service) GetCurrBaseUSD() (Curr, error) {
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

func (s *Service) GetCurrPair(base, symbol string) (string, error) {
	resp, err := http.Get("https://api.currencyfreaks.com/v2.0/rates/latest?apikey=" + apiKey + "&symbols=EUR,USD,BYN")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)

	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(buf, &data)
	rates, ok := data["rates"].(map[string]interface{})
	if !ok {
		log.Fatal("rates erreor")
	}

	bas, ok := rates[base].(string)
	if !ok {
		log.Fatal("rates erreor")
	}
	symb, ok := rates[symbol].(string)
	if !ok {
		log.Fatal("rates erreor")
	}
	basee, _ := strconv.ParseFloat(bas, 64)
	symbee, _ := strconv.ParseFloat(symb, 64)
	res := strconv.FormatFloat(symbee/basee, 'g', -1, 64)
	fmt.Println("1 " + symbol + " = " + res + " " + base)

	if err != nil {
		return "", err
	}

	return res, nil
}
