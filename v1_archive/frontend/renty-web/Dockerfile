# étape de build
FROM node:alpine as build-stage
WORKDIR /build
COPY package*.json yarn.lock ./
RUN yarn install
COPY . .
RUN ENV=production yarn build

# étape de production
FROM node:alpine as production-stage
USER node
WORKDIR /run
COPY --chown=node:node --from=build-stage /build/ .
# Run
CMD ["node", "__sapper__/build"]
