package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

// func Chunk(slice []int, size int){
// 	if size == 0 {
// 		fmt.Println()
// 		return
// 	}
// 	var chunks [][]int
// 	for i := 0; i < len(slice); i += size {
// 		end := i + size
// 		if end > len(slice) {
// 			end = len(slice)
// 		}
// 		chunks = append(chunks, slice[i:end])
// 	}
// 	fmt.Println(chunks)
// }

func Chunk(slice []int, size int) {
	// If the size is 0, print a newline and return
	if size == 0 {
		fmt.Println()
		return
	}

	// Initialize an empty slice of slices to hold the chunks
	var chunk [][]int

	// Iterate over the input slice in steps of 'size'
	for i := 0; i < len(slice); i += size {
		// Calculate the end index for the current chunk
		end := i + size
		// Ensure the end index does not exceed the slice length
		if end > len(slice) {
			end = len(slice)
		}
		// Append the chunk to the chunk list
		chunk = append(chunk, slice[i:end])
	}

	// Print the resulting chunks
	fmt.Println(chunk)
}
