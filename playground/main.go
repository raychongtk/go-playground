package main

import (
	"fmt"
	myArray "playground/datastructure/array"
	myQueue "playground/datastructure/queue"
	mySet "playground/datastructure/set"
	myMath "playground/math"
	myUrlShortener "playground/shortenurl"
	myText "playground/text"
	myUtil "playground/util"
)

func main() {
	numbers := []int{5, 2, 1, 4, 5, 6, 8, 9, 0, 10}
	singleNumbers := []int{4, 1, 2, 1, 2}
	singleNumber := myMath.SingleNumber(singleNumbers)
	min := myMath.FindMin(numbers)
	max := myMath.FindMax(numbers)
	top1, top2 := myMath.TopTwoElements(numbers)
	firstTop, secondTop, thirdTop := myMath.TopThreeElements(numbers)
	runningSum := myMath.RunningSum(numbers)
	sum := myMath.Sum(numbers)
	fmt.Printf("min=%d, max=%d\n", min, max)
	fmt.Println(top1, top2)
	fmt.Println(firstTop, secondTop, thirdTop)
	fmt.Println(runningSum)
	fmt.Println(sum)
	fmt.Println("single number:", singleNumber)

	text := "012345678999995"
	text2 := "       0123456789abcd       "
	fmt.Println(myText.IsNumber(text))
	fmt.Println(myText.IsNumber(text2))

	var str string
	str = "aaaa"
	fmt.Println(myText.IsEmpty(str))

	// trim spaces
	fmt.Println(myText.TrimLeft(text2))
	fmt.Println(myText.TrimRight(text2))

	// remove duplicated char
	fmt.Println(myText.RemoveDuplicateCharacter(text))

	// swapping value
	num1 := 10
	num2 := 20
	myUtil.Swap(&num1, &num2)
	fmt.Println(num1, num2)

	// join string
	var elements []string
	elements = append(elements, "test", "test1", "test2", "test3", "test4", "test5", "test6", "test7")
	fmt.Println(myText.Join(elements, ", "))

	fmt.Println(myUtil.Mask(text))
	fmt.Println(myUtil.Mask("1"))
	fmt.Println(myUtil.MaskLeft(text))
	fmt.Println(myUtil.MaskLeft("1"))
	fmt.Println("containing duplicate:", myText.IsContainingDuplicateCharacter("apple"))
	fmt.Println("containing duplicate:", myText.IsContainingDuplicateCharacter("ray"))
	fmt.Println("is palindrome:", myText.IsPalindrome("abccba"))
	fmt.Println(myText.ReverseString("ray"))
	fmt.Println(myArray.SameArray([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}))
	fmt.Println(myArray.SameArray([]int{1, 2, 4, 4, 5}, []int{1, 2, 3, 4, 5}))

	data := []int{1, 2, 4, 4, 5}
	myArray.ReverseArray(data)
	fmt.Println(data)

	myQueue.Add("test")
	myQueue.Add("test1")
	myQueue.Add("test2")
	myQueue.Poll()

	queue := myQueue.GetQueue()
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	mySet.Put("test")
	fmt.Println(mySet.Contains("test"))
	fmt.Println(mySet.Size())
	mySet.Remove("test")
	fmt.Println(mySet.Size())

	// url shortener
	shortener := myUrlShortener.NewShortener()
	fmt.Println(shortener.Encode("https://golang.org/doc/tutorial/handle-errors"))
	decodedUrl, err := shortener.Decode("https://short-url.com/d50fb44cc1")
	if err != nil {
		panic("url not found")
	}
	fmt.Println(decodedUrl)
}
