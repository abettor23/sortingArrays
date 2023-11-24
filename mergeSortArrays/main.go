package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const size1 int = 4
const size2 int = 5
const generalSize int = size1 + size2 //длина финального массива зависит от длин исходных массивов

var condition bool //условие на необходимость сортировки финального массива (не ясно из условий задачи)

func fillingElms(fillSlice []int) []int { //функция, наполняющая массив случайными числами
	seed := time.Now().UnixNano()
	rand := rand.New(rand.NewSource(seed))

	for i := 0; i < len(fillSlice); i++ {
		fillSlice[i] = rand.Intn(100)
	}
	return fillSlice
}

func sliceSorting(sortSlice []int) []int { //функция сортировки элементов в массиве по возрастанию
	for i := 0; i < len(sortSlice)-1; i++ {
		for j := 0; j < len(sortSlice)-i-1; j++ {
			if sortSlice[j] > sortSlice[j+1] {
				sortSlice[j], sortSlice[j+1] = sortSlice[j+1], sortSlice[j]
			}
		}
	}
	return sortSlice
}

func mergingArrs(s1, s2 []int) [generalSize]int { //функция, производящая слияние двух срезов...
	var mergeResult [generalSize]int

	if condition == true { //...учитывая условие решения
		tmp := append(s1, s2...)
		tmp = sliceSorting(tmp)
		copy(mergeResult[:], tmp)
	} else {
		tmp := append(s1, s2...)
		copy(mergeResult[:], tmp)
	}
	return mergeResult //...и предоставляющая на выходе массив с заданной длинной (константа)
}

func main() {

	var firstArray [size1]int //объявляем первый массив, переводим в срез, заполняем, сортируем
	slice1 := firstArray[:]
	slice1 = fillingElms(slice1)
	fmt.Println("Значения первого массива:", slice1)
	slice1 = sliceSorting(slice1)
	fmt.Printf("Первый массив отсортирован по возрастанию: %v\n\n", slice1)

	var secondArray [size2]int //второй аналогично первому
	slice2 := secondArray[:]
	slice2 = fillingElms(slice2)
	fmt.Println("Значения второго массива:", slice2)
	slice2 = sliceSorting(slice2)
	fmt.Printf("Второй массив отсортирован по возрастанию: %v\n\n", slice2)

	fmt.Println("Прежде чем произвести слияние массивов, уточните:") //выясняем условие решения
	fmt.Println("отсортировать получившийся массив? (да/нет)")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "да" {
			condition = true
			break
		} else if input == "нет" {
			condition = false
			break
		} else {
			fmt.Println("ОШИБКА! Введите \"да\" или \"нет\".")
		}
	}

	finalArray := mergingArrs(slice1, slice2) //объединяем срезы и формируем итоговый массив через функцию слияния
	fmt.Println("Ваш итоговый массив с числами:", finalArray)
}
