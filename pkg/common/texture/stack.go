package texture

import (
	"github.com/TheWozard/GoGenerate/pkg/generate/params"
)

func NewStack(factories ...TextureStackFactoryInfo) TextureFactory {
	totalFactor := 0
	for _, factory := range factories {
		totalFactor = totalFactor + factory.Factor
	}
	return textureStackFactory{
		factories:   factories,
		totalFactor: totalFactor,
	}
}

type TextureStackFactoryInfo struct {
	Factory TextureFactory
	Texture FactorTexture
	Factor  int
}

type textureStackFactory struct {
	factories   []TextureStackFactoryInfo
	totalFactor int
}

func (sf textureStackFactory) Texture(params *params.GenerationParams) FactorTexture {
	textures := make([]textureStackInfo, len(sf.factories))
	for i, factory := range sf.factories {
		if factory.Texture != nil {
			textures[i] = textureStackInfo{
				factor:  float64(factory.Factor),
				texture: factory.Texture,
			}
		} else {
			textures[i] = textureStackInfo{
				factor:  float64(factory.Factor),
				texture: factory.Factory.Texture(params),
			}
		}
	}
	return textureStack{
		textures:    textures,
		totalFactor: float64(sf.totalFactor),
	}
}

type textureStackInfo struct {
	texture FactorTexture
	factor  float64
}

type textureStack struct {
	textures    []textureStackInfo
	totalFactor float64
}

func (ts textureStack) FactorAt(x int, y int) float64 {
	factor := 0.0
	for _, texture := range ts.textures {
		factor = factor + (texture.texture.FactorAt(x, y) * texture.factor)
	}
	return factor / ts.totalFactor
}
