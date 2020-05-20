package main

/*
Необходимо написать функцию func Merge2Channels(f func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int)
в package main. Описание ее работы:
n раз сделать следующее
	прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
	вычислить f(x1) + f(x2)
	записать полученное значение в out
Функция Merge2Channels должна быть неблокирующей, сразу возвращая управление.
Функция f может работать долгое время, ожидая чего-либо или производя вычисления.
*/

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		for ; n > 0; n-- {
			x1 := <-in1
			x2 := <-in2
			out <- f(x1) + f(x2)
		}
	}()
}
