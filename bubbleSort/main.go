package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var genArray [6]int //объявляю и наполняю элементами массив
	seed := time.Now().UnixNano()
	rand := rand.New(rand.NewSource(seed))
	for i := 0; i < len(genArray); i++ {
		genArray[i] = rand.Intn(100)
	}

	fmt.Println("Ваш массив:", genArray)

	for i := 0; i < len(genArray)-1; i++ { //сортирую по убыванию
		for j := 0; j < len(genArray)-i-1; j++ {
			if genArray[j] < genArray[j+1] {
				genArray[j], genArray[j+1] = genArray[j+1], genArray[j]
			}
		}
	}

	fmt.Println("Ваш массив после сортировки пузырьком по убыванию:", genArray)
}
