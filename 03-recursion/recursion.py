# Infinite Loop Recursion
"""def count(n):
    print(n)
    return count(n-1)"""


# Recursion with base case and recursive case
def count(n):
    if(n == 1): # Base Case
        print(n)
        return 1
    print(n)
    return count(n-1) # Recursive Case


count(5)



# Factorial
def fact(n):
    if (n == 0): return 1
    return n * fact(n - 1)

print(fact(10))


# Fibonacci
def fibonacci(n):
    if(n <= 1):
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

print(fibonacci(5))