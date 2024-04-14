maxFibonacci = 30

counters = [0] * (maxFibonacci+1)

def fibonacci(n):
    global counters
    counters[n] += 1
    if n == 0:
        return 0
    if n == 1:
        return 1
    else:
        return fibonacci(n-1) + fibonacci(n-2)

fibonacci_result = fibonacci(maxFibonacci)


# print(f"Fibonacci od {lol}:", fibonacci_result)
# for i in range(len(counters)):
#     print(i, counters[i])

with open('silnia.txt', 'r') as file:
    silna_liczba = int(file.readline().strip())


closest = abs(silna_liczba - counters[0])
closest_index = 0

for i in range(len(counters)):
    if abs(silna_liczba - counters[i]) < closest:
        closest = abs(silna_liczba - counters[i])
        closest_index = i

print(f"Najbliższa liczba Twojej Silnej Liczbie to {counters[closest_index]}, czyli Twoja Słaba Liczba to {closest_index}.")