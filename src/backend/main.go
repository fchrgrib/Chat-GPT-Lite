package main

import "backend/utils"

func main() {
	//fmt.Println(algorithm.BuildLast("hahahahsgsdgsg"))
	//_db, err := db.GetDatabase()
	//fmt.Println(_db, err)

	print(utils.InsertQuestAns("Tambahkan pertanyaan Messi atau Ronaldo atau Neymar atau kamu atau saya? dengan jawaban Messi").Error())
}
