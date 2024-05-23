package entities

type Utugas struct {
	Id        int    `json:"id"`
	Nama      string `json:"nama"`
	Nim       string `json:"nim"`
	Kelas     string `json:"kelas"`
	Email     string `json:"email"`
	NoHp      string `json:"no_hp"`
	TitleFile string `json:"title_file"`
	FileTugas string `json:"file_tugas"`
}
