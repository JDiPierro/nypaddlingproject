from flask import Flask, jsonify, redirect, url_for, request
from flask_dance.contrib.facebook import make_facebook_blueprint, facebook
from flask_cors import CORS
from flask_pymongo import PyMongo
from pymongo import MongoClient
from bson.objectid import ObjectId
import gridfs
import os
import random
from datetime import datetime
import requests

import sentry_sdk
from sentry_sdk.integrations.flask import FlaskIntegration
sentry_sdk.init(dsn="https://9029e54585544861aaf3a574c79d0dc4@o410319.ingest.sentry.io/5284155",
                environment=os.environ.get("APP_ENV", "local"),
                integrations=[FlaskIntegration()])

#################
# configuration #
#################
DEBUG = True
MONGO_URI = os.environ.get("MONGO_URI", "mongodb://localhost:27017")
SECRET_KEY = os.environ.get("APP_SECRET_KEY", "omgwtfbbqdontguessthis")
FACEBOOK_APP_ID = os.environ.get("FB_APP_ID")
FACEBOOK_APP_SECRET = os.environ.get("FB_APP_SECRET")
APP_PROTOCOL = os.environ.get("APP_PROTOCOL", "https")
APP_DOMAIN = os.environ.get("APP_DOMAIN", "localhost:8080")

#################
# setup the app #
#################
app = Flask(__name__)
app.debug = DEBUG
app.secret_key = SECRET_KEY
app.config.from_object(__name__)
app.config['SERVER_NAME'] = APP_DOMAIN
app.config['MONGO_URI'] = MONGO_URI
app.config['PREFERRED_URL_SCHEME'] = APP_PROTOCOL

########
# Auth #
########
blueprint = make_facebook_blueprint(
  client_id=FACEBOOK_APP_ID,
  client_secret=FACEBOOK_APP_SECRET,
  scope="email"
)
app.register_blueprint(blueprint, url_prefix="/api/login")

# Mongo
mongo = PyMongo(app)

# enable CORS
CORS(app, resources={r'/*': {'origins': '*'}})

DB = mongo.cx.nypaddlingproject
LOCATIONS = DB.locations
USERS = DB.users


def find_or_create_user(fbid, name):
  user = {
    'fbid': fbid,
    'name': name,
  }

  found = USERS.find({'fbid': user['fbid']})
  if found.count() > 0:
    user = found[0]
  else:
    inserted = USERS.insert_one(user)
    user['id'] = inserted.inserted_id

  return user


@app.route('/api/ping', methods=['GET'])
def ping_pong():
  requests.get("http://echo")
  return jsonify('pong!')


@app.route('/api/login', methods=['GET'])
def login():
  if not facebook.authorized:
    return redirect(url_for("facebook.login", _external=True, _scheme='https'))

  me = facebook.get('/me').json()
  user = find_or_create_user(me['id'], me['name'])

  return jsonify(_strid(user))

##############################
# Paddling Application Logic #
##############################


@app.route('/api/location/save', methods=['POST'])
def save():
  location = request.json

  found = LOCATIONS.find({'name': location['name']})
  if found.count() > 0:
    print(f"Duplicate discarded: {location['name']}")
    return jsonify("Already Exists")

  inserted = LOCATIONS.insert_one(location)
  print(f"== Saved as: {inserted.inserted_id}")
  return jsonify(str(inserted.inserted_id))


@app.route('/api/locations', methods=['GET'])
def load_locations():
  if not facebook.authorized:
    return redirect(url_for("facebook.login", _external=True, _scheme='https'))
  itemcursor = LOCATIONS.find()

  return _resolve_items(itemcursor)


def _strid(it):
  if it.get('_id'):
    it['_id'] = str(it['_id'])
  if it.get('id'):
    it['id'] = str(it['id'])
  return it


def _resolve_items(cursor):
  return jsonify([_strid(item) for item in cursor])


def done():
  return jsonify("done")


if __name__ == '__main__':
  app.run(host='0.0.0.0')
