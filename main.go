package main

import "cart/w4"

var rocket = [13]byte{
    0b00111100,
    0b01111110,
    0b11111111,
    0b11111111,
    0b11000011,
    0b11000011,
    0b11111111,
    0b11111111,
    0b01111110,
    0b11111111,
    0b11111111,
    0b00111100,
    0b00011000,
}

var rocket_unthrust = [13]byte{
    0b00111100,
    0b01111110,
    0b11111111,
    0b11111111,
    0b11000011,
    0b11000011,
    0b11111111,
    0b11111111,
    0b01111110,
    0b11111111,
    0b11111111,
    0b00000000,
    0b00000000,
}
var car_pos_x = 0
var car_pos_y = 0

func start() {
    for i := range w4.FRAMEBUFFER {
        w4.FRAMEBUFFER[i] = 1 | (1 << 2) | (1 << 4) | (1 << 6)
    }
}

//go:export update
func update() {
    *w4.DRAW_COLORS = 2

    var gamepad = *w4.GAMEPAD1
    w4.PALETTE[2] = 0xff0000
    *w4.DRAW_COLORS = 0x31
    w4.Blit(&rocket_unthrust[0], car_pos_x, car_pos_y, 8, 13, w4.BLIT_1BPP)

    if gamepad & w4.BUTTON_LEFT != 0 {
        if (car_pos_x > 0) {
            car_pos_x--
        }
        w4.Blit(&rocket_unthrust[0], car_pos_x, car_pos_y, 8, 13, w4.BLIT_1BPP)
    }
    w4.Text("<", 0, 152)
    if gamepad & w4.BUTTON_RIGHT != 0 {
        if (car_pos_x < 150) {
            car_pos_x++
        }
        w4.Blit(&rocket_unthrust[0], car_pos_x, car_pos_y, 8, 13, w4.BLIT_1BPP)
    }
    w4.Text(">", 152, 152)
    if gamepad & w4.BUTTON_UP != 0 {
        if (car_pos_y > 0) {
            car_pos_y--
        }
        w4.Blit(&rocket[0], car_pos_x, car_pos_y, 8, 13, w4.BLIT_1BPP)
    }

    if gamepad & w4.BUTTON_DOWN != 0 {
        if (car_pos_y < 150) {
            car_pos_y++
        }
        w4.Blit(&rocket_unthrust[0], car_pos_x, car_pos_y, 8, 13, w4.BLIT_1BPP)
    }

}
