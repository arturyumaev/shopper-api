# shopper-api

### Запуск приложения локально

```shell
./run_local.sh
```

### Сборка в kubernetes кластере

Для запуска на MacOS с драйвером `virtualbox` я использовал [minikube](https://minikube.sigs.k8s.io/docs/start/) и [virtualbox](https://www.virtualbox.org/wiki/Downloads).

Установить [helm](https://helm.sh/), если он не был установлен. Запустить minikube с драйвером `virtualbox` командой

```shell
minikube start --memory=4096 --driver=virtualbox
```

### Загрузка манифестов и запуск приложения в кластере

Перейти в директорию `./k8s/` и запустить скрипт

```shell
cd ./k8s && ./apply.sh
```

Получить `ip` кластера в `minikube`:

```shell
minikube ip
```

После получения `ip` добавить следующую строчку в конец файла `/etc/hosts/` (`<ip>` заменить на полученный `ip`):

```shell
<ip> arch.homework
```

Проверить доступность сервиса:

```shell
curl arch.homework/health
```

### Полезные ссылки

https://kubernetes.github.io/ingress-nginx/

https://kubernetes.github.io/ingress-nginx/user-guide/basic-usage/

