package main

import (
	"flag" // Importa el paquete flag para manejar argumentos de línea de comando
	"log"  // Proporciona funciones de registro, útiles para informar de errores y otro tipo de mensajes
	"os"   // Permite interactuar con funciones del sistema operativo, incluyendo manejo de archivos

	"github.com/gorilla/websocket" // Biblioteca que facilita el trabajo con WebSockets
)

func main() {
	// Configura y analiza la dirección del servidor desde los argumentos de la línea de comando
	addr := flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()

	// Construye la URL completa a la que se conectará el cliente WebSocket
	u := "ws://" + *addr + "/upload"
	// Intenta conectar con el servidor WebSocket en la URL especificada
	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		log.Fatal("dial:", err) // Si hay un error al conectar, termina el programa y muestra el error
	}
	defer c.Close() // Asegura que la conexión WebSocket se cierra al final de la ejecución del programa

	// Lee el archivo 'prueba.txt' que deseas enviar
	data, err := os.ReadFile("prueba.txt") //ruta
	if err != nil {
		log.Fatal(err) // Si hay un error al leer el archivo, termina el programa y muestra el error
	}

	// Envia los datos del archivo al servidor como un mensaje binario a través del WebSocket
	err = c.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		log.Fatal("write:", err) // Si hay un error al enviar el mensaje, termina el programa y muestra el error
	}

	// Registra en la consola que el archivo ha sido enviado exitosamente
	log.Println("File sent successfully")
}
