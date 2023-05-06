//
// Паттерн "Стратегия" – это поведенческий паттерн проектирования, который определяет семейство
// схожих алгоритмов и помещает каждый из них в собственный класс, после чего алгоритмы можно
// взаимозаменять прямо во время исполнения программы.
//
// Паттерн "Стратегия" используют, когда:
// - Есть несколько родственных классов, которые отличаются поведением. Можно задать один
// основной класс, а разные варианты поведения вынести в отдельные классы и при необходимости их применять
// - Необходимо обеспечить выбор из нескольких вариантов алгоритмов, которые можно легко
// менять в зависимости от условий
// - Необходимо менять поведение объектов на стадии выполнения программы
// - Класс, применяющий определенную функциональность, ничего не должен знать о ее реализации
//
// Преимущества
// - Горячая замена алгоритмов на лету.
// - Изолирует код и данные алгоритмов от остальных классов.
// - Уход от наследования к делегированию.
// - Реализует принцип открытости/закрытости.
// Недостатки
// - Усложняет программу за счёт дополнительных классов.
// - Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
//
// Реализацию паттерна «Стратегия» отлично видно в приложении типа «навигатор». Пользователь
// выбирает начальную и конечную точки пути, а также вариант преодоления пути. То есть один и
// тот же путь может быть пройден пешком, на велосипеде, машине, поезде, самолете или смешанным
// видом транспорта. Выбор способа прохождения пути — это паттерн «Стратегия».
//

package pattern

import "fmt"

// Предположим, мы создаем In-Memory-Cache. Поскольку это кэш в памяти, он имеет ограниченный размер.
// Всякий раз, когда он достигает своего максимального размера, некоторые старые записи из кеша
// необходимо удалить. Это удаление может происходить по нескольким алгоритмам: fifo, lru, lfu.

// evictionAlgo - определяет метод evict(). Это общий интерфейс для всех реализующих его алгоритмов.
type evictionAlgo interface {
	evict(c *cache)
}

// FIFO – First In First Out: Удаляем запись, которая была создана раньше остальных.
type fifo struct{}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}

// LRU – Least Recently Used: Удаляем запись, которую использовали в последний раз.
type lru struct{}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

// LFU – Least Frequently Used: Удаляем запись, которую реже всего использовали.
type lfu struct{}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}

// --------------------- cache ---------------------

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func DemonstrateStrategy() {
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")
	lru := &lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")
	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")
}
