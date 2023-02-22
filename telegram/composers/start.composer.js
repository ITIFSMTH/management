const { Composer } = require('telegraf')
const keyboards = require('../helpers/keyboards')

const composer = new Composer()

composer.start((ctx) => {
    ctx.reply(
        ctx.i18n.t("greeting", ctx.session.operator), {
            "reply_markup": keyboards.getKeyboard(ctx.session.operator).reply_markup,
            "parse_mode": "HTML",
        }
    )
})

module.exports = composer