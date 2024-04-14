def factorial(n):
    result = 1
    for x in range(1, n + 1): #+1 bo przedzial zamkniety
        result *= x
    return result

def find_strong_number(name):
    ascii_values = [str(ord(letter)) for letter in name] # konwertujemy imię na tablicę znaków ASCII

    counter = 1 # szukamy Silnej Liczby, zaczynamy od 1
    while True:
        found = True
        factorial_str = str(factorial(counter))

        for ascii_letter in ascii_values:
            index = factorial_str.find(ascii_letter)
            if index == -1: # jeżeli nie znajdziemy kodu ASCII naszej litery w silni, wychodzimy z pętli
                found = False
                break
            factorial_str = factorial_str[index + len(ascii_letter):] # jeżeli znajdziemy, to ucinamy początek naszej silni i sprawdzamy kolejną literę

        if found: # jeżeli wszystkie litery zawierają się po kolei w naszej silni, to kończymy szukanie
            return counter

        counter += 1




def main():
    name = "dawkal"
    strong_number = find_strong_number(name)
    print("Twoja Silna Liczba:", strong_number)
    with open("silnia.txt", "w") as file:
        file.write(str(factorial(strong_number)))
    print("Silnia Twojej Silnej Liczby została zapisana w pliku silnia.txt")
    
if __name__ == "__main__":
    main()
# Silna liczba dla "dawkal" to 476
    # [100, 97, 119, 107, 97, 108] - kody ASCII mojego nicku
# 476! = 35867867064398478289209288875534387019028416258811836097155671133214919145444260935805432100557203526635297136372252064698892609981421602639632038857887829475680340687505785927330492899589147639440355071247941210396418976790496116638139844315723889910751888763214906402206414338520669922794130323497648626571442758115898723512911338574508975674744299325488079349804387736838497922290989218810430929600444317990195672429289253146597780504680119710765307542463279546249869881122681659239945050603874329534230560986949687181508415056078842404709174342374251031568128240156579745710039907957665407214016795792446584465884461641692036758184309865915561118301965962774560047028427934154513836663823856366754722964717185684148702128819913812101327095717508673881082348629325961023547740711742985105209962397132625285663757717550333733499916097970929256584070049965155727074714934204283094829846721551122905670975430021850311566169589525408067135863332586651648000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000