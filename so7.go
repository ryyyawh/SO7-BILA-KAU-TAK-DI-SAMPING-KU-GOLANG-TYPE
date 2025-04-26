package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Lirik lagu "Bila Kau Tak di Sampingku" oleh Sheila on 7
var lyrics = []string{
	"Tak seharusnya kita berpisah",
	"Tak semestinya kita bertengkar",
	"Kar'na diriku masih butuh kau",
	"Maafkanlah sikapku",
	"Lupakanlah salahku itu",
	"Terlalu bodoh untuk diriku",
	"Menahan berat jutaan rindu",
	"Apalagi menahan egoku",
	"Maafkanlah sikapku",
	"Lupakanlah salahku",
	"Luapkan kepadaku",
	"Tak'kan kubiarkan kau menangis",
	"Tak'kan kubiarkan kau terkikis",
	"Terluka perasaan oleh semua ucapanku",
	"Maafkanlah semua sifat kasarku",
	"Bukan maksud untuk melukaimu",
	"Aku hanyalah orang yang penuh rasa cemburu",
	"Bila kau tak di sampingku",
	"Terlalu bodoh untuk diriku",
	"Menahan berat jutaan rindu",
	"Apalagi menahan egoku",
	"Maafkanlah sikapku",
	"Lupakanlah salahku",
	"Luapkan kepadaku",
	"Tak'kan kubiarkan kau menangis",
	"Tak'kan kubiarkan kau terkikis",
	"Terluka perasaan oleh semua ucapanku",
	"Maafkanlah semua sifat kasarku",
	"Bukan maksud untuk melukaimu",
	"Aku hanyalah orang yang penuh rasa cemburu",
	"Bila kau tak di sampingku",
	"Bila kau tak di sampingku",
	"Bila kau tak di sampingku",
}

// Warna yang akan digunakan
var colors = []func(a ...interface{}) string{
	color.New(color.FgRed).SprintfFunc(),
	color.New(color.FgGreen).SprintfFunc(),
	color.New(color.FgYellow).SprintfFunc(),
	color.New(color.FgBlue).SprintfFunc(),
	color.New(color.FgMagenta).SprintfFunc(),
	color.New(color.FgCyan).SprintfFunc(),
}

func main() {
	// Menampilkan lirik dengan efek warna-warni
	for _, line := range lyrics {
		for _, c := range colors {
			fmt.Println(c(line))
			time.Sleep(500 * time.Millisecond) // Delay 0.5 detik untuk efek warna
		}
	}
}
