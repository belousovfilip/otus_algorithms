package visualization

import (
	tm "github.com/buger/goterm"
	"time"
)

type Sorting struct {
	slice    []int
	startY   int
	startX   int
	speed    time.Duration
	maxValue int
}

func InitSorting(l []int, speed time.Duration) *Sorting {
	tm.Clear()

	s := &Sorting{
		slice:  l,
		startY: 2,
		startX: 0,
		speed:  time.Millisecond * (1000 - speed),
	}
	s.setMaxValue()
	return s
}
func (s *Sorting) Compare(a, b int) {
	s.printLine(a, tm.BLUE)
	s.printLine(b, tm.GREEN)
	tm.Flush()
	s.sleep()
}
func (s *Sorting) Swap(a, b int) {
	s.printLine(a, tm.GREEN)
	s.printLine(b, tm.BLUE)
	tm.Flush()
	s.sleep()
}

func (s *Sorting) Print() {
	for i, _ := range s.slice {
		s.printLine(i, tm.WHITE)
	}
	tm.Flush()
	s.sleep()
}

func (s *Sorting) sleep() {
	time.Sleep(s.speed)
}

func (s *Sorting) setMaxValue() {
	var m int
	for _, v := range s.slice {
		if v > m {
			m = v
		}
	}
	s.maxValue = m
}

func (s *Sorting) clearLine(i int) {
	for j := 0; j < s.maxValue+10; j++ {
		tm.MoveCursor(j, s.startY+i)
		tm.Print(" ")
	}
}

func (s *Sorting) printLine(i int, color int) {
	s.clearLine(i)
	y := s.startY + i
	for j := 1; j <= s.slice[i]; j++ {
		tm.MoveCursor(j+s.startX, y)
		tm.Print(tm.Color("⬛", color))
	}
	tm.MoveCursor(s.slice[i]+1+s.startX, y)
	tm.Print(" ", s.slice[i])
	tm.Flush()
}
