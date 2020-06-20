PWD ?= $(shell pwd)

vue-build:
	cd ${PWD}/web/vue.js; yarn build

serve:
	cd ./web/vue.js; yarn serve

dev: vue-build
	docker-compose up -d
	cd ./web/vue.js; yarn serve

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

release:
	git push heroku master
