package main

import "container/ring"

// Compass - get the next or prev direction
type Compass interface {
	Move(string) string
	GetFacingCardinalDirection() string
}

type compass struct {
	compassRing *ring.Ring
}

// NewCompass returns a instance of new compass
func NewCompass() Compass {
	cardinalDirections := []string{"N", "E", "S", "W"}
	compassRing := ring.New(len(cardinalDirections))

	for _, direction := range cardinalDirections {
		compassRing.Value = direction
		compassRing = compassRing.Next()
	}

	return &compass{
		compassRing: compassRing,
	}
}

func (c *compass) Move(direction string) string {
	if direction == "R" {
		c.compassRing = c.compassRing.Next()
	} else {
		c.compassRing = c.compassRing.Prev()
	}

	return c.compassRing.Value.(string)
}

func (c *compass) GetFacingCardinalDirection() string {
	return c.compassRing.Value.(string)
}
