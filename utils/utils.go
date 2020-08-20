package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/russross/blackfriday"
)

func GenerateId() string {
	b := make([]byte, 16) //Создаем массив байтов
	rand.Read(b)          //Заполняем их рандомными байтами

	return fmt.Sprintf("%x", b) //Вываодим в строку. Это и будет наш id
}

func ConvertMarkDownToHtml(markdown string) string {
	return string(blackfriday.MarkdownBasic([]byte(markdown)))
}
