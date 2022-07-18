// Два, круговая очередь моделирования массива circlequeue
package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int    //4
	array   [5]int // Массив
	head    int    // Вставляем в начало очереди
	tail    int    // Вставить в конец очереди
}

// Добавляем метод в очередь AddQueue (push) GetQueue (pop)
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}

	//this.tail находится в конце очереди и не содержит последнего элемента
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// Вне очереди
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return -1, errors.New("queue empty")
	}

	// head относится к главе команды и содержит элементы главы команды
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

// Определяем, заполнена ли круговая очередь
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// Определяем, пусто ли оно
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// Удаляем количество элементов в круговой очереди
func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

// Отображение очереди
func (this *CircleQueue) ListQueue() {
	fmt.Println("Ситуация с круговой очередью следующая")
	// Удаляем количество элементов в текущей очереди
	size := this.Size()
	if size == 0 {
		fmt.Println("Очередь пуста")
	}
	// Создание вспомогательной переменной, указывающей на голову
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

func main() {
	// Сначала создаем круговую очередь
	circlequeue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}
	// добавляем данные
	var key string
	var val int
	for {
		fmt.Println("1, Добавить число в очередь")
		fmt.Println("2, Добавить число из очереди")
		fmt.Println("3. Показать очередь")
		fmt.Println("4, Выйти из очереди")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("Пожалуйста, введите номер, который вы хотите поставить в очередь")
			fmt.Scanln(&val)
			err := circlequeue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Добавлено")
			}
		case "2":
			val, err := circlequeue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("Данные из очереди: %d \n", val)
			}
		case "3":
			circlequeue.ListQueue()
		case "4":
			os.Exit(0)
		}
	}
}
