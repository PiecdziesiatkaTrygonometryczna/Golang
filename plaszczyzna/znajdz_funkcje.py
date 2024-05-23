import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from scipy.optimize import curve_fit

data = pd.read_csv('dane/srednie1.csv', header=None)
x = data[0].values
y = data[1].values

def sqrt_func(x, a, b):
    return a * np.sqrt(x) + b

params, _ = curve_fit(sqrt_func, x, y)

plt.figure(figsize=(12, 6))

plt.scatter(x, y)
plt.plot(x, sqrt_func(x, *params), color='red')
plt.legend()
plt.title('Funkcja')
plt.xlabel('X')
plt.ylabel('Y')
plt.show()

print(f"Funkcja: a = {params[0]}, b = {params[1]}")
