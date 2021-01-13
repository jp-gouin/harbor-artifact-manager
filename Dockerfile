FROM golang AS gobuilder
ENV Version="2.0.0" 
WORKDIR /go/src/github.com/jpgouin/i3s-backends-validator/
COPY backend/ .
ENV GOPATH=/go/src/github.com/jpgouin/
#RUN  go get -d -v github.com/gorilla/handlers \
#                 github.com/rs/cors \
#                 github.com/alexflint/go-arg \
#                 github.com/gorilla/mux  \
#                 go.etcd.io/etcd/clientv3  \
#                 google.golang.org/grpc/grpclog  \
#                 github.com/auth0/go-jwt-middleware \
#                 github.com/jtblin/go-ldap-client \
#                 github.com/coreos/etcd/client \
#                 github.com/bndr/gojenkins \
#                 github.com/google/uuid \
#                 github.com/gorilla/websocket \
#                 google.golang.org/api/oauth2/v2 \
 #                golang.org/x/oauth2  \
 #                google.golang.org/cloud/storage
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend-$Version *.go
                 
 
# Dockerfile for lighttpd
# build environment
#FROM node:12.2.0-alpine AS build
#WORKDIR /app
#COPY create-react-app-with-typescript/package.json /app/package.json
#ENV PATH /app/node_modules/.bin:$PATH
#RUN npm install --silent
#RUN npm install react-scripts@3.0.1 -g --silent
#COPY create-react-app-with-typescript /app
#COPY backend /tmp
#RUN npm run-script build

# Dockerfile for lighttpd
# build environment
FROM node:12.2.0-alpine AS buildvue
WORKDIR /app
COPY vuegui /app
RUN npm install -g @vue/cli && npm install --silent && npm run build

# production environment
FROM nginx:1.16.0-alpine
ENV HARBORURI="https://harbor.example.com" \
    HARBORUSER="admin" \
    HARBORPASS="myPassword" \
    Version="2.0.0" 
RUN apk add --no-cache ca-certificates 

COPY --from=gobuilder /go/src/github.com/jpgouin/i3s-backends-validator/backend-$Version /usr/share/
COPY backend/backends-save-images.sh /usr/share/
#COPY --from=build /app/build /usr/share/nginx/html/old
COPY --from=buildvue /app/dist /usr/share/nginx/html
COPY start.sh /usr/local/bin/
COPY default.conf /etc/nginx/conf.d/default.conf
EXPOSE 80 8080
CMD ["start.sh"]