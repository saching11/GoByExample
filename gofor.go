package main 
import "fmt"

func main() {
	i := 3
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	for j := range 3 {
		fmt.Println(j)
	}

	for {
		fmt.Println("For loop without condition")
		break
	}

	for n := range 6 {
		if n % 2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}