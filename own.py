def add(a, b, x, y):
    return a + b + x + y
def subtract(a, b, x, y):
    return a - b- x - y
def multiply(a, b, x, y):
    return a * b* x * y
def divide(a, b, x, y):
    return a / b / x / y
def exponential(a, b, x, y):
    return a ** b ** x ** y

print("Select operation.")
print("1.Add")
print("2.Subtract")
print("3.Multiply")
print("4.Divide")
print ("5.exponential")
while True:
    choice = input("Enter choice(1/2/3/4/5): ")

    if choice in ('1', '2', '3', '4','5'):
        # try:
        num1 = float(input("Enter first number: "))
        num2 = float(input("Enter second number: "))
        num3 = float(input("enter third number: "))
        num4 = float(input("enter fourth number: "))
            
        # except ValueError:
        #     print("Invalid input. Please enter a number:5")
        #     continue

        if choice == '1':
            print(num1, "+", num2, "+", num3, "+", num4, "=", add(num1, num2, num3, num4))

        elif choice == '2':
            print(num1, "-", num2, "-", num3, "-", num4, "=", subtract(num1, num2, num3, num4))

        elif choice == '3':
            print(num1, "*", num2, "*", num3, "*", num4, "=", multiply(num1, num2, num3, num4))

        elif choice == '4':
            print(num1, "/", num2, "/", num3, "/", num4, "=", divide(num1, num2, num3, num4))
        
        elif choice == '5':
            print(num1, "**", num2, "**", num3, "**", num4, "=",divide(num1, num2, num3, num4))
        next_calculation = input("Let's do next calculation? (yes/no): ")
        if next_calculation == "no":
          break
    else:
        print("Invalid Input")

