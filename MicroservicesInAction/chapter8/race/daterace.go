package singlegoroutine

var realNum = make(chan int)
var delta = make(chan int)

func SetNumber(n int) {
	realNum <- n
}
func ChangeByDelta(d int) {
	delta <- d
}
func GetNumber() int {
	return <- realNum
}
func monitor() {
	var i int
	for {
		select {
		case i = <- realNum:
		case d := <- delta:
			i += d
		case realName <- i:
		}
	}
}
func init() {
	go monitor()
}
