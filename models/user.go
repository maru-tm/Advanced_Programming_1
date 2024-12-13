package models

import "gorm.io/gorm"

type User struct {
    ID        uint           `gorm:"primaryKey;autoIncrement;type:serial" json:"id"`
    Name      string         `gorm:"type:varchar(100);not null" json:"name"`
    Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
    CreatedAt string         `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt string         `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
