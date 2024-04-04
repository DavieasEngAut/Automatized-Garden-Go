package main

import (
	"fmt"
	"log"
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

		command := []byte{0xC8}
		_, err := s.Write(command)
		if err != nil {
			log.Fatalf("Erro ao enviar comando Firmata: %v", err)
		}
		buffer := make([]byte, 1)
		_, err = s.Read(buffer)

		sensorValue := int(buffer[0])

		time.Sleep(1 * time.Second)

		if sensorValue == 850 {
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
