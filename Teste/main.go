package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Diretório onde estão os arquivos de vídeo
	dir := "D:\\GoExpert\\3-Fundação\\"

	// Abrir o diretório
	d, err := os.Open(dir)
	if err != nil {
		fmt.Println("Erro ao abrir o diretório:", err)
		return
	}
	defer d.Close()

	// Listar arquivos no diretório
	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println("Erro ao listar arquivos:", err)
		return
	}

	// Abrir o arquivo de saída
	outputFileName := "durations.txt"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Erro ao criar arquivo de saída:", err)
		return
	}
	defer outputFile.Close()

	// Iterar sobre os arquivos
	for _, file := range files {
		// Verificar se o arquivo é um vídeo
		if strings.HasSuffix(strings.ToLower(file.Name()), ".mp4") ||
			strings.HasSuffix(strings.ToLower(file.Name()), ".avi") ||
			strings.HasSuffix(strings.ToLower(file.Name()), ".mkv") {
			// Obter a duração do vídeo usando ffmpeg
			cmd := exec.Command("ffmpeg", "-i", dir+file.Name() /* , "-show_entries", "format=duration", "-v", "quiet", "-of", " " */)

			out, err := cmd.Output()

			fmt.Println(out)

			if err != nil {
				// fmt.Println("Erro ao obter a duração de", file.Name(), ":", err)
				continue
			}

			duration := strings.TrimSpace(string(out))

			// Escrever o nome do arquivo e a duração no arquivo de saída
			outputFile.WriteString(fmt.Sprintf("Arquivo: %s, Duração: %s segundos\n", file.Name(), duration))
		}
	}

	fmt.Println("Durações dos vídeos foram salvas em", outputFileName)
}
