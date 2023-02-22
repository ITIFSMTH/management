const { Markup } = require("telegraf");

const buttons = {
    "startShift": "🕑 Начать смену",
    "startTimeout": "🕑 Начать перерыв",
    "stopShift": "🏁 Закончить смену",
    "stopTimeout": "🏁 Закончить перерыв",
}

const keyboards = {
    "startShift": Markup.keyboard([[buttons.startShift]]).resize(),
    "onShift": Markup.keyboard([
        [buttons.startTimeout],
        [buttons.stopShift]
    ]).resize(),
    "stopTimeout": Markup.keyboard([[buttons.stopTimeout]]).resize(),
}

module.exports = {
    "buttons": buttons,
    "keyboards": keyboards,
    getKeyboard: (operator) => operator.onShift ? operator.onTimeout ? keyboards.stopTimeout : keyboards.onShift : keyboards.startShift
};