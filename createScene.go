package main

func CreateScene(name string, gameObjects ...*GameObj) *GameObj {
	scene := NewGameObject(name)
	scene.AddChildren(gameObjects...)

	return scene
}