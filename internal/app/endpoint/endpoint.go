package endpoint

import (
	"fmt"
	"log"
	"strconv"
	"test-task/internal/app/service"
)

type Endpoint struct {
	s service.Service
}

func New(s service.Service) *Endpoint {
	return &Endpoint{s: s}
}

func (e *Endpoint) Status() string {
	c, err := e.s.GetCurrBaseUSD()
	if err != nil {
		log.Println(err)
	}
	byn, _ := strconv.ParseFloat(c.Rates.BYN, 64)
	eur, _ := strconv.ParseFloat(c.Rates.EUR, 64)
	usd, _ := strconv.ParseFloat(c.Rates.USD, 64)
	return fmt.Sprintf("1 EUR = %.3f BYN\n1 USD = %.3f BYN\n", 1/eur*byn, usd*byn)
}

func (e *Endpoint) Pair(base, symbol string) string {
	s, err := e.s.GetCurrPair(base, symbol)
	if err != nil {
		log.Println(err)
	}
	sb, _ := strconv.ParseFloat(s, 64)
	return fmt.Sprintf("1 %s = %.3f %s", base, sb, symbol)

}
