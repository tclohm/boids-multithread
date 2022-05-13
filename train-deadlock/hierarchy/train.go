package hierarchy

import (
	. "github.com/tclohm/multithreading-go/train-deadlock/common"
	"time"
	"sort"
)

func lockIntersectionsInDistance(id, reserveStart, reserveEnd int, crossings []*Crossing) {
	var intersectionsToLock []*Intersection
	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position && reserveStart <= crossing.Position && crossing.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	// if the id is smaller we will lock it first
	sort.Slice(intersectionsToLock, func(i,j int) bool {
		return intersectionsToLock[i].Id < intersectionsToLock[j].Id
	})

	for _, it := range intersectionsToLock {
		it.Mutex.Lock()
		it.LockedBy = id
		// gives time to allow threads to lock
		time.Sleep(10 * time.Millisecond)
	}
}

func MoveTrain(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockIntersectionsInDistance(train.Id, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}