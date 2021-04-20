package main

import "fmt"

func main() {



	fmt.Println(HasUniqueChar("this is not unique"))
}









/*--------------------------------------------------------------------------------------------------
Write a program to determine if a string contains only unique characters. Return true if it does and false otherwise.
The string may contain any of the 128 ASCII characters. Characters are case-sensitive, e.g. 'a' and 'A' are considered different characters.

--------------------------------------------------------------------------------------------------*/
func HasUniqueChar (str string) bool {

	isUnique := true
	for i := 0; i < len(str); i++{
		for j := 0; j < len(str); j++{

		}
	}
	fmt.Println(string(str[0]))
	return isUnique

}







/*--------------------------------------------------------------------------------------------------

Your task, is to create NxN multiplication table, of size provided in parameter.
for example, when given size is 3:

the return value should be: [[1,2,3],[2,4,6],[3,6,9]]

--------------------------------------------------------------------------------------------------*/


func MultiplicationTable(size int) [][]int {

	returnArr := make([][]int,size)
	for i := 0; i < size; i++{
		returnArr[i] = make([]int, size)
		for j :=0; j < size; j++{
			returnArr[i][j] = (i+1)*(j+1)
		}
	}
	return returnArr
}





/*--------------------------------------------------------------------------------------------------
Create a function add(n)/Add(n) which returns a function that always adds n to any number
--------------------------------------------------------------------------------------------------*/

func Add(i int) func(int)int {
	return func(x int)int{
		return x + i
	}
}






/* ---------------------------------------------------------------------------------------------------
Write a function that when given a number >= 0, returns an Array of ascending length subarrays.

pyramid(0) => [ ]
pyramid(1) => [ [1] ]
pyramid(2) => [ [1], [1, 1] ]
pyramid(3) => [ [1], [1, 1], [1, 1, 1] ]

Note: the subarrays should be filled with 1s
----------------------------------------------------------------------------------------------------- */
func Pyramid(n int) [][] int{

	returnArr := make([][]int,n)

	for i := 0; i < n; i++{
		returnArr[i] = make([]int, i+1)
		for j := 0; j < len(returnArr[i]);

		j++{
			returnArr[i][j] = 1
		}
	}

	return returnArr
}
