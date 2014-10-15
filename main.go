package main

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {
	i18n.MustLoadTranslationFile("en-US.all.json")
	i18n.MustLoadTranslationFile("es-MX.all.json")
	Ten, err := i18n.Tfunc("en-US")
	if err == nil {
		fmt.Printf(Ten("hello_world") + "\n")
	}
	Tes, err := i18n.Tfunc("es-MX")
	if err == nil {
		fmt.Printf(Tes("hello_world") + "\n")
	}
}
