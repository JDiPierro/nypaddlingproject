import requests
import json

SERVER_URL = "http://localhost:5000"

data = None
with open("./locations.json") as locationjson:
  data = json.load(locationjson)

for location in data:
  resp = requests.post(f"{SERVER_URL}/location/save", json=location)
