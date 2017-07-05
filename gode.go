package GoDE

const (
	ShiftMask   = 1 << 0
	LockMask    = 1 << 1
	ControlMask = 1 << 2
	Mod1Mask    = 1 << 3
	Mod2Mask    = 1 << 4
	Mod3Mask    = 1 << 5
	Mod4Mask    = 1 << 6
	Mod5Mask    = 1 << 7
)

const (
	keyMask    = (1 << 24) - 1
	buttonMask = (1 << 8) - 1
)

type Keyboard uint32

func (k Keyboard) Key() uint32 {
	return uint32(k & keyMask)
}

func (k Keyboard) Mask() uint8 {
	return uint8(k >> 24)
}

type Point struct {
	X uint16
	Y uint16
}

type Mouse struct {
	buttons  uint16
	Position Point
}

func (m Mouse) Button() uint8 {
	return uint8(m.buttons & buttonMask)
}

func (m Mouse) Mask() uint8 {
	return uint8(m.buttons >> 8)
}
