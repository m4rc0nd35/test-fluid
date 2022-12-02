package toolkit

import "log"

func Error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Recover(msg string) {
	if r := recover(); r != nil {
		log.Println(msg, r)
	}
}
