import csv

def process_csv(file_name):
    with open(file_name, 'r') as file:
        reader = csv.reader(file)
        data = list(reader)

    processed_data = []
    for row in data:
        col1 = float(row[0])
        col2 = float(row[1])
        col3 = round(col1 * (col2/100), 2)
        col4 = round(col1 - col1 * col2 / 100, 2)
        processed_data.append([col1, col2, col3, col4])

    return processed_data

def save_to_csv(data, file_name):
    with open(file_name, 'w', newline='') as file:
        writer = csv.writer(file)
        writer.writerows(data)

file_name = 'sorted.csv'
processed_data = process_csv(file_name)
save_to_csv(processed_data, 'processed.csv')
