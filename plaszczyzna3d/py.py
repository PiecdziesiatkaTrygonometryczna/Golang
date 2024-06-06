import random
import math

NUM_OF_SIMS = 10_000
MAX_NUM_OF_STEPS = 20

sum = 0
for i in range(NUM_OF_SIMS):
    point = [0,0,0]
    for j in range(MAX_NUM_OF_STEPS):
        axis = random.randint(0,2) # 0,1,2
        point[axis] += (random.randint(0,1) * 2) - 1 # -1 or 1
    distance = math.sqrt(point[0]**2 + point[1]**2 + point[2]**2)
    sum += distance
average = sum / NUM_OF_SIMS
print(average)