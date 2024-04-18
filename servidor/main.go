package main

import (
	"log"      // Se utiliza para registrar mensajes en la consola, útil para depuración y registro de errores.
	"net/http" // Proporciona herramientas para construir servidores HTTP en Go.
	"os"       // Permite interactuar con funciones del sistema operativo, como crear archivos.

	"github.com/gorilla/websocket" // Importa el paquete gorilla/websocket que facilita trabajar con WebSockets.
)

// upgrader se configura con los tamaños de buffer para la lectura y escritura.
// Estos buffers determinan el tamaño en bytes para los buffers de datos de los WebSockets.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// uploadFile es un manejador HTTP que se ejecuta cuando se recibe una solicitud HTTP en el endpoint "/upload".
func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Intenta actualizar la conexión HTTP a una conexión WebSocket.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err) // Registra el error si la actualización falla.
		return
	}
	defer conn.Close() // Asegura que la conexión se cierre al finalizar la función.

	// Leer el mensaje (archivo) del cliente. WebSockets pueden enviar y recibir mensajes en múltiples formatos.
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println(err) // Registra el error si la lectura del mensaje falla.
		return
	}

	// Guardar el archivo recibido en el sistema de archivos del servidor.
	// El archivo se guarda con el nombre "prueba_recibida.txt", y los permisos 0666 permiten que el archivo sea leído y escrito por cualquier usuario.
	err = os.WriteFile("prueba_recibida.txt", msg, 0666)
	if err != nil {
		log.Println(err) // Registra el error si la escritura del archivo falla.
		return
	}

	log.Println("File received and saved successfully.") // Registra un mensaje de éxito cuando el archivo es guardado correctamente.
}

// La función main configura y ejecuta el servidor HTTP.
func main() {
	http.HandleFunc("/upload", uploadFile)                     // Registra el manejador uploadFile para la ruta "/upload".
	log.Println("Server starting on http://localhost:8080...") // Muestra un mensaje por consola indicando que el servidor está iniciado.
	log.Fatal(http.ListenAndServe(":8080", nil))               // Inicia el servidor en el puerto 8080 y registra fatalmente cualquier error que ocurra.
}
