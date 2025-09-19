package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func naturalParaPar(n int) int {
	return 2 * n
}

func parParaNatural(par int) int {
	return par / 2
}

func RunBijetividadeNatural() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o número de elementos para teste de bijetividade (máximo recomendado: 1000000): ")
	input, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || n < 1 || n > 1000000 {
		fmt.Println("Tamanho inválido! Tente um valor entre 1 e 1.000.000.")
		return
	}

	start := time.Now()

	for i := 0; i < n; i++ {
		par := naturalParaPar(i)
		nat := parParaNatural(par)
		if nat != i {
			panic("Erro de bijetividade!")
		}
	}
	elapsed := time.Since(start)

	fmt.Printf("Teste de bijetividade entre %d naturais e pares concluído.\n", n)
	fmt.Printf("Tempo de execução: %d ms\n", elapsed.Milliseconds())
}
