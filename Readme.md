#UserManager
Сервис UserManager позволяет добавлять и возвращать сущность User.

* `User` состоит из одного поля `full_name` и хранится в таблице `user`
  в базе данных POSTGRES. Информация о базе и
  сервере передается в файл конфигурации `config.json`
* Добавление `User` происходит с помощью метода POST с параметром `p`
    * Если параметр пустой, сервер выводит `Empty request`
    * Если пользователь с переданным именем уже существует, сервер выводит
      `User <name> already exists`
    * При успешном добавлении сервер выводит `Successful POST request: <name>`
* Получение `User` происходит с помощью метода GET с параметром `u`
    * Если параметр пустой, сервер выводит `Empty request`
    * Если пользователя с переданным именем не существует, сервер выводит
    `No user found`
    * При успешном получении сервер выводит `Successful GET request: <name>`
* Запуск сервера осуществляется функцией `StartUserManager`
  
Используемые сторонние пакеты:
* [https://github.com/go-chi/chi](https://github.com/go-chi/chi)
* [github.com/lib/pq](github.com/lib/pq)