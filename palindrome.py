no=int(input("Enter any number:"))
cpy=no
rev=0
while no!=0:
    r=no % 10;
    rev=rev*10+r
    no=int(no/10)
if cpy==rev:
    print("Given number is palindrome")
else:
    print("Not palindrome")   
        
    