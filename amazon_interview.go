package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    
)


/*
 * Complete the 'findRestaurants' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY allLocations
 *  2. INTEGER numRestaurants
 */
 func abs(x float32)float32{
     return absdiffInt(x,0)
 }
 func absdiffInt(x,y float32)float32{
     if x<y{
         return y - x
     }
     return x - y
 }
func root(num int32) int32{
    x1 := (num * 1.0)/2
    x2 := (x1 + (num/ x1))/2
    for((abs(float32(x1-x2)))>=0.00000001){
        x1 = x2
        x2 = (x1 + (num/x1)) /2
    }
    return x2
}
func swapp(a,b *int32){
    *a,*b = *b,*a
}
func sort(arr []int32)[]int32{
    // index := [][]int32{}
    for i:=0;i<len(arr);i++{
        var swapped bool
        for j:=i+1;j<len(arr);j++{
            if arr[i]>arr[j]{
                swapp(&arr[i],&arr[j])
                // index = append(index,arr[i])
                swapped = true
            }
        }
        if !swapped{
            break
        }
    }
    return arr
}
func findRestaurants(allLocations [][]int32, numRestaurants int32) []int32 {
    // Write your code here
    indexes := []int32{}
    
    for _,v := range allLocations{
        t := root(v[0]*v[0]+v[1]+v[1])
        indexes = append(indexes,t)
    }
    arr := sort(indexes)
    var i int32
    result := []int32{}
    for i=0;i<numRestaurants;i++{
        result = append(result,arr[i])  
    }
    return result

}
func main() {