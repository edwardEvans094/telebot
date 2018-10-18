package main

import (
	"fmt"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var mCommand map[string]string
var mCreate map[string][]string

// Bot object
type Bot struct {
	bot *tb.Bot
	// storage  *QuestionStorage
	deadline int64
}

type Option struct {
	id     string
	name   string
	url    string
	voters []string
}

type Poll struct {
	id               string
	title            string
	options          []Option
	end              int
	admin            string
	creatorId        string
	isMultipleChoice bool
}

func updateCurrentCommand(command string, m *tb.Message) {
	if len(mCommand) == 0 {
		mCommand = map[string]string{}
	}
	mCommand[fmt.Sprintf("%d_%d", m.Chat.ID, m.Sender.ID)] = command
}

func (b Bot) handleCreatePoll(m *tb.Message) {
	if len(mCreate) == 0 {
		mCreate = map[string][]string{}
	}
	pollData := mCreate[fmt.Sprintf("%d_%d", m.Chat.ID, m.Sender.ID)]
	pollData = append(pollData, m.Text)

	mCreate[fmt.Sprintf("%d_%d", m.Chat.ID, m.Sender.ID)] = pollData

	if len(pollData) == 0 {
		b.bot.Reply(m, `đúng rồi đấy, tao đang tạo poll, đưa bố title nào`)
	} else if len(pollData) == 1 {
		b.bot.Reply(m, `DM title dài vler, thế Options đâu ?`)
	} else {
		optionIndex := len(pollData)
		b.bot.Reply(m, fmt.Sprintf("Okey, option %d của mày là gì ?", optionIndex))
	}

	fmt.Println(pollData)
	fmt.Println(mCreate[fmt.Sprintf("%d_%d", m.Chat.ID, m.Sender.ID)])

}

func (b Bot) handleDefault(m *tb.Message) {
	if m.Private() {
		b.bot.Send(m.Chat, `Mày nói clgv, éo hiểu!!`)
	}
}

func (b Bot) handleText(m *tb.Message) {
	switch mCommand[fmt.Sprintf("%d_%d", m.Chat.ID, m.Sender.ID)] {
	case "createPoll":
		b.handleCreatePoll(m)
	default:
		b.handleDefault(m)
	}
}

func (b Bot) handleDone(m *tb.Message) {

	inlineKeys := [][]tb.InlineButton{}

	pollData := mCreate[fmt.Sprintf("%d_%d", m.Chat.ID, m.Sender.ID)]
	fmt.Println(pollData)
	if len(pollData) > 0 {
		for n := 1; n < len(pollData); n++ {
			fmt.Println(pollData[n])
			inlineBtn := tb.InlineButton{
				Unique: fmt.Sprintf("%d", n),
				Text:   pollData[n],
			}
			b.bot.Handle(&inlineBtn, func(c *tb.Callback) {
				fmt.Println("--------------", c)
				// option := 0
				// for i, v := range questionOptions {
				// 	if v == replyBtn.Text {
				// 		option = i
				// 	}
				// }
				// b.handleAnswer(c)
				b.bot.Respond(c, &tb.CallbackResponse{
					Text: fmt.Sprintf("fuck you, mày vừa chọn phương án %d phải không ?", n-1),
				})

			})
			inlineKeysRow := []tb.InlineButton{inlineBtn}
			inlineKeys = append(inlineKeys, inlineKeysRow)
		}

		// for index, value := range textArr {
		// inlineBtn := tb.InlineButton{
		// 	Unique: "",
		// 	Text: value,
		// }
		// I
		// 	inlineKeys[index][0].Text = textArr[index+1]
		// }
	}

	b.bot.Send(m.Sender, "Hello!", &tb.ReplyMarkup{
		// ReplyKeyboard: replyKeys,
		InlineKeyboard: inlineKeys,
	})
}

// func (b Bot) handleAnswer(c *tb.Callback) {
// 	// b.bot.Send(m.Sender, fmt.Sprintf("Mày vừa chọn phương án %d phải không", option))
// 	b.bot.Respond(c, &tb.CallbackResponse{...})
// }

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: "454550958:AAHd8FQnm-x6uHjcIIKBHOfQhmh6TqRtsBY",
		Poller: &tb.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	mybot := Bot{
		bot: b,
		// storage:  storage,
		// deadline: botConfig.Deadline,
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	mybot.bot.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello world")
	})

	mybot.bot.Handle("/createPoll", func(m *tb.Message) {
		updateCurrentCommand("createPoll", m)
		// mybot.handleCreatePoll(m)
		b.Reply(m, "đúng rồi đấy, tao đang tạo poll, đưa bố title nào")
	})

	mybot.bot.Handle("/done", func(m *tb.Message) {
		mybot.handleDone(m)
	})

	mybot.bot.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
		mybot.handleText(m)
		// fmt.Println("in handle text, %d, %d", m.Chat.ID, m.Sender.ID)
	})

	b.Start()
}
