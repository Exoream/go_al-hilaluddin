package main

import "fmt"

func pascalTriangle(numRows int) [][]int {
	result := make([][]int, numRows)
    
    for i := 0; i < numRows; i++ {
        row := make([]int, i+1)
        row[0] = 1
		row[i] = 1
        
        for j := 1; j < i; j++ {
            row[j] = result[i-1][j-1] + result[i-1][j]
        }    
        result[i] = row
    }
    return result
}


func main() {
	result := pascalTriangle(5)
	fmt.Println(result)
	
}