from flask import Flask, jsonify, request
from flask_cors import CORS
from flask_pymongo import PyMongo
from pymongo import MongoClient
from bson.objectid import ObjectId
import gridfs
import os
import random
from datetime import datetime

# configuration
DEBUG = True
MONGO_URI = os.environ.get("MONGO_URI", "mongodb://localhost:27017")

# instantiate the app
app = Flask(__name__)
app.config.from_object(__name__)
app.config['MONGO_URI'] = MONGO_URI


# Mongo
mongo = PyMongo(app)

# enable CORS
CORS(app, resources={r'/*': {'origins': '*'}})


IMGDB = mongo.cx.nypaddlingproject
LOCATIONS = IMGDB.locations

# sanity check route
@app.route('/ping', methods=['GET'])
def ping_pong():
  return jsonify('pong!')


@app.route('/location/save', methods=['POST'])
def save():
  location = request.json

  found = LOCATIONS.find({'name': location['name']})
  if found.count() > 0:
    print(f"Duplicate discarded: {location['name']}")
    return jsonify("Already Exists")

  inserted = LOCATIONS.insert_one(location)
  print(f"== Saved as: {inserted.inserted_id}")
  return jsonify(str(inserted.inserted_id))


@app.route('/locations', methods=['GET'])
def load_locations():
  itemcursor = LOCATIONS.find()

  return _resolve_items(itemcursor)


def _strid(it):
  if it.get('_id'):
    it['_id'] = str(it['_id'])
  if it.get('fsid'):
    it['fsid'] = str(it['fsid'])
  return it


def _resolve_items(cursor):
  return jsonify([_strid(item) for item in cursor])


def done():
  return jsonify("done")


if __name__ == '__main__':
  app.run(host='0.0.0.0')
