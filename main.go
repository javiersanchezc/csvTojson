package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// CSVtoJSONConverter es una estructura para manejar la conversión de CSV a JSON
type CSVtoJSONConverter struct {
	CSVFilePath  string
	JSONFilePath string
}

// Convert realiza la conversión de CSV a JSON
func (c *CSVtoJSONConverter) Convert() error {
	// Abrir el archivo CSV
	file, err := os.Open(c.CSVFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Leer el archivo CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Convertir registros de CSV a JSON
	var jsonData []map[string]string
	headers := records[0]
	for _, record := range records[1:] {
		entry := make(map[string]string)
		for i, value := range record {
			entry[headers[i]] = value
		}
		jsonData = append(jsonData, entry)
	}

	// Codificar los datos JSON
	jsonDataBytes, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		return err
	}

	// Escribir los datos JSON en un archivo
	outputFile, err := os.Create(c.JSONFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = outputFile.Write(jsonDataBytes)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Ejemplo de uso
	converter := CSVtoJSONConverter{
		CSVFilePath:  "C:/data/scotiabank_cpulse_invitations.csv",
		JSONFilePath: "C:/data/output.json",
	}
	if err := converter.Convert(); err != nil {
		fmt.Println("Error al convertir CSV a JSON:", err)
		return
	}
	fmt.Println("CSV convertido a JSON correctamente.")
}
