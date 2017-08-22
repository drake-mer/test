package main

import (
    "fmt"
    "strconv"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

type FizzBuzz struct {
    fizz string
    buzz string
    fizz_multiplicator int
    buzz_multiplicator int
    max_range int
}

type fizzbuzz interface {
    Generator() []string  // generates the FizzBuzz list
    Validator() bool  // Validate that the FizzBuzz is contained in a given set of parameters
    Serializer(httprouter.Params)  // transform URI parameters into a FizzBuzz
    Stringify(int)  // convert an integer value according to the FizzBuzz parameters
}


func (fizzBuzz FizzBuzz) Generator() (result []string){
    for i:=1 ; i < fizzBuzz.max_range ; i=i+1 {
        fizz_buzz_result := fizzBuzz.Stringify(i)
        result = append(result, fizz_buzz_result)
    }
    return
}

func (fizzBuzz FizzBuzz) Validator() (ok bool, msg string) {
    /*
    This method validate if the set of parameters inside the struct
    is valid to perform a FizzBuzz on it
    */

    // check if we have only positive numbers
    ok = !( fizzBuzz.fizz_multiplicator < 1 || 
            fizzBuzz.buzz_multiplicator < 1 ||
            fizzBuzz.max_range < 1 )
    if !ok {
        msg = "In " + fmt.Sprint(fizzBuzz) 
        msg += ", a FizzBuzz integer parameter is not > 0"
        return
    }

    // check if integer range is no bigger than 10000
    ok = !( fizzBuzz.fizz_multiplicator > 10000 ||
            fizzBuzz.buzz_multiplicator > 10000 ||
            fizzBuzz.max_range > 10000 )
    if !ok {
        msg = "One of the FizzBuzz integer parameter is bigger than 10000"
        return
    }

    // check the string length
    ok = ( len(fizzBuzz.buzz) < 50 || len(fizzBuzz.fizz) < 50 ||
           len(fizzBuzz.buzz) == 0 || len(fizzBuzz.fizz) == 0 )
    if !ok {
        msg = "One of the input word is larger than 50 bytes or empty string"
        return
    }
    msg = "OK"
    return
}


func (fizzBuzz * FizzBuzz) Serializer(ps httprouter.Params) {
    /*
    This method transform the set of parameters inside a
    structure of type httprouter.Params into a convenient 
    FizzBuzz structure.
    */
    
    fizz_int, err := strconv.ParseInt(ps.ByName("fizz_int"), 10, 64)
    if err == nil {
        fizzBuzz.fizz_multiplicator = int(fizz_int)
    }

    buzz_int, err := strconv.ParseInt(ps.ByName("buzz_int"), 10, 64)
    if err == nil {
        fizzBuzz.buzz_multiplicator = int(buzz_int)
    }

    max_range, err := strconv.ParseInt(ps.ByName("max_range"), 10, 64)
    if err == nil {
        fizzBuzz.max_range = int(max_range)
    }

    fizzBuzz.buzz = ps.ByName("buzz_string")
    fizzBuzz.fizz = ps.ByName("fizz_string")
}


func (parameters FizzBuzz) Stringify(input int) (string) {
    /* 
    This method takes an int and return a string.
    If the spec says Fizz, it returns Fizz
    If the spec says Buzz, it returns Buzz
    And so on...
    */
    is_fizz := (input%parameters.fizz_multiplicator == 0)
    is_buzz := (input%parameters.buzz_multiplicator == 0)
    var output = ""
    if is_fizz || is_buzz {
        if is_fizz {
            output = output + parameters.fizz
        }
        if is_buzz {
            output = output + parameters.buzz
        }
    } else {
        // in this case, it is neither a buzz nor a fizz
        output = fmt.Sprint(input)
    }
    return output
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome to your FizzBuzz Generator!\n")
}

func FizzBuzzAnswer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    myFizzBuzz := FizzBuzz{}
    myFizzBuzz.Serializer(ps)
    ok, err_msg := myFizzBuzz.Validator()
    if !ok {
        fmt.Println("incorrect parameters: "+err_msg)
        return
    }
    fmt.Println("result:\n" + fmt.Sprint(myFizzBuzz.Generator()))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/:fizz_int/:buzz_int/:max_range/:fizz_string/:buzz_string", FizzBuzzAnswer)
    log.Fatal(http.ListenAndServe(":8080", router))
}
