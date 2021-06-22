package main

import "fmt"

type Dog struct{}

func (d Dog) Speak() string {
	return "woof"
}

type Cat struct { }

func (c Cat) Speak() string {
	return "meow"
}

type Speaker interface {
	Speak() string
}

func haveDogSpeak(d Dog) {
	fmt.Println(d.Speak())
}

func haveSpeakerSpeak(s Speaker) {
	fmt.Println(s.Speak())
}
