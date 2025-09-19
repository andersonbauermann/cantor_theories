# Teoria dos Conjuntos e Algoritmos de Cantor em C# e Go

Este projeto reúne exemplos de algoritmos inspirados na Teoria dos Conjuntos de Georg Cantor, implementados em **C#** e **Go**. Aqui você verá, na prática, ideias fundamentais da matemática do infinito, como o famoso "Hotel Infinito", o argumento diagonal, a bijetividade entre conjuntos infinitos e a geração do conjunto potência.

---

## ✨ **Conceitos Fundamentais de Cantor**

### 1. O Infinito Contável e o Hotel de Hilbert (Cantor)

Imagine um hotel com infinitos quartos, todos ocupados. Chega um novo hóspede! O gerente pede para cada hóspede se mover para o quarto seguinte (n → n+1), liberando o quarto 1. Assim, sempre é possível acomodar mais pessoas em um “infinito contável”.  
**Moral:** O conjunto dos números naturais é infinito, mas ainda podemos listar todos seus elementos (1, 2, 3, ...).

### 2. Argumento Diagonal de Cantor

Cantor mostrou que nem todo infinito é igual!  
Ao tentar listar todos os números reais entre 0 e 1 em formato decimal (ou binário), é possível construir um novo número **diferença** de cada número da lista em pelo menos uma casa decimal.  
**Conclusão:** O conjunto dos números reais é “maior” que o dos naturais — é um infinito **incontável**.

### 3. Bijetividade entre Naturais e Pares

Cantor provou que conjuntos infinitos podem ter o mesmo tamanho!  
Existe uma correspondência perfeita entre os números naturais (1, 2, 3, ...) e os números pares (2, 4, 6, ...): basta multiplicar por 2.  
**Moral:** Nem todo infinito é maior que outro — alguns são “iguais” em tamanho.

### 4. O Conjunto Potência

O conjunto potência de um conjunto com n elementos tem 2ⁿ subconjuntos. Cantor mostrou que sempre podemos criar um conjunto “maior” que qualquer dado, mesmo infinito.  
**Na prática:** Para um conjunto com 20 elementos, temos mais de um milhão de subconjuntos!

---

## 🧮 **Algoritmos Implementados**

### 1. **Argumento Diagonal (Cantor)**

- Gera uma lista de números binários e cria um novo número que difere de cada número da lista em uma posição específica (a diagonal).
- **Mostra:** Não dá para listar todos os números reais — sempre existe um fora da lista!

### 2. **Bijetividade entre Naturais e Pares**

- Mapeia naturais para pares e vice-versa.
- **Mostra:** Existem infinitos do mesmo tamanho, mesmo que pareçam diferentes.

### 3. **Geração do Conjunto Potência**

- Gera todos os subconjuntos possíveis de um conjunto finito.
- **Mostra:** O tamanho do conjunto potência cresce exponencialmente, ilustrando infinitos maiores.

---

## 🚀 **Como Clonar e Testar o Projeto**

### **Pré-requisitos**

- [Git](https://git-scm.com/)
- [.NET SDK](https://dotnet.microsoft.com/) para C#
- [Go](https://golang.org/) para Go

### **Clone o repositório**

```bash
git clone https://github.com/andersonbauermann/cantor_theories.git
cd data_structure_and_algorithms
```

### **Para rodar os exemplos em C#**

1. Navegue até o diretório do exemplo desejado.
2. Compile e execute:

   ```bash
   dotnet run -c Release
   ```

3. Siga as instruções do programa para inserir o tamanho do conjunto ou lista.

### **Para rodar os exemplos em Go**

1. Navegue até o diretório do exemplo desejado.
2. Compile e execute:

   ```bash
   go run nome_do_arquivo.go
   ```

3. Siga as instruções do programa para inserir o tamanho do conjunto ou lista.

### **Sobre os Benchmarks**

- Os exemplos incluem medições de tempo simples.
- Para benchmarks completos, utilize [BenchmarkDotNet](https://benchmarkdotnet.org/) em C# e o framework de benchmarks do Go (`go test -bench .`).

---

## 📚 **O que cada algoritmo tenta mostrar?**

- **Diagonal de Cantor:** Você nunca conseguirá listar todos os números reais — sempre existe um “fora da lista”.
- **Bijetividade Naturais ↔ Pares:** Nem todo infinito é maior que outro; alguns são “iguais”.
- **Conjunto Potência:** Sempre existe um conjunto maior, até entre infinitos!

---

## 🤔 **Reflexão Final**

Sabe o que isso prova? Nada — mas esse código exemplifica um pouco sobre teoria dos conjuntos e como diferentes linguagens se adaptam a diferentes problemas.

---

**Divirta-se descobrindo os infinitos com Cantor e algoritmos!**
