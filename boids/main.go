package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"fmt"
	"sync"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount = 500
	viewRadius = 13
	adjRate = 0.015
)

// rgba(79, 230, 219, 0.8)

var (
	green = color.RGBA{R: 79, G: 230, B: 219, A: 255}
	boids [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
	rWLock = sync.RWMutex{}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x + 1), int(boid.position.y), green)
		screen.Set(int(boid.position.x - 1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y - 1), green)
		screen.Set(int(boid.position.x), int(boid.position.y + 1), green)

	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	fmt.Println("starting up...")
	fmt.Println("generating 2D Matrix with Map")
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}
	fmt.Println("✅ 2D Matrix Created")
	for i := 0 ; i < boidCount ; i++ {
		createBoid(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("ebiten initialized and running")
}