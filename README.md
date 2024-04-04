# Sistema Embarcado para Horticultura Automatizada

Este é um projeto de sistema embarcado escrito em Go para uma horta automatizada. O sistema é projetado para monitorar a umidade do solo e regar a horta automaticamente, conforme necessário.

## Componentes

Este projeto usa os seguintes componentes:

- Arduino: Placa microcontroladora usada para controlar o sistema e se comunicar com os sensores e atuadores.
- Sensor de umidade do solo: Detecta a umidade do solo na horta.
- Bomba de água: Atuador usado para regar a horta.
- Porta serial: Usada para comunicação entre o Arduino e o computador que executa o código Go.

## Pré-requisitos

Antes de executar este projeto, certifique-se de ter instalado o seguinte:

- Go (https://golang.org/)
- Firmware Firmata no Arduino (https://github.com/firmata/arduino)
- Biblioteca "github.com/tarm/serial" para comunicação serial em Go

## Instalação e Uso

1. Clone este repositório em sua máquina local:
´´´bash 
git clone https://github.com/seuusuario/sistema-horta-automatizada.git
´´´
