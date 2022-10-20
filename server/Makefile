# help ターゲットをデフォルトのターゲットにする
.DEFAULT_GOAL := help

# env
ENV_LOCAL_FILE := .env.local
ENV_LOCAL       = $(shell cat $(ENV_LOCAL_FILE))

# go
GO           := go run
GO_MAIN_PATH := ./src

# docker
DOCKER_FILE_DIR            := ./docker
DOCKER_COMPOSE_BASE        := $(DOCKER_FILE_DIR)/docker-compose.base.yml
DOCKER_COMPOSE_DEV_DB      := $(DOCKER_FILE_DIR)/docker-compose.dev.db.yml
DOCKER_COMPOSE_DEV_SERVER  := $(DOCKER_FILE_DIR)/docker-compose.dev.server.yml
DOCKER_EXEC                := docker exec -it
DB_CONTAINER_NAME          := f_2205-db

# dir
SRC_DIR := ./src

# volume
DATA_DIR := ./db/data

# rm
RM:=rm -rf

.PHONY: up
up: ## docker環境を立ち上げる
	$(ENV_LOCAL) docker-compose \
	-f $(DOCKER_COMPOSE_BASE) \
	-f $(DOCKER_COMPOSE_DEV_SERVER) up

.PHONY: down
down: ## dockerイメージを削除し、docker環境を閉じる
	docker-compose \
	-f $(DOCKER_COMPOSE_BASE) \
	-f $(DOCKER_COMPOSE_DEV_SERVER) down \
	--rmi all --volumes --remove-orphans

.PHONY: fclean
fclean:down del-volumes ## マウントしたデータを削除、またdockerイメージも削除する

.PHONY: re
re:fclean up ## 完全に初期化した状態でdocker環境を立ち上げる

.PHONY: del-volumes
del-volumes:del-data

.PHONY: del-data
del-data:
	$(RM) $(DATA_DIR)


.PHONY: attach-db
attach-db: ## dockerのdbコンテナにアクセスする
	$(DOCKER_EXEC) $(DB_CONTAINER_NAME) bash

.PHONY: lint
lint: ## goのコードを整形する
	cd $(SRC_DIR) && gofmt -l -w .

.PHONY: deploy
deploy: ## cloud run にdeployする
	cd src && gcloud run deploy f-2205-server --source=. --region=asia-east1 --project=serious-timer-366017 --set-env-vars=ENVIROMENT=PRODUCTION --port=8080 --allow-unauthenticated

.PHONY: help
help: ## コマンドの一覧を標示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'