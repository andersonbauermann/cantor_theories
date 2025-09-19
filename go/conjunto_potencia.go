package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func powerSet(set []int) [][]int {
	n := len(set)
	total := 1 << n
	result := make([][]int, 0, total)

	for mask := 0; mask < total; mask++ {
		subset := make([]int, 0, n)
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, set[i])
			}
		}
		result = append(result, subset)
	}

	return result
}

func RunConjuntoPotencia() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o tamanho do conjunto (máximo recomendado: 20): ")
	input, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || n < 0 || n > 20 {
		fmt.Println("Tamanho inválido! Tente um valor entre 0 e 20.")
		return
	}

	set := make([]int, n)

	for i := 0; i < n; i++ {
		set[i] = i + 1
	}

	start := time.Now()
	powerSet := powerSet(set)
	elapsed := time.Since(start)

	fmt.Printf("Total de subconjuntos gerados: %d (2^%d)\n", len(powerSet), n)
	fmt.Printf("Tempo de execução: %d ms\n", elapsed.Milliseconds())
	fmt.Println("Exemplo de subconjuntos:")
	exibe := 8

	if len(powerSet) < exibe {
		exibe = len(powerSet)
	}

	for i := 0; i < exibe; i++ {
		fmt.Println(powerSet[i])
	}
}
