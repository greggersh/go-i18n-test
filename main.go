package main

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {
	i18n.MustLoadTranslationFile("en-US.all.json")
	i18n.MustLoadTranslationFile("es-MX.all.json")
  fmt.Println("Test English")
	enT, err := i18n.Tfunc("en-US")
	if err == nil {
		fmt.Println(enT("hello_world"))
	}
  fmt.Println("Test Spanish")
	esT, err := i18n.Tfunc("es-MX")
	if err == nil {
		fmt.Println(esT("hello_world"))
    fmt.Println("Test fallback (key is not in ES dictionary, but is in English)")
    fmt.Println(esT("fallback"))
	}
}
