name = "dawkal"
arr = []
for letter in name:
    arr.append(ord(letter))


print(arr)

for n in range(1,10000):
   
    silnia = 1
    for x in range(1, n):
        silnia = silnia * x

    flag = True
    for element in arr:
        if str(element) not in str(silnia):
            flag = False
            break
    if flag == True:
        print(str(n) + "! = " + str(silnia))
        break