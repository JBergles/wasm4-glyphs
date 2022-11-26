package main

import (
	"cart/w4"
	"math/rand"
)

var randSeed int64 = 0
var frameCount int64 = 0
var gamepadBounce uint8 = 0

func writeChar() {
	var buf [1]byte
	op := rand.Intn(100)
	x, y := rand.Intn(20)*8, rand.Intn(20)*8
	if op < 10 {
		//Clear a character by filling with background
		*w4.DRAW_COLORS = 0x11
		w4.Rect(x, y, 8, 8)
	} else if op < 15 {
		//Write a character with a transparent background
		buf[0] = byte(rand.Intn(0x5D) + 0x21)
		*w4.DRAW_COLORS = 0x02 + uint16(rand.Intn(3))
		w4.Text(string(buf[:]), x, y)
	} else {
		//Write a character with a solid background
		buf[0] = byte(rand.Intn(0x5D) + 0x21)
		*w4.DRAW_COLORS = 0x12 + uint16(rand.Intn(3))
		w4.Text(string(buf[:]), x, y)
	}
}

//go:export start
func start() {
	*w4.SYSTEM_FLAGS = w4.SYSTEM_PRESERVE_FRAMEBUFFER
	randSeed = int64(*w4.MOUSE_X) * int64(*w4.MOUSE_Y)
	rand.Seed(randSeed)
}

//go:export update
func update() {
	frameCount += 1
	gamepad := *w4.GAMEPAD1
	if gamepadBounce != gamepad && gamepad&w4.BUTTON_1 != 0 {
		//Reseed RNG with frame count
		randSeed = randSeed ^ frameCount
		rand.Seed(randSeed)
		//Clear the screen by filling with background color
		*w4.DRAW_COLORS = 0x11
		w4.Rect(0, 0, 160, 160)
	}
	if gamepadBounce != gamepad && gamepad&w4.BUTTON_2 != 0 {
		//Randomize color palette
		w4.PALETTE[0] = rand.Uint32()
		w4.PALETTE[1] = rand.Uint32()
		w4.PALETTE[2] = rand.Uint32()
		w4.PALETTE[3] = rand.Uint32()
	}
	for i := 0; i < 2; i++ {
		writeChar()
	}
	gamepadBounce = gamepad
}
