"""
This python script is doing multi-threaded calls
to the golang fizzbuzz rest API.


"""

def fizzbuzz_alike( fizz_multiplicator, buzz_multiplicator, integer_max, fizz, buzz ):
    """
    Generate the list of integers with fizz, buzz or fizzbuzz 
    intercalated according to specifications.
    """
    def is_fizz(integer):
        nonlocal fizz_multiplicator
        return integer % fizz_multiplicator == 0

    def is_buzz(integer):
        nonlocal buzz_multiplicator
        return integer % buzz_multiplicator == 0

    result = []
    for x in range(1, integer_max + 1):
        output = ''
        is_fizz_bool, is_buzz_bool = is_fizz(x), is_buzz(x)
        if not is_fizz_bool and not is_buzz_bool:
            output = str(x)
        else:
            output = (fizz if is_fizz_bool else '') + (buzz if is_buzz_bool else '')

        result.append(output)
    return result
from pprint import pprint
pprint(fizzbuzz_alike(3,5,100,'fizz','buzz'))

