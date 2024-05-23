import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv('processed.csv', header=None)

x = df[0]
y = df[1]

plt.plot(x, y, marker='o', linestyle='-')
plt.xlabel('% zalesienia')
plt.ylabel('% spalenia')
plt.title('% Zalesienia vs % spalonych drzew ')
plt.grid(True)
plt.savefig("wykres.png")
