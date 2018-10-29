from pymongo import MongoClient
from bson.son import SON

con_str = "mongodb+srv://ciromoraismedeiros:m0ng0db*@myfirstcluster-or3vu.mongodb.net/test?retryWrites=true"

client = MongoClient()
db = client.jesus

print(db)


'''db.inventory.insert_many([
   # MongoDB adds the _id field with an ObjectId if _id is not present
   { "item": "journal", "qty": 25, "status": "A",
       "size": { "h": 14, "w": 21, "uom": "cm" }, "tags": [ "blank", "red" ] },
   { "item": "notebook", "qty": 50, "status": "A",
       "size": { "h": 8.5, "w": 11, "uom": "in" }, "tags": [ "red", "blank" ] },
   { "item": "paper", "qty": 100, "status": "D",
       "size": { "h": 8.5, "w": 11, "uom": "in" }, "tags": [ "red", "blank", "plain" ] },
   { "item": "planner", "qty": 75, "status": "D",
       "size": { "h": 22.85, "w": 30, "uom": "cm" }, "tags": [ "blank", "red" ] },
   { "item": "postcard", "qty": 45, "status": "A",
       "size": { "h": 10, "w": 15.25, "uom": "cm" }, "tags": [ "blue" ] }
])'''

for doc in db.inventory.find(
    {"size": SON([("h", 14), ("w", 21), ("uom", "cm")])}):
    print(doc)
    
# db.inventory.update({}):

