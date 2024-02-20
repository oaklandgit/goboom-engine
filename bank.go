package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// could also repel if force is negative
type Bank struct {
	GameObj *GameObj
	Balance int
}

func (*Bank) Id() string {
	return "bank"
}

type BankOption func(*Bank)

func (obj *GameObj) NewBank(
	opts ...BankOption) *Bank {

	bank := &Bank{
		GameObj: obj,
	}

	for _, opt := range opts {
		opt(bank)
	}

	obj.AddComponents(bank)

	return bank
}

func (b *Bank) Deposit(amount int) *Bank {
	b.Balance += amount
	return b
}

func (b *Bank) Update() {
	// no op
}

func (b *Bank) Draw() {

	formatter := message.NewPrinter(language.English)
	text := formatter.Sprintf("$%d", b.Balance)

	DrawText(text, 400, 16, 42, 12, rl.Green, Center)
	
}