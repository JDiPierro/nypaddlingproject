PWD ?= $(shell pwd)

vue-build:
	cd ${PWD}/web/vue.js; yarn build

serve:
	cd ./web/vue.js; yarn serve

server:
	python3 pyserver/server.py

dev: dbs
	docker-compose up web app

test_ci:
	echo "TODO: Write Tests..."

dbs:
	docker-compose up -d mongo mongo-express

build:
	docker build -t nypaddlingproject:latest .

heroku_push:
	heroku container:login
	docker tag nypaddlingproject registry.heroku.com/nypaddlingproject/web
	heroku container:push -a nypaddlingproject web

release-live:
	git push heroku master


release-dev:
	git push heroku-dev master

logs:
	heroku logs --tail
