import csv

with open('results.csv', 'r') as file:
    reader = csv.reader(file)
    data = list(reader)

sorted_data = sorted(data, key=lambda x: float(x[1]))

with open('sorted.csv', 'w', newline='') as file:
    writer = csv.writer(file)
    writer.writerows(sorted_data)

