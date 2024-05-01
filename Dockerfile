FROM golang:1.21.2-alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache make bash

WORKDIR /src

COPY . .

# Build
RUN make build

# Using a distroless image from https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static-debian11

COPY --from=builder /src/bin/app /
# COPY --from=builder /src/config/firebase/Credentials/school-management-b6bd7-firebase-adminsdk-v1z7j-d7301ecd03.json /config/firebase/Credentials/
COPY --from=builder /src/config /config

EXPOSE 8000

CMD ["/app"]
