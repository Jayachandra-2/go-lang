import boto3,pprint,sqlite3
client = boto3.client('ec2')
response =client.describe_instances()
pprint.pprint(response)
con = sqlite3.connect("inst.db")    
con.execute("create table inst(InstanceType TEXT NOT NULL, InstanceId TEXT UNIQUE NOT NULL, VpcId TEXT NOT NULL)")  
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
con = sqlite3.connect("inst.db")   
cur = con.cursor()  
cur.execute("INSERT into inst(InstanceType, InstanceId, VpcId) values (?,?,?)",(d,c,e))  
con.commit() 
con.close()