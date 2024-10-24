package main

import (
	"fmt"
	"log"
	"time"
	"github.com/d2r2/go-hd44780"
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
	// Inicializa
	garden := &Garden{}

	// Configuração da comunicação serial com o Arduino
	c := &serial.Config{Name: "COM4", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println("Erro ao abrir porta serial:", err)
		return
	}
	defer s.Close()

	lcd, err := hd44780.NewLcd(2, 16, hd44780.LCD_5x8DOTS, serialPort, 0x27, hd44780.HD44780_4BITMODE)
	if err != nil {
		log.Fatalf("Erro ao inicializar o LCD: %v", err)
	}
	defer lcd.Close()

	for {
		// Comando Firmata enviado para Arduino
		command := []byte{0xC8}
		_, err := s.Write(command)
		if err != nil {
			log.Fatalf("Erro ao enviar comando Firmata: %v", err)
		}

		// buffer armazenador do valor lido pelo sensor
		buffer := make([]byte, 1)
		_, err = s.Read(buffer)

		// valor do sensor
		sensorValue := int(buffer[0])

		lcd.Clear()
		lcd.SetCursorPosition(0, 0)
		lcd.Print(fmt.Sprintf("Umidade do solo: %d", sensorValue))

		time.Sleep(1 * time.Second)

		// verificação valor do sensor
		if sensorValue == 85 {
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
