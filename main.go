package main

import (
	"github.com/JanainaLudwig/caesar-cipher-tdd/caesar"
	"log"
)

func main() {
	log.Println("Hello")
}


func Crypt(text string, key int) string {
	c := caesar.Caesar{Key: key}
	return c.Crypt(text)
}

func Decrypt(text string, key int) string {
	c := caesar.Caesar{Key: key}
	return c.Decrypt(text)
}