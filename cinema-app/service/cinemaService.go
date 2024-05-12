package cinemaService

import (
	"time"

	GORM "com.cinema/cinema-app/common"
	cinemaSessionModel "com.cinema/cinema-app/model"
)

func CreateCinemaSession() {
	cinemaSession := &cinemaSessionModel.CinemaSessionModel{
		FilmName:      "Nome do Filme",
		SessionNumber: 1,
		SessionTime:   time.Now(),
		TotalSeats:    100,
	}

	GORM.GetORMConnection().Create(cinemaSession)
}
