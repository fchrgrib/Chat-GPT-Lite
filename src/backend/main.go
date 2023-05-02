package main

import (
	"backend/utils"
	"fmt"
)

func main() {
	//fmt.Println(algorithm.BuildLast("hahahahsgsdgsg"))
	//_db, err := db.GetDatabase()
	//fmt.Println(_db, err)

	//text := "I need some help with my computer and some advice on buying a new one."
	//pattern := `I need (.+) with my computer and (.+) on buying a new one\.`
	//
	//re := regexp.MustCompile(pattern)
	//matches := re.FindStringSubmatch(text)
	//
	//if len(matches) > 0 {
	//	fmt.Printf("What would it mean to you if you got %s with your computer and %s on buying a new one?\n", matches[1], matches[2])
	//}

	pattern := "20/21/2001*2^6"

	fmt.Println(utils.CheckCalculate(pattern))

	//text := "The year 2022 is coming soon, but it's not as far away as 1999 or 2000."
	//
	//re := regexp.MustCompile(`\d{4}`)
	//matches := re.FindAllString(text, -1)
	//
	//fmt.Println(matches)
}
