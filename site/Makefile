# Yandex Container Registry
REGISTRY     := cr.yandex/crpagjmg3tp0j9f768ui
IMAGE        := $(REGISTRY)/pocketdoctor.medclever.com
VERSION      := $(shell cat version.md)
HUGO_VERSION := 0.163.0
HUGO_IMAGE   := $(REGISTRY)/hugo:$(HUGO_VERSION)
PLATFORM     := linux/amd64

.PHONY: dev build push hugo

# Сервер разработки (включает черновики)
dev:
	hugo server -D

# Сборка образа под платформу сервера (теги latest и текущей версии из version.md)
build:
	docker build --platform $(PLATFORM) -t $(IMAGE):latest -t $(IMAGE):$(VERSION) .

# Сборка + публикация обоих тегов в реестр
push: build
	docker push $(IMAGE):latest
	docker push $(IMAGE):$(VERSION)

# Зеркалирование builder-образа Hugo в реестр (повторить при смене HUGO_VERSION)
hugo:
	docker build -f Dockerfile.hugo --build-arg HUGO_VERSION=$(HUGO_VERSION) --platform $(PLATFORM) -t $(HUGO_IMAGE) .
	docker push $(HUGO_IMAGE)
