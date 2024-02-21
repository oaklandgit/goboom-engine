package main

type Tween struct {
	GameObj *GameObj
	Step float32
	Count float32
	Field func(*GameObj) *float32
	Callback func()
}

func (*Tween) Id() string {
	return "tween"
}

type TweenOption func(*Tween)

func (obj *GameObj) NewTween(
	step float32,
	field func(*GameObj) *float32,
	callback func(),
	opts ...TweenOption) *Tween {

	tween := &Tween{
		GameObj: obj,
		Step: 0.01,
		Field: field,
		Callback: callback,
	}

	for _, opt := range opts {
		opt(tween)
	}

	obj.AddComponents(tween)

	return tween
}

func (t *Tween) Update() {

	t.Count += t.Step
	
	if t.Count < 1 {
		*t.Field(t.GameObj) -= t.Step
		return
	}

	t.Callback()

}

func (t *Tween) Draw() {
	// no op
}
	
