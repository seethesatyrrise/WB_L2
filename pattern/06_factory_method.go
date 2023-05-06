//
// Паттерн "Фабричный метод" (Factory Method) - это паттерн, который определяет интерфейс
// для создания объектов некоторого класса, но непосредственное решение о том, объект какого
// класса создавать происходит в подклассах. То есть паттерн предполагает, что базовый класс
// делегирует создание объектов классам-наследникам.
//
// Паттерн "Фабричный метод" используют, когда:
// - Заранее неизвестно, объекты каких типов необходимо создавать
// - Система должна быть независимой от процесса создания новых объектов и расширяемой: в нее можно
// легко вводить новые классы, объекты которых система должна создавать.
// - Создание новых объектов необходимо делегировать из базового класса классам наследникам
//
// Преимущества:
// - Вы избегаете тесной связи между классом создателя и конкретными классами продуктов.
// - Принцип единственной ответственности. Вы можете переместить код создания продукта в
// одно место в программе, что упростит поддержку кода.
// - Принцип открытости/закрытости. Вы можете вводить в программу новые типы продуктов,
// не нарушая существующий клиентский код.
// Недостатки:
// - Код может стать более сложным, поскольку вам нужно ввести много новых подклассов для
// реализации шаблона. В идеале вы вводите этот паттерн в существующую иерархию классов-создателей.
//

package pattern

import "fmt"

// Есть два типа котов - Хейзел и другие коты. У них есть характеристика "Уровень милоты".
// Фабрика котов создает эти два типа.

// ICat - определяет интерфейс класса, объекты которого надо создавать.
type iCat interface {
	setName(name string)
	setCuteness(power int)
	getName() string
	getCuteness() int
}

type cat struct {
	name     string
	cuteness int
}

func (c *cat) setName(name string) {
	c.name = name
}

func (c *cat) getName() string {
	return c.name
}

func (c *cat) setCuteness(cuteness int) {
	c.cuteness = cuteness
}

func (c *cat) getCuteness() int {
	return c.cuteness
}

// hazel, anyOtherCat - конкретные классы, представляющие реализацию класса cat.

type hazel struct {
	cat
}

func newHazel() iCat {
	return &hazel{
		cat: cat{
			name:     "Hazel",
			cuteness: 100,
		},
	}
}

type anyOtherCat struct {
	cat
}

func newAnyOtherCat() iCat {
	return &anyOtherCat{
		cat: cat{
			name:     "Any other cat",
			cuteness: 75,
		},
	}
}

// getCat - фабрика котов :3

func getCat(catType string) (iCat, error) {
	if catType == "Hazel" {
		return newHazel(), nil
	}
	if catType == "any other cat" {
		return newAnyOtherCat(), nil
	}
	return nil, fmt.Errorf("Wrong cat type passed")
}

func DemonstrateFactoryMethod() {
	hazel, _ := getCat("Hazel")
	anyOtherCat, _ := getCat("any other cat")
	printDetails(hazel)
	printDetails(anyOtherCat)
}

func printDetails(g iCat) {
	fmt.Println("Cat:", g.getName())
	fmt.Println("Cuteness:", g.getCuteness())
}
