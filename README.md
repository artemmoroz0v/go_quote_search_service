# Сервис для поиска статьи по точному совпадению наименования с импортом данных цитат википедии.

Данный HTTP API микросервис принимает JSON и отвечает в формате JSON. Для разработки микросервиса был использован микрофреймворк ***gin***, объектно-реляционная СУБД ***PostgreSQL***, а также http-клиент для тестирования API ***Postman***. Написан на языке ***Golang***.

### Подробное описание проделанной работы
1. В самом начале мною был разархивирован файл с цитатами, на выходе был получен ***JSON-файл***.

2. После этого в файле ***main.go*** идет блок с кодом подключения к базе данных и последующим ***импортом*** всех данных, полученных из ***JSON***.

3. Далее, при запуске программы, вместе с базой данных поднимается и локальный ***http***-сервер.

4. Были реализованы три метода: два метода ***GET***, отвечающие за получение статьи по названию и за получение информации о статьях по категории, а также один метод ***PUT***, отвечающий за редактирование статьи пользователем.

### Как запускать и инициализировать базу данных?

В переменной ***connection_line*** хранятся условия подключения к базе данных PostgreSQL. Чтобы запустить с другой машины, нужно будет поменять поле ***password***. При запуске автоматически создается база данных и в нее импортируются все объекты, полученные в результате парсинга ***JSON***. Запуск: ***go run main.go***.

### Пример работы сервиса
Разберем все 3 функции на примерах:
1. Метод ***GET***, поиск статьи по точному совпадению наименования.

Запрос в программе ***Postman***:

![Screenshot](https://github.com/artemmoroz0v/go_quote_search_service/blob/main/screenshots/1.png)


Результат поиска статьи:

![Screenshot](https://github.com/artemmoroz0v/go_quote_search_service/blob/main/screenshots/2.png)



2. Метод ***GET***, получение информации о статьях по названию категории

Запрос в программе ***Postman***:

![Screenshot](https://github.com/artemmoroz0v/go_quote_search_service/blob/main/screenshots/3.png)


Результат поиска статьи:

![Screenshot](https://github.com/artemmoroz0v/go_quote_search_service/blob/main/screenshots/4.png)



3. Метод ***PUT***, редактирование статьи пользователем.

До изменений:

![Screenshot](https://github.com/artemmoroz0v/go_quote_search_service/blob/main/screenshots/temp.png)

Запрос в программе ***Postman***:

![Screenshot](https://github.com/artemmoroz0v/go_quote_search_service/blob/main/screenshots/5.png)

Результат изменений в базе данных ***PostgreSQL***:

![Screenshot](https://github.com/artemmoroz0v/go_quote_search_service/blob/main/screenshots/6.png)

