# Build the frontend so we can serve it statically
FROM node:14.4.0-alpine as vue

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./web/vue.js/package*.json ./
RUN yarn install

COPY ./web/vue.js .

FROM vue as build-vue

RUN yarn build

# Build the Server image
FROM nginx:stable-alpine as build-server

WORKDIR /app

RUN apk update && apk add --no-cache python3 git && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --upgrade pip setuptools && \
    if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi && \
    if [[ ! -e /usr/bin/python ]]; then ln -sf /usr/bin/python3 /usr/bin/python; fi && \
    rm -r /root/.cache
RUN apk update && apk add gcc python3-dev musl-dev

COPY --from=build-vue /app/dist /usr/share/nginx/html
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf

COPY ./pyserver/requirements.pip ./
RUN pip install -r requirements.pip
RUN pip install gunicorn

COPY ./pyserver .

CMD nginx && gunicorn -b 0.0.0.0:5000 server:app --reload

FROM build-server as heroku

CMD gunicorn -b 0.0.0.0:5000 server:app --daemon && \
          sed -i -e 's/$PORT/'"$PORT"'/g' /etc/nginx/conf.d/default.conf && \
          nginx -g 'daemon off;'
