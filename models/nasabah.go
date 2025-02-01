package models

type Nasabah struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Nik       string     `gorm:"type:varchar(20);unique;not null" json:"nik"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	NoHp      string     `gorm:"type:varchar(15);unique;not null" json:"no_hp"`
	CreatedAt string     `gorm:"type:timestamptz;default:null" json:"created_at"`
	UpdatedAt string     `gorm:"type:timestamptz;default:null" json:"updated_at"`
	Tabungans []Tabungan `gorm:"foreignKey:IdNasabah" json:"tabungans"`
}
