package cinemaService
import (
	cinemaSessionModel "com.cinema/cinema-app/model"
	GORM "com.cinema/cinema-app/common"
	"time"
)

func CreateCinemaSession() {
	GORM.GetORMConnection().Create(&cinemaSessionModel.CinemaSession{
        FilmName: "Nome do Filme",
        SessionNumber: 1,
        SessionTime: time.Now(),
        TotalSeats: 100,
    })
}