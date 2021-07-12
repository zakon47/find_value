## Пакет разбивает строку до ближайшего(ей) числа/строки

### Описание:
Метод пытается найти запрашиваемое значение и результат помещает в [0] ячейку - остальной код в [1] ячейку ответа

### Примеры:
Поиск числа слева-направо
```
Number("name33") => ["", "name33"]
Number("33name") => ["33", "name"]
```
Поиск числа справа-налево
```
NumberReverse("name33") => ["33", "name"]
NumberReverse("33name") => ["", "33name"]
```
Поиск строки слева-направо
```
String("name33") => ["name", "33"]
String("33name") => ["", "33name"]
```
Поиск строки справа-налево
```
StringReverse("name33") => ["", "name33"]
StringReverse("33name") => ["name", "33"]
```
