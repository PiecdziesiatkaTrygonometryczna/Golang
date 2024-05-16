import random
import math
import time

MAX_ILOSC_ITERACJI = 1000
srednie = []
ILOSC_PROB = 1000

def pitagoras(a, b):
    return math.sqrt(pow(a, 2) + pow(b, 2))

def avg(arr):
    sum = 0
    for i in range(len(arr)):
        sum += arr[i]
    return round(sum / len(arr), 2)

start_time = time.time()

for i in range(1, MAX_ILOSC_ITERACJI + 1):
    licznik_prob = 0
    odleglosci = []

    while licznik_prob < ILOSC_PROB:
        punkt = [0, 0]
        licznik_iteracji = 0
        while licznik_iteracji < i:
            x_or_y = random.choice([0, 1])
            plus_or_minus = random.choice([-1, 1])
            punkt[x_or_y] += plus_or_minus
            licznik_iteracji += 1
        odleglosc = pitagoras(punkt[0], punkt[1])
        odleglosci.append(odleglosc)
        licznik_prob += 1

    srednie.append(avg(odleglosci))

end_time = time.time()
elapsed_time = end_time - start_time
print("Czas wykonania:", elapsed_time)

print(srednie)
