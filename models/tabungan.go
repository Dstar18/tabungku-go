package models

type Tabungan struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	IdNasabah  int    `gorm:"type:bigint;not null" json:"id_nasabah"` //foregn key ID Nasabah
	NoRekening string `gorm:"type:varchar(50);unique;not null" json:"no_rekening"`
	Saldo      int    `gorm:"type:bigint;not null" json:"saldo"`
}
