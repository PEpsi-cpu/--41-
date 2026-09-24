package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Текущая дата и время:", now.Format("02.01.2006 15:04:05"))
	fmt.Println("Дата:", now.Format("02 January 2006"))
	fmt.Println("Время:", now.Format("15:04:05"))
}
