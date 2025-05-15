package main
import "fmt"
const NMAX int = 100
func main(){
	var n, i, j, idx, compalier int
	var A[NMAX]int
	fmt.Scan(&n)
	for i = 0; i< n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		idx = i
		for j = i+1; j < n; j++{
			if A[idx] > A[j]{
				idx = j
			}
		}
		compalier = A[i]
		A[i] = A[idx]
		A[idx] = compalier
	}
	for i = 0; i < n; i++{
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

