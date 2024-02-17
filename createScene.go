package main

func CreateScene(name string, width, height float32, gameObjects ...*GameObj) *GameObj {
	scene := NewGameObject(name)
	scene.AddChildren(gameObjects...)

	return scene
}