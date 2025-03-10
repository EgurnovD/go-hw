**Цель:** Создать HTTP-сервер на языке Go, который будет обрабатывать заявки студентов на поступление в университет. Сервер должен принимать данные о студентах, проверять их баллы и выводить список поступивших студентов.  

**Задачи:**
1. Создание структуры данных: Определите структуру `Student`, которая будет содержать следующие поля:  
`FullName` (строка) — полное имя студента.  
`MathScore` (целое число) — балл по математике.  
`InformaticsScore` (целое число) — балл по информатике.  
`EnglishScore` (целое число) — балл по английскому языку.  
2. Создание HTTP-сервера:
Реализуйте HTTP-сервер, который будет слушать на порту 8080.
- Обработчик для поступления:
Создайте обработчик для POST-запросов на маршрут `/apply`, который будет принимать JSON с данными студента.
В обработчике проверьте, если сумма баллов по трем предметам (математика, информатика, английский) больше или равна 14, то добавьте студента в список поступивших. В противном случае, верните сообщение о том, что студент не поступил.
- Создание студентов: Создайте трех студентов (клиентов) с заранее определенными баллами:  
Два студента должны иметь общую сумму баллов >= 14.  
Один студент должен иметь общую сумму баллов < 14.  
- Обработчик для вывода поступивших студентов: Создайте новый маршрут `/admitted`, который будет возвращать список всех студентов, которые поступили. Список должен быть представлен в формате JSON.
