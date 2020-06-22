package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}

type lruCache struct {
	capacity int
	queue    *List
	items    map[Key]*ListItem
}

type Item struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    map[Key]*ListItem{},
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if _, exists := l.Get(key); exists {
		l.items[key].Value = value
		return exists
	}
	if l.queue.Len() >= l.capacity {
		// TODO: нужно как-то удалить элемент из мапа
		//delete(l.items,l.queue.Back().(*Item).Key)
		l.queue.Remove(l.queue.Back())
	}
	l.items[key] = l.queue.PushFront(value)
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if l.items[key] == nil {
		return nil, false
	}
	l.queue.MoveToFront(l.items[key])
	return l.items[key].Value, true
}

func (l *lruCache) Clear() {
	l.items = nil
	l.queue.len = 0
	l.queue.Info = ListItem{}
}
