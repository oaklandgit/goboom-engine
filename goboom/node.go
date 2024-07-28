package goboom

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node struct {
	Visible  bool
	Tags     []string
	Velocity rl.Vector2
	Position rl.Vector2
	Scale    rl.Vector2
	Origin   rl.Vector2
	Rotation float32
	Alpha    float32
	Children []*Node
	Parent   *Node
	Layer    int

	OnUpdate  func(*Node)
	OnDraw    func(*Node)
	GetBounds func(*Node) rl.Rectangle
}

func CreateRootNode(w, h float32) *Node {

	node := &Node{
		Origin:  rl.Vector2{X: 0, Y: 0},
		Visible: true,
		Alpha:   1,
		Scale:   rl.Vector2{X: 1, Y: 1},
		OnDraw: func(n *Node) {
			rl.DrawRectangle(0, 0, int32(w), int32(h), rl.Blue)
		},
		GetBounds: func(n *Node) rl.Rectangle {
			return rl.Rectangle{X: n.Position.X, Y: n.Position.Y, Width: w, Height: h}
		},
	}

	return node
}

func (n *Node) AddChildren(children ...*Node) {
	for _, c := range children {
		// remove from other parents
		if c.Parent != nil {
			c.Parent.RemoveChild(c)
		}

		c.Parent = n
	}
	n.Children = append(n.Children, children...)
}

func (n *Node) RemoveChild(child *Node) {
	if child.Parent != n {
		return
	}

	for i, c := range n.Children {
		if c == child {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			child.Parent = nil
			break
		}
	}
}

func (n *Node) GetPosGlobal() rl.Vector2 {
	if n.Parent != nil {
		return rl.Vector2Add(n.Position, n.Parent.GetPosGlobal())
	}
	return n.Position
}

func (n *Node) GetLayer() int {
	return n.Layer
}

func (n *Node) SetLayer(l int) {
	n.Layer = l
	if n.Parent != nil {
		sort.Slice(n.Parent.Children, func(i, j int) bool {
			return n.Parent.Children[i].Layer < n.Parent.Children[j].Layer
		})
	}
}

func (n *Node) Render() {
	if !n.Visible {
		return
	}

	rl.PushMatrix()

	globalPos := n.GetPosGlobal()
	rl.Translatef(globalPos.X, globalPos.Y, 0)

	if n.Parent != nil {
		rl.Rotatef(n.Rotation, 0, 0, 1)
		rl.Scalef(n.Scale.X, n.Scale.Y, 1)
	}

	if n.OnDraw != nil {
		originOffset := rl.Vector2{X: n.GetBounds(n).Width * n.Origin.X, Y: n.GetBounds(n).Height * n.Origin.Y}
		rl.Translatef(-originOffset.X, -originOffset.Y, 0)
		n.OnDraw(n)
		// draw centerpoint for debugging
		rl.DrawCircle(int32(originOffset.X), int32(originOffset.Y), 2, rl.Black)
	}

	rl.PopMatrix()

	for _, child := range n.Children {
		child.Render()
	}

}
