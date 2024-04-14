import math

def findStrongNumber(name):
    ascii_values = [str(ord(letter)) for letter in name] # konwertujemy imię na tablicę znaków ASCII
    
    counter = 1 # przeszukujemy wszystkie możliwe silnie, poczynając od 1
    while True:
        factorial_str = str(math.factorial(counter)) # konwertujemy silnię na string

        current_index = 0 # na razie szukamy od indeksu 0 (czyli w całej silni)
        found = True # flaga
        for ascii_letter in ascii_values:
            index = factorial_str.find(ascii_letter, current_index) #szukamy kodu ASCII naszej litery w silni
            if index == -1:  # jeżeli nie znajdziemy, wychodzimy z pętli
                found = False
                break
            current_index = index + 1 # jeżeli znajdziemy, zapisujemy indeks i szukamy kolejnej litery w przedziale (indeks, koniec)
        
        if found:
            return counter # jeżeli znaleźliśmy wszystkie litery, kończymy działanie funkcji

        counter += 1

def main():
    name = "dawkal"
    strongNumber = findStrongNumber(name)
    print("Twoja Silna Liczba:", strongNumber)
    with open("silnia.txt", "w") as file:
        file.write(str(strongNumber) + "\n")
        file.write(str(math.factorial(strongNumber)) + "\n")
    print("Twoja Silna Liczba, oraz jej silnia, zostały zapisane w pliku silnia.txt")

if __name__ == "__main__":
    main()