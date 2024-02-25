package main

type PointAt struct {
	GameObj *GameObj
	Target *GameObj
}

func (*PointAt) Id() string {
	return "pointat"
}

func (obj *GameObj) NewPointAt(target *GameObj) *PointAt {

	pointAt := &PointAt{
		GameObj: obj,
		Target: target,
	}

	obj.AddComponents(pointAt)

	return pointAt
}

func (p *PointAt) Update() {

	angle := calculateAngle(p.GameObj.PosGlobal(), p.Target.PosGlobal())
	p.GameObj.Angle = angle

}

func (s *PointAt) Draw() {
	// no op
}