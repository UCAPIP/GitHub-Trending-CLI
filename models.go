package main

// Repo — структура репозитория из ответа GitHub
type Repo struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Language    string `json:"language"`
	URL         string `json:"html_url"`
}

// Duration — период поиска
type Duration string

const (
	Day   Duration = "day"
	Week  Duration = "week"
	Month Duration = "month"
	Year  Duration = "year"
)

// Days возвращает количество дней для периода
func (d Duration) Days() int {
	switch d {
	case Day:
		return 1
	case Week:
		return 7
	case Month:
		return 30
	case Year:
		return 365
	default:
		return 7
	}
}

// IsValid проверяет валидность периода
func (d Duration) IsValid() bool {
	switch d {
	case Day, Week, Month, Year:
		return true
	}
	return false
}
