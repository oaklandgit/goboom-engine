package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// could also repel if force is negative
type Mine struct {
	GameObj *GameObj
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
	if !m.GameObj.HasTag("docking") { return }
	if len(m.Resources) == 0 { return }

	r := m.Resources[0]
	
	if r.Remaining > 0 {
		r.Remaining--
		m.Resources[0] = r
	}

	// remove if depleted
	if r.Remaining == 0 {
		m.Resources = m.Resources[1:]
	}
}

func (m *Mine) Draw() {

	if !m.GameObj.HasTag("docking") {
		return
	}
	
	for i, r := range m.Resources {

		text := fmt.Sprintf("%s: %d of %d", r.Name, r.Remaining, r.Amount)

		rl.DrawText(
			text,
			int32(m.GameObj.Position.X - m.GameObj.Width() / 2),
			int32(m.GameObj.Position.Y) + int32(i * 14) - int32(m.GameObj.Height()),
			12, rl.White)
	}
	
}