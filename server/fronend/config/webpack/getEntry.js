const config = require('../config'),
    path = require('path'),
    constants = require('./constants');

module.exports = function () {
    const NODE_ENV = config.get('NODE_ENV');

    if (NODE_ENV == 'production') {
        return [
            "babel-polyfill",
            path.resolve(constants.SRC_PATH, 'app.js')
        ];
    } else {
        return [
            "babel-polyfill",
            `webpack-dev-server/client?http://0.0.0.0:${config.get('server:port')}`,
            'webpack/hot/only-dev-server',
            path.resolve(constants.SRC_PATH, 'app.js')
        ];
    }
};
