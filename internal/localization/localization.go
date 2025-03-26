package localization

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/erupshis/tg-bot/locales"
	"github.com/spf13/viper"
)

type Localizer struct {
	lang  string
	viper *viper.Viper
	cache map[string]string
	mutex sync.RWMutex
}

// New создает новый экземпляр локализатора для указанного языка
func New(lang string) (*Localizer, error) {
	l := &Localizer{
		lang:  lang,
		viper: viper.New(),
		cache: make(map[string]string),
	}

	// Настраиваем Viper для чтения YAML файлов
	l.viper.SetConfigName(lang)
	l.viper.SetConfigType("yaml")
	l.viper.AddConfigPath(filepath.Join("locales"))

	// Загружаем файл локализации
	if err := l.viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read locale file: %w", err)
	}

	return l, nil
}

// Get возвращает локализованное сообщение по ключу
func (l *Localizer) Get(key locales.MessageKey) string {
	// Проверяем кэш сначала
	l.mutex.RLock()
	if val, ok := l.cache[string(key)]; ok {
		l.mutex.RUnlock()
		return val
	}
	l.mutex.RUnlock()

	// Получаем значение из Viper
	val := l.viper.GetString(string(key))
	if val == "" {
		return fmt.Sprintf("[[translation missing: %s]]", key)
	}

	// Сохраняем в кэш
	l.mutex.Lock()
	l.cache[string(key)] = val
	l.mutex.Unlock()

	return val
}

// Getf возвращает форматированное локализованное сообщение
func (l *Localizer) Getf(key locales.MessageKey, args ...interface{}) string {
	format := l.Get(key)
	return fmt.Sprintf(format, args...)
}
