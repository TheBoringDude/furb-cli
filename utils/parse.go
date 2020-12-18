package utils

// ReverseSlice reverses a slice interface.
// Based from: https://stackoverflow.com/questions/28058278/how-do-i-reverse-a-slice-in-go
func ReverseSlice(s []interface{}) []interface{} {
	slice := make([]interface{}, len(s))
    copy(slice, s)

    for i := len(slice)/2 - 1; i >= 0; i-- {
        opp := len(slice) - 1 - i
        slice[i], slice[opp] = slice[opp], slice[i]
    }

    return slice
}