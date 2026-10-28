# GitHub Trending CLI

A simple and fast command-line interface (CLI) tool written in Go that fetches and displays trending repositories from GitHub. It allows you to filter repositories by time range and limit the number of results, presenting the data in a clean, readable table.

🇬🇧 [English](#english) | 🇷 [Русский](#russian)

---

## English

### ✨ Features

- **Real-time Data**: Fetches live data directly from the GitHub Search API.
- **Time Filtering**: Filter repositories by creation date (`day`, `week`, `month`, or `year`).
- **Customizable Limits**: Control how many repositories are displayed (up to 100).
- **Sorted by Stars**: Automatically sorts results by the number of stars (descending).
- **Beautiful Output**: Displays data in a well-formatted table using `text/tabwriter`.
- **Robust Error Handling**: Gracefully handles invalid inputs, network errors, and API rate limits.
- **Zero Dependencies**: Built entirely with the Go standard library.

###  Tech Stack

- **Language**: Go (1.22+)
- **API**: GitHub REST API (Search Repositories)
- **Libraries**: Only standard library (`net/http`, `encoding/json`, `flag`, `text/tabwriter`, `context`)

### 📁 Project Structure

```text
trending-repos/
├── main.go          # Entry point, flag parsing, and orchestration
├── models.go        # Data structures (Repo) and domain logic (Duration)
── github.go        # GitHub API client and HTTP request logic
├── render.go        # Terminal output formatting (tabwriter)
├── go.mod           # Go module definition
└── README.md        # This file
```

###  Getting Started

**Prerequisites:**
- Go 1.22 or newer installed on your machine.

**1. Clone the repository:**
```bash
git clone <your-repo-url>
cd trending-repos
```

**2. Build and run:**
```bash
go build -o trending-repos .
./trending-repos
```

*(Optional) To avoid GitHub API rate limits (10 requests/min for unauthenticated users), set your personal access token as an environment variable:*
```bash
export GITHUB_TOKEN=your_github_token_here
```

### 📖 Usage

The tool supports two main flags:
- `--duration`: Time range to filter repositories (`day`, `week`, `month`, `year`). Defaults to `week`.
- `--limit`: Number of repositories to display (1-100). Defaults to `10`.

**Examples:**

Show top 10 trending repos for the current week (default):
```bash
./trending-repos
```

Show top 20 trending repos for the current month:
```bash
./trending-repos --duration month --limit 20
```

Show top 5 trending repos for the current day:
```bash
./trending-repos --duration day --limit 5
```

**Example Output:**
```text
STARS   LANGUAGE   REPOSITORY                     DESCRIPTION
-----   --------   ----------                     -----------
15420   Go         owner/awesome-project          A brief description of the project...
12300   Python     owner/another-cool-repo        Another brief description...
```

### ⚙️ How It Works

1. The CLI parses the `--duration` and `--limit` flags.
2. It calculates the `since` date (e.g., `today - 7 days` for `week`) and formats it as `YYYY-MM-DD`.
3. It sends a GET request to `https://api.github.com/search/repositories` with the query `created:>YYYY-MM-DD`, sorted by stars.
4. The JSON response is parsed into Go structs.
5. The results are formatted and printed to the terminal using `text/tabwriter` for perfect column alignment.

###  Error Handling

- **Invalid Duration**: If an unsupported duration is passed, the tool exits with a clear error message.
- **Invalid Limit**: Limits outside the 1-100 range are rejected.
- **API Errors**: Network failures, 403 (Rate Limit), and 422 (Invalid Query) errors are caught and displayed in `stderr`.

### 🚧 Possible Improvements

- Add support for filtering by programming language (`--language go`).
- Implement a caching layer to avoid hitting the API on repeated requests.
- Add a `--json` flag to output raw JSON for scripting/automation.
- Add interactive mode using a library like `bubbletea`.

---

## Русский

### ✨ Возможности

- **Актуальные данные**: Получает данные напрямую из GitHub Search API в реальном времени.
- **Фильтрация по времени**: Фильтрация репозиториев по дате создания (`day`, `week`, `month` или `year`).
- **Настраиваемый лимит**: Управление количеством отображаемых репозиториев (до 100).
- **Сортировка по звездам**: Автоматическая сортировка результатов по количеству звезд (по убыванию).
- **Красивый вывод**: Отображение данных в виде аккуратной таблицы с помощью `text/tabwriter`.
- **Надежная обработка ошибок**: Корректная обработка невалидного ввода, сетевых ошибок и лимитов API.
- **Ноль зависимостей**: Написан исключительно на стандартной библиотеке Go.

### 🛠 Технологии

- **Язык**: Go (1.22+)
- **API**: GitHub REST API (Search Repositories)
- **Библиотеки**: Только стандартная библиотека (`net/http`, `encoding/json`, `flag`, `text/tabwriter`, `context`)

### 📁 Структура проекта

```text
trending-repos/
├── main.go          # Точка входа, парсинг флагов и оркестрация
├── models.go        # Структуры данных (Repo) и доменная логика (Duration)
├── github.go        # Клиент GitHub API и логика HTTP-запросов
├── render.go        # Форматирование вывода в терминал (tabwriter)
├── go.mod           # Определение Go-модуля
└── README.md        # Этот файл
```

### 🚀 Запуск

**Требования:**
- Установленный Go 1.22 или новее.

**1. Клонировать репозиторий:**
```bash
git clone <your-repo-url>
cd trending-repos
```

**2. Собрать и запустить:**
```bash
go build -o trending-repos .
./trending-repos
```

*(Опционально) Чтобы избежать лимитов GitHub API (10 запросов/мин для неавторизованных пользователей), задайте токен как переменную окружения:*
```bash
export GITHUB_TOKEN=ваш_github_токен
```

###  Использование

Инструмент поддерживает два основных флага:
- `--duration`: Временной диапазон для фильтрации (`day`, `week`, `month`, `year`). По умолчанию `week`.
- `--limit`: Количество отображаемых репозиториев (1-100). По умолчанию `10`.

**Примеры:**

Показать топ-10 трендовых репо за текущую неделю (по умолчанию):
```bash
./trending-repos
```

Показать топ-20 трендовых репо за текущий месяц:
```bash
./trending-repos --duration month --limit 20
```

Показать топ-5 трендовых репо за текущий день:
```bash
./trending-repos --duration day --limit 5
```

**Пример вывода:**
```text
STARS   LANGUAGE   REPOSITORY                     DESCRIPTION
-----   --------   ----------                     -----------
15420   Go         owner/awesome-project          A brief description of the project...
12300   Python     owner/another-cool-repo        Another brief description...
```

### ⚙️ Как это работает

1. CLI парсит флаги `--duration` и `--limit`.
2. Вычисляется дата `since` (например, `сегодня - 7 дней` для `week`) и форматируется как `YYYY-MM-DD`.
3. Отправляется GET-запрос к `https://api.github.com/search/repositories` с запросом `created:>YYYY-MM-DD`, отсортированным по звездам.
4. JSON-ответ парсится в Go-структуры.
5. Результаты форматируются и выводятся в терминал с помощью `text/tabwriter` для идеального выравнивания колонок.

### 🛡 Обработка ошибок

- **Неверный Duration**: Если передан неподдерживаемый период, инструмент завершится с понятным сообщением об ошибке.
- **Неверный Limit**: Лимиты вне диапазона 1-100 отклоняются.
- **Ошибки API**: Сетевые сбои, 403 (Rate Limit) и 422 (Invalid Query) перехватываются и выводятся в `stderr`.

### 🚧 Возможные улучшения

- Добавить фильтрацию по языку программирования (`--language go`).
- Реализовать кэширование, чтобы не дергать API при повторных запросах.
- Добавить флаг `--json` для вывода сырого JSON (для скриптов и автоматизации).
- Добавить интерактивный режим с использованием библиотеки вроде `bubbletea`.

---

## License / Лицензия

This project is created for learning purposes. / Проект создан в учебных целях.