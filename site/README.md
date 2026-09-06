# Сайт «Карманный доктор»

Статическая версия сайта [pocketdoctor.medclever.com](https://pocketdoctor.medclever.com),
перенесённая со старой технологии на генератор статических сайтов [Hugo](https://gohugo.io/).

Две языковые версии — русская на `/` и английская на `/en` — пока повторяют текущий сайт.
Дальше сюда же добавляются статьи (см. «Как добавить статью»).

## Структура

```
site/
├── hugo.toml            # конфигурация: языки (ru — по умолчанию, en — в подпапке /en)
├── version.md           # текущая версия образа (semver, одна строка)
├── Makefile             # dev / build / push / hugo
├── content/
│   ├── ru/_index.html   # русская главная (разметка перенесена со старого сайта)
│   └── en/_index.html   # английская главная
├── layouts/index.html   # общий шаблон страниц: head, счётчик Метрики, подключение CSS/JS
├── static/              # ассеты, отдаются как есть: css, js, fonts, images
├── Dockerfile           # двухэтапная сборка: Hugo (из реестра) → nginx
├── Dockerfile.hugo      # зеркало builder-образа Hugo в нашем реестре
└── nginx.conf           # конфиг nginx для раздачи собранного сайта
```

## Локальная разработка

Нужен [Hugo](https://gohugo.io/installation/) (Extended не обязателен, sass не используется):

```sh
make dev           # hugo server -D — http://localhost:1313, /en — английская версия
```

Разовая сборка статики в `public/` (каталог в .gitignore): `hugo`.

## Docker и публикация

Всем управляет `Makefile`:

```sh
make build   # собрать …:latest и …:<version> (версия из version.md)
make push    # build + публикация обоих тегов в Yandex Container Registry
make hugo    # пере-зеркалировать builder-образ Hugo (только при апгрейде версии)
```

- Образ собирается под **`linux/amd64`** — так и задумано: сервер x86_64,
  локальный Mac arm64. Вручную `docker build` без `--platform` не запускать.
- Builder — зеркало `hugomods/hugo` в нашем реестре (`Dockerfile.hugo`),
  финальная стадия — nginx.
- **Апгрейд Hugo**: поменять `HUGO_VERSION` в `Makefile` и тег в `FROM cr.yandex/…/hugo:<версия>`
  в `Dockerfile`, затем `make hugo`.
- Для `make push` нужна авторизация в реестре (`yc container registry configure-docker`).

Проверить собранный образ локально:

```sh
docker run -d --name pd-site -p 8080:80 cr.yandex/crpagjmg3tp0j9f768ui/pocketdoctor.medclever.com:latest
# сайт: http://localhost:8080
```

## Версионирование

- Источник истины — `version.md` в корне (одна строка, semver). Текущая версия — `1.0.0`
  (переезд и замена старого сайта завершены).
- При любых изменениях в проекте поднимай версию **в том же коммите** и включай `version.md` в коммит:
  - **patch** (`0.1.x`) — правки существующего: опечатки, фиксы вёрстки, правки инфраструктуры
    сборки (Dockerfile, Makefile, nginx.conf);
  - **minor** (`0.x.0`) — новый контент (статья, раздел) и новая функциональность сайта;
  - **major** (`1.0.0`) — переезд/замена старого сайта, ломающие изменения.
- `make push` публикует `:<version>` и `:latest` — история выпусков и откат по версионным тегам.

## Аналитика

Подключён только счётчик Яндекс.Метрики № 26736234 (сниппет — в `layouts/index.html`).
Google Analytics со старого сайта не переносился (Universal Analytics больше не работает).

Цели, отправляемые из `static/js/main.js` (функции `onTarget*`, вызываются инлайновыми
`onclick` в разметке): `GO_TO_STORE`, `GO_TO_GOOGLE_PLAY`, `GO_TO_APPSTORE`,
`READ_PAYMENT`, `READ_PAYMENT_GOOGLE_PLAY`, `READ_PAYMENT_APPSTORE`.

## Отличия от старого сайта

- Google Fonts подключается по `https` (по `http` современные браузеры шрифты
  со старого сайта не загружали — заголовки отображались запасным шрифтом).
- jQuery не переносился: аккордеон «Как приобрести» и отправка целей написаны
  на чистом JS (`static/js/main.js`).
- Функция переключения вкладок оплаты `paymentShowTab` была мёртвым кодом
  (кнопок-табов на странице нет, обе вкладки видны одновременно) — не перенесена.

## Как добавить статью

Статья — это пара файлов-переводов с одинаковым путём в деревьях двух языков:

```
content/ru/articles/<slug>/index.md     # русская версия
content/en/articles/<slug>/index.md     # английская версия
content/ru/articles/<slug>/illustration.png   # иллюстрации лежат рядом (page bundle)
```

Шаблон статьи — «светофор»: четыре секции `## Необходимо`, `## Можно`, `## Нельзя`,
`## Информация`, в каждой — короткие инструкции по одной на строку (списком).
Front matter:

```toml
+++
title = 'Носовое кровотечение'
weight = 18          # позиция в содержании (1–55)
description = '...'
+++
```

Ссылки на другие статьи: `[текст](/articles/<slug>/)`.

Когда статьи появятся, нужно будет добавить шаблоны `layouts/articles/single.html`
(в `static/css/style.css` уже есть готовые стили `.article-page`) и, при желании,
генерировать блок «Содержание» на главной из дочерних страниц вместо статического списка.
