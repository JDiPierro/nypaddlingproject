FROM node:10 as ui

WORKDIR /biome

ADD . .

RUN cd web/vue.js && yarn install && yarn build

## Build the server, copy the built UI over
FROM golang:1.12 as builder

WORKDIR /starter

# Get Delve for interactive debugging
RUN go get -u github.com/go-delve/delve/cmd/dlv

# Get `esc` for file embedding
RUN go get -u github.com/mjibson/esc

ADD . .
COPY --from=ui /biome/web/vue.js/dist ./web/vue.js/dist

RUN make build
RUN cp ./bin/* /bin

CMD ["starter"]
