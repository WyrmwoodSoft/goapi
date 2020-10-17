FROM golang:1.15

WORKDIR C:\Users\metro\Desktop\Wyrmwood\GoApi

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["GoApi"]