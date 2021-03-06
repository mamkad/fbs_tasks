# Тестовое задание Junior Golang Developer
## Фибоначчи
*Реализовать сервис, возвращающий срез последовательности чисел из ряда Фибоначчи.
Сервис должен отвечать на запросы и возвращать ответ. В ответе должны быть перечислены все числа, последовательности Фибоначчи с порядковыми номерами от x до y.*

Требования:

1. Требуется реализовать два протокола: HTTP REST и GRPC
2. Код должен быть выложен в репозиторий с возможность предоставления доступа (например github.com, bitbucker.org, gitlab.com). Решение предоставить ссылкой на этот репозиторий.
3. Необходимо продумать и описать в readme развертку сервиса на другом компьютере
4. (Опционально) Кэширование. Сервис не должен повторно вычислять числа из ряда Фибоначчи. Значения необходимо сохранить в Redis или Memcache.
5. (Опционально) Код должен быть покрыт тестами. 

## HTTP REST
Исходный код: [ссылка](https://github.com/mamkad/fbs_tasks/tree/main/rest_http)

Для компиляции достаточно `go build main.go`

После запуска main сервер будет запущен, в браузере необохдимо зайти на адрес: `http://localhost:8080/input`. Откроется форма ввода, куда можно ввести два числовых значения через запятую. В случае неправильного ввода пользователь будет перенаправлен обратно в форму ввода. После нажатия кнопки отправки запроса будет загружена другая страница, содержащая таблицу выходных значений.

## GRPC
Исходный код: [ссылка](https://github.com/mamkad/fbs_tasks/tree/main/gRPC)

Необходимы следующие зависимости:

    go get -u google.golang.org/grpc
    go get github.com/golang/protobuf/protoc-gen-go
    github.com/protocolbuffers/protobuf
    
Для компиляции достаточно `go build main.go`

После запуска main сервер будет запущен. Для проверки работы был создан клиент, он находится в файле: [client.go](https://github.com/mamkad/fbs_tasks/blob/main/gRPC/client.go). Клиент компилируется отдельно и запускается после запуска сервера. Клиент соединяется по каналу 8080 с сервером и посылает ему два числа. В ответе будут получены номера и числа фибоначчи.

## Реализованные пункты
Были реализованы все пункты, кроме 4.
