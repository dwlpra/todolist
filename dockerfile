FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o todo-api
# Set the environment variable for production mode
ENV FIBER_ENV=production


EXPOSE 3030

CMD ./todo-api