import random
import math
import time

def random_walk(steps):
    x, y = 0, 0
    for _ in range(steps):
        direction = random.choice(['left', 'right', 'up', 'down'])
        if direction == 'left':
            x -= 1
        elif direction == 'right':
            x += 1
        elif direction == 'up':
            y += 1
        elif direction == 'down':
            y -= 1
    distance = math.sqrt(x**2 + y**2)
    return distance

def average_distance(steps, trials):
    distances = []
    for _ in range(trials):
        x, y = 0, 0
        current_distance = 0
        for _ in range(steps):
            direction = random.choice(['left', 'right', 'up', 'down'])
            if direction == 'left':
                x -= 1
            elif direction == 'right':
                x += 1
            elif direction == 'up':
                y += 1
            elif direction == 'down':
                y -= 1
            current_distance = math.sqrt(x**2 + y**2)
            distances.append(current_distance)
    avg_distances = []
    for i in range(steps):
        avg_distance = sum(distances[i::steps]) / trials
        avg_distances.append(avg_distance)
    return avg_distances

trials = 10000
start_time = time.time()
results = average_distance(10000, trials)
end_time = time.time()
elapsed_time = end_time - start_time
print("Czas wykonania:", elapsed_time)

# Zapis wyników do pliku
with open("average_distances.csv", "w") as file:
    for steps, avg_dist in enumerate(results, start=1):
        file.write(f"{steps},{avg_dist}\n")

print("Wyniki zapisano do pliku 'average_distances.csv'")
