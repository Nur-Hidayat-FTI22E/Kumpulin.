package controllers

import (
	"html/template"
	"io/ioutil"
	"koding/entities"
	"koding/models"

	// "log"
	"net/http"
)

func CreateTugas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
	} else if r.Method == "POST" {
		var tugas entities.Utugas

		// Mengambil nilai form
		tugas.Nama = r.FormValue("nama")
		tugas.Nim = r.FormValue("nim")
		tugas.Kelas = r.FormValue("kelas")
		tugas.Email = r.FormValue("email")
		tugas.NoHp = r.FormValue("no_hp")

		// Mengambil file yang diunggah
		file, header, err := r.FormFile("file_tugas")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Membaca isi file
		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Mengisi struktur Utugas dengan judul dan isi file
		tugas.TitleFile = header.Filename
		tugas.FileTugas = string(fileContent)

		// Mengirim data ke database
		err = models.SendData(tugas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect kembali ke halaman utama setelah berhasil mengirim data
		http.Redirect(w, r, "/?success=true", http.StatusSeeOther)
	}
}
