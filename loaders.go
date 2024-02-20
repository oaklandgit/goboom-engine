package main

import rl "github.com/gen2brain/raylib-go/raylib"

func LoadTextures(textures ...string) map[string]rl.Texture2D {
	textureMap := make(map[string]rl.Texture2D)
	for _, tex := range textures {
		textureMap[tex] = rl.LoadTexture(tex)
	}
	return textureMap
}

func LoadFonts(fonts ...string) map[string]rl.Font {
	fontMap := make(map[string]rl.Font)
	for _, f := range fonts {
		fontMap[f] = rl.LoadFont(f)
	}
	return fontMap
}