package main

import (
	"fmt"
	"time"

	"github.com/tarm/serial" // Biblioteca para comunicação serial
)

type Garden struct {
	isWatered bool
}

func (g *Garden) water() {
	fmt.Println("Regando a horta...")

	time.Sleep(2 * time.Second)
	g.isWatered = true
	fmt.Println("Horta regada!")
}

func (g *Garden) needsWater() bool {
	return !g.isWatered
}

func main() {
	// Inicializa a horta
	garden := &Garden{}

	// Configuração da comunicação serial com o Arduino
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("Erro ao abrir porta serial:", err)
		return
	}
	defer s.Close()

	for {
		if garden.needsWater() {
			garden.water()
			_, err := s.Write([]byte("L"))
			if err != nil {
				fmt.Println("Erro ao enviar comando para ligar a bomba de água:", err)
			}
		} else {
			fmt.Println("A horta já está regada.")
			_, err := s.Write([]byte("D"))
			if err != nil {
				fmt.Println("Erro ao enviar comando para desligar a bomba de água:", err)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
