# GoGenerate
[![test](https://github.com/TheWozard/GoGenerate/actions/workflows/test.yml/badge.svg)](https://github.com/TheWozard/GoGenerate/actions/workflows/test.yml)  
An exploration into the world of procedural image generation

## Getting started
Run
```
$ go run ./cmd/generate [generator]
```

## Current Generators

| Name | Output |
| - | - |
| Blank | Creates a blank image of the specified color |
| Gradient | Creates a basic left to right linear gradient |
| Ramp | Creates a basic left to right color ramp gradient |
| Perlin | A basic perlin texture at a scale of 50px grid |
| Ramp-Tex | Texture based gradient implementation |
| Worley | A basic worley texture of 5x5 cells each with a single point |
| Tiles | WIP making generated stone pavers |
| Stack-Worley | implementation of multiple worley textures stacked |

## Future Plans
Deploy as a serverless application to make easier to demo