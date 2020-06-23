from flask import Flask, jsonify, redirect, url_for, request, abort, session
from flask_dance.contrib.facebook import make_facebook_blueprint, facebook
from flask_cors import CORS
from flask_pymongo import PyMongo
from bson import ObjectId
from datetime import datetime
import os
from functools import wraps
import sentry_sdk
from sentry_sdk.integrations.flask import FlaskIntegration


###################
# Error Reporting #
###################
sentry_sdk.init(dsn="https://9029e54585544861aaf3a574c79d0dc4@o410319.ingest.sentry.io/5284155",
                environment=os.environ.get("APP_ENV", "local"),
                integrations=[FlaskIntegration()])

#################
# Configuration #
#################
DEBUG = os.environ.get("FLASK_DEBUG", False)
MONGO_URI = os.environ.get("MONGO_URI", "mongodb://mongo:27017")
SECRET_KEY = os.environ.get("APP_SECRET_KEY", "omgwtfbbqplzdontguessthis")
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
CORS(app, resources={r'/api/login*': {'origins': '*'}})

#########
# Mongo #
#########
mongo = PyMongo(app)
DB = mongo.cx.nypaddlingproject
LOCATIONS = DB.locations
USERS = DB.users
CLAIMS = DB.location_claims


def login_required(f):
  @wraps(f)
  def decorator(*args, **kwargs):
    if not facebook.authorized:
      return redirect(url_for("facebook.login", _external=True, _scheme='https'))
    return f(*args, **kwargs)

  return decorator


def find_or_create_user(fbid, name):
  user = {
    'fbid': fbid,
    'name': name,
    "created_at": datetime.now(),
    "updated_at": datetime.now(),
    "role": "user"
  }

  found = USERS.find_one({'fbid': user['fbid']})
  if found is not None:
    user = found[0]
  else:
    inserted = USERS.insert_one(user)
    user['id'] = inserted.inserted_id

  return user


@app.route('/api/ping', methods=['GET'])
def ping_pong():
  return jsonify('pong!')


@app.route('/api/login', methods=['GET'])
def login_redirect():
  return redirect(url_for("facebook.login", _external=True, _scheme='https'))


@app.route('/api/me', methods=['GET'])
def me():
  if not facebook.authorized:
    return abort(404)

  me = facebook.get('/me').json()
  user = find_or_create_user(me['id'], me['name'])
  login(user)

  return jsonify(_strid(user))


def login(user):
  session['user'] = user


##############################
# Paddling Application Logic #
##############################


@app.route('/api/location/save', methods=['POST'])
def save():
  location = request.json

  found = LOCATIONS.find_one({'name': location['name']})
  if found is None:
    print(f"Duplicate discarded: {location['name']}")
    return jsonify("Already Exists")

  inserted = LOCATIONS.insert_one(location)
  print(f"== Saved as: {inserted.inserted_id}")
  return jsonify(str(inserted.inserted_id))


@app.route('/api/locations', methods=['GET'])
@login_required
def load_locations():
  itemcursor = LOCATIONS.find()

  return _resolve_items(itemcursor)


@app.route('/api/locations/<location_id>', methods=['GET'])
@login_required
def location_details(location_id):
  location = LOCATIONS.find_one({"_id": ObjectId(location_id)})

  claims = CLAIMS.find({"location_id": location["_id"]})
  location['claims'] = claims

  return _resolve(location)


@app.route('/api/locations/<location_id>/claim', methods=['POST'])
@login_required
def claim_location(location_id):
  user = session.get("user")
  if user is None or user.get("_id") is None:
    return abort(400, "User not found in session")

  location = LOCATIONS.find_one({"_id": ObjectId(location_id)})
  if location is None:
    return abort(404, f"Location  doesn't exist")

  inserted = CLAIMS.insert_one({
    "location_id": location_id,
    "user_id": user.get("_id"),
    "created_at": datetime.now(),
    "updated_at": datetime.now(),
    "status": "pending"
  })

  return jsonify(str(inserted.inserted_id))


def _strid(it):
  if it.get('_id'):
    it['_id'] = str(it['_id'])
  if it.get('id'):
    it['id'] = str(it['id'])
  return it


def _resolve(item):
  return jsonify(_strid(item))


def _resolve_items(cursor):
  return jsonify([_strid(item) for item in cursor])


def done():
  return jsonify("done")


if __name__ == '__main__':
  app.run(host='0.0.0.0')
