import requests
import json
import os

#SERVER_URL = "https://nypaddlingproject.herokuapp.com/"

MONGO_URI = os.environ.get("MONGO_URI", "mongodb://localhost:27017")
print(f"MONGO @ {MONGO_URI}")

import pymongo
client = pymongo.MongoClient(MONGO_URI)
db = client.nypaddlingproject

LOCATIONS = db.locations

data = None
with open("./locations.json") as locationjson:
  data = json.load(locationjson)

for location in data:
  if location.get("county") == None:
    continue

  location['paddling_id'] = location['paddlingid']
  del location['paddlingid']
  found = LOCATIONS.find({'paddling_id': location['paddling_id']})
  if found.count() > 0:
    print(f"Duplicate discarded: {location['name']}")

  inserted = LOCATIONS.insert_one(location)
  print(f"== Saved as: {inserted.inserted_id}")

  #resp = requests.post(f"{SERVER_URL}/api/location/save", json=location)
