package main

import (
	"fmt"

	service "github.com/bootcamp-go/desafio-go-bases/internal/services"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/bootcamp-go/desafio-go-bases/internal/utils"
)

func main() {
	data, err := utils.FromCSVtoTickets()
	if err == nil {
		tickets.Tickets = data
	}

	fmt.Println(service.GetCountByPeriod("Tarde"))
	fmt.Println(service.AverageDestination("Colombia"))
	//fmt.Println(data, err)
}
