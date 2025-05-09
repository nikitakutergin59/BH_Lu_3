# BH_Lu_3

# Описание
Распределённый вычислитель арифметических выражений с реализованной афторизацие
пользователей. Представляет из себя три микро сервера

1: BH - отвечает за афторизацию, выдачю JWT и взаимодейтвия с базой данных, а так-же
управляет работой всего проекта. Запущен на порту 9090.

2: Orchestrator - отвечает за дробления выражений на Tasks для последующего вычисления. Запущен на порту 8080.

3: Lu - отвечает за вычисление Task и отправку результата orchestratoru. Запущен на порту 8081.

## Запуск
1. Установленный [Docker](https://www.docker.com/).

2. Перейдите в папку с проектом.
    ```bash
    cd BH_Lu3/BH_Lu_3
    ```
3. Постройте докер образ.
    ```bash
    docker build -t bh_lu_3_app .
    ```
4. Запустите созданный образ.
   ```bash
    docker run -d -p 9090:9090 -p 8080:8080 -p 8081:8081 bh_lu_3_app
    ```
   -**Вы увидете container_id, пример**:
    ```bash
    da43020c2034700eb4e537074a2baba7dc018d829277040d26e45d66349fad14
    ```
    
5. Проверьте всё ли коректно запустилось.
    ```bash
    docker logs <container_id>
    ```
    -**Вы должны увидеть**:
    ```bash
    2025/05/09 11:52:33 Демон запущен на порту 8081
    2025/05/09 11:52:33 Оркестратор запущен на порту 8080
    2025/05/09 11:52:33 Таблица user_data успешно создана или уже существует.
    2025/05/09 11:52:33 База данных созданна
    2025/05/09 11:52:33 Сервер запущен на порту 9090
    ```
    
## Взаимодействия и примеры запросов
Откройте cmd (Win+R)
1. Регистраци
    Отправте запрос(придумайте логин и пароль)
    ```bash
    curl -X POST -H "Content-Type: application/json" -d "{\"login\":\"<your_login>\", \"password\":\"<your_password>\"}" http://localhost:9090/register
    ```
    -**Вы получите ответ, пример**:
    ```                     
    {"hash_login":"dda844144da6949024dc7cc621a47f73d0b3357fd1f5c5c4f9167bf026414aa7","hash_password":"b5b6f2b7707de42254c0e13ce2a2f53ce9e4bbbd282c35cb61ba70470c331440",
    "message_0":"Данные успешно сохранены"}
    ```
2. Авторизация
    Отправте запрос
    ```bash
    curl -X POST -H "Content-Type: application/json" -d "{\"login\":\"your_hash_login\", \"password\":\"your_hash_password\"}" http://localhost:9090/login
    ```
    -**Например**:
    ```bash
    curl -X POST -H "Content-Type: application/json" -d "{\"login\":\"dda844144da6949024dc7cc621a47f73d0b3357fd1f5c5c4f9167bf026414aa7\",         \"password\":\"b5b6f2b7707de42254c0e13ce2a2f53ce9e4bbbd282c35cb61ba70470c331440\"}" http://localhost:9090/login
    ```
    -**Вы получите ответ пример**:
    ```
    {"message":"успешный вход","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkZGE4NDQxNDRkYTY5NDkwMjRkYzdjYzYyMWE0N2Y3M2QwYjMzNTdmZDFmNWM1YzRmOTE2N2JmMDI2NDE0YWE3IiwiZXhwIjoxNzQ2NzI5NzIwLCJpYXQiOjE3NDY2NDMzMjB9.m3mwt3mgpcnoTQ23cEeVZUkyDP5eZEhA03jqiMgwAY0"}
    ```
    //Да уж он получился очень длинный(это нормально)
   
    -**Важно**:
   
    Токен действует только 24 часа!
   
    Если вы столкнулись с тем что вы всё ввели правильно, а BH_Lu_3 выдаёт ошибку
    -**неверный логин или пароль**
    то придумайте новый логин и пароль и попробуйте занова(смотреть пункт 1)

3. Вычисления. Поздравляю вы афторизировались и теперь можно пользоваться всем доступный функционалом.
    
    Создайте запрос на вычисление
    ```bash
    curl -X POST -H "Authorization: Bearer <your_jwt_token>" -H "Content-Type: application/json" -d "{\"expr\": \"<your_expression>\"}" http://localhost:9090/calculator
    ```
    -**Пример**:
    ```bash
    curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkZGE4NDQxNDRkYTY5NDkwMjRkYzdjYzYyMWE0N2Y3M2QwYjMzNTdmZDFmNWM1YzRmOTE2N2JmMDI2NDE0YWE3IiwiZXhwIjoxNzQ2NzI5NzIwLCJpYXQiOjE3NDY2NDMzMjB9.m3mwt3mgpcnoTQ23cEeVZUkyDP5eZEhA03jqiMgwAY0" -H "Content-Type: application/json" -d "{\"expr\": \"2*(5+3)-(4+6)/2\"}" http://localhost:9090/calculator
    ```

    -**В ответ вы получите, пример**:
    ```
    {"message":"выражение обработано","result":"11","token_id":"dda844144da6949024dc7cc621a47f73d0b3357fd1f5c5c4f9167bf026414aa7"}
    ```
    -**Вот ещё несколько примеров**:
    ```bash
    curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkZGE4NDQxNDRkYTY5NDkwMjRkYzdjYzYyMWE0N2Y3M2QwYjMzNTdmZDFmNWM1YzRmOTE2N2JmMDI2NDE0YWE3IiwiZXhwIjoxNzQ2NzI5NzIwLCJpYXQiOjE3NDY2NDMzMjB9.m3mwt3mgpcnoTQ23cEeVZUkyDP5eZEhA03jqiMgwAY0" -H "Content-Type: application/json" -d "{\"expr\": \"3+(2*(7-4))\"}" http://localhost:9090/calculator
    ```
    -**Ответ**:
    ```
    {"message":"выражение обработано","result":"1","token_id":"dda844144da6949024dc7cc621a47f73d0b3357fd1f5c5c4f9167bf026414aa7"}
    ```
    -**Пример**:
    ```bash
    curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkZGE4NDQxNDRkYTY5NDkwMjRkYzdjYzYyMWE0N2Y3M2QwYjMzNTdmZDFmNWM1YzRmOTE2N2JmMDI2NDE0YWE3IiwiZXhwIjoxNzQ2NzI5NzIwLCJpYXQiOjE3NDY2NDMzMjB9.m3mwt3mgpcnoTQ23cEeVZUkyDP5eZEhA03jqiMgwAY0" -H "Content-Type: application/json" -d "{\"expr\": \"(10+5)*(6-2)\"}" http://localhost:9090/calculator
    ```
    -**Ответ**:
    ```
    {"message":"выражение обработано","result":"60","token_id":"dda844144da6949024dc7cc621a47f73d0b3357fd1f5c5c4f9167bf026414aa7"}
    ```
    -**Пример**:
    ```bash
    curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkZGE4NDQxNDRkYTY5NDkwMjRkYzdjYzYyMWE0N2Y3M2QwYjMzNTdmZDFmNWM1YzRmOTE2N2JmMDI2NDE0YWE3IiwiZXhwIjoxNzQ2NzI5NzIwLCJpYXQiOjE3NDY2NDMzMjB9.m3mwt3mgpcnoTQ23cEeVZUkyDP5eZEhA03jqiMgwAY0" -H "Content-Type: application/json" -d "{\"expr\": \"((8+2)*3-4)/2\"}" http://localhost:9090/calculator
    ```
    -**Ответ**:
    ```
    {"message":"выражение обработано","result":"2","token_id":"dda844144da6949024dc7cc621a47f73d0b3357fd1f5c5c4f9167bf026414aa7"}
    ```
    **Если вы введёте выражение которое не может быть посчитано или введёте нечего то поле result будет пустым, но ошибка обработанна и информация об этом храниться в логах**

## Получение выражений и результатов по JWT
Отправте запрос 
 
    curl -X GET -H "Authorization: Bearer <your_jwt_token>" http://localhost:9090/my_data
 
-**Пример**: 

    curl -X GET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkZGE4NDQxNDRkYTY5NDkwMjRkYzdjYzYyMWE0N2Y3M2QwYjMzNTdmZDFmNWM1YzRmOTE2N2JmMDI2NDE0YWE3IiwiZXhwIjoxNzQ2NzI5NzIwLCJpYXQiOjE3NDY2NDMzMjB9.m3mwt3mgpcnoTQ23cEeVZUkyDP5eZEhA03jqiMgwAY0" http://localhost:9090/my_data
    
-**Ответ**:
    
    {"expressions":"2*(5+3)-(4+6)/2;3+(2*(7-4));(10+5)*(6-2);;10/0;((8+2)*3-4)/2;","results":"11;1;60;;;2;"}

## Заключение
Настоятельно рекомендую читать логи, там очень много информации помогающей понять как работает сервер. BH
