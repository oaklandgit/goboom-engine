package main

type Follow struct {
	GameObj *GameObj
	Target *GameObj
}

func (*Follow) Id() string {
	return "follow"
}


type FollowOptions func(*Follow)

func (obj *GameObj) NewFollow(target *GameObj, opts ...FollowOptions) *Follow {

	follow := &Follow{
		GameObj: obj,
		Target: target,
	}

	for _, opt := range opts {
		opt(follow)
	}

	obj.AddComponents(follow)

	return follow
}

func (f *Follow) Update() {
	f.GameObj.Position = f.Target.Position
}

func (f *Follow) Draw() {
	// no op
}