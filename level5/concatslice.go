package piscine

func ConcatSlice(slice1, slice2 []int) []int {
    if len(slice1) == 0 && len(slice2) == 0 {
        return nil // Return nil if both slices are empty
    }

    // Create a new slice with a length equal to the sum of both slices
    result := make([]int, len(slice1) + len(slice2))
    
    // Copy elements from the first slice
    copy(result, slice1)
    
    // Copy elements from the second slice
    copy(result[len(slice1):], slice2)
    
    return result
}
