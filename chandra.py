# This is extend dict: 

  
# marks = {
#     "year1": [80, 90, 85],
#     "year2": [75, 95, 82],
#     "year3": [72, 88, 94],
# }
# marks["year4"] = [89, 92, 95, 87]
# print(marks)


# this is insert dict:


# years = {
#        "1st year": ["45,89,67,89"],
#        "2nd year":["56,67,89,78"],
#        "3rd year":["68,89,90,96"],
# }
# years.update({"4th year": ["89, 92, 95, 87"]})
# print(years)




# employ = {
# "FirstName": "jaya",
# "LastName": "devarinti",
# "employeeI-D": 6281958978,
# "Designation": "DevOps Engineer",
# "LanguageExpertise": ["AWS","Python"],
# "bike": {
# "Make&Model": "pulser yamaha",
# "MakeYear": 2023,
# "Color": "black",
# }
# }
# {
# "FirstName": "venkat",
# "LastName": "reddy",
# "employeeID": 7995910832,
# "Designation": "python developer",
# "LanguageExpertise": ["python", "AWS"],
# "Car": {
# "Make&Model": "Hyundai Verna",
# "MakeYear": 2023,
# "Color": "white",
# }
# }

# print(employ)




# a={
# "FirstName": "jaya",
# "LastName": "chandra",
# "employeeID": 7995910832,
# "Designation": "python developer",
# }
# e=len(a)
# for i in range(e):
#     var=input("Enter  a key :")
#     if var in a:
#         print("Key exist in JSON data")
#         print(a[var])
#     else:
#         print("Key doesn't exist in Json data")


# dict1={
#     "name":"jaya",
#     "age":23,
#     "gender":"male",
#     "number":"6281958978",
#     "adress":"rayachoti",
# }
# fun=len(dict1)
# for i in range(fun):
#     var=input("enter any key:")
#     if var in dict1:
#         print("key is existing ")
#         print(dict1[var])
#     else:
#         print("key is not existing ")
    




# user_list = []

# for i in range(10):
#     user_input = input(f"Enter item {i + 1}: ")
#     user_list.append(user_input)

# print("The list is: ", user_list)

# import random

# my_list = [1, 2, 3, 4, 5]

# random_item = random.choice(my_list)
# print("Random item: ", random_item)

# index = my_list.index(random_item)
# print(f"Index of the random item: {index}")



# 09/02/2023
# ---------

# spam = ['cats', 'dogs', 'moose']
# bacon = ['dogs', 'moose', 'cats']
# spam == bacon
# print(spam==bacon)



# eggs ={'name': 'Zophie', 'species': 'cat', 'age': '8'}
# ham = {'species': 'cat', 'age': '8', 'name': 'Zophie'}
# eggs == ham
# print(eggs==ham)


# spam = {'name': 'Zophie', 'age': 7}
# spam['age']
# print(spam['age'])

# spam = {}
# spam['first key'] = 'value'
# spam['second key'] = 'value'
# spam['third key'] = 'value'
# list(spam)
# print(list(spam))


# dictinaries are supported three types[keys(),values(),items()]

# spam = {'color': 'red', 'age': 42}
# for i in spam.keys():
#    print(i)


# spam = {'color': 'red', 'age': 42}
# for i in spam.values():
#     print(i)


# spam = {'color': 'red', 'age': 42}
# for i in spam.items():
#     print(i)


# spam = {'color': 'red', 'age': 42}
# for k, v in spam.items():
#     print('Key: ' + k + ' Value: ' + str(v))


# birthdays = {'Alice': 'Apr 1', 'Bob': 'Dec 12', 'Carol': 'Mar 4'}
# while True:
#  print('Enter a name: (blank to quit)')
#  name = input()
# if name == '':
#    break
# if name in birthdays:
#   print(birthdays[name] + ' is the birthday of ' + name)
# else:
#  print('I do not have birthday information for ' + name)
#  print('What is their birthday?')
#  bday = input()
# birthdays[name] = bday
# print('Birthday database updated.')

#  get() method:
# picnicItems = {'apples': 5, 'cups': 2}
# # print('I am bringing ' + str(picnicItems.get('cups',0)) + ' cups.')
# print('i am bringing' + str(picnicItems.get('apples',0)) +' apples ')


# setdefault()method:
# spam = {'name': 'jaya', 'age': 5}
# spam.setdefault('color', 'black')
# print(spam)

# count() in dictionary
# import pprint
# message = 'It was a bright cold day in April, and the clocks were striking thirteen.'
# count = {}
# for character in message:
#   count.setdefault(character, 0)
#   count[character] = count[character] + 1
# pprint.pprint(count)



# 10/02/2023

