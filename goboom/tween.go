package goboom

// a GameObj can have multiple tweens
type Tweening struct {
	GameObj *GameObj
	Tweens []*Tween
}

type Easing struct {
}

type Tween struct {
	GameObj *GameObj
	Start float32
	End float32
	Duration float32 // in frames
	Easing string
	Yoyo bool
	DelayBeforeRepeat float32 // between yoyos
	//
	Playing bool
	Step float32
	Count float32
	Property func(*GameObj) *float32
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
	opts ...TweenOption) *GameObj {

	tween := &Tween{
		GameObj: obj,
		Step: 0.01,
		Property: field,
		Callback: callback,
	}

	for _, opt := range opts {
		opt(tween)
	}

	obj.AddComponents(tween)

	return obj
}

func WithCallback(callback func()) TweenOption {
	return func(t *Tween) {
		t.Callback = callback
	}
}

func (t *Tween) Play() {
	t.Playing = true
}

func (t *Tween) Pause() {
	t.Playing = false
}

func (t *Tween) Update() {

	if !t.Playing { return }

	t.Count += t.Step
	
	if t.Count < 1 {
		*t.Property(t.GameObj) -= t.Step
		return
	}

	t.Callback()

}

func (t *Tween) Draw() {
	// no op
}
	
