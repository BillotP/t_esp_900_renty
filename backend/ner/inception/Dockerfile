FROM alpine:latest
RUN apk --no-cache add openjdk11-jdk gcompat
WORKDIR /run
COPY inception-app-standalone-0.17.0.jar .
ENTRYPOINT ["java","-jar","/run/inception-app-standalone-0.17.0.jar"]