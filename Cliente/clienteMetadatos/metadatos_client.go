package clientemetadatos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const urlMetadatos = "http://localhost:8081"

/**
 * @brief Obtiene la lista de tipos de audio disponibles.
 * 
 * Realiza una petición GET al servidor de metadatos.
 * 
 * @return Una lista de TipoAudio y un error si ocurre un fallo.
 */
func ObtenerTiposAudio() ([]TipoAudio, error) {
	resp, err := http.Get(urlMetadatos + "/tipos-audio")
	if err != nil {
		return nil, fmt.Errorf("⚠️ Error al conectar con ServidorDeMetadatos: %v", err)
	}
	defer resp.Body.Close()

	var data RespuestaTiposAudio
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("⚠️ Error decodificando respuesta JSON: %v", err)
	}

	return data.Tipos, nil
}

/**
 * @brief Obtiene los audios disponibles según el tipo de audio.
 * 
 * @param idTipo Identificador del tipo de audio.
 * @return Una lista de ItemAudio y un error en caso de fallo.
 */
func ObtenerAudiosPorTipo(idTipo int) ([]ItemAudio, error) {
	var endpoint string
	switch idTipo {
	case 1:
		endpoint = "/musicas"
	case 2:
		endpoint = "/ruidosblancos"
	case 3:
		endpoint = "/podcasts"
	case 4:
		endpoint = "/audiolibros"
	default:
		return nil, fmt.Errorf("⚠️ Tipo de audio #%d no soportado", idTipo)
	}

	resp, err := http.Get(urlMetadatos + endpoint)
	if err != nil {
		return nil, fmt.Errorf("⚠️ Error al obtener lista de audios: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var rawMap map[string]interface{}
	json.Unmarshal(bodyBytes, &rawMap)

	var items []ItemAudio
	for key, val := range rawMap {
		if key != "codigo" && key != "mensaje" {
			bytesSlice, _ := json.Marshal(val)
			json.Unmarshal(bytesSlice, &items)
		}
	}

	return items, nil
}

/**
 * @brief Obtiene los detalles específicos de un audio en formato JSON.
 * 
 * @param idTipo Identificador del tipo de audio.
 * @param idAudio Identificador único del audio.
 * @return Una cadena en formato JSON indentado con los detalles del audio y un error si falla.
 */
func ObtenerDetalleAudio(idTipo int, idAudio int) (string, error) {
	var endpoint string
	switch idTipo {
	case 1:
		endpoint = fmt.Sprintf("/musica/%d", idAudio)
	case 2:
		endpoint = fmt.Sprintf("/ruidoblanco/%d", idAudio)
	case 3:
		endpoint = fmt.Sprintf("/podcast/%d", idAudio)
	case 4:
		endpoint = fmt.Sprintf("/audiolibro/%d", idAudio)
	}

	resp, err := http.Get(urlMetadatos + endpoint)
	if err != nil || resp.StatusCode != 200 {
		return "", fmt.Errorf("⚠️ No se pudo obtener el detalle del audio #%d", idAudio)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var prettyJSON map[string]interface{}
	json.Unmarshal(bodyBytes, &prettyJSON)
	detalles, _ := json.MarshalIndent(prettyJSON, "", "  ")

	return string(detalles), nil
}
