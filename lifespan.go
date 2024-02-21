package main

// could also repel if force is negative
type Lifespan struct {
	GameObj *GameObj
	Total int
}

func (*Lifespan) Id() string {
	return "lifespan"
}

type LifespanOption func(*Lifespan)

func (obj *GameObj) NewLifespan(total int, opts ...LifespanOption) *Lifespan {

	lifespan := &Lifespan{
		GameObj: obj,
		Total: total,
	}

	for _, opt := range opts {
		opt(lifespan)
	}

	obj.AddComponents(lifespan)

	return lifespan
}


func (l *Lifespan) Update() {
	l.Total--
	if l.Total <= 0 {
		l.GameObj.Delete()
	}
}

func (l *Lifespan) Draw() {
	// no op
}