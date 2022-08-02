## **shopper-api** microservice

___

* 🇬🇧 [English version](www.github.com/arturyumaev/shopper-api/README.md)
* 🇷🇺 [Russian version](www.github.com/arturyumaev/shopper-api/README_rus.md)

___

## Запуск проекта в **kubernetes** кластере

[Ссылка на корневой репозиторий](www.github.com/arturyumaev/shopper)

## Запуск приложения локально

Для запуска приложения локально должны бытьу становлены переменные окружение. Они устанавливают параметры подключения к базе данных postgres.

```
POSTGRES_USER=<user>
POSTGRES_PASSWORD=<password>
POSTGRES_HOST=localhost
POSTGRES_DB=<database>
```

Далее запустить

```shell
./run_local.sh
```

Проверить доступность сервиса:

```shell
curl localhost:8080/health
```

## Полезные ссылки

https://kubernetes.github.io/ingress-nginx/

https://kubernetes.github.io/ingress-nginx/user-guide/basic-usage/
