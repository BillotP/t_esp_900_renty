### STEP 1 : Build
FROM node:lts-alpine as build
WORKDIR /build
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build:local
## STEP 2 : Serve
FROM nginx:stable-alpine as production-stage
COPY --from=build /build/dist /usr/share/nginx/html
COPY ./config/nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
