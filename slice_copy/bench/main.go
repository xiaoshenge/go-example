package main

func InitSlice() []int {
	s := make([]int , 0, 100)
	for i :=0;i< 100;i++{
		s = append(s, i)
	}
	return s
}

func SliceWithCopy(){
	s := InitSlice()
	for j:=0;j<1000;j++{
		for i:=0;i<100;i++{
			copy(s,s[1:])
			s = s[:len(s)-1]
		}
		for i:=0;i<100;i++{
			s = append(s,i)
		}
	}
	

}

func SliceWithSlice()  {
	s := InitSlice()
	for j:=0;j<1000;j++{

		for i:=0;i<100;i++{
			s = s[1:]
		}
		for i:=0;i<100;i++{
			s = append(s,i)
		}
	}
}