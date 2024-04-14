import time
import matplotlib.pyplot as plt

def fibonacci(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    else:
        return fibonacci(n-1) + fibonacci(n-2)

wartosci_n = []
czasy = []
od = 1
do = 46
for n in range(od, do + 1):
    start_time = time.time()
    result = fibonacci(n)
    end_time = time.time()
    czas_obliczen = end_time - start_time
    wartosci_n.append(n)
    czasy.append(czas_obliczen)
    print("Fibonacci od", n, ":", result)
    print("Czas:", czas_obliczen, "s")
    print()

plt.figure(figsize=(10, 6))
plt.plot(wartosci_n, czasy, marker='o', linestyle='-')
plt.title('Czas liczenia rekurencyjnie ciągu Fibonacciego')
plt.xlabel('n wyraz ciagu')
plt.ylabel('Czas (sekundy)')
plt.grid(True)
plt.xticks(range(od, do + 1))
plt.tight_layout()
plt.savefig("plot.png")
