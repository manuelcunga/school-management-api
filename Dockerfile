FROM golang:1.21.2-alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache make bash

WORKDIR /app

COPY . .

# Build
RUN make build

# Using a distroless image from https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static-debian11

COPY --from=builder /app/server /server



# Copy the credentials file to the directory
COPY --from=builder /app/config/firebase/Credentials/school-management-b6bd7-firebase-adminsdk-v1z7j-d7301ecd03.json /app/config/firebase/Credentials
COPY --from=builder /app/bin/server /server

# Copy the .env file

EXPOSE 8000

CMD ["/server"]
