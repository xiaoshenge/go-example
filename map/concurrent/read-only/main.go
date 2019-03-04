package main

func main()  {
	m := map[int]int{
		1:1,
		2:2,
	}

	for i := 0; i < 100; i++ {
		go func ()  {
			for j := 0; j < 100000; j++ {
				_ = m[1]
			}
		}()
	}
}