import time
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker

def fibonacci(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    else:
        return fibonacci(n-1) + fibonacci(n-2)
    
def fibonacci_iter(n):
    if n == 0:
        return 0
    elif n == 1:
        return 1
    else:
        fib = [0, 1]
        for i in range(2, n + 1):
            fib.append(fib[i-1] + fib[i-2])
        return fib[n]

def plot_fibonacci_times(od, do):
    wartosci_n = []
    czasy = []
    for n in range(od, do + 1):
        start_time = time.time()
        result = fibonacci(n)
        end_time = time.time()
        czas_obliczen = end_time - start_time
        wartosci_n.append(n)
        czasy.append(czas_obliczen)
        print("Fibonacci od", n, ":", result)
        print("Czas:", czas_obliczen, "s")
        print()

    plt.figure(figsize=(10, 6))
    plt.plot(wartosci_n, czasy, marker='o', linestyle='-')
    plt.title('Czas liczenia rekurencyjnie ciągu Fibonacciego')
    plt.xlabel('n wyraz ciagu')
    plt.ylabel('Czas (sekundy)')
    plt.grid(True)
    plt.xticks(range(od, do + 1))
    plt.tight_layout()
    plt.savefig("plot.png")

def plot_fibonacci_iterative(od, do):
    wartosci_n = []
    czasy = []
    for n in range(od, do + 1):
        start_time = time.time()
        result = fibonacci_iter(n)
        end_time = time.time()
        czas_obliczen = end_time - start_time
        wartosci_n.append(n)
        czasy.append(czas_obliczen)
        print("Fibonacci od", n, ":", result)
        print("Czas:", czas_obliczen, "s")
        print()

    plt.figure(figsize=(10, 6))
    plt.plot(wartosci_n, czasy, marker='o', linestyle='-')
    plt.title('Czas liczenia iteracyjnie ciągu Fibonacciego')
    plt.xlabel('n-ty wyraz ciągu')
    plt.ylabel('Czas (sekundy)')
    plt.grid(True)
    plt.xticks(range(od, do, 100))
    plt.gca().yaxis.set_major_formatter(ticker.FormatStrFormatter('%.5f'))
    plt.tight_layout()
    plt.savefig("plot_iterative.png")


def calculate_time(POTEZNA_LICZBA):
    time = 466.0715682507
    ratio = 1.618033988749
    n = 46

    while True:
        time *= ratio
        n += 1
        if n >= POTEZNA_LICZBA:
            break

    print("Po {} wykonaniach czas wynosi {:.0f} sekund.".format(POTEZNA_LICZBA - 46, time))

    minutes_per_second = 1 / 60
    hours_per_second = minutes_per_second / 60
    days_per_second = hours_per_second / 24
    years_per_second = days_per_second / 365.25 # a co tam, policzymy z latami przestępnymi 

    time_in_minutes = time * minutes_per_second
    time_in_hours = time * hours_per_second
    time_in_days = time * days_per_second
    time_in_years = time * years_per_second

    print("Czas w minutach:", '{:.0f}'.format(time_in_minutes))
    print("Czas w godzinach:", '{:.0f}'.format(time_in_hours))
    print("Czas w dniach:", '{:.0f}'.format(time_in_days))
    print("Czas w latach:", '{:.0f}'.format(time_in_years))

    print("Ile wszechświatów by zajęło liczenie: " + str(int(time_in_years / 13_787_000_000)))

def calculate_avg():
    times = [
        0.0000007153, 0.0000009537, 0.0000009537, 0.0000007153, 0.0000009537,
        0.0000030994, 0.0000028610, 0.0000042915, 0.0000069141, 0.0000112057,
        0.0000181200, 0.0000288486, 0.0000441074, 0.0000724792, 0.0001153946,
        0.0002038479, 0.0003063679, 0.0005016327, 0.0007898807, 0.0012705326,
        0.0020649433, 0.0033750534, 0.0054795742, 0.0088250637, 0.0145370960,
        0.0236260891, 0.0381035805, 0.0617091656, 0.0998525619, 0.1628260612,
        0.2572720051, 0.4154846668, 0.7016372681, 1.3900573254, 2.4276556969,
        3.9783451557, 6.3765065670, 10.1886794567, 16.3728692532, 26.5859353542,
        42.4968469143, 64.1382308006, 112.1614754200, 175.1904096603, 292.0188045502,
        466.0715682507
    ]





    ratios = [times[i] / times[i-1] for i in range(1, len(times))]
    sorted_ratios = sorted(ratios)


    average = sum(ratios) / len(ratios)

    print("Średnia:", average)
    n = len(sorted_ratios)
    if n % 2 == 0:
        median = (sorted_ratios[n//2 - 1] + sorted_ratios[n//2]) / 2
    else:
        median = sorted_ratios[n//2]

    print("Mediana: ", median)

# plot_fibonacci_times(1, 30)
# plot_fibonacci_iterative(1, 476)
calculate_time(476)
calculate_avg()

# print(341301280657442590686492882549513967819153224652743659108293029741489293429064219215897034752/5.173683166503906e-05) // 6.5968724731181295e+96