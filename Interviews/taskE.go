package main

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	for ; n > 0; n-- {

		x1 := <-in1
		x2 := <-in2

		go func(x1, x2 int) {
			out <- f(x1) + f(x2)
		}(x1, x2)
	}
}
