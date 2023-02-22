package cron

import (
	"bytes"
	"management-backend/db"
	"management-backend/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/steambap/captcha"
	"gorm.io/gorm"
)

var StopTimeoutKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🏁 Закончить перерыв"),
	),
)

func CheckShift() {
	// If hour is 23 return because all shifts closed
	if time.Now().Hour() == 23 {
		return
	}

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

	// Get all opened shifts
	var shifts []models.Shift
	db.Model(&shifts).Where("end_date = ?", time.Time{}).Find(&shifts)

	// For each shift
	for _, shift := range shifts {
		isOnTimeout := len(shift.Timeouts) > 0 && shift.Timeouts[len(shift.Timeouts)-1].EndDate.IsZero()
		isLastNotifyZero := shift.LastNotify.IsZero()

		// If shift on timeout then skip
		// If shift lastNotify  isn't zero then skip
		// Check is it this time for notifying
		if !isOnTimeout && isLastNotifyZero && shift.NextNotify.Format("15:04") == time.Now().Format("15:04") {
			// Get operator telegram id
			var telegramId int64
			db.Model(&models.Operator{}).Where(shift.OperatorID).Select("telegram_id").Scan(&telegramId)

			// Send operator captcha
			message := tgbotapi.NewPhoto(telegramId, tgbotapi.FileBytes{
				Name:  "captcha",
				Bytes: updateCaptcha(&shift),
			})

			message.Caption = "⏰ <b>Пожалуйста, введите капчу в течении 15 минут</b>"
			if shift.Delays > 0 {
				message.Caption = "⏰ <b>Пожалуйста, введите капчу в течении 15 минут</b>"
			}
			message.ParseMode = tgbotapi.ModeHTML
			bot.Send(message)

			continue
		}

		// Check is was a notify
		if !isLastNotifyZero {
			if shift.Delays == 0 && shift.LastNotify.Add(15*time.Minute).Format("15:04") == time.Now().Format("15:04") {
				// Add delay
				db.Model(&shift).Update("delays", gorm.Expr("delays + 1"))

				// Get operator telegram id
				var telegramId int64
				db.Model(&models.Operator{}).Where(shift.OperatorID).Select("telegram_id").Scan(&telegramId)

				// Send message
				message := tgbotapi.NewMessage(telegramId, "⏰ <b>Вы не ответили в течении 15 минут!</b>\n\n<b>Вам присвоена <code>1</code> задержка</b>")
				message.ParseMode = tgbotapi.ModeHTML
				bot.Send(message)

				// Send operator captcha
				captchaMessage := tgbotapi.NewPhoto(telegramId, tgbotapi.FileBytes{
					Name:  "captcha",
					Bytes: updateCaptcha(&shift),
				})

				captchaMessage.Caption = "⏰ <b>Пожалуйста, введите капчу в течении 10 минут</b>"
				captchaMessage.ParseMode = tgbotapi.ModeHTML
				bot.Send(captchaMessage)
			} else if shift.Delays > 0 && shift.LastNotify.Add(10*time.Minute).Format("15:04") == time.Now().Format("15:04") {
				// Add delay
				db.Model(&shift).Update("delays", gorm.Expr("delays + 1"))

				// Get operator telegram id
				var telegramId int64
				db.Model(&models.Operator{}).Where(shift.OperatorID).Select("telegram_id").Scan(&telegramId)

				// Create message
				message := tgbotapi.NewMessage(telegramId, "<b>Вы не ответили на капчу, вам присвоена еще одна задежка!</b>")
				message.ParseMode = tgbotapi.ModeHTML

				// Create timeout
				if !(len(shift.Timeouts) > 0 && shift.Timeouts[len(shift.Timeouts)-1].EndDate.IsZero()) {
					// Create timeout
					timeout := models.Timeout{
						StartDate: time.Now(),
						EndDate:   time.Time{},
					}

					if err := db.Model(&shift).Association("Timeouts").Append(&timeout); err != nil {
						continue
					}

					// Set null last notify and captcha-answer
					db.Model(&shift).Select("LastNotify", "NextNotify", "CaptchaAnswer").Updates(&models.Shift{
						LastNotify:    time.Time{},
						NextNotify:    time.Time{},
						CaptchaAnswer: "",
					})

					message.Text = "<b>Вы не ответили на капчу, вам присвоена еще одна задежка и присвоен режим перерыва!</b>"
					message.ReplyMarkup = StopTimeoutKeyboard
				}

				bot.Send(message)
			}
		}
	}
}

func updateCaptcha(shift *models.Shift) []byte {
	// Get DB
	db := db.GetDB()

	// Generate a captcha
	c, _ := captcha.New(800, 400, func(options *captcha.Options) {
		options.Noise = 0.0
	})

	captchaBuff := new(bytes.Buffer)
	c.WriteImage(captchaBuff)

	// Set captcha answer to shift
	db.Model(&shift).Select("CaptchaAnswer", "LastNotify", "NextNotify").Updates(&models.Shift{
		CaptchaAnswer: c.Text,
		LastNotify:    time.Now(),
		NextNotify:    time.Time{},
	})

	return captchaBuff.Bytes()
}
