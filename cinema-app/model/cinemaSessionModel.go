package cinemaSessionModel

import (
	"time"
)

type CinemaSessionModel struct {
	FilmName      string `gorm:"column:film_name"`
	SessionNumber uint `gorm:"column:session_number"`
	SessionTime   time.Time `gorm:"column:session_time"`
	TotalSeats    uint `gorm:"column:total_seats"`
}

func (CinemaSessionModel) TableName() string {
	return "cinema_session"
}
