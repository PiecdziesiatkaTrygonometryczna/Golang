import random
import math
import time

MAX_ILOSC_ITERACJI = 100
srednie = []
ILOSC_PROB = 10000

def pitagoras(a, b):
    return math.hypot(a, b)  # Użyj math.hypot do obliczenia długości wektora

def avg(arr):
    return round(sum(arr) / len(arr), 2)  # Wykorzystaj funkcje wbudowane do obliczenia średniej

start_time = time.time()

for i in range(1, MAX_ILOSC_ITERACJI + 1):
    odleglosci = []

    for _ in range(ILOSC_PROB):  # Użyj pętli for zamiast while
        punkt = [random.randint(-i, i) for _ in range(2)]  # Wygeneruj punkt za pomocą listy składanej
        odleglosc = pitagoras(*punkt)  # Rozpakuj punkt jako argumenty dla funkcji pitagoras
        odleglosci.append(odleglosc)

    srednie.append(avg(odleglosci))

end_time = time.time()
elapsed_time = end_time - start_time
print("Czas wykonania:", elapsed_time)

print(srednie)
