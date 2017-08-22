# Exercise: Write a simple fizz-buzz REST server.

The original fizz-buzz consists in writing all numbers from 1 to 100, 
and just replacing all multiples of 3 by “fizz”, all multiples of 5 
by “buzz”, and all multiples of 15 by “fizzbuzz”. 

The output would look like this:

    1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,fizz,...

More specifically :

• Expose a REST API endpoint that accepts five parameters : 
  two strings (say, string1 and string2), 
  and three integers (say, int1, int2 and limit), and returns a JSON
• It must return a list of strings with numbers from 1 to limit, where:
• all multiples of int1 are replaced by string1,
• all multiples of int2 are replaced by string2,
• all multiples of int1 and int2 are replaced by string1string2

# Answer

A simple REST API is provided that returns the FizzBuzz data as JSON.

The input parameters are expressed in the URI as follow:

http://localhost:8080/{int1}/{int2}/{limit}/{string1}/{string2}

* int1, int2 and limit are only allowed in the range [1,10000]
* string1 and string2 must not be larger than 50 bytes

Deviating parameters will result in a 404 error.

The API support only the GET HTTP verb. Any other verb will result
in a 404 error.

A simple set of automated tests may be run using pytest. 
They illustrate very well the behaviour of the API.

To execute them, you will need
* a working docker installation
* a working python3 environment

## How To Test (v1)

A set of automated tests are written with `pytest`
and the `requests` library.

Before moving forward, please ensure that all the dependancies
to this project are met. Please run

    pip install --user pytest pytest-docker requests

and make sure that all the installed lib are correctly set up in your path.

Then you can do, from the root folder of the application :

    pytest

It will launch the pytest session and it will report accordingly.

## How To Test (v2)

A simple shell script using curl can be set up to 
launch the web server as a process and to retrieve results
from it.

Simply run, from the tests/ folder:

    ./curl_tests.sh

It will launch the webserver and execute a series of test
on it.

