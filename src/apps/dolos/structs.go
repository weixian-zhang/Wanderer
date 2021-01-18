package main

import(
	"time"
)

type Book struct {
	Name string				`fake:"{sentence:3}"`
	Author string			`fake:"{firstname}"`
	PublishDate time.Time 
}