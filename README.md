# shopper-api

Для запуска на MacOS с драйвером `virtualbox` я использовал [minikube](https://minikube.sigs.k8s.io/docs/start/) и [virtualbox](https://www.virtualbox.org/wiki/Downloads).

### Сборка в kubernetes кластере

Запустить minikube с драйвером `virtualbox` командой

```
$ minikube start --memory=4096 --driver=virtualbox
```

Перейти в директорию `./k8s/` и установить ingress controller через `helm` (заранее установить `helm`, если не установлен)

```
$ kubectl create namespace m
$ helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/
$ helm repo update
$ helm install nginx ingress-nginx/ingress-nginx --namespace m -f nginx-ingress.yaml
```

Применить манифесты

```
$ kubectl apply -f deployment.yaml
$ kubectl apply -f service.yaml
$ kubectl apply -f ingress.yaml
```

Получить `ip` кластера в `minikube`:

```
$ minikube ip
```

После получения `ip` добавить следующую строчку в конец файла `/etc/hosts/` (`<ip>` заменить на полученный `ip`):

```
<ip> arch.homework
```

Проверить доступность сервиса запросом и добавить необходимый путь:

```
$ curl arch.homework
```

### Полезные ссылки

https://kubernetes.github.io/ingress-nginx/

https://kubernetes.github.io/ingress-nginx/user-guide/basic-usage/

