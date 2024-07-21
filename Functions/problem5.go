/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */

package main

import "fmt"

func main() {
	fmt.Println(sumInt(1, 2, 3, 4, 5))
}

func sumInt(a ...int) (int, int) {
	var sum int = 0

	for _, elem := range a {
		sum += elem
	}

	return len(a), sum
}
