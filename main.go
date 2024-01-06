// https://stackoverflow.com/questions/58793857/robotgo-for-windows-10-fatal-error-zlib-h-no-such-file-or-directory
// https://sourceforge.net/projects/mingw-w64/files/
package main

import (
	"math"

	"os"
	"time"

	"github.com/andybrewer/mack"
	"github.com/go-vgo/robotgo"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func main() { mainthread.Init(mainThread) }
func mainThread() {
	unspacex := func() {}

	mousePushX, mousePushY := pusher(
		func(dx int, dy int, k int) {
			// println("M", dx, dy)
			if dx != 0 || dy != 0 {
				robotgo.MoveRelative(dx, dy)
			}
		},
		1080, 1080,
		8192, 8192,
		false,
	)
	arrowPushX, arrowPushY := pusher(
		func(dx int, dy int, k int) {
			println("arrow", dx, dy, k)
			mods := modsDecode(k)
			tap := func(o string, t int) {
				for i := 0; i < t; i++ {
					if len(mods) == 0 {
						robotgo.KeyTap(o)
					} else {
						robotgo.KeyTap(o, mods)
					}
				}
			}
			if dx > 0 {
				tap("right", dx)
			}
			if dx < 0 {
				tap("left", -dx)
			}
			if dy > 0 {
				tap("up", dy)
			}
			if dy < 0 {
				tap("down", -dy)
			}
		},
		96, 96,
		2048, 2048,
		true,
	)
	wheelPushX, wheelPushY := pusher(
		func(dx int, dy int, k int) {
			// println("W", dx, dy)
			if dx != 0 || dy != 0 {
				if k&int(hotkey.ModShift) != 0 {
					robotgo.Scroll(dx, 0)
				} else {
					robotgo.Scroll(0, -dy)

				}

			}
			// robotgo.MoveRelative(dx, dy)
			// mods := modsDecode(k)
		},
		192, 192,
		2147483647, 2147483647,
		false,
	)
	// downer
	// ()
	for {
		escaped := false
		unspacex = spacex(func() {
			unspacex()
			robotgo.KeyTap("space")
			escaped = true
		},
			mousePushX, mousePushY,
			arrowPushX, arrowPushY,
			wheelPushX, wheelPushY,
		)
		for !escaped {
			time.Sleep(time.Millisecond * time.Duration(100))
		}
	}
	// return unspacex()
}
func spacex(tap func(),
	mousePushX func(float64, int), mousePushY func(float64, int),
	arrowPushX func(float64, int), arrowPushY func(float64, int),
	wheelPushX func(float64, int), wheelPushY func(float64, int),
) func() {
	unclxedit := func() {}
	unclxdesktop := func() {}
	unclxmouse := func() {}
	acted := false
	act := func() { acted = true }
	unreg := myreg([]hotkey.Modifier{}, hotkey.KeySpace,
		func() {
			acted = false
			unclxedit = clxedit(act, arrowPushX, arrowPushY)
			unclxdesktop = clxdesktop(act)
			unclxmouse = clxmouse(act,
				mousePushX, mousePushY,
				wheelPushX, wheelPushY,
			)
		},
		func() {
			unclxedit()
			unclxdesktop()
			unclxmouse()
			if !acted {
				tap()
			}
		})
	return unreg
}

func clxmouse(
	act func(),
	mousePushX func(float64, int), mousePushY func(float64, int),
	wheelPushX func(float64, int), wheelPushY func(float64, int),
) func() {
	unregs := []func(){
		modsreg(hotkey.KeyA,
			func(k int, d int) {
				act()
				println("mousePushX")
				mousePushX(-1, 0)
			},
			func() {
				println("mousePushX")
				mousePushX(0, 0)
			}),
		modsreg(hotkey.KeyD,
			func(k int, d int) {
				act()
				println("mousePushX")
				mousePushX(1, 0)
			},
			func() {
				println("mousePushX")
				mousePushX(0, 0)
			}),
		modsreg(hotkey.KeyW,
			func(k int, d int) {
				act()
				println("mousePushY")
				mousePushY(-1, 0)
			},
			func() {
				println("mousePushY")
				mousePushY(0, 0)
			}),
		modsreg(hotkey.KeyS,
			func(k int, d int) {
				act()
				println("mousePushY")
				mousePushY(1, 0)
			},
			func() {
				println("mousePushY")
				mousePushY(0, 0)
			}),
		modsreg(hotkey.KeyR,
			func(k int, d int) { act(); wheelPushY(-10, 0) },
			func() { wheelPushY(0, 0) }),
		modsreg(hotkey.KeyF,
			func(k int, d int) { act(); wheelPushY(10, 0) },
			func() { wheelPushY(0, 0) }),
		// TODO HOLD
		modsreg(hotkey.KeyE,
			func(k int, taps int) {
				act()
				robotgo.Toggle("left")
			},
			func() {
				act()
				robotgo.Toggle("left", "up")
			}),
		modsreg(hotkey.KeyQ,
			func(k int, taps int) {
				act()
				robotgo.Toggle("right")
			},
			func() {
				act()
				robotgo.Toggle("right", "up")
			}),
	}
	return func() {
		for _, unreg := range unregs {
			unreg()
		}
	}
}

func clxdesktop(act func()) func() {
	unregs := []func(){
		modsreg(hotkey.KeyX,
			func(k int, taps int) { act(); robotgo.KeyTap("w", "command ") },
			func() {},
		),
		// kVK_ANSI_1
		modsreg(0x12,
			func(k int, taps int) {
				// robotgo.KeyTap("left", "control");
				mack.Tell("System Events", "key code 123 using {control down}")
				act()
			},
			func() {},
		),
		// kVK_ANSI_2
		modsreg(0x13,
			func(k int, taps int) {
				// robotgo.KeyTap("right", "control");
				mack.Tell("System Events", "key code 124 using {control down}")
				act()
			},
			func() {},
		),
		// kVK_ANSI_Slash
		modsreg(0x2C,
			func(k int, taps int) {
				// robotgo.KeyTap("right", "control");
				os.Exit(0)
				// touchMain()
				act()
			},
			func() {},
		),
		// kVK_ANSI_Backslash
		modsreg(0x2A,
			func(k int, taps int) {
				// robotgo.KeyTap("right", "control");
				// touchMain()
				os.Exit(0)
				act()
			},
			func() {},
		),
	}
	return func() {
		for _, unreg := range unregs {
			unreg()
		}
	}
}

func clxedit(act func(),
	arrowPushX func(float64, int), arrowPushY func(float64, int),
) func() {
	unregs := []func(){
		turboKey(hotkey.KeyT, "delete", act),
		turboKey(hotkey.KeyG, "enter", act),
		//
		turboKey(hotkey.KeyH, "left", act),
		turboKey(hotkey.KeyJ, "down", act),
		turboKey(hotkey.KeyK, "up", act),
		turboKey(hotkey.KeyL, "right", act),
		//
		turboKey(hotkey.KeyY, "home", act),
		turboKey(hotkey.KeyO, "end", act),
		turboKey(hotkey.KeyU, "pagedown", act),
		turboKey(hotkey.KeyI, "pageup", act),
		//
		turboTap(hotkey.KeyP, func(k int, taps int) {
			robotgo.KeyTap("tab", "shift")
		}, act),
		turboTap(hotkey.KeyN, func(k int, taps int) {
			robotgo.KeyTap("tab")
		}, act)}
	return func() {
		for _, unreg := range unregs {
			unreg()
		}
	}
}
func turboKey(i hotkey.Key, o string, act func()) func() {
	return turboTap(i, func(k int, taps int) {
		mods := modsDecode(k)
		if len(mods) == 0 {
			robotgo.KeyTap(o)
		} else {
			robotgo.KeyTap(o, mods)
		}
	}, act)
}

func turboTap(i hotkey.Key, tap func(k int, taps int), act func()) func() {
	taps := 0
	unreg := modsreg(i,
		func(kk int, k int) {
			taps = 0
			act()
			go func() {
				for taps >= 0 {
					tap(k, taps)
					ms := math.Max(0, 120*(math.Pow(0.5, 0.5*float64(taps))))
					time.Sleep(time.Millisecond * time.Duration(ms))
					taps++
				}
			}()
		},
		func() { taps = -100 })
	return func() {
		taps = -100
		unreg()
	}
}

// func turboMove(i hotkey.Key, move func(k int, distance int), act func()) func() {
// 	t := int64(0)
// 	unreg := modsreg(i,
// 		func(kk int, k int) { vb bb
// 			if t != 0 {
// 				return
// 			}
// 			t = time.Now().UnixNano() / int64(time.Millisecond)
// 			act()
// 			go func()

// 				tracking := 0
// 				for t != 0 {
// 					ct := time.Now().UnixNano() / int64(time.Millisecond)
// 					dt := (ct - t)
// 					P := 0.2
// 					B := 1.1
// 					E := 0.1
// 					distance := math.Max(0.0, math.Min(
// 						P*(math.Pow(B, 1+E*float64(dt))),
// 						2147483647.0))
// 					diff := int(distance) - tracking
// 					tracking += diff
// 					move(k, diff)
// 					time.Sleep(time.Millisecond * time.Duration(1))
// 				}
// 			}()
// 		},
// 		func() { t = 0 })
// 	return func() {
// 		t = 0
// 		unreg()
// 	}
// }

/*
	enum {
		kVK_ANSI_A                    = 0x00,
		kVK_ANSI_S                    = 0x01,
		kVK_ANSI_D                    = 0x02,
		kVK_ANSI_F                    = 0x03,
		kVK_ANSI_H                    = 0x04,
		kVK_ANSI_G                    = 0x05,
		kVK_ANSI_Z                    = 0x06,
		kVK_ANSI_X                    = 0x07,
		kVK_ANSI_C                    = 0x08,
		kVK_ANSI_V                    = 0x09,
		kVK_ANSI_B                    = 0x0B,
		kVK_ANSI_Q                    = 0x0C,
		kVK_ANSI_W                    = 0x0D,
		kVK_ANSI_E                    = 0x0E,
		kVK_ANSI_R                    = 0x0F,
		kVK_ANSI_Y                    = 0x10,
		kVK_ANSI_T                    = 0x11,
		kVK_ANSI_1                    = 0x12,
		kVK_ANSI_2                    = 0x13,
		kVK_ANSI_3                    = 0x14,
		kVK_ANSI_4                    = 0x15,
		kVK_ANSI_6                    = 0x16,
		kVK_ANSI_5                    = 0x17,
		kVK_ANSI_Equal                = 0x18,
		kVK_ANSI_9                    = 0x19,
		kVK_ANSI_7                    = 0x1A,
		kVK_ANSI_Minus                = 0x1B,
		kVK_ANSI_8                    = 0x1C,
		kVK_ANSI_0                    = 0x1D,
		kVK_ANSI_RightBracket         = 0x1E,
		kVK_ANSI_O                    = 0x1F,
		kVK_ANSI_U                    = 0x20,
		kVK_ANSI_LeftBracket          = 0x21,
		kVK_ANSI_I                    = 0x22,
		kVK_ANSI_P                    = 0x23,
		kVK_ANSI_L                    = 0x25,
		kVK_ANSI_J                    = 0x26,
		kVK_ANSI_Quote                = 0x27,
		kVK_ANSI_K                    = 0x28,
		kVK_ANSI_Semicolon            = 0x29,
		kVK_ANSI_Backslash            = 0x2A,
		kVK_ANSI_Comma                = 0x2B,
		kVK_ANSI_Slash                = 0x2C,
		kVK_ANSI_N                    = 0x2D,
		kVK_ANSI_M                    = 0x2E,
		kVK_ANSI_Period               = 0x2F,
		kVK_ANSI_Grave                = 0x32,
		kVK_ANSI_KeypadDecimal        = 0x41,
		kVK_ANSI_KeypadMultiply       = 0x43,
		kVK_ANSI_KeypadPlus           = 0x45,
		kVK_ANSI_KeypadClear          = 0x47,
		kVK_ANSI_KeypadDivide         = 0x4B,
		kVK_ANSI_KeypadEnter          = 0x4C,
		kVK_ANSI_KeypadMinus          = 0x4E,
		kVK_ANSI_KeypadEquals         = 0x51,
		kVK_ANSI_Keypad0              = 0x52,
		kVK_ANSI_Keypad1              = 0x53,
		kVK_ANSI_Keypad2              = 0x54,
		kVK_ANSI_Keypad3              = 0x55,
		kVK_ANSI_Keypad4              = 0x56,
		kVK_ANSI_Keypad5              = 0x57,
		kVK_ANSI_Keypad6              = 0x58,
		kVK_ANSI_Keypad7              = 0x59,
		kVK_ANSI_Keypad8              = 0x5B,
		kVK_ANSI_Keypad9              = 0x5C
	  };
*/
func pusher(
	ctrl func(dx int, dy int, k int),
	px float64, py float64,
	maxVx float64, maxVy float64,
	rush bool,
) (
	func(fx float64, kx int), func(fy float64, ky int),
) {
	now := func() float64 {
		t := float64(time.Now().UnixMicro()) / float64(1000000)
		// println(int64(t)) // seconds
		return t
	}
	// pusher1d := func() (func(dt float64), func()) {
	// 	x := float64(0)
	// 	v := float64(0)
	// 	a := float64(0)
	// 	return func(dt float64) {
	// 		v = dt * a
	// 		x = dt * v
	// 	},
	// 	func(f float64) { a = f },
	// };
	// iterX, pusherX := pusher1d()
	x := float64(0)
	vx := float64(0)
	ax := float64(0)
	fx := float64(0)
	kx := int(0) // mod key

	y := float64(0)
	vy := float64(0)
	ay := float64(0)
	fy := float64(0)
	ky := int(0) // mod key
	t := float64(0)
	go func() {
		escaped := false
		for !escaped {
			ct := now()
			dt := (ct - t)
			t = ct

			pow := float64(1.8)

			ax = fx
			vx = vx + ax*math.Pow(dt*px, pow*1.2)*0.5
			vx = math.Max(-maxVx, math.Min(vx, maxVx))
			frictionx := math.Pow(0.95, dt*300)
			if ax == 0 {
				vx = vx * frictionx
			}
			x1 := x + vx*dt
			dx := int(x1 - x)
			x = x + float64(dx)

			ay = fy
			vy = vy + ay*math.Pow(dt*py, pow*1.2)*0.5
			vy = math.Max(-maxVy, math.Min(vy, maxVy))
			frictiony := math.Pow(0.95, dt*300)
			if ay == 0 {
				vy = vy * frictiony
			}
			y1 := y + vy*dt
			dy := int(y1 - y)
			y = y + float64(dy)
			if dx != 0 || dy != 0 {
				ctrl(dx, dy, kx|ky)
			}
			// model debug
			// println(dx, int(x), int(vx), int(ax), int(dt*1000))
			time.Sleep(time.Millisecond * time.Duration(10))
		}
	}()
	return func(pfx float64, pkx int) {
			// println("pf", pkx, pfx)
			if pfx != 0 {
				t = now()
			}
			fx = pfx
			kx = pkx
		},
		func(pfy float64, pky int) {
			// println("pf", pky, pfy)
			if pfy != 0 {
				t = now()
			}
			fy = pfy
			ky = pky
		}
}

// func touchMain() {
// 	f, err := os.OpenFile("main.go",
// 		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		// log.Println(err)
// 	}
// 	defer f.Close()
// 	if _, err := f.WriteString("// touched\n"); err != nil {
// 		// log.Println(err)
// 	}
// }
