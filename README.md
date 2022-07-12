# shopper-api

Для запуска на MacOS с драйвером `virtualbox` я использовал [minikube](https://minikube.sigs.k8s.io/docs/start/) и [virtualbox](https://www.virtualbox.org/wiki/Downloads).

### Сборка в kubernetes кластере

Запустить minikube с драйвером `virtualbox` командой

```
$ minikube start --memory=4096 --driver=virtualbox
```

Установить nginx ingress controller через minikube, он установится в namespace `ingress-nginx`

```
$ minikube addons enable ingress
```

Проверить, что под с ingress контроллером запустился

```
$ kubectl get pods -n ingress-nginx
```

Перейти в директорию `/k8s` в проекте и раскатить манифесты сервиса командой:

```
$ kubectl apply -f .
```

Дождаться запуска подов:

```
$ kubectl get pods --watch
```

Дождаться получения `ip` в поле `ADDRESS` у ингресса

```
$ kubectl get ing shopper-api-ingress --watch
```

После получения `ip` добавить следующую строчку в конец файла `/etc/hosts/` (`<ip>` заменить на полученный `ip`):

```
<ip> arch.homework
```

Проверить доступность сервиса запросом:

```
$ curl arch.homework
```
