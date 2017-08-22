#!(whereis env) bash

# basic tests.
# the fizzbuzz server must be already running on the local machine:
# execute:
# go run ../src/fizzbuzz.go &

CURL_OPTIONS=""
TEST_URL="http://localhost:8080"

msg="Fetching the root URL at localhost"
curl ${CURL_OPTIONS} ${TEST_URL}

msg="Fetching the FizzBuzz with incorrect number of parameters"
curl ${CURL_OPTIONS} ${TEST_URL}/bar/baz/3

msg="Fetching the FizzBuzz with inverted parameter order"
curl ${CURL_OPTIONS} ${TEST_URL}/bar/baz/10/2/3

msg="Fetching the FizzBuzz with invalid parameters"
curl ${CURL_OPTIONS}\
 ${TEST_URL}/3/4/56/thisafiftyandmorelongsentencebutyoudonthavetocount/a

curl ${CURL_OPTIONS} ${TEST_URL}/10001/45/80/Fizz/Buzz
curl ${CURL_OPTIONS} ${TEST_URL}/3/-4/18/Fizz/Buzz

msg="Authentic FizzBuzz"
curl ${CURL_OPTIONS} ${TEST_URL}/3/5/50/Fizz/Buzz

msg="Fancy BeerPong"
curl ${CURL_OPTIONS} ${TEST_URL}/6/7/42/Beer/Pong


