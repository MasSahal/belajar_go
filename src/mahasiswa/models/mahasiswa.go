package models

import "time"

type Mahasiswa struct {
	ID        int       `json:"id" gorm:"primary_key"`
	NIM       int       `json:"nim"`
	Nama      string    `json:"nama"`
	Prodi     string    `json:"prodi"`
	Semester  int       `json:"semester"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
