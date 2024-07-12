package utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func FromCSVtoTickets() ([]tickets.Ticket, error) {
	file, err := os.Open("tickets.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var ticketsSlice []tickets.Ticket
	for _, record := range records {

		id, intErr := strconv.ParseInt(record[0], 10, 32)
		price, err := strconv.ParseFloat(record[5], 64)

		if err != nil || intErr != nil {
			return nil, err
		}

		ticket := tickets.Ticket{
			Id:          int(id),
			Name:        record[1],
			Email:       record[2],
			Destination: record[3],
			FlightTime:  record[4],
			Price:       float64(price),
		}
		ticketsSlice = append(ticketsSlice, ticket)
	}
	return ticketsSlice, nil
}
