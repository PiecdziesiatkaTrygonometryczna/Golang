import matplotlib.pyplot as plt
import csv

# Wczytaj dane z pliku CSV
steps = []
avg_distances = []
with open("dane/average_distances.csv", "r") as file:
    reader = csv.reader(file)
    for row in reader:
        steps.append(int(row[0]))
        avg_distances.append(float(row[1]))

# Utwórz wykres
plt.figure(figsize=(10, 5))
plt.plot(steps, avg_distances, marker='o', color='blue', linestyle='-')

# Dodaj etykiety osi
plt.xlabel('Ilość iteracji')
plt.ylabel('Odległość')

# Dodaj tytuł
plt.title('Średnia odległość od punktu startowego')

# Zapisz wykres do pliku obrazu
plt.savefig('wykres2.png')