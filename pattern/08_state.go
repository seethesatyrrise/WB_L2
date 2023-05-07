//
// Паттерн "Состояние" – это поведенческий паттерн проектирования, который позволяет объектам
// менять поведение в зависимости от своего состояния. Извне создаётся впечатление, что изменился класс объекта.
//
// Паттерн "Состояние" используют, когда:
// - Поведение объекта должно зависеть от его состояния и может изменяться динамически во время выполнения
// - В коде методов объекта используются многочисленные условные конструкции, выбор которых зависит
// от текущего состояния объекта
//
// Преимущества
// - Избавляет от множества больших условных операторов машины состояний.
// - Концентрирует в одном месте код, связанный с определённым состоянием.
// - Упрощает код контекста.
// Недостатки
// - Может неоправданно усложнить код, если состояний мало и они редко меняются.
//
// Пример:
// Торговые автоматы могут иметь различные состояния в зависимости от наличия товаров, суммы полученных
// монет, возможности размена денег и т.д. После того как покупатель выбрал и оплатил товар, возможны
// следующие ситуации (состояния):
// - Выдать покупателю товар, выдавать сдачу не требуется.
// - Выдать покупателю товар и сдачу.
// - Покупатель товар не получит из-за отсутствия достаточной суммы денег.
// - Покупатель товар не получит из-за его отсутствия.
//

package pattern

import (
	"fmt"
	"log"
)

// --------------------- vendingMachine ---------------------

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

// --------------------- state ---------------------

// state - интерфейс, который определяет сигнатуры функций, представляющих действие в контексте торгового автомата.
type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

// --------------------- noItemState ---------------------
// состояние, когда товара нет в наличии

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *noItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *noItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

// --------------------- hasItemState ---------------------
// состояние ожидания добавления/запроса товара

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requested\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (i *hasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}

// --------------------- itemRequestedState ---------------------
// состояние, когда товар запрошен (ожидается внесение суммы)

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

// --------------------- hasMoneyState ---------------------
// состояние после внесения необходимой суммы (выдача товара)

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

func DemonstrateState() {
	vendingMachine := newVendingMachine(1, 10)
	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
