# First stage: build the application

FROM golang:1.21-alpine AS app_builder

RUN mkdir /project

ADD . /project

WORKDIR /project

RUN go build -o myoauth


# Second stage: copy only the executable and the environment files

FROM alpine:3.14

WORKDIR /root

COPY --from=app_builder /project/myoauth .

RUN mkdir /root/env

COPY --from=app_builder /project/env/.env /root/env

RUN echo -e "\nMYSQLDB_HOST=host.docker.internal" >> /root/env/.env

COPY --from=app_builder /project/public /root/public

EXPOSE $APP_PORT

CMD ["/root/myoauth"]