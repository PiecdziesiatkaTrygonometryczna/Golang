counterOfZeros = 0
counterOfOnes = 0

def fibbonaci(n):
    global counterOfZeros, counterOfOnes

    if n == 0:
        counterOfZeros += 1
        return 0
    if n == 1:
        counterOfOnes += 1
        return 1
    else:
        return fibbonaci(n-1) + fibbonaci(n-2)
    
print(fibbonaci(30))
print("0: " + str(counterOfZeros))
print("1: " + str(counterOfOnes))
