const UglifyJsPlugin = require("uglifyjs-webpack-plugin"),
    getPlugins = require('./getPlugins'),
    getEntry = require('./getEntry'),
    getOutput = require('./getOutput'),
    getModule = require('./getModule'),
    config = require('../config'),
    constants = require('./constants'),
    path = require('path');

const NODE_ENV = config.get('NODE_ENV');

module.exports = {
    mode: NODE_ENV === 'production' ? 'production' : 'development',
    optimization: {
        minimizer: [
            new UglifyJsPlugin({
                cache: true,
                parallel: true,
                sourceMap: false
            })
        ]
    },
    devtool: NODE_ENV == 'production' ? 'none' : 'cheap-module-eval-source-map',
    entry: getEntry(),
    resolve: {
        extensions: ['.js', '.jsx', '.css', '.scss', '.less'],
        alias: {
            actions: path.resolve(constants.SRC_PATH, 'actions'),
            components: path.resolve(constants.SRC_PATH, 'components'),
            pages: path.resolve(constants.SRC_PATH, 'pages'),
            styles: path.resolve(constants.SRC_PATH, 'styles'),
            utils: path.resolve(constants.SRC_PATH, 'utils'),
            resources: path.resolve(constants.SRC_PATH, 'resources'),
            providers: path.resolve(constants.SRC_PATH, 'providers')
        }
    },
    output: getOutput(),
    plugins: getPlugins(),
    module: getModule()
};
