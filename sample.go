package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
	"sort"
)

type Dog struct{
	Name string
	Age int
}
func (d *Dog) getName() string{
	return d.Name
}


func main() {
	var s string
	var x string
	fmt.Println("Enter a boolean and a string:")
	fmt.Scanln(&s,&x)
	str1,_:=strconv.ParseBool(s)
	fmt.Println(str1,x)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text")
	str,_ := reader.ReadString('\n')
	fmt.Println(str)
	fmt.Println("Enter a number:")
	str,_ = reader.ReadString('\n')
	f,_:=strconv.ParseFloat(strings.TrimSpace(str),8)
	fmt.Println("You entered:",f)
	a:=float64(1)
	fmt.Println(math.Sqrt(a))
	var b *float64
	b=&a
	fmt.Println("value of b:",*b)
	var arr [5]int
	arr[0]=0
	arr[1]=1
	arr[2]=2
	arr[3]=3
	arr[4]=4
	var slice = []string{"Red","Green","Blue"}
	slice=append(slice,"Purple")
	slice=append(slice[1:])
	defer fmt.Println(slice)
	//make is like arraylist - used for memory allocation
	numbers := make([]int,2,5)
	numbers[0]=1
	numbers[1]=-1
	sort.Ints(numbers)
	fmt.Println(numbers)
	map1:=make(map[string]int)
	map1,_ = mapTest(map1)
	for k,v := range map1{
		fmt.Println("Key:",k,"Value:",v)
	}
	maltese:=Dog{Name:"Woofie",Age:2}
	fmt.Println(maltese.getName(),"lives at",&maltese.Name)
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}

func mapTest(input map[string]int)(map[string]int,int){
	input["num3"] = 3
	i,j,a:=0,0,0
	for a<=5{
		input[strconv.Itoa(i)]=j+1
		i++
		j++
		a++
	}
	return input,len(input)
}

func WordCount(s string) map[string]int {
	fields:=strings.Fields(s)
	wordmap:=make(map[string]int)
	for j:=range fields{
		wordmap[fields[j]]=1
	}
	for i:=0;i<len(fields);i++{
		for j:=i+1;j<len(fields);j++{
			if fields[i]==fields[j]{
				wordmap[fields[i]]+=1
			}
		}
	}
	return wordmap
}