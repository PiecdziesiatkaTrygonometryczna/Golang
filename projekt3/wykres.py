import matplotlib.pyplot as plt
import csv

steps = []
avg_distances = []
with open("srednie.csv", "r") as file:
    reader = csv.reader(file)
    for row in reader:
        steps.append(int(row[0]))
        avg_distances.append(float(row[1]))

plt.figure(figsize=(10, 5))
plt.plot(steps, avg_distances, marker='o', color='blue', linestyle='-')

plt.xlabel('Ilość iteracji')
plt.ylabel('Odległość')

plt.title('Średnia odległość od punktu startowego')

plt.savefig('wykres.png')