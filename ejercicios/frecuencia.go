package ejercicios

import (
	"guia6/dictionary"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func Frecuencia(s string) *dictionary.Dictionary[string, int] {

	var frecuencia = dictionary.NewDictionary[string, int]()

	for _, letra := range s {
		if stringUtils.IsAlpha(string(letra)) {
			if frecuencia.Contains(string(letra)) {
				frecuencia.Put(string(letra), frecuencia.Get(string(letra))+1)

			} else {
				frecuencia.Put(string(letra), 1)
			}
		}
	}
	return &frecuencia
}
