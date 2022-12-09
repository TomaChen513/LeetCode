package hashset

// 242  Valid Anagram
func isAnagram(s string, t string) bool {
    if len(s)!=len(t){
        return false
    }
	letterMap:=make(map[byte]int)
	for i := 0; i < len(s); i++ {
		letterMap[s[i]]+=1
	}
	for i := 0; i < len(t); i++ {
		if letterMap[t[i]]-=1;letterMap[t[i]]<0 {
			return false
		}
	}
	return true
}

// 349  Intersection of Two Arrays
func intersection(nums1 []int, nums2 []int) []int {
	res:=make([]int,0)
	numMap:=make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		numMap[nums1[i]]=struct{}{}
	}
	for i := 0; i < len(nums2); i++ {
		if _,ok:=numMap[nums2[i]];ok {
			delete(numMap,nums2[i])
			res = append(res, nums2[i])
		}
	}
	return res
}

// 202 Happy Number
func isHappy(n int) bool {
	numMap:=make(map[int]struct{})
	getNextNum:=func (n int)  int{
		sum:=0
		for n>0 {
			sum+=(n%10)*(n%10)
			n/=10
		}
		return sum
	}
	for  {
		_,ok:=numMap[n]
		if ok {
			return false
		}
        numMap[n]=struct{}{}
		n=getNextNum(n)
		if n==1 {
			return true
		}
	}
}

// 1 Two Sum
func twoSum(nums []int, target int) []int {
	numMap:=make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if key,ok:=numMap[target-nums[i]];ok{
			return []int{key,i}
		}else{
			numMap[nums[i]]=i
		}
	}
	return []int{0,0}
}

