### Первая задача для домашки
Напишите программу, которая принимает строки через канал и подсчитывает количество слов в каждой строке. Используйте несколько горутин для обработки строк и один канал для передачи результатов.

Условно, на вход строка:
```
Всем привет!
Следующая лекция в среду
Увидимся на лекции!
```

Результат
```
Word count: 2
Word count: 4
Word count: 3
```

### Вторая задача
Реализуйте простую модель многопользовательского чата, где несколько пользователей могут отправлять сообщения в общий канал. Каждое сообщение должно содержать имя отправителя и текст сообщения. Создайте несколько горутин для имитации пользователей, которые отправляют сообщения.

Примерный вывод:
```
[User3]: Message 1 from User3
[User1]: Message 1 from User1
[User2]: Message 1 from User2
[User2]: Message 2 from User2
[User1]: Message 2 from User1
[User3]: Message 2 from User3
[User3]: Message 3 from User3
[User2]: Message 3 from User2
[User1]: Message 3 from User1
[User2]: Message 4 from User2
[User3]: Message 4 from User3
[User1]: Message 4 from User1
[User1]: Message 5 from User1
[User3]: Message 5 from User3
[User2]: Message 5 from User2
```
