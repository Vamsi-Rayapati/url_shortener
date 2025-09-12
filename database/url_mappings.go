package database

import (
	"time"
)

// shorturl
// longurl
//

type UrlMapping struct {
	ID        uint      `gorm:"autoIncrement;primaryKey"`
	ShortKey  string    `gorm:"type:varchar(20);uniqueIndex"`
	LongURL   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	ExpiresAt time.Time `gorm:"type:datetime;not null"`
	Count     int       `gorm:"type:int;not null;default:0"`
}
