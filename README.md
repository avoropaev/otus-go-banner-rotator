## Финальный проект курса [OTUS. Golang Developer. Professional](https://otus.ru/lessons/golang-professional/)

[![Go Report Card](https://goreportcard.com/badge/avoropaev/otus-go-banner-rotator)](https://goreportcard.com/report/avoropaev/otus-go-banner-rotator)
![ci](https://github.com/avoropaev/otus-go-banner-rotator/actions/workflows/tests.yaml/badge.svg)

## Техническое задание

[ТЗ на разработку сервиса "Ротация баннеров"](https://github.com/OtusGolang/final_project/blob/master/02-banners-rotation.md)

[Обязательные требования для проекта](https://github.com/OtusGolang/final_project)

## Пример запуска

```shell
git clone git@github.com:avoropaev/otus-go-banner-rotator.git
cd otus-go-banner-rotator
cp .env.dist .env
cp .env.test.dist .env.test
make run
```

Потыкать можно например с помощью [BloomRPC](https://github.com/bloomrpc/bloomrpc)

- .proto лежит в `api/app.proto`
- адрес сервера `localhost:8889`
- rabbitmq ui `localhost:15672` (логин: guest, пароль: guest)
- креды для postgresql - хост: localhost, порт: 5435, логин: postgres, пароль: password, бд: postgres

## Makefile

- `make build` - создать бинарник и сбилдить все образы
- `make up` - запустить только базу и кролика
- `make down` - отключить базу и кролика
- `make migrate` - запустить миграции
- `make run` - запустить все контейнеры (база, кролик и приложение)
- `make stop` - отключить все контейнеры (база, кролик и приложение)
- `make test` - запустить тесты
- `make lint` - запустить линтер
- `make generate` - сгенерировать протобаф
- `make integration-tests` - запустить интеграционные тесты (будет поднято отдельное окружение, которое удалится после выполнения тестов)
