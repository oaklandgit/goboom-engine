package goboom

import rl "github.com/gen2brain/raylib-go/raylib"

type Shoot struct {
	GameObj *GameObj
	Target *GameObj
	Ordnance *GameObj
	Cooldown float32
	OriginPos rl.Vector2
}

func (*Shoot) Id() string {
	return "shoot"
}

type ShootOptions func(*Shoot)

func (obj *GameObj) NewShoot(target *GameObj) *GameObj {

	shoot := &Shoot{
		GameObj: obj,
	}

	obj.AddComponents(shoot)

	return obj
}

func WithTarget(target *GameObj) ShootOptions {
	return func(s *Shoot) {
		s.Target = target
	}
}

func WithOrdnance(ordnance *GameObj) ShootOptions {
	return func(s *Shoot) {
		s.Ordnance = ordnance
	}
}


func (s *Shoot) Pew() {
	
}

func (s *Shoot) Update() {


}

func (s *Shoot) Draw() {
	// no op
}