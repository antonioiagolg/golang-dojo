package main

func Sum(nums []int) int {

    var sum int
    for _, i := range nums {
        sum += i
    }
    return sum
}


func SumAll(lis ...[]int) []int {
    
    result := []int{}
    for _, aList := range lis {
        result = append(result, Sum(aList))
    }

   return result
}
