package tcptestd

// Tester  - интерфейс для задачи, которая вызывается демоном и возвращает свой результат в виде лога и уровня срочности доставки.
type Tester interface {
	Run() (log []byte, urgency int)
}