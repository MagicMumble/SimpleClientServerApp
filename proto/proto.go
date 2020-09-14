package proto

import "encoding/json"

// Request -- запрос клиента к серверу.
type Request struct {
	// Поле Command может принимать три значения:
	// * "quit" - прощание с сервером (после этого сервер рвёт соединение);
	// * "set new array" - передача нового массива на сервер;
	// * "sort" - отсортировать массив
	Command string `json:"command"`

	// Если Command == "set new array", в поле Data должен лежать массив
	// В противном случае, поле Data пустое.
	Data *json.RawMessage `json:"data"`
}

// Response -- ответ сервера клиенту.
type Response struct {
	// Поле Status может принимать три значения:
	// * "ok" - успешное выполнение команды "quit" или "sort";
	// * "failed" - в процессе выполнения команды произошла ошибка;
	// * "result" - массив отсортирован.
	Status string `json:"status"`

	// Если Status == "failed", то в поле Data находится сообщение об ошибке.
	// Если Status == "result", в поле Data должна лежать дробь
	// в виде структуры Fraction.
	// В противном случае, поле Data пустое.
	Data *json.RawMessage `json:"data"`
}

