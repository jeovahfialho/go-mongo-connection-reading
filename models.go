package main

import (
	"time"
)

type Fornecedor struct {
	Index        int       `bson:"index"`
	CPF          string    `bson:"cpf"`
	DataAbertura time.Time `bson:"dataAbertura"`
	Email        string    `bson:"email"`
	Nome         string    `bson:"nome"`
	Status       string    `bson:"status"`
}
