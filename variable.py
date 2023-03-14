# Israiny = input("enter your raining yes or no :")
# if Israiny=="yes":
#     print("have umbrulla")
# elif Israiny=="no":
#     print("Go outside")
# else:
#     print("enter current raining Yes or no ")
    
# lst=[1,2,3,4,5,6,7,8,9,10]
# for num in lst:
#         if num % 2 != 0:
#            print(num)
           
           
# lst=input("enter any number :")
# for num in range(20):
#         if num % 2 != 0:
#            print(num)


# lst=[636,787,78,347,587]
# lst.sort()
# print("second highest number of the lst is:" ,lst[-2])
# print(lst)


# import random
# secretNumber = random.randint(1, 20)
# print('I am thinking of a number between 1 and 20.')
# # Ask the player to guess 6 times.
# for guessesTaken in range(1, 7):
#  print('Take a guess.')
#  guess = int(input())
#  if guess < secretNumber:
#     print('Your guess is too low.')
#  elif guess > secretNumber:
#     print('Your guess is too high.')
#  else:
#    break # This condition is the correct guess!
# if guess == secretNumber:
#  print('Good job! You guessed my number in ' + str(guessesTaken) + 'guesses!')
# else:
#  print('Nope. The number I was thinking of was ' + str(secretNumber))



# lst=input("enter any number :")
# for i in range(20):
#      i.sort()
# print("second highest number of the lst is:" ,lst[-2])
# print(lst)


# lst = []
# # number of elemetns as input
# n = int(input("Enter number of elements : "))
# # iterating till the range
# for i in range(0, n):
#             ele = int(input)
#             lst.append(ele) # adding the element
# print(lst)







# import mysql.connector

# # Connect to the database
# conn = mysql.connector.connect(
#   host="localhost",
#   user="jaya",
#   password="@Jay2000",
#   database="jaya"
# )

# # Check if the connection was successful
# if conn.is_connected():
#   print("Connected to MySQL Database")
  
  
  
  
import boto3,pprint,sqlite3
client = boto3.client('ec2')
response =client.describe_instances()
pprint.pprint(response)
con = sqlite3.connect("instance.db")    
con.execute("create table instance(InstanceType TEXT NOT NULL, InstanceId TEXT UNIQUE NOT NULL, VpcId TEXT NOT NULL)")  
for i in response.values():
     if type(i)==list:
        a=i[0]
        for j in a.values():
            if len(j)==0:
                continue
            else:
                b=j[0]
                # print(type(b))
                # pprint.pprint(b)
                c=b['InstanceId']
                d=b['InstanceType']
                e=b['VpcId']
                break
con = sqlite3.connect("instance.db")   
cur = con.cursor()  
cur.execute("INSERT into instance(InstanceType, InstanceId, VpcId) values (?,?,?)",(d,c,e))  
con.commit()        
con.close()
    
    
    

    

    
    
    
    
    
    
    
    
    
    
    
    
    
    
    