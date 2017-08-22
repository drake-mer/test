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


func fizz_buzz(input int, parameters FizzBuzz) (string){
    is_fizz := (input%parameters.fizz_multiplicator == 0)
    is_buzz := (input%parameters.buzz_multiplicator == 0)
    var output = ""
    if is_fizz || is_buzz {
        if is_fizz{
            output = output + parameters.fizz
        }    
        if is_buzz {
            output = output + parameters.buzz
        }
    } else {
        output = fmt.Sprint(input)
    }
    return output
}


func FizzBuzzGenerator(fizzBuzzParams FizzBuzz) (result []string){
    for i:=1;i<fizzBuzzParams.max_range;i=i+1 {
        fizz_buzz_result := fizz_buzz(i, fizzBuzzParams)
        result = append(result, fizz_buzz_result)
    }
    return
}


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func FizzBuzzAnswer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Println("in FizzBuzzAnswer")
    fizz_int, err1 := strconv.ParseInt(ps.ByName("fizz_int"), 10, 32)
    buzz_int, err2 := strconv.ParseInt(ps.ByName("buzz_int"), 10, 32)
    max_range, err3 := strconv.ParseInt(ps.ByName("max_range"), 10, 32)
    check_parameters_ok := max_range > 0 && fizz_int > 0 && buzz_int > 0
    check_parameters_ok = check_parameters_ok && max_range < 10000 && fizz_int < 10000 && buzz_int < 10000

    if err1==nil && err2==nil && err3==nil && check_parameters_ok{
        fizzBuzzParams := FizzBuzz{
            ps.ByName("fizz_string"),
            ps.ByName("buzz_string"),
            int(fizz_int),
            int(buzz_int),
            int(max_range),
        }
        fmt.Println("URI parameters are", fizzBuzzParams)
        fmt.Println("Output results are", FizzBuzzGenerator(fizzBuzzParams))
    } else {
        fmt.Println("Invalid Parameters in URI request")
        fmt.Println("Please make sure that the URI has format",
        "http://hostname/fizzbuzz/{int}/{int}/{int}/{string}/{string}",
        "where the first int is the fizz multiplicator",
        "the second int is the buzz multiplicator",
        "the third int is the maximum range for the fizzbuzz generator",
        "the first and second string are respectively the fizz and buzz strings",
        "The integer parameters must not be lesser than 1 and must not be greater than 9999")
    }
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/:fizz_int/:buzz_int/:max_range/:fizz_string/:buzz_string", FizzBuzzAnswer)
    log.Fatal(http.ListenAndServe(":8080", router))
}
