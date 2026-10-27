stargazers_count это общее число звёзд, а не прирост за период. Настоящий тренд (кто набрал больше всего за неделю) API не отдаёт.
Search API живёт по своим лимитам: без авторизации можно сделать до 10 запросов в минуту, с авторизацией до 30, и это отдельный пул от основного лимита в 60 запросов в час. Для CLI хватает, но ошибку «лимит исчерпан» надо обработать красиво. 
github
Поиск отдаёт не больше 1000 результатов, а на странице помещается максимум 100 записей. Значит, для --limit разумный потолок первой версии 100. 
github
github
Если запрос не успел выполниться за отведённое время, API возвращает найденное и ставит флаг incomplete_results в true. Можно предупреждать об этом пользователя. 
github

trending-repos/
├── cmd/
│   └── trending-repos/
│       └── main.go          5-10 строк: os.Exit(cli.Run(...))
├── internal/
│   ├── domain/
│   │   ├── repo.go          Repo, SearchQuery
│   │   ├── duration.go      Duration (day/week/month/year) и расчёт начальной даты
│   │   └── errors.go        UsageError, RateLimitError, APIError
│   ├── github/
│   │   └── client.go        запрос к API, разбор JSON, ошибки
│   ├── trending/
│   │   └── service.go       Top(ctx, duration, limit): период -> запрос -> сортировка
│   ├── render/
│   │   └── table.go         красивый вывод в io.Writer
│   └── cli/
│       ├── args.go          флаги и валидация
│       └── cli.go           Run(): склеивает всё, выбирает код выхода
├── go.mod
└── README.md