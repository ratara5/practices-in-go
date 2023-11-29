# FROM golang:latest

# RUN mkdir /build
# WORKDIR /build

# RUN export GO111MODULE=on
# RUN go install github.com/AdminTurnedDevOps/GoWebAPI@latest
# RUN cd /build && git clone https://github.com/AdminTurnedDevOps/GoWebAPI.git

# RUN cd /build/GoWebAPI/main && build

# EXPOSE 8080

# ENTRYPOINT [ "/build/GoWebAPI/main/main" ]

FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN cd /app/main
RUN go build -o main/main ./main
CMD ["/app/main/main"]
