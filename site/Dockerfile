# Сборка статического сайта (зеркало hugomods/hugo в Yandex Container Registry, см. make hugo)
FROM cr.yandex/crpagjmg3tp0j9f768ui/hugo:0.163.0 AS build
WORKDIR /src
COPY . /src
RUN hugo --destination /public

# Раздача статики через nginx.
FROM nginx:1.29-alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /public /usr/share/nginx/html
EXPOSE 80
