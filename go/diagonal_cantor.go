package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func diagonalArgument(binaryList []string) string {
	n := len(binaryList)
	result := make([]byte, n)

	for i := 0; i < n; i++ {
		if binaryList[i][i] == '0' {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}

	return string(result)
}

func RunDiagonalCantor() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o tamanho da lista binária (máximo recomendado: 30): ")
	input, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || n < 1 || n > 30 {
		fmt.Println("Tamanho inválido! Tente um valor entre 1 e 30.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	binaryList := make([]string, n)

	for i := 0; i < n; i++ {
		s := make([]byte, n)
		for j := 0; j < n; j++ {
			if rand.Intn(2) == 0 {
				s[j] = '0'
			} else {
				s[j] = '1'
			}
		}
		binaryList[i] = string(s)
	}

	start := time.Now()
	diagonal := diagonalArgument(binaryList)
	elapsed := time.Since(start)

	fmt.Println("Exemplo de lista binária:")

	for i := 0; i < n && i < 6; i++ {
		fmt.Println(binaryList[i])
	}

	fmt.Printf("Novo número gerado pela diagonalização: %s\n", diagonal)
	fmt.Printf("Tempo de execução: %d ms\n", elapsed.Milliseconds())
}
