package domain

type Pizza struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(255); not null"`
	Description string  `gorm:"type:text; not null"`
	Price       float64 `gorm:"type:decimal; not null"`
	ImageURL    string  `gorm:"type:varchar(255)"`
}
