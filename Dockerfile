FROM ubuntu:rolling

LABEL maintainer="BlueTeam" version="1.0"

WORKDIR /app

COPY server-pulsa-app /app/server-pulsa-app

RUN chmod +x /app/server-pulsa-app

ENV DB_HOST=167.172.91.111
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=rahasia
ENV DB_NAME=server_pulsa_db
ENV DB_DRIVER=postgres
ENV API_PORT=8080
ENV TOKEN_ISSUE=Enigma Camp Incubation Class
ENV TOKEN_SECRET=Golang Incubation Class
ENV TOKEN_EXPIRE=120
ENV BASE_URL_MIDTRANS=https://app.sandbox.midtrans.com/snap/v1/transactions
ENV SERVER_KEY_MIDTRANS='U0ItTWlkLXNlcnZlci1FaWtzTGtwb2VRNkJ3UmFvQkFPTzhXZVI='

ENTRYPOINT ["/bin/sh", "-c", "/app/server-pulsa-app"]