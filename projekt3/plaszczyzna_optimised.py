import math
import random
import time
import csv

MAX_ILOSC_ITERACJI = 10_000
ILOSC_PROB = 1_000

def pitagoras(a, b):
    return math.sqrt(a**2 + b**2)

def avg(arr):
    return sum(arr) / len(arr)

def main():
    srednie = []
    punkty = []
    odleglosci = [[] for _ in range(ILOSC_PROB)]

    start = time.time()

    for i in range(1, MAX_ILOSC_ITERACJI + 1):
        if i == 1:
            punkty = [[0, 0] for _ in range(ILOSC_PROB)]
            for j in range(ILOSC_PROB):
                for k in range(i):
                    if random.randint(0, 1) == 0:
                        punkty[j][0] += random.randint(0, 1) * 2 - 1
                    else:
                        punkty[j][1] += random.randint(0, 1) * 2 - 1
                odleglosci[j].append(pitagoras(punkty[j][0], punkty[j][1]))
        else:
            for j in range(ILOSC_PROB):
                if random.randint(0, 1) == 0:
                    punkty[j][0] += random.randint(0, 1) * 2 - 1
                else:
                    punkty[j][1] += random.randint(0, 1) * 2 - 1
                odleglosci[j].append(pitagoras(punkty[j][0], punkty[j][1]))

        suma = sum(odleglosci[j][i-1] for j in range(ILOSC_PROB))
        srednie.append(suma / ILOSC_PROB)

    print(f"Czas wykonania: {time.time() - start:.2f} sekund")

    with open('srednie.csv', 'w', newline='') as file:
        writer = csv.writer(file)
        writer.writerow(['n', 'wynik'])
        for i, srednia in enumerate(srednie):
            writer.writerow([i + 1, srednia])

    print("Średnie zostały zapisane do pliku srednie.csv.")

if __name__ == "__main__":
    main()
