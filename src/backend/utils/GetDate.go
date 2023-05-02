package utils

import "time"

func SearchDay(querry string) string {
	// Format tanggal
	const layout = "02/01/2006"

	// Mengubah querry ke type time
	tm, _ := time.Parse(layout, querry)
	day := tm.Weekday().String()
	var hari string

	if day == "Monday" {
		hari = "Senin"
	}
	if day == "Tuesday" {
		hari = "Selasa"
	}
	if day == "Wednesday" {
		hari = "Rabu"
	}
	if day == "Thursday" {
		hari = "Kamis"
	}
	if day == "Friday" {
		hari = "Jumat"
	}
	if day == "Saturday" {
		hari = "Sabtu"
	}
	if day == "Sunday" {
		hari = "Minggu"
	}

	// Mengembalikan harinya
	return hari
}
