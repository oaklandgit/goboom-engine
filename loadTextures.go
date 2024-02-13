package main

import rl "github.com/gen2brain/raylib-go/raylib"

func LoadTextures(textures ...string) map[string]rl.Texture2D {
	textureMap := make(map[string]rl.Texture2D)
	for _, tex := range textures {
		textureMap[tex] = rl.LoadTexture(tex)
	}
	return textureMap
}