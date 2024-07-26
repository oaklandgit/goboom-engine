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
	DrawFunc Drawable
}

type Drawable struct {
	Draw    func()
	GetSize func() rl.Vector2
}

func getRootSize() rl.Vector2 {
	return rl.Vector2{X: 600, Y: 400}
}

func drawRoot() {
	rl.DrawRectangle(0, 0, int32(getRootSize().X), int32(getRootSize().Y), rl.Blue)
}

func CreateRootNode() *Node {
	node := &Node{
		Origin:  rl.Vector2{X: 0, Y: 0},
		Visible: true,
		Alpha:   1,
		// Position: rl.Vector2{X: -getRootSize().X / 2, Y: -getRootSize().Y / 2},
		Scale: rl.Vector2{X: 1, Y: 1},
		DrawFunc: Drawable{
			Draw:    drawRoot,
			GetSize: getRootSize,
		},
		// all other defaults are zero
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

// func (n *Node) RenderRoot() {

// 	if n.DrawFunc.Draw != nil {
// 		n.DrawFunc.Draw()
// 	}

// 	for _, c := range n.Children {
// 		c.Render()
// 	}

// }

func (n *Node) Render() {
	if !n.Visible {
		return
	}

	width := n.DrawFunc.GetSize().X
	height := n.DrawFunc.GetSize().Y
	origin := rl.Vector2{X: width * n.Origin.X, Y: height * n.Origin.X}
	// width := getRootSize().X
	// height := getRootSize().Y

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

	if n.DrawFunc.Draw != nil {
		n.DrawFunc.Draw()
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

	// rl.Translatef(-300, -200, 0)

	rl.PopMatrix()
}
