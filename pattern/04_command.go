//
// Паттерн "Команда" позволяет инкапсулировать запрос на выполнение определенного
// действия в виде отдельного объекта. Этот объект запроса на действие и называется командой.
// При этом объекты, инициирующие запросы на выполнение действия, отделяются от объектов, которые
// выполняют это действие. Команды могут использовать параметры, которые передают ассоциированную
// с командой информацию. Кроме того, команды могут ставиться в очередь и также могут быть отменены.
//
// Паттерн "Команда" используют, когда:
// - надо передавать в качестве параметров определенные действия, вызываемые в ответ на другие
// действия. То есть когда необходимы функции обратного действия в ответ на определенные действия.
// - необходимо обеспечить выполнение очереди запросов, а также их возможную отмену.
// - надо поддерживать логгирование изменений в результате запросов. Использование логов может
// помочь восстановить состояние системы - для этого необходимо будет использовать последовательность
// запротоколированных команд.
//
// Преимущества
// - Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
// - Позволяет реализовать простую отмену и повтор операций.
// - Позволяет реализовать отложенный запуск операций.
// - Позволяет собирать сложные команды из простых.
// - Реализует принцип открытости/закрытости.
// Недостатки
// - Усложняет код программы из-за введения множества дополнительных классов.
//
// Аналогия из жизни:
// Вы заходите в ресторан и садитесь у окна. К вам подходит вежливый официант и принимает заказ,
// записывая все пожелания в блокнот. Откланявшись, он уходит на кухню, где вырывает лист из блокнота
// и клеит на стену. Далее лист оказывается в руках повара, который читает содержание заказа и
// готовит заказанные блюда.
// В этом примере вы являетесь отправителем, официант с блокнотом – командой, а повар – получателем.
// Как и в паттерне, вы не соприкасаетесь напрямую с поваром. Вместо этого вы отправляете заказ с
// официантом, который самостоятельно «настраивает» повара на работу. С другой стороны, повар не
// знает, кто конкретно послал ему заказ. Но это ему безразлично, так как вся необходимая информация
// есть в листе заказа.
//

package pattern

import (
	"WB-L2/pattern/utils"
	"fmt"
)

// Создадим пульт к телевизору, имеющий две кнопки: включить и повысить уровень громкости.
// Также предоставим возможность отменять выполненные действия.

// ICommand - интерфейс, представляющий команду. Обычно определяет метод Execute() для
// выполнения действия, а также нередко включает метод Undo(), реализация
// которого должна заключаться в отмене действия команды
type ICommand interface {
	Execute()
	Undo()
}

type NoCommand struct{}

func NewNoCommand() *NoCommand {
	return &NoCommand{}
}

func (com *NoCommand) Execute() {}
func (com *NoCommand) Undo()    {}

// Получатели команды. Определяют действия, которые должны выполняться в результате запроса.

// TV - команды: включить, выключить
type TV struct{}

func NewTV() *TV {
	return &TV{}
}

func (tv *TV) On() {
	fmt.Println("Телевизор включен")
}

func (tv *TV) Off() {
	fmt.Println("Телевизор выключен")
}

// Volume - команды: повысить, понизить
type Volume struct {
	off   int
	high  int
	level int
}

func NewVolume() *Volume {
	return &Volume{off: 0, high: 20}
}

func (v *Volume) RaiseLevel() {
	if v.level < v.high {
		v.level++
	}
	fmt.Println("Уровень звука:", v.level)
}

func (v *Volume) DropLevel() {
	if v.level > v.off {
		v.level--
	}
	fmt.Println("Уровень звука:", v.level)
}

// Конкретные реализации команд, реализуют методы Execute(), Undo(),
// в которых вызывается определенный метод, определенный в соответствующем классе

// TVOnCommand - использует методы, определенные в TV
type TVOnCommand struct {
	tv *TV
}

func NewTVOnCommand(tvSet *TV) *TVOnCommand {
	return &TVOnCommand{tv: tvSet}
}

func (com *TVOnCommand) Execute() {
	com.tv.On()
}

func (com *TVOnCommand) Undo() {
	com.tv.Off()
}

// VolumeCommand - использует методы, определенные в Volume
type VolumeCommand struct {
	volume *Volume
}

func NewVolumeCommand(v *Volume) *VolumeCommand {
	return &VolumeCommand{volume: v}
}

func (com *VolumeCommand) Execute() {
	com.volume.RaiseLevel()
}

func (com *VolumeCommand) Undo() {
	com.volume.DropLevel()
}

// MultiPult - инициатор команд - вызывает команду для выполнения определенного запроса
type MultiPult struct {
	buttons         []ICommand
	commandsHistory *utils.Stack
}

func NewMultiPult() *MultiPult {
	buttons := make([]ICommand, 2)
	for i := 0; i < len(buttons); i++ {
		buttons[i] = NewNoCommand()
	}
	commandsHistory := utils.NewStack()
	return &MultiPult{
		buttons:         buttons,
		commandsHistory: commandsHistory,
	}
}

func (mp *MultiPult) SetCommand(number int, com ICommand) {
	mp.buttons[number] = com
}

func (mp *MultiPult) PressButton(number int) {
	mp.buttons[number].Execute()
	mp.commandsHistory.Push(number)
}

func (mp *MultiPult) PressUndoButton() {
	if mp.commandsHistory.Count() > 0 {
		number := mp.commandsHistory.Pop().(int)
		mp.buttons[number].Undo()
	}
}

func DemonstrateCommand() {
	tv := NewTV()
	volume := NewVolume()
	mPult := NewMultiPult()
	mPult.SetCommand(0, NewTVOnCommand(tv))
	mPult.SetCommand(1, NewVolumeCommand(volume))
	// on
	mPult.PressButton(0)
	// volume up
	mPult.PressButton(1)
	mPult.PressButton(1)
	mPult.PressButton(1)
	// undo
	mPult.PressUndoButton()
	mPult.PressUndoButton()
	mPult.PressUndoButton()
	mPult.PressUndoButton()
}
