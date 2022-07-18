package main

import (
	"errors"
	"fmt"
	"os"
)

// Один, очередь моделирования массива
// Используем структуру для управления очередью
type Queue struct {
	maxSize int
	array   [5]int // Очередь моделирования массива
	front   int    // означает указание на начало очереди
	back    int    // означает указание на конец очереди
}

// Добавляем данные из очереди
func (this *Queue) AddQueue(val int) (err error) {
	// Сначала определяем, заполнена ли очередь
	if this.back == this.maxSize-1 {
		return errors.New("queue full")
	}
	this.back++ // задняя часть движется назад
	this.array[this.back] = val
	return
}

// Удаляем данные из очереди
func (this *Queue) GetQueue() (val int, err error) {
	// Сначала определяем, пуста ли очередь
	if this.back == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

// Отображение очереди, нахождение заголовка строки, а затем переход до конца строки
func (this *Queue) ShowQueue() {
	fmt.Println("Текущая ситуация в очереди:")
	for i := this.front + 1; i <= this.back; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

func main() {
	// Сначала создаем очередь
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		back:    -1,
	}
	// добавляем данные
	var key string
	var val int
	for {
		fmt.Println("1, добавить число в очередь")
		fmt.Println("2, получить число из очереди")
		fmt.Println("3. показать очередь отображения")
		fmt.Println("4, выйти")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("Пожалуйста, введите номер, который вы хотите поместить в очередь")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Добавлено в очередь")
			}
		case "2":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Удалено из очереди =", val)
			}
		case "3":
			queue.ShowQueue()
		case "4":
			os.Exit(0)
		}
	}
}
