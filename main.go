package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

type tablaPeriodica struct{
	NumAt int `json:"numat"`
	Elemento string `json:"Elemento"`
}

var elementos =  []Elemento {
	{NumAt: 1, Elemento: "Hidrógeno"},
	{NumAt: 2, Elemento: "Helio"},
	{NumAt: 3, Elemento: "Litio"},
	{NumAt: 4, Elemento: "Berilio"},
	{NumAt: 5, Elemento: "Boro"},
	{NumAt: 6, Elemento: "Carbono"},
	{NumAt: 7, Elemento: "Nitrógeno"},
	{NumAt: 8, Elemento: "Oxígeno"},
	{NumAt: 9, Elemento: "Flúor"},
	{NumAt: 10, Elemento: "Neón"},
	{NumAt: 11, Elemento: "Sodio"},
	{NumAt: 12, Elemento: "Magnesio"},
	{NumAt: 13, Elemento: "Aluminio"},
	{NumAt: 14, Elemento: "Silicio"},
	{NumAt: 15, Elemento: "Fósforo"},
	{NumAt: 16, Elemento: "Azufre"},
	{NumAt: 17, Elemento: "Cloro"},
	{NumAt: 18, Elemento: "Argón"},
	{NumAt: 19, Elemento: "Potasio"},
	{NumAt: 20, Elemento: "Calcio"},
	{NumAt: 21, Elemento: "Escandio"},
	{NumAt: 22, Elemento: "Titanio"},
	{NumAt: 23, Elemento: "Vanadio"},
	{NumAt: 24, Elemento: "Cromo"},
	{NumAt: 25, Elemento: "Manganeso"},
	{NumAt: 26, Elemento: "Hierro"},
	{NumAt: 27, Elemento: "Cobalto"},
	{NumAt: 28, Elemento: "Níquel"},
	{NumAt: 29, Elemento: "Cobre"},
	{NumAt: 30, Elemento: "Zinc"},
	{NumAt: 31, Elemento: "Galio"},
	{NumAt: 32, Elemento: "Germanio"},
	{NumAt: 33, Elemento: "Arsénico"},
	{NumAt: 34, Elemento: "Selenio"},
	{NumAt: 35, Elemento: "Bromo"},
	{NumAt: 36, Elemento: "Kriptón"},
	{NumAt: 37, Elemento: "Rubidio"},
	{NumAt: 38, Elemento: "Estroncio"},
	{NumAt: 39, Elemento: "Itrio"},
	{NumAt: 40, Elemento: "Circonio"},
	{NumAt: 41, Elemento: "Niobio"},
	{NumAt: 42, Elemento: "Molibdeno"},
	{NumAt: 43, Elemento: "Tecnecio"},
	{NumAt: 44, Elemento: "Rutenio"},
	{NumAt: 45, Elemento: "Rodio"},
	{NumAt: 46, Elemento: "Paladio"},
	{NumAt: 47, Elemento: "Plata"},
	{NumAt: 48, Elemento: "Cadmio"},
	{NumAt: 49, Elemento: "Indio"},
	{NumAt: 50, Elemento: "Estaño"},
	{NumAt: 51, Elemento: "Antimonio"},
	{NumAt: 52, Elemento: "Telurio"},
	{NumAt: 53, Elemento: "Yodo"},
	{NumAt: 54, Elemento: "Xenón"},
	{NumAt: 55, Elemento: "Cesio"},
	{NumAt: 56, Elemento: "Bario"},
	{NumAt: 57, Elemento: "Lantano"},
	{NumAt: 58, Elemento: "Cerio"},
	{NumAt: 59, Elemento: "Praseodimio"},
	{NumAt: 60, Elemento: "Neodimio"},
	{NumAt: 61, Elemento: "Prometio"},
	{NumAt: 62, Elemento: "Samario"},
	{NumAt: 63, Elemento: "Europio"},
	{NumAt: 64, Elemento: "Gadolinio"},
	{NumAt: 65, Elemento: "Terbio"},
	{NumAt: 66, Elemento: "Disprosio"},
	{NumAt: 67, Elemento: "Holmio"},
	{NumAt: 68, Elemento: "Erbio"},
	{NumAt: 69, Elemento: "Tulio"},
	{NumAt: 70, Elemento: "Iterbio"},
	{NumAt: 71, Elemento: "Lutecio"},
	{NumAt: 72, Elemento: "Hafnio"},
	{NumAt: 73, Elemento: "Tántalo"},
	{NumAt: 74, Elemento: "Wolframio"},
	{NumAt: 75, Elemento: "Renio"},
	{NumAt: 76, Elemento: "Osmio"},
	{NumAt: 77, Elemento: "Iridio"},
	{NumAt: 78, Elemento: "Platino"},
	{NumAt: 79, Elemento: "Oro"},
	{NumAt: 80, Elemento: "Mercurio"},
	{NumAt: 81, Elemento: "Talio"},
	{NumAt: 82, Elemento: "Plomo"},
	{NumAt: 83, Elemento: "Bismuto"},
	{NumAt: 84, Elemento: "Polonio"},
	{NumAt: 85, Elemento: "Astato"},
	{NumAt: 86, Elemento: "Radón"},
	{NumAt: 87, Elemento: "Francio"},
	{NumAt: 88, Elemento: "Radio"},
	{NumAt: 89, Elemento: "Actinio"},
	{NumAt: 90, Elemento: "Torio"},
	{NumAt: 91, Elemento: "Protactinio"},
	{NumAt: 92, Elemento: "Uranio"},
	{NumAt: 93, Elemento: "Neptunio"},
	{NumAt: 94, Elemento: "Plutonio"},
	{NumAt: 95, Elemento: "Americio"},
	{NumAt: 96, Elemento: "Curio"},
	{NumAt: 97, Elemento: "Berkelio"},
	{NumAt: 98, Elemento: "Californio"},
	{NumAt: 99, Elemento: "Einstenio"},
	{NumAt: 100, Elemento: "Fermio"},
	{NumAt: 101, Elemento: "Mendelevio"},
	{NumAt: 102, Elemento: "Nobelio"},
	{NumAt: 103, Elemento: "Lawrencio"},
	{NumAt: 104, Elemento: "Rutherfordio"},
	{NumAt: 105, Elemento: "Dubnio"},
	{NumAt: 106, Elemento: "Seaborgio"},
	{NumAt: 107, Elemento: "Bohrio"},
	{NumAt: 108, Elemento: "Hassio"},
	{NumAt: 109, Elemento: "Meitnerio"},
	{NumAt: 110, Elemento: "Darmstadtio"},
	{NumAt: 111, Elemento: "Roentgenio"},
	{NumAt: 112, Elemento: "Copernicio"},
	{NumAt: 113, Elemento: "Nihonio"},
	{NumAt: 114, Elemento: "Flerovio"},
	{NumAt: 115, Elemento: "Moscovio"},
	{NumAt: 116, Elemento: "Livermorio"},
	{NumAt: 117, Elemento: "Tenesino"},
	{NumAt: 118, Elemento: "Oganesón"},
}
//generamos el principio del quiz
func generateQuiz() (Elemento, []string) {
	rand.Seed(time.Now().UnixNano())
	correct := elementos[rand.Intn(len(elementos))]
	//seleccionamos las 3 opciones aleatorias
	var options []string
	options = append(options, correct.Elemento)
	for len(options) < 4 {
		incorrect := elementos[rand.Intn(len(elementos))].Elemento
		if !=contains(options, incorrect){
			options = append(options, incorrect)
		}
	}
	//Haciendo que las opciones aparezcan en orden aleatorio
	rand.Shuffle(len(options), func(i,j int) {
		options[i], options[j]= options[j], options[i]
	})
	return correct, options
}

//verificar si la opcion esta en la lista
func contains(s []string, e string) bool {
	for _, a:= range s {
		if a == e {
			return true
		}
	}
	return false
}