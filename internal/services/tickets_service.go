package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

const (
	MadrugadaInicio = 0
	MadrugadaFim    = 6
	MañanaInicio    = 7
	MañanaFin       = 12
	TardeInicio     = 13
	TardeFin        = 19
	NocheInicio     = 20
	NocheFin        = 23
)

func GetTotalTickets(destination string) (int, error) {
	data := tickets.Tickets

	count := 0
	for _, ticket := range data {
		if ticket.Destination == destination {
			count++
		}
	}

	return count, nil
}

func GetCountByPeriod(timeMoment string) (int, error) {
	data := tickets.Tickets

	count := 0

	for _, ticket := range data {

		timeInStr := strings.Split(ticket.FlightTime, ":")

		intValue, err := strconv.Atoi(timeInStr[0])
		if err != nil {
			return 0, err
		}
		switch timeMoment {
		case "Madrugada":
			if intValue >= MadrugadaInicio && intValue <= MadrugadaFim {
				count++
			}
		case "Mañana":
			if intValue >= MañanaInicio && intValue <= MañanaFin {
				count++
			}
		case "Tarde":
			if intValue >= TardeInicio && intValue <= TardeFin {
				count++
			}
		case "Noche":
			if intValue >= NocheInicio && intValue <= NocheFin {
				count++
			}
		default:
			fmt.Println("Invalid time moment")
		}
	}

	return count, nil
}

func AverageDestination(destination string) (float64, error) {
	data := tickets.Tickets
	total := len(data)

	count := 0

	for _, ticket := range data {
		if ticket.Destination == destination {
			count++
		}
	}

	return float64(count) / float64(total), nil
}
