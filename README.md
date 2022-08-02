## shopper-api microservice

___

* 🇬🇧 [English version](https://www.github.com/arturyumaev/shopper-api/blob/main/README.md)
* 🇷🇺 [Russian version](https://www.github.com/arturyumaev/shopper-api/blob/main/README_rus.md)

___

#### Run project in **kubernetes** cluster

[Root repository](https://www.github.com/arturyumaev/shopper)

#### Run application locally

To run application locally environment variables must be set. They specify connection to postgres database.

```
POSTGRES_USER=<user>
POSTGRES_PASSWORD=<password>
POSTGRES_HOST=localhost
POSTGRES_DB=<database>
```

Then run

```shell
./run_local.sh
```

Check service availability:

```shell
curl arch.homework/health
```

#### Usefull links

- https://kubernetes.github.io/ingress-nginx/
- https://kubernetes.github.io/ingress-nginx/user-guide/basic-usage/

