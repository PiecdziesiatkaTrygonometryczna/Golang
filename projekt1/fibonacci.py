lol = 30

counters = [0] * (lol+1)

def fibonacci(n):
    global counters
    counters[n] += 1
    if n == 0:
        return 0
    if n == 1:
        return 1
    else:
        return fibonacci(n-1) + fibonacci(n-2)

fibonacci_result = fibonacci(lol)

print(f"Fibonacci od {lol}:", fibonacci_result)

for i in range(len(counters)):
    print(i, counters[i])
