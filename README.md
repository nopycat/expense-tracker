# Expense Tracker

Простой CLI-трекер расходов на Go.

Приложение позволяет добавлять, изменять, удалять и просматривать расходы, а также получать общую сумму расходов или сумму за определённый месяц.

Данные сохраняются в JSON-файл и восстанавливаются при следующем запуске программы.

## Возможности

* Добавление расхода
* Изменение расхода
* Удаление расхода
* Просмотр всех расходов
* Подсчёт общей суммы расходов
* Подсчёт расходов за определённый месяц
* Сохранение данных в JSON
* Загрузка данных из JSON при запуске

## Технологии

* Go
* Cobra
* JSON
* `encoding/json`

## Установка

Для запуска из исходного кода требуется Go.

Клонируйте репозиторий и перейдите в директорию проекта:

```bash
git clone <repository-url>
cd expenseTracker
```

Установите зависимости:

```bash
go mod download
```

Запустите приложение:

```bash
go run ./cmd/expense-tracker
```

## Сборка

Для создания исполняемого файла:

```bash
go build -o expense-tracker ./cmd/expense-tracker
```

После этого программу можно запускать напрямую:

```bash
./expense-tracker
```

В Windows:

```powershell
.\expense-tracker.exe
```

## Использование

### Добавление расхода

```bash
expense-tracker add --description "Lunch" --amount 20
```

Пример результата:

```text
Expense added successfully with ID: 1
```

### Просмотр расходов

```bash
expense-tracker list
```

Пример:

```text
ID  Date        Description  Amount
1   2026-08-08  Lunch        $20.00
2   2026-08-08  Coffee       $3.50
```

### Изменение расхода

Можно изменить описание, сумму или дату:

```bash
expense-tracker update --id 1 --description "Dinner"
```

```bash
expense-tracker update --id 1 --amount 25
```

```bash
expense-tracker update --id 1 --date 2026-08-10
```

Можно изменить несколько полей одновременно:

```bash
expense-tracker update --id 1 --description "Dinner" --amount 25
```

### Удаление расхода

```bash
expense-tracker delete --id 1
```

### Общая сумма расходов

```bash
expense-tracker summary
```

Пример:

```text
Total expenses: $48.50
```

### Сумма за месяц

```bash
expense-tracker summary --month 8
```

Пример:

```text
Total expenses for August: $48.50
```

## Хранение данных

Все расходы сохраняются в файл:

```text
data.json
```

При запуске приложения данные загружаются из JSON в `Storage`.

После выполнения команды изменения сохраняются обратно в файл.

Схема работы:

```text
data.json
    ↓
   Load
    ↓
 Storage
    ↓
   CLI
    ↓
 Storage
    ↓
   Save
    ↓
data.json
```

## Архитектура

Проект разделён на несколько частей.

### `expense`

Содержит модели данных:

* `Expense`
* `UpdateData`
* `FileData`

### `storage`

Отвечает за работу с расходами в памяти:

* `Add`
* `Update`
* `Delete`
* `List`
* `Summary`
* `SummaryByMonth`

Для хранения расходов используется:

```go
map[int]expense.Expense
```

### `cli`

Отвечает за взаимодействие с пользователем через командную строку.

Для каждой команды используется отдельный Cobra command:

```text
add
delete
list
summary
update
```

## Цель проекта

Проект создан для практики разработки CLI-приложений на Go и изучения:

* работы с пакетами и структурой проекта;
* методов и структур Go;
* работы с `map`;
* обработки ошибок;
* указателей;
* JSON;
* работы с файлами;
* Cobra;
* разделения ответственности между слоями приложения.

