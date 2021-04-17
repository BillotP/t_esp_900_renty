FROM python:3.7.3-alpine as base

FROM base as builder

RUN apk add --no-cache python3 python3-dev py-pip build-base

RUN mkdir /install/
WORKDIR /install

COPY requirements.txt requirements.txt

RUN pip install -r requirements.txt

FROM base

COPY --from=builder /usr/local /usr/local
RUN mkdir /app/

COPY models /app/models
COPY photos /app/photos
COPY main.py /app/main.py

CMD ["python", "/app/main.py"]
