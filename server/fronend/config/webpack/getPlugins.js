const webpack = require('webpack'),
    HtmlWebpackPlugin = require('html-webpack-plugin'),
    MiniCssExtractPlugin = require('mini-css-extract-plugin'),
    config = require('../config'),
    path = require('path'),
    constants = require('./constants');

module.exports = function () {
    const NODE_ENV = config.get('NODE_ENV');
    let plugins = [];
    plugins.push(
        new HtmlWebpackPlugin({
            template: path.resolve(constants.SRC_PATH, 'index.tpl.html'),
            inject: 'body',
            filename: 'index.html'
        })
    );
    if (NODE_ENV === 'production') {
        plugins.push(new MiniCssExtractPlugin({
            filename: 'app.css'
        }));
        plugins.push(new webpack.DefinePlugin({
            'process.env.NODE_ENV': JSON.stringify(NODE_ENV)
        }));
    }
    if (NODE_ENV !== 'production') {
        plugins.push(new webpack.HotModuleReplacementPlugin());
    }
    return plugins;
};
