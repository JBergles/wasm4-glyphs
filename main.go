package main

import (
	"cart/w4"
	"math/rand"
)

func writeChar() {
	var buf [1]byte
	op := rand.Intn(100)
	x, y := rand.Intn(20)*8, rand.Intn(20)*8
	if op < 10 {
		*w4.DRAW_COLORS = 0x11
		w4.Rect(x, y, 8, 8)
	} else if op < 15 {
		buf[0] = byte(rand.Intn(0x5D) + 0x21)
		*w4.DRAW_COLORS = 0x02 + uint16(rand.Intn(3))
		w4.Text(string(buf[:]), x, y)
	} else {
		buf[0] = byte(rand.Intn(0x5D) + 0x21)
		*w4.DRAW_COLORS = 0x12 + uint16(rand.Intn(3))
		w4.Text(string(buf[:]), x, y)
	}
}

//go:export start
func start() {
	*w4.SYSTEM_FLAGS = w4.SYSTEM_PRESERVE_FRAMEBUFFER
	rand.Seed(int64(*w4.MOUSE_X) * int64(*w4.MOUSE_Y))
}

//go:export update
func update() {
	gamepad := *w4.GAMEPAD1
	if gamepad&w4.BUTTON_1 != 0 {
		*w4.DRAW_COLORS = 0x11
		w4.Rect(0, 0, 160, 160)
	}
	if gamepad&w4.BUTTON_2 != 0 {
		w4.PALETTE[0] = rand.Uint32()
		w4.PALETTE[1] = rand.Uint32()
		w4.PALETTE[2] = rand.Uint32()
		w4.PALETTE[3] = rand.Uint32()
	}
	for i := 0; i < 2; i++ {
		writeChar()
	}
}
