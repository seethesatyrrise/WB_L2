//
// Паттерн "Цепочка обязанностей" -  это поведенческий паттерн проектирования,
// который позволяет передавать запросы последовательно по цепочке обработчиков.
// Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит
// ли передавать запрос дальше по цепи.
//
// Паттерн "Цепочка обязанностей" используют, когда:
// - Имеется более одного объекта, который может обработать определенный запрос
// - Надо передать запрос на выполнение одному из нескольких объектов, точно не определяя, какому именно объекту
// - Набор объектов задается динамически
//
// Преимущества
// - Уменьшает зависимость между клиентом и обработчиками.
// - Реализует принцип единственной обязанности.
// - Реализует принцип открытости/закрытости.
// Недостатки
// - Запрос может остаться никем не обработанным.
//
// Одним из примеров паттерна «Цепь ответственности» является банкомат. Пользователь вводит сумму
// для выдачи, и автомат выдает сумму в виде банкнот определенной валюты, например, 50$, 20$, 10$ и т.д.
// Если пользователь вводит сумму, не кратную 10, выдается ошибка. Дальше он высчитывает, сколько
// купюр номиналом в 50$ можно выдать, после этого для остатка рассчитывает количество купюр в 20$ и т.д.
//

package pattern

import "fmt"

// Проходим квест "Сходить в больницу". Сначала проходим регистратуру, далее идем ко врачу,
// после чего фармацевт выдает нам лекарства, за которые мы в итоге должны заплатить кассиру.

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// department - определяет интерфейс для обработки запроса и ссылку на следующий обработчик запроса
type department interface {
	execute(*patient)
	setNext(department)
}

// reception, doctor, medical, cashier - конкретные обработчики, которые реализуют функционал
// для обработки запроса. При невозможности обработки и наличия ссылки на следующий обработчик,
// передают запрос этому обработчику

// --------------------- reception ---------------------

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

// --------------------- doctor ---------------------

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

// --------------------- medical ---------------------

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

// --------------------- cashier ---------------------

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

func (c *cashier) setNext(next department) {
	c.next = next
}

func DemonstrateChainOfResp() {
	cashier := &cashier{}
	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)
	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)
	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)
	patient := &patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
