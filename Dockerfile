from golang:1.18-alpine

WORKDIR /app
COPY . ./
RUN apk add git

RUN go mod download

RUN go build -o icanhazquotes

EXPOSE 9999

CMD [ "./icanhazquotes" ]