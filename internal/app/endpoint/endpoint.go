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

func (e *Endpoint) Print() string {
	c, err := e.s.GetCurr()
	if err != nil {
		log.Println(err)
	}
	byn, _ := strconv.ParseFloat(c.Rates.BYN, 64)
	eur, _ := strconv.ParseFloat(c.Rates.EUR, 64)
	usd, _ := strconv.ParseFloat(c.Rates.USD, 64)
	fmt.Println("D")
	return fmt.Sprintf("1 BYN = %.3f EUR\n1 BYN = %.3f USD\n", 1/eur*byn, usd*byn)
}
