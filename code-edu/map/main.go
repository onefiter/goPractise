package main


func main() {
	m := make(map[int]int)

	go func() {
		for {
			_ = m[1]
		}
	}()


	go func() {
		for {
			m[2] = 2
		}
	}()

	select {}
}