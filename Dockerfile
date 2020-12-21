FROM golang

RUN mkdir /app
WORKDIR /app

COPY ./go.mod /app
RUN go mod download

COPY . /app
RUN go build -o get_github_commit ./src

ENTRYPOINT [ "./get_github_commit" ]