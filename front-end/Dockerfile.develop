FROM node:16

SHELL ["/bin/bash", "-c"]

ARG PROJECT_NAME='nuxt-app'

WORKDIR /var/www/html/${PROJECT_NAME}

RUN apt-get update && \
    npm install -g npm 

ARG UID=1000
USER ${UID}

EXPOSE 3000
ENV HOST 0.0.0.0