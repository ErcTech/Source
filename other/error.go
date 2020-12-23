package erctech

type Erro struct {
	// Código do Erro
	Codigo int `json:"codigo"`
	// Mensagem do Erro
	Erro string `json:"msg"`
}
