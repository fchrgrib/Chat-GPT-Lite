package main

import (
	"backend/db"
	"fmt"
)

func main() {
	//fmt.Println(algorithm.BuildLast("hahahahsgsdgsg"))
	_db, err := db.GetDatabase()
	fmt.Println(_db, err)
}
