package main

func createWelcomeKeyboard() ReplyKeyboard {
	b1 := KeyboardButton{"привет"}
	b2 := KeyboardButton{"команды"}
	b3 := KeyboardButton{"нуклеотиды"}
	b4 := KeyboardButton{"/start"}

	l1 := []KeyboardButton{b1, b2}
	l2 := []KeyboardButton{b3, b4}
	layers := [][]KeyboardButton{l1, l2}

	kb := ReplyKeyboard{layers, true, false}
	return kb
}

func createFirstLayerKeyboard() InlineKeyboard {
	b1 := InlineKeyboardButton{
		Text:         "Аденин",
		CallbackData: "layer1button1",
	}
	b2 := InlineKeyboardButton{
		Text:         "Гуанин",
		CallbackData: "layer1button2",
	}
	b3 := InlineKeyboardButton{
		Text:         "Тимин",
		CallbackData: "layer1button3",
	}
	b4 := InlineKeyboardButton{
		Text:         "Цитазин",
		CallbackData: "layer1button4",
	}
	l1 := []InlineKeyboardButton{b1, b2, b3, b4}

	layers := [][]InlineKeyboardButton{l1}

	kb := InlineKeyboard{layers}
	return kb
}

func createSecondLayerKeyboard1() InlineKeyboard {
	b1 := InlineKeyboardButton{
		Text:         "Назад",
		CallbackData: "layer1",
	}
	b2 := InlineKeyboardButton{
		Text:         "Подробнее",
		CallbackData: "layer2adenin",
	}
	l1 := []InlineKeyboardButton{b1, b2}

	layers := [][]InlineKeyboardButton{l1}

	kb := InlineKeyboard{layers}
	return kb
}

func createSecondLayerKeyboard2() InlineKeyboard {
	b1 := InlineKeyboardButton{
		Text:         "Назад",
		CallbackData: "layer1",
	}
	b2 := InlineKeyboardButton{
		Text:         "Подробнее",
		CallbackData: "layer2guanin",
	}
	l1 := []InlineKeyboardButton{b1, b2}

	layers := [][]InlineKeyboardButton{l1}

	kb := InlineKeyboard{layers}
	return kb
}

func createSecondLayerKeyboard3() InlineKeyboard {
	b1 := InlineKeyboardButton{
		Text:         "Назад",
		CallbackData: "layer1",
	}
	b2 := InlineKeyboardButton{
		Text:         "Подробнее",
		CallbackData: "layer2timin",
	}
	l1 := []InlineKeyboardButton{b1, b2}

	layers := [][]InlineKeyboardButton{l1}

	kb := InlineKeyboard{layers}
	return kb
}

func createSecondLayerKeyboard4() InlineKeyboard {
	b1 := InlineKeyboardButton{
		Text:         "Назад",
		CallbackData: "layer1",
	}
	b2 := InlineKeyboardButton{
		Text:         "Подробнее",
		CallbackData: "layer2citozin",
	}
	l1 := []InlineKeyboardButton{b1, b2}

	layers := [][]InlineKeyboardButton{l1}

	kb := InlineKeyboard{layers}
	return kb
}
