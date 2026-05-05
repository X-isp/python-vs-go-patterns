# Python-vs-Go-patterns

Проект представляет собой структурированное сравнение языков Python и Go через примеры кода, алгоритмы и небольшие задачи.

Основной акцент сделан на различиях в синтаксисе, подходах к решению задач и производительности.

## Структура проекта
```text
python-vs-go-patterns/
│  
├── python/          # Примеры на Python  
│   ├── basics/  
│   ├── algorithms/  
│   └── projects/  
│  
├── go/              # Примеры на Go  
│   ├── basics/  
│   ├── algorithms/  
│   └── projects/  
│  
├── tasks/           # Одинаковые задачи на двух языках  
│  
├── benchmarks/      # Сравнение производительности  
│  
└── README.md  
```
## Идея проекта

Этот репозиторий создан для изучения различий между Python и Go на практике.

## Навигатор
- [Как запустить проект](#как-запустить-проект)  
- [Основы языка Python vs Go](#основы-языка-Python-vs-Go)  
    - [Переменные и типы](#переменные-и-типы)  
    - [Условия if/else](#условия-ifelse)  
    - [Циклы](#циклы)  
    - [Функции](#функции)  
    - [Списки и массивы](#списки-и-массивы)  
    - [Словари](#словари)
- [Ключевая идея](#ключевая-идея)  

## Как запустить проект
#### Python

1. Перейти в папку проекта:
```bash
cd python-vs-go-patterns
```

2. Активировать виртуальное окружение:
```bash
source venv/bin/activate   # Linux / macOS
venv\Scripts\activate      # Windows
```

3. Запустить файл:
```bash
python python/basics/variables.py
```

#### Go

1. Перейти в папку проекта:
```bash
cd python-vs-go-patterns
```

2. Инициализировать модуль (если ещё не сделано):
```bash
go mod init python-vs-go-patterns
```

3. Запустить файл:
```bash
go run go/basics/variables.go
```

## Основы языка Python vs Go

В этом разделе показаны базовые конструкции двух языков: переменные, условия и циклы.


### Переменные и типы

#### Python
```python
x = 10
name = "Alex"
is_active = True

print(x, name, is_active)
```

#### Go
```go
package main

import "fmt"

func main() {
    var x int = 10
    var name string = "Alex"
    var isActive bool = true

    fmt.Println(x, name, isActive)
}
```

Отличия:
- Python — динамическая типизация
- Go — строгая статическая типизация

---

### Условия if/else

#### Python
```python
x = 10

if x > 5:
    print("больше 5")
else:
    print("меньше или равно 5")
```

#### Go
```go
package main

import "fmt"

func main() {
    x := 10

    if x > 5 {
        fmt.Println("больше 5")
    } else {
        fmt.Println("меньше или равно 5")
    }
}
```

Отличия:
- Python использует отступы
- Go использует `{}` и не требует скобок в `if`

---

### Циклы

#### Python
```python
for i in range(5):
    print(i)
```

#### Go
```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

Отличия:
- Python: `for in range`
- Go: классический цикл с инициализацией и шагом

---

### Функции

Функции — это основной способ организации логики в программе.

#### Python

```python
def greet(name):
    return f"Привет, {name}"

def add(a, b):
    return a + b
```

#### Go

```go
func greet(name string) string {
    return "Привет, " + name
}

func add(a int, b int) int {
    return a + b
}
```

Отличия:
- Python не требует указания типов
- Go требует строгую типизацию входных и выходных данных
- Go более “явный”, Python более “гибкий”

---

### Списки и массивы

Списки / массивы — одна из базовых структур данных, используемых для хранения коллекций элементов.

#### Python

```python
numbers = [1, 2, 3, 4, 5]

# добавление элемента
numbers.append(6)

# обход списка
for n in numbers:
    print(n)

# генерация нового списка (квадраты)
squares = [n * n for n in numbers]
print(squares)
```

#### Go

```go
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}

    // добавление элемента
    numbers = append(numbers, 6)

    // обход слайса
    for _, n := range numbers {
        fmt.Println(n)
    }

    // создание нового слайса (квадраты)
    squares := make([]int, len(numbers))

    for i, n := range numbers {
        squares[i] = n * n
    }

    fmt.Println(squares)
}
```

Отличия:
- Python: встроенные удобные конструкции (list comprehensions)
- Go: более явная работа с памятью и индексами
- Python проще и короче
- Go более строгий и контролируемый

### Словари

Словарь — структура данных для хранения пар ключ-значение.

#### Python (dict)

```python
user = {
    "name": "Alex",
    "age": 25,
    "active": True
}

print(user["name"])

# добавление нового ключа
user["city"] = "Vilnius"

print(user)
```

#### Go (map)

```go
package main

import "fmt"

func main() {
    user := map[string]interface{}{
        "name":   "Alex",
        "age":    25,
        "active": true,
    }

    fmt.Println(user["name"])

    // добавление значения
    user["city"] = "Vilnius"

    fmt.Println(user)
}
```

Отличия:
- Python: проще синтаксис (`dict`)
- Go: нужно явно объявлять тип `map[key]value`
- Go более строгий, Python более гибкий

## Ключевая идея

- Python — читаемость и простота
- Go — явность и строгая структура

Один и тот же алгоритм в Go обычно более “явный”, а в Python — более “компактный”.