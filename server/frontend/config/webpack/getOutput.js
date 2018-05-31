const config = require('../config'),
    path = require('path'),
    constants = require('./constants');

module.exports = function () {
    const NODE_ENV = config.get('NODE_ENV');

    if (NODE_ENV == 'production') {
        return {
            path: path.resolve(constants.ROOT_PATH, 'dist'),
            filename: 'app.min.js',
            publicPath: config.get('config:prod:public_path')
        };
    } else {
        return {
            path: path.resolve(constants.ROOT_PATH, 'dist'),
            filename: 'app.js',
            publicPath: config.get('config:dev:public_path')
        };
    }
};