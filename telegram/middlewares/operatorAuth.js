const authCheck = () => (ctx, next) => {
    require('../grpc/client').ReadOperator({
        "username": ctx.from.username
    }, (error, data) => {
        if (!error) {
            if (ctx.session === undefined) ctx.session = {};
            ctx.session.operator = data.operator;

            if (ctx.session.operator.telegramId == 0) {
                require('../grpc/client').WriteOperatorTelegramID({
                    "username": ctx.from.username,
                    "telegramId": ctx.from.id.toString(),
                }, () => {})
            }

            return next();
        }
    });
};

module.exports = authCheck;