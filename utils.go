package main

import (
	"crypto/rand"
	"fmt"
)

func GenerateId() string {
	b := make([]byte, 16) //Создаем массив байтов
	rand.Read(b)          //Заполняем их рандомными байтами

	return fmt.Sprintf("%x", b) //Вываодим в строку. Это и будет наш id
}
