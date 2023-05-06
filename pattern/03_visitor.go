//
// Паттерн Посетитель позволяет определить операцию для объектов
// других классов без изменения этих классов.
// При использовании паттерна Посетитель определяются две иерархии классов:
// одна для элементов, для которых надо определить новую операцию, и вторая
// иерархия для посетителей, описывающих данную операцию.
//
// Паттер "Посетитель" используют, когда:
// - Имеется много объектов разнородных классов с разными интерфейсами,
// и требуется выполнить ряд операций над каждым из этих объектов
// - Классам необходимо добавить одинаковый набор операций без изменения этих классов
// - Часто добавляются новые операции к классам, при этом общая структура
// классов стабильна и практически не изменяется
//
// Аналогия из жизни:
// Представьте начинающего страхового агента, жаждущего получить новых клиентов.
// Он беспорядочно посещает все дома в округе, предлагая свои услуги. Но для каждого
// из посещаемых типов домов у него имеется особое предложение:
// - Придя в дом к обычной семье, он предлагает оформить медицинскую страховку.
// - Придя в банк, он предлагает страховку от грабежа.
// - Придя на фабрику, он предлагает страховку предприятия от пожара и наводнения.
//
// Преимущества
// - Упрощает добавление операций, работающих со сложными структурами объектов.
// - Объединяет родственные операции в одном классе.
// - Посетитель может накапливать состояние при обходе структуры элементов.
// Недостатки
// - Паттерн не оправдан, если иерархия элементов часто меняется.
// - Может привести к нарушению инкапсуляции элементов.
//

package pattern

import "fmt"

// Employee - определяет метод Accept(), в котором в качестве параметра принимается объект Visitor
type Employee interface {
	FullName()
	Accept(Visitor)
}

// Developer, Director - конкретные работники, которые реализуют метод Accept()

// --------------------- Developer ---------------------

type Developer struct {
	FirstName string
	LastName  string
	Income    int
}

func (d Developer) FullName() {
	fmt.Println("Developer ", d.FirstName, " ", d.LastName)
}

func (d Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}

// --------------------- Director ---------------------

type Director struct {
	FirstName     string
	LastName      string
	Income        int
	BlockOfShares int
}

func (d Director) FullName() {
	fmt.Println("Director ", d.FirstName, " ", d.LastName)
}

func (d Director) Accept(v Visitor) {
	v.VisitDirector(d)
}

// Visitor - интерфейс посетителя, который определяет метод Visit() для объектов Developer, Director
type Visitor interface {
	VisitDeveloper(d Developer)
	VisitDirector(d Director)
}

// CalculIncome - конкретный класс посетителей, реализует интерфейс, определенный в Visitor
type CalculIncome struct {
	bonusRate int
}

func (c CalculIncome) VisitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}

func (c CalculIncome) VisitDirector(d Director) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100 + d.BlockOfShares)
}

func DemonstrateVisitor() {
	var backend Employee = Developer{"Iana", "Lapina", 1000}
	var boss Employee = Director{"Hazel", "Lapina", 2000, 500}

	backend.FullName()
	backend.Accept(CalculIncome{20})

	boss.FullName()
	boss.Accept(CalculIncome{10})
}
