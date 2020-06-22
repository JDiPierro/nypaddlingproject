from flask import Flask, jsonify, redirect, url_for, request, abort
from flask_dance.contrib.facebook import make_facebook_blueprint, facebook
from flask_cors import CORS
from flask_pymongo import PyMongo
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
@login_required
def load_locations():
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
