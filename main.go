// Dependendo do tema escolhido, gere um JSON que atenda as seguintes chaves de acordo
// com o tema.
// Os produtos variam por id, nome, cor, preço, estoque, código (alfanumérico), publicação
// (sim-não), data de criação.
// Os usuários variam por id, nome, sobrenome, e-mail, idade, altura, ativo (sim-não), data de
// criação.
// Transações: id, código da transação (alfanumérico), moeda, valor, emissor (string), receptor
// (string), data da transação.
// 1. Dentro da pasta go-web crie um arquivo theme.json, o nome tem que ser o tema
// escolhido, ex: products.json.
// 2. Dentro dele escrevi um JSON que permite ter uma matriz de produtos, usuários ou
// transações com todas as suas variantes.

// Crie dentro da pasta go-web um arquivo chamado main.go
// 2. Crie um servidor web com Gin que retorne um JSON que tenha uma chave
// “mensagem” e diga Olá seguido do seu nome.
// 3. Acesse o end-point para verificar se a resposta está correta.

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Jessica!",
	})
}

func usuariosHandler(c *gin.Context) {
	jsonFile, err := os.Open("usuarios.json")
	jsonText, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	var usuarios []Usuario
	json.Unmarshal([]byte(jsonText), &usuarios)
	c.JSON(200, usuarios)

}

type Usuario struct {
	Id            int     `json:"id"`
	Nome          string  `json:"nome"`
	Sobrenome     string  `json:"sobrenome"`
	Email         string  `json:"email"`
	Idade         int     `json:"idade"`
	Altura        float32 `json:"altura"`
	Ativo         bool    `json:"ativo"`
	DataDeCriacao string  `json:"dataDeCriacao"`
}

func main() {
	router := gin.Default()

	router.GET("/nome", helloHandler)

	router.GET("/usuarios", usuariosHandler)

	router.Run()

}
