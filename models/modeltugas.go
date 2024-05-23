package models

import (
	"koding/config"
	"koding/entities"
)

func SendData(lol entities.Utugas) error {
	query := `
			INSERT INTO Tugas1 (nama, nim, kelas, email, nohp, titlefiletugas, filetugas)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`

	_, err := config.DB.Exec(query, lol.Nama, lol.Nim, lol.Kelas, lol.Email, lol.NoHp, lol.TitleFile, lol.FileTugas)

	return err
}
