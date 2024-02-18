package main

type Rotate struct {
	GameObj *GameObj
	Speed float32 // can be negative
}

func (*Rotate) Id() string {
	return "rotate"
}


func (r *Rotate) Draw() {
	// no op
}

type RotateOption func(*Rotate)

func (o *GameObj) NewRotate(speed float32, opts ...RotateOption) *Rotate {

	rotate := &Rotate{
		GameObj: o,
		Speed: speed,
	}

	for _, opt := range opts {
		opt(rotate)
	}

	o.AddComponents(rotate)

	return rotate
}

func (r *Rotate) Update() {

	r.GameObj.Angle += r.Speed
	if r.GameObj.Angle > 360 {
		r.GameObj.Angle = 0
	}
	if r.GameObj.Angle < 0 {
		r.GameObj.Angle = 360
	}
	
}

