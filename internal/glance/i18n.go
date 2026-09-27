package glance

// russianWidgetTitles contains translations for Glance's built-in widget titles.
// User-provided titles remain unchanged, so existing configurations keep full control.
var russianWidgetTitles = map[string]string{
	"Calendar":            "Календарь",
	"Change Detection":    "Отслеживание изменений",
	"Clock":               "Часы",
	"Custom API":          "Пользовательский API",
	"DNS Stats":           "Статистика DNS",
	"Docker Containers":   "Контейнеры Docker",
	"Hacker News":         "Новости Hacker News",
	"IFrame":              "Встроенная страница",
	"Markets":             "Рынки",
	"Monitor":             "Мониторинг",
	"RSS Feed":            "RSS-лента",
	"Releases":            "Релизы",
	"Repository":          "Репозиторий",
	"Search":              "Поиск",
	"Server Stats":        "Статистика сервера",
	"Split Column":        "Разделённая колонка",
	"To-do":               "Задачи",
	"Top games on Twitch": "Популярные игры на Twitch",
	"Twitch Channels":     "Каналы Twitch",
	"Videos":              "Видео",
	"Weather":             "Погода",
	"Bookmarks":           "Закладки",
	"Lobsters":            "Lobsters",
}

func translateDefaultTitle(title string) string {
	if translated, ok := russianWidgetTitles[title]; ok {
		return translated
	}
	return title
}
