// Imports
const { Composer } = require('telegraf');
const keyboards = require('../helpers/keyboards');
const moment = require('moment');
const client = require('../grpc/client');

// Create new composer
const composer = new Composer();

// On StartShift
composer.hears(keyboards.buttons.startShift, (ctx) => {
    // Start new shift
    client.StartOperatorShift({
        "username": ctx.from.username,
    }, (error, data) => {
        if (error) {
            if (error.message === "6 ALREADY_EXISTS: already_exists") {
                ctx.replyWithHTML(ctx.i18n.t("shiftAlreadyExistToday"))
                return;
            } else if (error.message === "14 UNAVAILABLE: not_today") {
                ctx.replyWithHTML(ctx.i18n.t("shiftNotToday"))
                return
            }
        }
    
        ctx.reply(
            ctx.i18n.t("startShift", {
                "startDate": moment.unix(data.shift.startDate.seconds).format("DD.MM.YYYY HH:mm"),  
            }), {
                "reply_markup": keyboards.keyboards.onShift.reply_markup,
                "parse_mode": "HTML"
            }
        );
    });
});

// On StopShift
composer.hears(keyboards.buttons.stopShift, (ctx) => {
    // Stop shift
    client.StopOperatorShift({
        "username": ctx.from.username,
    }, (error, data) => {
        if (error) {
            return;
        };

        const startDate = moment.unix(data.shift.startDate.seconds);
        const endDate = moment.unix(data.shift.endDate.seconds);

        let timeouts = "⏱ <b>Перерывы:</b>\n";
        let timeoutsTime = 0;
        for (const [timeoutI, timeout] of data.timeouts.entries()) {
            const timeoutStartDate = moment.unix(timeout.startDate.seconds);
            const timeoutEndDate = moment.unix(timeout.endDate.seconds);
            timeoutsTime += timeoutEndDate.diff(timeoutStartDate);
            timeouts += `<b>${timeoutI+1}. С</b> <code>${timeoutStartDate.format("DD.MM.YYYY HH:mm")}</code> <b>до</b> <code>${timeoutEndDate.format("DD.MM.YYYY HH:mm")}</code>\n<b>(</b><code>${moment.duration(timeoutStartDate.diff(timeoutEndDate)).humanize()}</code><b>)</b>\n`;
        }

        ctx.reply(
            ctx.i18n.t("stopShift", {
                "startDate": startDate.format("DD.MM.YYYY HH:mm"),
                "endDate": endDate.format("DD.MM.YYYY HH:mm"),
                "duration": moment.duration(endDate.diff(startDate)).humanize(),
                "workDuration": moment.duration(endDate.diff(startDate) - timeoutsTime).humanize(),
                "timeoutDuration": moment.duration(timeoutsTime).humanize(),
                "delays": data.shift.delays,
                "timeouts": data.timeouts.length > 0 ? timeouts : "",
            }), {
                "reply_markup": keyboards.keyboards.startShift.reply_markup,
                "parse_mode": "HTML"
            }
        );
    });
});

// On StartTimeout
composer.hears(keyboards.buttons.startTimeout, (ctx) => {
    // Start timeout
    client.StartOperatorTimeout({
        "username": ctx.from.username,
    }, (error, data) => {
        if (error) {
            if (error === "10 ABORTED: first_provide_captcha") ctx.reply(ctx.i18n.t("firstProvideCaptcha"), {
                parse_mode: "HTML"
            })
            return;
        };

        ctx.reply(
            ctx.i18n.t("startTimeout", {
                "startDate": moment.unix(data.timeout.startDate.seconds).format("DD.MM.YYYY HH:mm"),
            }), {
                "reply_markup": keyboards.keyboards.stopTimeout.reply_markup,
                "parse_mode": "HTML"
            }
        );
    });
});

// On StopTimeout
composer.hears(keyboards.buttons.stopTimeout, (ctx) => {
    // Stop timeout
    client.StopOperatorTimeout({
        "username": ctx.from.username,
    }, (error, data) => {
        if (error) {
            return;
        };

        const startDate = moment.unix(data.timeout.startDate.seconds);
        const endDate = moment.unix(data.timeout.endDate.seconds);

        ctx.reply(
            ctx.i18n.t("stopTimeout", {
                "startDate": startDate.format("DD.MM.YYYY HH:mm"),
                "endDate": endDate.format("DD.MM.YYYY HH:mm"),
                "duration": moment.duration(endDate.diff(startDate)).humanize()
            }), {
                "reply_markup": keyboards.keyboards.onShift.reply_markup,
                "parse_mode": "HTML"
            }
        );
    });
});

// On captcha answer
composer.hears(/.{4}/gm, (ctx) => {
    if (!ctx.session.operator.onShift) return

    client.WriteOperatorCaptcha({
        "username": ctx.from.username,
        "captcha": ctx.message.text,
    }, (error, data) => {
        if (error && error.message == "5 NOT_FOUND: no_shift") return

        return ctx.reply(ctx.i18n.t(error ? 'badAnswer' : 'goodAnswer'), {
            parse_mode: 'HTML',
        })
    })
})

// Export composer
module.exports = composer;
