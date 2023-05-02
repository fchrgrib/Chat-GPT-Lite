package algorithm

import "time"

func SearchDay(querry string) string {
	// Format tanggal
	const layout = "02/01/2006"

	// Mengubah querry ke type time
	tm, _ := time.Parse(layout, querry)

	// Mengembalikan harinya
	return tm.Weekday().String()
}
