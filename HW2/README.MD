### Задача 1:  
*Описание задачи:*  
Создайте программу для управления библиотекой. Каждый экземпляр книги должен иметь название, автора, год издания и статус (доступна или на руках у читателя). Добавьте возможность добавления новой книги, поиска книги по названию, выдачи книги читателю и возврата книги.

*Требования:*  
- Используйте структуры для представления книги.
- Напишите методы для добавления книги, выдачи книги читателю и возврата книги.
- Используйте функции для поиска книги по названию.
- Используйте циклы для вывода списка всех книг.

*Подсказка:*
- Метод `Issue` меняет статус книги на "на руках у читателя".
- Метод `Return` меняет статус книги на "доступна".

### Задача 2:  
*Описание задачи:*
Создайте программу, которая поможет пользователям учитывать свои ежемесячные расходы. Используйте карту, где ключом будет название категории расходов (например, "Продукты", "Транспорт", "Развлечения"), а значением - сумма расходов по этой категории.

*Требования:*
Используя функции, пользователь должен иметь возможность добавлять новые категории и записывать расходы по каждой из них.
Также добавьте функцию для подсчета общей суммы расходов и вывода ее на экран

### Задача 3:  
*Описание задачи:*
1. Объяснить вывов при defer tryTest()()
2. Объяснить вывод при defer tryTest()
```
package main

import (
  "fmt"
  "time"
)

func tryTest() func() {
  fmt.Println("tryTest")
  return func() {
    fmt.Println("tryTest2")
  }
}

func main() {
  defer fmt.Println("Первое время:", time.Now())
  defer tryTest()()
  time.Sleep(2 * time.Second)
  defer fmt.Println("Второе время", time.Now())

}
```
