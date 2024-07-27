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
	Rotation float32
	Scale    rl.Vector2
	Alpha    float32
	Children []*Node
	Parent   *Node
	Layer    int
	Origin   rl.Vector2
	Update   func(*Node)
	Draw     func(*Node)
	Bounds   func(*Node) rl.Rectangle
}

func CreateRootNode(w, h float32) *Node {

	node := &Node{
		Origin:  rl.Vector2{X: 0, Y: 0},
		Visible: true,
		Alpha:   1,
		Scale:   rl.Vector2{X: 1, Y: 1},
		Draw: func(n *Node) {
			rl.DrawRectangle(0, 0, int32(w), int32(h), rl.Blue)
		},
		Bounds: func(n *Node) rl.Rectangle {
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

	width := n.Bounds(n).Width
	height := n.Bounds(n).Height
	origin := rl.Vector2{X: width * n.Origin.X, Y: height * n.Origin.X}

	// Apply parent's global position
	if n.Parent != nil {
		rl.PushMatrix()
		parentPos := n.Parent.GetPosGlobal()
		rl.Translatef(parentPos.X+n.Position.X, parentPos.Y+n.Position.Y, 0)
	}

	rl.PushMatrix()
	rl.Translatef(-origin.X, -origin.Y, 0)

	rl.Rotatef(n.Rotation, 0, 0, 1)
	rl.Scalef(n.Scale.X, n.Scale.Y, 1)

	if n.Draw != nil {
		n.Draw(n)
		rl.DrawCircle(int32(origin.X), int32(origin.Y), 2, rl.Black) // should be the centerpoint
	}

	// Render children
	for _, c := range n.Children {
		c.Render()
	}

	// Reset
	if n.Parent != nil {
		rl.Translatef(origin.X, origin.Y, 0)
		rl.PopMatrix()
	}

	rl.PopMatrix()
}
