/*package main

import "fmt"

func main() {
	var a = 10
	//fmt.Println(a)
	var b = 20.0
	//fmt.Println(b)
	c := "jaya"
	//fmt.Println(c)
	var s, j = 30, 776
	//fmt.Println(s, j)
	var slice = "chandra"
	slice = "devops"
	//fmt.Println(slice)
	h := 50.0
	//fmt.Println(h)
	var f, k, t, u = 13, 45, 67, 90
	//fmt.Println(f, k, t, u)
	fmt.Println(a, b, c, s, j, slice, h, f, k, t, u)

}

//--------------------------------------------------------------------------------------------------------------------


/*
declare a string
int
float
slice
map
*/
//----------------------------------------------------------------------------------------------------------------------
/*package main

  import "fmt"

  func main() {
  	fmt.Println("enter i value:")
  	var n1 int
  	fmt.Scanln(&n1)
  	fmt.Println("enter j value:")
  	var n2 int
  	fmt.Scanln(&n2)
  	fmt.Println("enter k value:")
  	var n3 int
  	fmt.Scanln(&n3)
  	var x int
  	for i := 1; i <= n1; i++ {
  		for j := 1; j <= n2; j++ {
  			for k := 1; k <= n3; k++ {
  				fmt.Println((i * j) + k)
  				x = x + 1
  			}
  		}
  	}
  	fmt.Println("count=", x)

  }*/

// i=1 j=2  j=10 k=1*2+3 ==3
//----------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	course := "DevOps"
// 	if course == "DevOps" {
// 		fmt.Println("Its True")
// 	} else {
// 		fmt.Println("Its False")
// 	}
// 	for i := 1; i <= 100; i++ {
// 		if i%2 != 0 {
// 			fmt.Println(i)
// 			fmt.Println("oddnumbers")
// 		} else {
// 			fmt.Println("even numbers")
// 		}

// 	}
// 	decimal := 100.00
// 	if decimal == 0 {
// 		fmt.Println("it is float")
// 	} else {
// 		fmt.Println("It is not decimal number")
// 	}
// 	if 9%2 == 0 {
// 		fmt.Println("9 is even")
// 	} else {
// 		fmt.Println(" 9 is odd number")
// 	}
// 	if 200%4 == 0 {
// 		fmt.Println("200.0 divisible by 4")
// 	} else {
// 		fmt.Println("200.0 is not divisible by 4")
// 	}
// 	marks := 100
// 	fmt.Println("Please enter the marks:")
// 	fmt.Scanln(&marks)

// 	if marks >= 45 {
// 		fmt.Println("pass marks")
// 	} else {
// 		fmt.Println("fail marks")
// 	}

// }
//-------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%2 == 0 {
// 			continue
// 			fmt.Println("even numbers", i)

// 		} else {

// 			fmt.Println("odd numbers", i)
// 		}
// 	}
// }
//-------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	age := 22
// 	married := false
// 	fmt.Println(age, married)
// }
//-------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {

// 	// fmt.Print("Enter the Number to Print Even's = ")
// 	// fmt.Scanln(&evnum)

// 	// fmt.Println("Even Numbers from 1 to ", evnum, " are = ")
// 	for i := 1; i <= 100; i++ {
// 		if i%2 == 0 {
// 			fmt.Print(i, "\t")
// 		}
// 	}
// 	fmt.Println()
// }

// switch statement
//
//------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	a := 10
// 	b := 20
// 	switch {
// 	case a+b == 30:
// 		fmt.Println("a+b is 30")
// 	case a-b == 30:
// 		fmt.Println("a+b is 40")
// 	default:
// 		fmt.Println("unknown")
// 	}
// }
//---------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var grade string
// 	fmt.Println("enter the grade")
// 	fmt.Scanln(&grade)
// 	switch {
// 	case grade == "A":
// 		fmt.Println("grade is equal to A")
// 	case grade == "B":
// 		fmt.Println("grade is equal to B")
// 	default:
// 		fmt.Println("unknown")
// 	}
// }
//-----------------------------------------------------------------------------------------------------------------
/*package main

  import "fmt"

  func main() {
  	var number int
  	fmt.Println("enter the number :")
  	fmt.Scanln(&number)
  	switch {
  	case number <= 5:
  		fmt.Println("number correct less than 5 ")

  	case number >= 5:
  		fmt.Println("incorrect graeter than 5")
  	}
  }*/
//-----------------------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var num int
// 	number := []int{}
// 	fmt.Print("Enter a number: ")
// 	for {
// 		fmt.Scanln(&num)

// 		if num == 5 {
// 			break
// 		}
// 		number = append(number, num)
// 	}

// 	fmt.Println("You entered:", number)
// }
//-----------------------------------------------------------------------------------------------------------------
/*package main
  import "fmt"
  func main(){
  	for i:=1;i<=50;i++{
  		if i%2==0{
  			fmt.Print(i,'\t')
  		count :=0
  		}
  		fmt.Print(i)
  	}

  }*/
//-------------------------------------------------------------------------------------------------------------------
/*package main

  import "fmt"

  func main() {
  	var a = 10
  	var b = 13
  	var _ = 34
  	var _ = 79
  	fmt.Println(a, b)
  }*/
//----------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	a := []int{1, 2, 3}
// 	fmt.Println(a)

// }
//-----------------------------------------------------------------------------------------------------------------
// multiple return values

// package main

// import "fmt"

// func hello() (int, int) {
// 	return 6, 29
// }
// func main() {
// 	a, b := hello()
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	_, c := hello()
// 	fmt.Println(c)

// }
//-----------------------------------------------------------------------------------------------------------------

//interview process
//---------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var jobs string
// 	fmt.Println("Hey Is there any vacancies in your company please say yes/no: ")
// 	fmt.Scanln(&jobs)
// 	switch {
// 	case jobs == "yes":
// 		fmt.Println("yes i have vacancies")
// 	case jobs == "no":
// 		fmt.Println("no I don't have any vacancies")
// 		if jobs == "yes" {
// 			fmt.Println("Please send the Resume")
// 		} else {
// 			fmt.Println("no vacancies for this time", jobs)
// 		}
// 	}
// }
//---------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func add(a, b int) int {
// 	sum := a + b
// 	return sum
// }
// func main() {
// 	fmt.Println("addition the numbers :", add(79, 4))

// }
//------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	switch dayOfWeek := 6; dayOfWeek {
// 	case 1:
// 		fmt.Println("Monday")
// 	case 2:
// 		fmt.Println("Tuesday")
// 	case 3:
// 		fmt.Println("Wednesday")
// 	case 4:
// 		fmt.Println("Thursday")
// 	case 5:
// 		fmt.Println("Friday")
// 	case 6:
// 		{
// 			fmt.Println("Saturday")
// 			fmt.Println("Weekend. Yaay!")
// 		}
// 	case 7:
// 		{
// 			fmt.Println("Sunday")
// 			fmt.Println("Weekend. Yaay!")
// 		}
// 	default:
// 		fmt.Println("Invalid day")
// 	}
// }
//----------------------------------------------------------------------------------------------------------------

// // Golang program to make a Simple
// // Calculator using Switch Case
// package main

// // fmt package provides the function to print anything
// import "fmt"

// // start the main() function
// func main() {

// 	// Declare amd initialize the variables
// 	var number1 int = 20
// 	var number2 int = 10
// 	var choice int = 0

// 	// choice of the input calculation
// 	var x int // the result variable
// 	fmt.Println("number 1 = ", number1, "\nnumber 2 =", number2)
// 	fmt.Println(" choice 1: Addition of the two numbers")
// 	fmt.Println(" choice 2: Subtraction of the two numbers")
// 	fmt.Println(" choice 3: Multiplication of the two numbers")
// 	fmt.Println(" choice 4: Division of the two numbers")
// 	fmt.Scanln(&choice)
// 	// print the choice of calculation using switch case
// 	switch choice {
// 	case 1:
// 		x = number1 + number2
// 		fmt.Printf("Addition of the two numbers is: %d", x)
// 	case 2:
// 		x = number1 - number2
// 		fmt.Printf("Subtraction of the two numbers is: %d", x)
// 	case 3:
// 		x = number1 * number2
// 		fmt.Printf("Multiplication of the two numbers is: %d", x)
// 	case 4:
// 		x = number1 / number2
// 		fmt.Printf("Division of the two numbers is: %d", x)
// 	default:
// 		fmt.Println("Invalid number")
// 	}
// 	// Print the result using built-in function fmt.Println()
// }

//------------------------------------------------------------------------------------------------------------------

//interview process program

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println("Hello!  I am looking for a job in our company?")
// 	fmt.Print("Please answer 'yes' or 'no': ")
// 	answer, _ := reader.ReadString('\n')
// 	answer = strings.TrimSpace(answer)

// 	if strings.ToLower(answer) == "yes" {
// 		fmt.Println("We have vacancies in our company.")
// 		fmt.Println("Please send your resume")
// 		fmt.Println("your resume is shortlisted, we will schedule an interview.")
// 		fmt.Println("after complete your interview you are selected or not i will decide")
// 	} else if strings.ToLower(answer) == "no" {
// 		fmt.Println("Thank you for your interest, but we currently do not have any vacancies.")
// 	} else {
// 		fmt.Println("Sorry, I did not understand your response. Please answer 'yes' or 'no'.")
// 	}
// }
//-----------------------------------------------------------------------------------------------------------------

// package main

// import "fmt"

// func main() {

// 	var jobs string
// 	fmt.Println("Is there any vacancies in your company yes or no")
// 	fmt.Scanln(&jobs)
// 	switch string(jobs) {
// 	case "yes":
// 		fmt.Println("We have vacancies in our company.")
// 		fmt.Println("Please send your resume .")
// 		fmt.Println("If your resume is shortlisted, we will schedule the interview.")
// 	case "no":
// 		fmt.Println("Thank you for your interest, but we currently do not have any vacancies.")
// 	default:
// 		fmt.Println("Sorry, I can't  understand . Please answer 'yes' or 'no'.")
// 	}
// }

//---------------------------------------------------------------------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	var patient string
// 	var age int
// 	fmt.Println("whats your name:")
// 	fmt.Scanln(&patient)
// 	fmt.Println("Please enter your age")
// 	fmt.Scanln(&age)
// 	fmt.Println("what type of disease")
// 	fmt.Println("select disease")
// 	choice := " select option"
// 	fmt.Println(" choice 1: cardiologist")
// 	fmt.Println(" choice 2: Alopecia Areata")
// 	fmt.Println(" choice 3: Ankylosing Spondylitis ")
// 	fmt.Println(" choice 4: Arthritis")
// 	fmt.Println(" choice 5: Atopic Dermatitis")
// 	fmt.Scanln(&choice)
// 	if choice == "cardiologist" {
// 		fmt.Println("you are suffering from this disease=,choice")
// 	} else {
// 		fmt.Println("please consult this doctor")
// 	}

// }

// switch choice {
// case "a":
// 	patient = "cardilogist"
// 	fmt.Println("Please go that roomnum1,patient")
// case "b":
// 	patient = "Alopecia Areata"
// 	fmt.Println("please go to roomnum2,patient")
// case "c":
// 	patient = "Ankylosing Spondylitis"
// 	fmt.Println("pleasde go first floor in that go to 2 room,patient")
// case "d":
// 	patient = "Arthritis"
// 	fmt.Println("pleasde go second floor in that go to 7 room,patient")
// case "e":
// 	patient =
// 		"Atopic Dermatitis"
// 	fmt.Println("pleasde go fifth floor in that go to 4 room,patient")
// default:
// 	fmt.Println("Thank you for visiting the hospital.Hope you doing good")
// }
//-------------------------------------------------------------------------------------------------------------
/*package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter the age ")
	fmt.Scanln(&num)
	if num <= 18 {
		fmt.Println("your not eligible for vote")
	} else {
		fmt.Println("your eligible for vote")

	}
}*/

//-----------------------------------------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// type Bike struct {
// 	model      string
// 	price      int
// 	bikeType   string
// 	complaints []string
// }

// type Sales struct {
// 	bikes []Bike
// }

// func (s *Sales) addBike(bike Bike) {
// 	s.bikes = append(s.bikes, bike)
// }

// func (s Sales) totalSales() int {
// 	total := 0
// 	for _, bike := range s.bikes {
// 		total += bike.price
// 	}
// 	return total
// }

// func main() {
// 	s := Sales{}
// 	s.addBike(Bike{model: "Mountain Bike", price: 500, bikeType: "Off-Road"})
// 	s.addBike(Bike{model: "Road Bike", price: 700, bikeType: "Street"})
// 	s.addBike(Bike{model: "BMX Bike", price: 300, bikeType: "Stunt"})
// 	s.addBike(Bike{model: "Hybrid Bike", price: 600, bikeType: "Multi-Purpose"})
// 	s.addBike(Bike{model: "Electric Bike", price: 1200, bikeType: "Electric"})

// 	var model string
// 	var complaints []string

// 	fmt.Print("Enter bike model: ")
// 	fmt.Scanln(&model)

// 	fmt.Print("Enter any complaints (separated by comma): ")
// 	fmt.Scanln(&complaints)

// 	for i, bike := range s.bikes {
// 		if bike.model == model {
// 			s.bikes[i].complaints = complaints
// 		}
// 	}

// 	fmt.Println("Total sales:", s.totalSales())
// 	fmt.Println("Bike list:")
// 	for _, bike := range s.bikes {
// 		fmt.Println("Model:", bike.model, "| Price:", bike.price, "| Type:", bike.bikeType, "| Complaints:", bike.complaints)
// 	}
// }

//--------------------------------------------------------------------------------------------------------------------

//variadic function

// package main

// import "fmt"

// func main() {
// 	variadicExample(1, 2, 3, 4, 5)
// }

// func variadicExample(s ...int) {
// 	fmt.Println(s[0])
// 	fmt.Println(s[3])
// }

//-----------------------------------------------------------------------------------------------------------

//anonymous

// package main

// import "fmt"

// func main() {
// 	result := func(word1, word2 string) string {
// 		return word1 + " " + word2
// 	}("hello", "world")
// 	fmt.Println(result)
// }

// package main

// import "fmt"

// func main() {
// 	value := func(word1, word2, word3 string) string {
// 		return word1 + " " + word2 + " " + word3
// 	}("jaya", "chandra", "reddy")
// 	fmt.Println(value)
// }

//-------------------------------------------------------------------------------------------------------

//   slices and array

// package main

// import "fmt"

// func main() {
// 	numbers := []string{"jaya", "chandra", "reddy", "DevOps", "Jenkins", "Docker"}
// 	var index int = 3
// 	numbers = append(numbers[:index], numbers[index+1:]...)
// 	// a:=numbers[4]
// 	fmt.Println(numbers)
// }

//---------------------------------------------------------------------------------------------------------

// package main
// import (
// 	"fmt"
// 	"math"
// )
// func primenumbers(num1 ,num2 int){
// 	if num1<2 || num2<2{
// 		fmt.Println("numbers must be greater than 2")
// 		return
// 	}
// 	for num1 <= num2{
// 		isprime := true
// 		for i:=1;i<=int(math.Sqrt(float64(num1)));i++ {
// 			if num1 % i ==0  {
// 				isprime = false
// 				break
// 			}

// 		}
// 		if isprime {
// 			fmt.Printf("%d",num1)
// 		}
// 		num1++
// 	}
// 	fmt.Println()
// }
// func main(){
// 	fmt.Println PrimeNumbers(5,19)
// 	fmt.Println PrimeNumbers(0,2)
// 	fmt.Println PrimeNumbers(7,11)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func printPrimeNumbers(num1, num2 int) {
// 	if num1 < 2 || num2 < 2 {
// 		fmt.Println("Numbers must be greater than 2.")
// 		return
// 	}
// 	for num1 <= num2 {
// 		isPrime := true
// 		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
// 			if num1%i == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			fmt.Printf("%d ", num1)
// 		}
// 		num1++
// 	}
// 	fmt.Println()
// }

// func main() {
// 	printPrimeNumbers(5, 19)
// 	printPrimeNumbers(0, 2)
// 	printPrimeNumbers(13, 100)
// }

//
//----------------------------------------------------------------------------------------------------------------

//prime numbers

// package main

// import "fmt"

// func main() {
// 	var isprime bool
// 	for i := 2; i <= 100; i++ {
// 		isprime = true
// 		for j := 2; j <= i/2; j++ {
// 			if i%j == 0 {
// 				isprime = false
// 				break
// 			}
// 		}
// 		if isprime {
// 			fmt.Println(i)
// 		}
// 	}
// }

//------------------------------------------------------------------------------------------------------------------

//Palindrome

// package main

// import "fmt"

// func main() {
// 	company := []string{"jaya", "chandra", "reddy", "devops"}
// 	company = append(company, "kubernetes", "jenkins", "ansible")
// 	fmt.Println(company)

// }

//

// package main
// import "fmt"
// func main(){

// }

// package main

// import (
//     "github.com/aws/aws-sdk-go/aws"
//     "github.com/aws/aws-sdk-go/aws/session"
//     "github.com/aws/aws-sdk-go/service/s3"
//     "log"
//     "strings"
// )

// // Downloads an item from an S3 Bucket in the region configured in the shared config
// // or AWS_REGION environment variable.
// //
// // Usage:
// //    go run s3_download.go
// func main() {
//     bucket := "myBucket"
//     key := "TestFile.txt"

//     // Initialize a session in us-west-2 that the SDK will use to load
//     // credentials from the shared credentials file ~/.aws/credentials.
//     sess, err := session.NewSession(&aws.Config{
//         Region: aws.String("us-west-2")},
//     )

//     // Create S3 service client
//     svc := s3.New(sess)

//     _, err = svc.CreateBucket(&s3.CreateBucketInput{
//         Bucket: &bucket,
//     })
//     if err != nil {
//         log.Println("Failed to create bucket", err)
//         return
//     }

//     if err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket: &bucket}); err != nil {
//         log.Printf("Failed to wait for bucket to exist %s, %s\n", bucket, err)
//         return
//     }

//     _, err = svc.PutObject(&s3.PutObjectInput{
//         Body:   strings.NewReader("Hello World!"),
//         Bucket: &bucket,
//         Key:    &key,
//     })
//     if err != nil {
//         log.Printf("Failed to upload data to %s/%s, %s\n", bucket, key, err)
//         return
//     }

//     log.Printf("Successfully created bucket %s and uploaded data with key %s\n", bucket, key)
// }

// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	numbers = numbers[6:7]

	fmt.Println(numbers)
}
