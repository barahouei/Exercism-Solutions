package clock

import "fmt"

type Clock int

const (
	hourMinutes = 60
	dayMinutes  = 1440
)

func New(h, m int) Clock {
	hm := ((h*hourMinutes+m)%dayMinutes + dayMinutes) % dayMinutes

	return Clock(hm)
}

func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/hourMinutes, c%hourMinutes)
}
