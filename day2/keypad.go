package main

import "strconv"

// Keypad move and get target button
type Keypad interface {
	Move(string) string
	GetTargetButton() string
}

type keypad struct {
	keyMatrix     map[Element]string
	targetElement Element
}

// Element i,j matrix
type Element struct {
	i int
	j int
}

// NewKeypad returns instance of new Keypad
func NewKeypad() Keypad {
	keyMatrix := make(map[Element]string)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			keyMatrix[Element{i: i, j: j}] = strconv.Itoa(3*i + j + 1)
		}
	}

	targetElement := Element{i: 1, j: 1}

	return &keypad{
		keyMatrix:     keyMatrix,
		targetElement: targetElement,
	}
}

// NewActualKeypad returns instance of new actual Keypad
func NewActualKeypad() Keypad {
	keyMatrix := make(map[Element]string)

	// TODO this can be refactored
	keyMatrix[Element{i: 0, j: 2}] = "1"
	keyMatrix[Element{i: 1, j: 1}] = "2"
	keyMatrix[Element{i: 1, j: 2}] = "3"
	keyMatrix[Element{i: 1, j: 3}] = "4"
	keyMatrix[Element{i: 2, j: 0}] = "5"
	keyMatrix[Element{i: 2, j: 1}] = "6"
	keyMatrix[Element{i: 2, j: 2}] = "7"
	keyMatrix[Element{i: 2, j: 3}] = "8"
	keyMatrix[Element{i: 2, j: 4}] = "9"
	keyMatrix[Element{i: 3, j: 1}] = "A"
	keyMatrix[Element{i: 3, j: 2}] = "B"
	keyMatrix[Element{i: 3, j: 3}] = "C"
	keyMatrix[Element{i: 4, j: 2}] = "D"

	targetElement := Element{i: 1, j: 1}

	return &keypad{
		keyMatrix:     keyMatrix,
		targetElement: targetElement,
	}
}

func (k *keypad) Move(direction string) string {
	switch direction {
	case "U":
		if _, ok := k.keyMatrix[Element{i: k.targetElement.i - 1, j: k.targetElement.j}]; ok {
			k.targetElement.i--
		}
	case "R":
		if _, ok := k.keyMatrix[Element{i: k.targetElement.i, j: k.targetElement.j + 1}]; ok {
			k.targetElement.j++
		}
	case "D":
		if _, ok := k.keyMatrix[Element{i: k.targetElement.i + 1, j: k.targetElement.j}]; ok {
			k.targetElement.i++
		}
	case "L":
		if _, ok := k.keyMatrix[Element{i: k.targetElement.i, j: k.targetElement.j - 1}]; ok {
			k.targetElement.j--
		}
	}

	return k.keyMatrix[k.targetElement]
}

func (k *keypad) GetTargetButton() string {
	return k.keyMatrix[k.targetElement]
}
