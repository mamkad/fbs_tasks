/*
 * Функция для вычисления чисел фибоначчи на заданном промежутке
 */

package fibonacci

import "errors"

func Fibonacci(startNumber, endNumber int64) (map[int64]int64, error) {
	//проверяем корректность диапазона

	if endNumber < 0 || startNumber < 0 {
		return nil, errors.New("range boundaries are less than 0")
	}

	if startNumber >= endNumber {
		return nil, errors.New("beginning of the range should be less than the end")
	}

	//вычисление чисел фибоначчи

	fibSeries := make(map[int64]int64, endNumber-startNumber+1) //map
	var fir, sec, N int64 = 0, 1, 1                             //числа для вычисления и номер в ряду фибоначчи

	//вычисляем без сохранения числа фибоначчи до нашего диапазона

	for ; N <= startNumber; N++ {
		fir, sec = sec, fir+sec
	}

	//вычисляем числа фибоначчи, входящие в наш диапазон

	for i := int64(startNumber); i <= endNumber; i += 1 {
		fibSeries[i] = fir
		fir, sec = sec, fir+sec
	}

	return fibSeries, nil
}
