# marks=int(input("enter the no. of students: "))
# dict={"a":"enter the student hall ticket number",
#       "b":"enter the student name",
#       "c":"enter the student marks"
#      }
# for marks in range(30):
#     if(marks<=50):
#         print("pass")
#     else:
#         print("fail")  
#     print()  
        

# # for marks in range(30):
# if(marks<=50):
#       print("Pass")
# else:
#     print("fail")
#     # print(marks, end="")
# # print()



# #Student class
# class Student():

# #Details method to get the student marks and name
#   def details(self,mark1,mark2,mark3,name):
#       self.m1=mark1
#       self.m2=mark2
#       self.m3=mark3
#       self.n=name

# #printing the marks and name
#   def _print(self):
#       print("Enter the mark1:",self.m1)
#       print("Enter the mark2:",self.m2)
#       print("Enter the mark3:",self.m3)
#       print("Enter the name:",self.n)

#   def loop(self,passmark =40):
#       self.pm=passmark
#       if(self.m1>=self.pm):
#         print("pass")
#       else:
#         print("Fail")
#         if(self.m2>=self.pm):
#           print("Pass")
#         else:
#           print("Fail")

#         if(self.m3>=self.pm): 
#           print( "Pass")
#         else:
#           print( "Fail")

# s=Student()

# s.details(input('enter the mark1:'),input('enter the mark2:'),input('enter the mark3:'),input('enter the name:'))

# s._print()
# s.loop("")

marks ={}
threshold = 50
pass_count = 0
fail_count = 0

for i in range(3):
    HTno = input("Enter the student's Hall ticket number :")
    name = input("Enter the student's name: ")
    mark = float(input("Enter the student's mark: "))
    marks[name] = mark
for name, mark in marks.items():
    if mark >= threshold:
        pass_count += 1
    else:
        fail_count += 1
print("Number of students who passed: ", pass_count)
print("Number of students who failed: ", fail_count)














