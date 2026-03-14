package main

import ( 
	"fmt" 
)

func main() {
	arr := []int {1,2,4,5,6}
	arr = reverse(arr)
	fmt.Println(arr)
	fmt.Println(findAnagrams("cbaebabacd","abc"))
}

func reverse(arr []int) []int{
	beg:= 0
	end:= len(arr)-1
	for end>beg{
		arr[beg], arr[end] = arr[end], arr[beg]
		beg++
		end--
	}
	return arr
}

func findAnagrams(s string, p string) []int {
    sliceS:= []rune(s)
    sliceP:= []rune(p)
    var countP, countW [26]int
    l:=0
    r:=0
    index:=[]int{}
    for _,ch := range(sliceP){
        countP[ch-'a']++
    }

    for r<=len(sliceS) && l<(len(sliceS)-len(sliceP)+1){
		if(r-l)<len(sliceP){
            countW[sliceS[r]-'a']++
            r++
        } else{
            if countP==countW{
                index = append(index,l)
            }
			countW[sliceS[l]-'a']--
			l++
        }
    }
    return index
}