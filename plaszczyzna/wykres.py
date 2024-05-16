import matplotlib.pyplot as plt

# Wczytaj dane z pliku
with open("srednie.txt", "r") as file:
    lines = file.readlines()
    values = [float(line.strip()) for line in lines]

# Utwórz wykres
plt.figure(figsize=(10, 5))
plt.scatter(range(1, len(values) + 1), values, marker='o', color='blue')

# Dodaj etykiety osi
plt.xlabel('Ilość iteracji')
plt.ylabel('Odległość')

# Dodaj tytuł
plt.title('Średnia odległość od punktu startowego')

# Zapisz wykres do pliku obrazu
plt.savefig('wykres.png')

# Wyświetl wykres
plt.show()
