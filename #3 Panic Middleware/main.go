package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the http server")

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/add", add)
	mux.HandleFunc("/substract", substract)
	mux.HandleFunc("/multiply", multiply)
	mux.HandleFunc("/divide", divide)
	log.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(":3000", recoverMw(mux, true)))
}

func recoverMw(app http.Handler, dev bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				if !dev {
					http.Error(w, "Something went wrong :(", http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, string(stack))
			}
		}()

		nw := &responseWriter{ResponseWriter: w}
		app.ServeHTTP(nw, r)
		nw.flush()
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome in calculator!</h1>")
}

func add(w http.ResponseWriter, r *http.Request) {
	numbers := getNumbers(r)
	result := numbers[0] + numbers[1]
	printResult(w, "Addition", "+", numbers, result)
}

func substract(w http.ResponseWriter, r *http.Request) {
	numbers := getNumbers(r)
	result := numbers[0] - numbers[1]
	printResult(w, "Substraction", "-", numbers, result)
}

func multiply(w http.ResponseWriter, r *http.Request) {
	numbers := getNumbers(r)
	result := numbers[0] * numbers[1]
	printResult(w, "Multiplication", "*", numbers, result)
}

func divide(w http.ResponseWriter, r *http.Request) {
	numbers := getNumbers(r)

	if numbers[1] == 0 {
		panic("You can't divide by 0!")
	}

	result := numbers[0] / numbers[1]
	printResult(w, "Division ", "/", numbers, result)
}

func printResult(w http.ResponseWriter, header string, calculation string, numbers []int, result int) {
	fmt.Fprintf(w, "<h1>%s</h1><p>%v %s %v = %v</p>", header, numbers[0], calculation, numbers[1], result)
}

func getNumbers(r *http.Request) []int {
	numbers, ok := r.URL.Query()["numbers"]

	if !ok || len(numbers) != 2 {
		panic("Two numbers required!")
	}

	intNumbers := []int{}

	for _, i := range numbers {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		intNumbers = append(intNumbers, num)
	}
	return intNumbers
}

type responseWriter struct {
	http.ResponseWriter
	writes [][]byte
	status int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.writes = append(rw.writes, b)
	return len(b), nil
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
}

func (rw *responseWriter) flush() error {
	if rw.status != 0 {
		rw.ResponseWriter.WriteHeader(rw.status)
	}
	for _, write := range rw.writes {
		_, err := rw.ResponseWriter.Write(write)
		if err != nil {
			return err
		}
	}
	return nil
}
