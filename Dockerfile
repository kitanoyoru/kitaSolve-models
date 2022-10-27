FROM golang:1.16 AS build

LABEL maintainer="Alexandr Rutkovskij <kitanoyoru@protonmain.com>"

ENV PROJECT_DIR /app/

WORKDIR $PROJECT_DIR
COPY /src/ $PROJECT_DIR

RUN go mod tidy
RUN go build -o /bin/kitasolve_models app.go
RUN rm -rf $PROJECT_DIR

FROM debian:buster-slim AS production

ENV USER kitanoyoru

RUN useradd -m -U -d /home/$USER -s /bin/bash

RUN set -ex; \
  BUILD_DEPS='vim'; \
  apt-get update; \
  apt-get install -y --no-install-recommends $BUILD_DEPS;

RUN mkdir /var/run/kitasolve_models
RUN chown $USER:$USER /var/run/kitasolve_models

COPY --from=build /bin/kitasolve_models /usr/local/bin/kitasolve_models

USER $USER
WORKDIR /home/$USER

EXPOSE 5000

CMD ["kitasolve_models"]
