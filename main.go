package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const message = "Hello"

func main()  {
	var randomNumbers []int
	randomNumbers = createRandomNumbers()
	fmt.Println(randomNumbers)

	mux := http.NewServeMux()
	mux.HandleFunc("/prime", func(writer http.ResponseWriter, request *http.Request) {

		var primeNumbers []int

		for i := range randomNumbers{

			if isPrime(i){
				primeNumbers = append(primeNumbers, randomNumbers[i])
			}
		}
		fmt.Fprint(writer,primeNumbers)
	})

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello!"))
	})
	http.ListenAndServe(":8080",mux)
}

func isPrime(number int) (b bool){

	if number == 0 || number == 1{
		b = false
		return
	}

	for i := 2; i < number; i++ {

		if number % i == 0{
			b = false
			return
		}
	}
	b = true
	return
}

func createRandomNumbers() (randomNumbers []int){
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000000; i++ {
		random := rand.Intn(2500 - 2) + 2

		randomNumbers = append(randomNumbers, random)
	}
	fmt.Println(randomNumbers)
	return
}



