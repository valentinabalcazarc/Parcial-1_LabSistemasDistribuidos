package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	clientemetadatos "cliente/clienteMetadatos"
	clientestreaming "cliente/clienteStreaming"
)

func IniciarMenu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==========================================")
		fmt.Println("       🎵 SISTEMA DE STREAMING DE AUDIO   ")
		fmt.Println("==========================================")
		fmt.Println("1. Ver tipos de audio disponibles")
		fmt.Println("2. Salir de la aplicación")
		fmt.Print("Seleccione una opción: ")

		if !scanner.Scan() {
			break
		}
		opcion := strings.TrimSpace(scanner.Text())

		switch opcion {
		case "1":
			menuTiposAudio(scanner)
		case "2":
			fmt.Println("¡Gracias por usar la aplicación! Hasta luego.")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

func menuTiposAudio(scanner *bufio.Scanner) {
	tipos, err := clientemetadatos.ObtenerTiposAudio()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for {
		fmt.Println("\n------------------------------------------")
		fmt.Println("       📁 TIPOS DE AUDIO DISPONIBLES       ")
		fmt.Println("------------------------------------------")
		for _, tipo := range tipos {
			fmt.Printf("%d. %s\n", tipo.ID, tipo.Nombre)
		}
		fmt.Println("0. Volver al menú principal")
		fmt.Print("Seleccione un tipo de audio: ")

		if !scanner.Scan() {
			return
		}
		itemSel := strings.TrimSpace(scanner.Text())
		if itemSel == "0" {
			return
		}

		idTipo, err := strconv.Atoi(itemSel)
		if err != nil {
			fmt.Println("Por favor ingrese un número válido.")
			continue
		}

		var tipoEncontrado *clientemetadatos.TipoAudio
		for _, t := range tipos {
			if t.ID == idTipo {
				tipoEncontrado = &t
				break
			}
		}

		if tipoEncontrado == nil {
			fmt.Println("Tipo de audio no encontrado.")
			continue
		}

		menuListaAudios(scanner, tipoEncontrado)
	}
}

func menuListaAudios(scanner *bufio.Scanner, tipo *clientemetadatos.TipoAudio) {
	audios, err := clientemetadatos.ObtenerAudiosPorTipo(tipo.ID)
	if err != nil {
		fmt.Printf("Error al obtener audios: %v\n", err)
		return
	}

	for {
		fmt.Printf("\n------------------------------------------\n")
		fmt.Printf("   🎧 LISTA DE AUDIOS: %s\n", strings.ToUpper(tipo.Nombre))
		fmt.Printf("------------------------------------------\n")

		if len(audios) == 0 {
			fmt.Println("No hay audios registrados en este tipo.")
			return
		}

		for _, item := range audios {
			fmt.Printf("ID: %d | Título: %s\n", item.ID, item.ObtenerNombre())
		}
		fmt.Println("0. Volver")
		fmt.Print("Seleccione un ID de audio para ver detalles: ")

		if !scanner.Scan() {
			return
		}
		inputID := strings.TrimSpace(scanner.Text())
		if inputID == "0" {
			return
		}

		idAudio, err := strconv.Atoi(inputID)
		if err != nil {
			fmt.Println("Ingrese un ID entero válido.")
			continue
		}

		menuDetalleAudio(scanner, tipo, idAudio)
	}
}

func menuDetalleAudio(scanner *bufio.Scanner, tipo *clientemetadatos.TipoAudio, idAudio int) {
	detalles, err := clientemetadatos.ObtenerDetalleAudio(tipo.ID, idAudio)
	if err != nil {
		fmt.Println("No se pudo obtener el detalle del audio.")
		return
	}

	fmt.Println("\n==========================================")
	fmt.Println("         📝 FICHA TÉCNICA / DETALLES       ")
	fmt.Println("==========================================")
	fmt.Println(detalles)
	fmt.Println("==========================================")

	for {
		fmt.Println("1. ▶️  Reproducir audio mediante Streaming")
		fmt.Println("2. ↩️  Volver al menú anterior")
		fmt.Print("Seleccione una opción: ")

		if !scanner.Scan() {
			return
		}
		opcion := strings.TrimSpace(scanner.Text())

		switch opcion {
		case "1":
			nombreArchivo := fmt.Sprintf("%d.mp3", idAudio)
			clientestreaming.ReproducirAudio(scanner, nombreArchivo)
		case "2":
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}
