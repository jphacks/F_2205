# help ターゲットをデフォルトのターゲットにする
.DEFAULT_GOAL := help

# go
GO           := go run
GO_MAIN_PATH := ./src

# docker
DOCKER_COMPOSE_FRONT := docker-compose.yml
DOCKER_COMPOSE_FRONT_CONTAINER_NAME := nuxt

# rm
RM:=rm -rf

.PHONY: all ## 開発用サーバーの起動をする
all:up install dev

.PHONY: up
up: ## docker環境を立ち上げる
	docker-compose -f $(DOCKER_COMPOSE_FRONT) up -d

.PHONY: down
down: ## dockerイメージを削除し、docker環境を閉じる
	docker-compose -f $(DOCKER_COMPOSE_FRONT) down \
	--rmi all --volumes --remove-orphans

.PHONY: install
install: ## 依存関係をインストール
	docker-compose exec $(DOCKER_COMPOSE_FRONT_CONTAINER_NAME) yarn install

.PHONY: dev
dev: ## 開発用サーバを起動
	docker-compose exec $(DOCKER_COMPOSE_FRONT_CONTAINER_NAME) yarn dev

.PHONY: del-nodemodules
del-nodemodules:
	cd nuxt-app && rm -rf node_modules

.PHONY: fclean
fclean:down del-nodemodules

.PHONY: re
re:fclean all ## 完全に初期化した状態でdocker環境を立ち上げる

# TODO nodemodulesの削除コマンドを追加する
.PHONY: del-packages
del-packages:

.PHONY: help
help: ## コマンドの一覧を標示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
