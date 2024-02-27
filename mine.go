package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mine struct {
	GameObj *GameObj
	MinedBy *GameObj
	Resources []Resource
}

type Resource struct {
	Name string
	Amount int
	Remaining int
	Price int // per unit
}

func (*Mine) Id() string {
	return "mine"
}

type MineOption func(*Mine)

func (obj *GameObj) NewMine(
	opts ...MineOption) *Mine {

	mine := &Mine{
		GameObj: obj,
	}

	for _, opt := range opts {
		opt(mine)
	}

	obj.AddComponents(mine)

	return mine
}

func (m *Mine) AddResource(name string, amount int, price int) *Mine {
	m.Resources = append(m.Resources, Resource{
		Name: name,
		Amount: amount,
		Remaining: amount,
		Price: price,
	})

	return m
}

func (m *Mine) Update() {

	// deplete if docked
	// if !m.GameObj.HasTag("docking") { return }

	if m.MinedBy == nil { return }
	if len(m.Resources) == 0 { return }

	minerBank := m.MinedBy.Components["bank"].(*Bank)

	r := m.Resources[0]
	
	if r.Remaining > 0 {
		r.Remaining--
		minerBank.Balance += r.Price
		m.Resources[0] = r
	}

	// remove if depleted
	if r.Remaining == 0 {

		sound := sounds["sounds/collected.wav"]
		rl.PlaySound(sound)

		m.Resources = m.Resources[1:]
	}
}

func (m *Mine) Draw() {

	if m.MinedBy == nil { return }
	if len(m.Resources) == 0 { return }

	for i, r := range m.GameObj.Components["mine"].(*Mine).Resources {

		itemSpacing := 38
		progressBarWidth := 160
		text := fmt.Sprintf(
			"%s: %d of %d @ $%d",
			r.Name,
			r.Remaining,
			r.Amount,
			r.Price)

		DrawProgressBar(
			22,
			int32(80 + (i * itemSpacing)),
			int32(progressBarWidth),
			int32(r.Amount - r.Remaining),
			int32(r.Amount),
			text,
		)	
	}
	
}