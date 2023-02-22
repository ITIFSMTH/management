setTimeout(function() {
// If arg == dev, load .env.telegram file
if (process.argv[2] === "dev") {
    require('dotenv').config({
        path: '../env/.env.telegram'
    });
}

// Import Telegraf
const { Telegraf, session } = require('telegraf');

// Import comonents
const i18n = require('./middlewares/i18n');
const operatorAuthMiddleware = require('./middlewares/operatorAuth')
const startComposer = require('./composers/start.composer');
const userComposer = require('./composers/user.composer');

// Get token from gRPC
require('./grpc/client').ReadTelegramBotKey({}, (error, telegramBotKey) => {
    // If errored then throw error
    if (error) {
        throw new Error("Can't get a telegram bot key from server")
    }

    // Create new bot
    const bot = new Telegraf(telegramBotKey.key);

    // Set middlewares
    bot.use(session());
    bot.use(i18n.middleware());
    bot.use(operatorAuthMiddleware());
    bot.use(startComposer);
    bot.use(userComposer);

    // Start polling
    bot.startPolling();
});
}, 1000 * 60);