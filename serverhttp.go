package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func isPrime(n int) bool {
	fmt.Println(strconv.Itoa(n))
	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Println(strconv.Itoa(n) + " " + strconv.Itoa(i))
			return false
		}
	}
	return true
}

func primes(n int) []int {
	var resp []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			resp = append(resp, i)
		}
	}
	return resp
}

func toString(li []int) string {
	var resp string = "["
	for i := 0; i < len(li); i++ {
		resp = resp + strconv.Itoa(li[i]) + ","
	}
	return resp[:len(resp)-1] + "]"
}

func primeshttp(w http.ResponseWriter, r *http.Request) {
	var param string = r.URL.Query().Get("n")
	n, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//fmt.Println(primes(int(n)))
	fmt.Fprintf(w, toString(primes(int(n))))
}
func main() {
	http.HandleFunc("/prime", primeshttp)
	http.ListenAndServe(":80", nil)
}
