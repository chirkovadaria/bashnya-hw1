package main

import "fmt"

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	for num < 123 {
		if num%13 == 0 && num%9 == 0 {
			fmt.Println("service error")
			break
		} else if num < 0 {
			num *= -1
		} else if num%7 == 0 {
			num *= 39
		} else if num%9 == 0 {
			num *= 13
			num += 1
			continue
		} else {
			num += 2
			num *= 3
		}
	}
	fmt.Printf("После всех изменений мы получили: %d ", num)

	/*
	    Написание кода заняло 3 минуты.

	   Осознание, что в bashnya-hw1 не должен находиться файл bashnya-hw2,
	   распутывание структуры проектов и понимание важности
	   регулярного сохранения через "Save All" заняло 2,5 часа.

	   Было больно, но это бесценный опыт.
	   (Toxic positivity)
	*/
}
