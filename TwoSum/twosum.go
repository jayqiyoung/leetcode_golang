//Two Sum 
package main 
import "fmt"

func twoSun(nums []int ,target int) []int {
	var result []int 
	for i := 0 ;i < len(nums); i++{
		for j := i + 1; j < len(nums);j++{
			if nums[i] == target - nums[j]{
				result = append(result, i)
				result = append(result,j)
			}
		}
	}
	return result 
}

func main(){
	nums := []int{1,7,8,11,15}
	target := 9
	result := twoSun(nums,target)
	fmt.Println("result : \n",result)
}
