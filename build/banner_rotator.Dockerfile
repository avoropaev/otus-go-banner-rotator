# Собираем в гошке
FROM golang:1.17.10 as build

ENV BIN_FILE /opt/banner-rotator/banner-rotator-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} .

# На выходе тонкий образ
FROM alpine:3.9

WORKDIR "/opt/banner-rotator"

LABEL ORGANIZATION="OTUS Online Education"
LABEL SERVICE="banner-rotator"
LABEL MAINTAINERS="awz.voropaev@gmail.com"

ENV BIN_FILE "/opt/banner-rotator/banner-rotator-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV WAITFORIT_VERSION="v2.4.1"
ENV WAIT_FOR_IT_PATH "/usr/local/bin/waitforit"
RUN wget -q -O $WAIT_FOR_IT_PATH https://github.com/maxcnunes/waitforit/releases/download/$WAITFORIT_VERSION/waitforit-linux_amd64 \
    && chmod +x $WAIT_FOR_IT_PATH

ENV CONFIG_FILE /etc/banner-rotator/config.yaml
COPY ./config/banner_rotator_config.yaml ${CONFIG_FILE}

CMD ${BIN_FILE} serve-http --config ${CONFIG_FILE}
