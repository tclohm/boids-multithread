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
	avgPosition, avgVelocity, separation := Vector2D{x: 0, y: 0}, Vector2D{x: 0, y: 0}, Vector2D{x: 0, y: 0} 

	count := 0.0
	// Readers lock
	rWLock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth) ; i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight) ; j++ {
			if otherId := boidMap[int(i)][int(j)] ; otherId != -1 && otherId != b.id {
				if dist := boids[otherId].position.Distance(b.position) ; dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherId].velocity)
					avgPosition = avgPosition.Add(boids[otherId].position)
					separation = separation.Add(b.position.Subtract(boids[otherId].position).DivideValue(dist))
				}
			}
		}
	}
	rWLock.RUnlock()

	accl := Vector2D{x:b.borderBounce(b.position.x, screenWidth), y: b.borderBounce(b.position.y, screenHeight)}

	if count > 0 {
		avgVelocity, avgPosition = avgVelocity.DivideValue(count), avgPosition.DivideValue(count)
		acclAlignment := avgVelocity.Subtract(b.velocity).MultiplyValue(adjRate)
		acclCohesion := avgPosition.Subtract(b.position).MultiplyValue(adjRate)
		acclSeparation := separation.MultiplyValue(adjRate)
		accl = accl.Add(acclAlignment).Add(acclCohesion).Add(acclSeparation)
	}

	return accl
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos - viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	// Writer Lock
	rWLock.Lock()
	b.velocity = b.velocity.Add(acceleration).Limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	rWLock.Unlock()
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