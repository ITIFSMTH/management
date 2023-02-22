package cron

import (
	"fmt"
	"management-backend/db"
	"management-backend/models"
	"sort"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nleeper/goment"
)

var StartShiftKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🕑 Начать смену"),
	),
)

func CloseShifts() {
	// Get DB
	db := db.GetDB()

	// Get bot token
	var token string
	db.Model(&models.Setting{}).Select("telegram_bot_key").Scan(&token)

	// Get Bot
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return
	}

	// Get all unclosed timeouts
	var timeouts []models.Timeout
	db.Model(&timeouts).Where("end_date = ?", time.Time{}).Find(&timeouts)
	// Close all timeouts
	db.Model(&timeouts).Update("end_date", time.Now())

	// Get all unclosed shifts
	var shifts []models.Shift
	db.Model(&shifts).Where("end_date = ?", time.Time{}).Preload("Timeouts").Find(&shifts)

	if len(shifts) == 0 {
		return
	}

	// Close all shifts
	db.Model(&shifts).Update("end_date", time.Now())

	// Get operators ids
	var operatorsIds []uint
	for _, shift := range shifts {
		operatorsIds = append(operatorsIds, shift.OperatorID)
	}

	// Get this operators
	var operators []models.Operator
	db.Find(&operators, operatorsIds)

	sort.Slice(operators, func(i, j int) bool { return operators[i].ID < operators[j].ID })
	sort.Slice(shifts, func(i, j int) bool { return shifts[i].OperatorID < shifts[j].OperatorID })

	// Notify every operator that his shift was closed
	for operatorI, operator := range operators {
		message := tgbotapi.NewMessage(
			operator.TelegramID,
			getMessageFromShifts(&shifts[operatorI]),
		)
		message.ReplyMarkup = StartShiftKeyboard
		message.ParseMode = tgbotapi.ModeHTML

		bot.Send(message)
	}
}

func getMessageFromShifts(shift *models.Shift) string {
	message := "🥱 <b>Уже поздно, мы закроем Вашу смену...</b>\n\n<b>Начало: </b><code>%s</code>\n<b>Конец</b>: <code>%s</code>\n\n⏳ <b>Продолжительность</b>: <code>%s</code>\n🎧 <b>Рабочее время</b>: <code>%s</code>\n🛏 <b>Время перерывов</b>: <code>%s</code>\n🏃 <b>Число задержек</b>: <code>%d</code>\n\n%s"

	startDate, _ := goment.New(shift.StartDate)
	endDate, _ := goment.New(shift.EndDate)

	timeouts := "⏱ <b>Перерывы:</b>\n"
	timeoutsTime := 0
	for timeoutI, timeout := range shift.Timeouts {
		timeoutStartDate, _ := goment.New(timeout.StartDate)
		timeoutEndDate, _ := goment.New(timeout.EndDate)
		timeoutsTime += timeoutEndDate.Diff(timeoutStartDate)
		timeouts += fmt.Sprintf("<b>%d. С</b> <code>%s</code> <b>до</b> <code>%s</code>\n<b>(</b><code>%s</code><b>)</b>\n",
			timeoutI+1,
			timeoutStartDate.Format("DD.MM.YYYY HH:mm"),
			timeoutEndDate.Format("DD.MM.YYYY HH:mm"),
			humanizeDuration(timeoutEndDate.Diff(timeoutStartDate)),
		)
	}

	timeoutsForSend := ""
	if len(shift.Timeouts) > 0 {
		timeoutsForSend = timeouts
	}

	return fmt.Sprintf(
		message,
		startDate.Format("DD.MM.YYYY HH:mm"),
		endDate.Format("DD.MM.YYYY HH:mm"),
		humanizeDuration(endDate.Diff(startDate)),
		humanizeDuration(endDate.Diff(startDate)-timeoutsTime),
		humanizeDuration(timeoutsTime),
		0,
		timeoutsForSend,
	)
}

func humanizeDuration(seconds int) string {
	hours := seconds / 60 / 60
	minutes := seconds / 60 % 60
	secondsRem := seconds % 60

	if hours > 0 {
		return fmt.Sprintf("%dч. %dм.", hours, minutes)
	}
	return fmt.Sprintf("%dм. %dc.", minutes, secondsRem)
}
