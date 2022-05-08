package main

import (
	"math/rand"
	"time"
	"math"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id int
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddValue(viewRadius), b.position.AddValue(-viewRadius)
	avgVelocity := Vector2D{x: 0, y: 0}

	count := 0.0

	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth) ; i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight) ; j++ {
			if otherId := boidMap[int(i)][int(j)] ; otherId != -1 && otherId != b.id {
				if dist := boids[otherId].position.Distance(b.position) ; dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherId].velocity)
				}
			}
		}
	}

	accl := Vector2D{x:0, y:0}

	if count > 0 {
		avgVelocity = avgVelocity.DivideValue(count)
		accl = avgVelocity.Subtract(b.velocity).MultiplyValue(adjRate)
	}

	return accl
}

func (b *Boid) moveOne() {
	b.velocity = b.velocity.Add(b.calcAcceleration()).Limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	next := b.position.Add(b.velocity)
	if next.x >= screenWidth || next.x < 0 {
		b.velocity = Vector2D{x: -b.velocity.x, y: b.velocity.y}
	}

	if next.y >= screenHeight || next.y < 0 {
		b.velocity = Vector2D{x: b.velocity.x, y: -b.velocity.y}
	}

}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		velocity: Vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id: bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}