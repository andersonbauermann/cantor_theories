package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("Selecione a operação:")
		fmt.Println("1 - Teste de bijetividade natural")
		fmt.Println("2 - Gerar conjunto potência")
		fmt.Println("3 - Diagonalização de Cantor")
		fmt.Println("0 - Sair")
		fmt.Print("Opção: ")

		var opcao int
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Entrada inválida! Tente novamente.")
			continue
		}

		if opcao == 0 {
			fmt.Println("Saindo...")
			break
		}

		switch opcao {
		case 1:
			RunBijetividadeNatural()
		case 2:
			RunConjuntoPotencia()
		case 3:
			RunDiagonalCantor()
		default:
			fmt.Println("Opção inválida!")
		}

		fmt.Println()
	}
}
