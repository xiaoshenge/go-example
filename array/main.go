package main

func main()  {
	arr := [...]int{0,1,2,3,4,5}

	for i := 0; i < 10000; i++ {
		go func ()  {
			for j := 0; j < 5; j++ {
				arr[j] = j
			}
		}()
	}
}