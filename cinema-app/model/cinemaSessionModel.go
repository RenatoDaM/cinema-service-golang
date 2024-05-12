package cinemaSessionModel

import (
	"gorm.io/gorm"
	"time"
)

type CinemaSession struct {
	gorm.Model
	FilmName string
	SessionNumber uint
	SessionTime time.Time
	TotalSeats uint
}