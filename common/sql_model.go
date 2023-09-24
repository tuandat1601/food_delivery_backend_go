package common
import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_ad" gorm:"column:updated_at"`
}
type OperatingHours struct{
	DayOfWeek   string    `json:"day_of_week" gorm:"column:day_of_week"`
	OpeningTime string   `json:"opening_time" gorm:"column:opening_time"`
	ClosingTime string   `json:"closing_time" gorm:"column:closing_time"`
}
