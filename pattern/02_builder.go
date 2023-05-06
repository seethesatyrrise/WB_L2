//
// Паттерн "Строитель" — это паттерн проектирования, который позволяет поэтапно создавать сложные
// объекты с помощью четко определенной последовательности действий. Строительство контролируется
// объектом-распорядителем (Director), которому нужно знать только тип создаваемого объекта.
//
// Паттерн "Строитель" используют, когда:
// - Процесс создания нового объекта не должен зависеть от того, из каких частей
// этот объект состоит и как эти части связаны между собой
// - Необходимо обеспечить получение различных вариаций объекта в процессе его создания
//
// Преимущества
// - Позволяет создавать продукты пошагово.
// - Позволяет использовать один и тот же код для создания различных продуктов.
// - Изолирует сложный код сборки продукта от его основной бизнес-логики.
// Недостатки
// - Усложняет код программы из-за введения дополнительных классов.
// - Клиент будет привязан к конкретным классам строителей, так как в интерфейсе
// директора может не быть метода получения результата.
//

package pattern

import "fmt"

// Person - представляет объект, который должен быть создан.
type Person struct {
	name, address, pin             string
	workAddress, company, position string
	salary                         int
}

// --------------------- PersonBuilder ---------------------

type PersonBuilder struct {
	person *Person
}

// NewPersonBuilder - конструктор PersonBuilder
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
}

// Lives возвращает *PersonAddressBuilder, связанный с *PersonBuilder
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

// Works возвращает *PersonJobBuilder, связанный с *PersonBuilder
func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

// Build builds a person from PersonBuilder
func (b *PersonBuilder) Build() *Person {
	return b.person
}

// --------------------- PersonAddressBuilder ---------------------

// PersonAddressBuilder - потомок PersonBuilder для добавления информации об адресе
type PersonAddressBuilder struct {
	PersonBuilder
}

// At - добавляет адрес к Person
func (a *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	a.person.address = address
	return a
}

// WithPostalCode - добавляет почтовый индекс к Person
func (a *PersonAddressBuilder) WithPostalCode(pin string) *PersonAddressBuilder {
	a.person.pin = pin
	return a
}

// --------------------- PersonJobBuilder ---------------------

// PersonJobBuilder - потомок PersonBuilder для добавления информации о работе
type PersonJobBuilder struct {
	PersonBuilder
}

// As - добавляет название рабочей позиции к Person
func (j *PersonJobBuilder) As(position string) *PersonJobBuilder {
	j.person.position = position
	return j
}

// For - добавляет название компании к Person
func (j *PersonJobBuilder) For(company string) *PersonJobBuilder {
	j.person.company = company
	return j
}

// In - добавляет адрес компании к Person
func (j *PersonJobBuilder) In(companyAddress string) *PersonJobBuilder {
	j.person.workAddress = companyAddress
	return j
}

// WithSalary - добавляет зарплату к Person
func (j *PersonJobBuilder) WithSalary(salary int) {
	j.person.salary = salary
}

func DemonstrateBuilder() {
	pb := NewPersonBuilder()
	pb.Lives().
		At("Bangalore").
		WithPostalCode("560102").
		Works().
		As("Software Engineer").
		For("IBM").
		In("Bangalore").
		WithSalary(150000)

	person := pb.Build()

	fmt.Println(person)
}
