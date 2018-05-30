const MiniCssExtractPlugin = require('mini-css-extract-plugin'),
    constants = require('./constants'),
    config = require('../config'),
    path = require('path'),
    NODE_ENV = config.get('NODE_ENV');

module.exports = function () {
    return {
        rules: [
            {
                test: /\.(js|jsx)$/,
                use: [
                    'babel-loader'
                ],
                exclude: /node_modules/,
            },
            {
                test: /\.scss$/,
                use: NODE_ENV === 'production' ? [
                    MiniCssExtractPlugin.loader,
                    {
                        loader: 'css-loader',
                        options: {
                            modules: true,
                            camelCase: true,
                            localIdentName: '[name]__[local]__[hash:base64:5]',
                            minimize: true
                        }
                    },
                    {
                        loader: "sass-loader",
                        options: {
                            includePaths: [
                                path.resolve(constants.SRC_PATH, 'styles')
                            ]
                        }
                    }
                ] : [
                    'style-loader',
                    {
                        loader: 'css-loader',
                        options: {
                            modules: true,
                            camelCase: true,
                            localIdentName: '[name]__[local]__[hash:base64:5]'
                        }
                    },
                    {
                        loader: "sass-loader",
                        options: {
                            includePaths: [
                                path.resolve(constants.SRC_PATH, 'styles')
                            ]
                        }
                    }
                ]
            },
            {
                test: /\.less/,
                use: NODE_ENV === 'production' ? [
                    MiniCssExtractPlugin.loader,
                    'css-loader',
                    {
                        loader: 'less-loader',
                        options: {
                            javascriptEnabled: true
                        }
                    }
                ] : [
                    'style-loader',
                    'css-loader',
                    {
                        loader: 'less-loader',
                        options: {
                            javascriptEnabled: true
                        }
                    }
                ]
            },
            {
                test: /\.css/,
                use: NODE_ENV === 'production' ? [
                    MiniCssExtractPlugin.loader,
                    'css-loader'
                ] : [
                    'style-loader',
                    'css-loader'
                ]
            },
            {
                test: /\.((woff(2)?)|ttf|eot|otf)(\?[a-z0-9#=&.]+)?$/,
                use: [
                    {
                        loader: 'file-loader',
                        options: {
                            name: 'fonts/[name].[ext]'
                        }
                    }
                ]
            },
            {
                test: /\.hbs$/,
                use: [
                    {
                        loader: 'handlebars-loader'
                    }
                ]
            },
            {
                test: /\.(png|jpg|gif|svg)$/,
                use: [
                    {
                        loader: 'file-loader',
                        options: {
                            name: 'resources/[name].[ext]'
                        }
                    }
                ]
            }
        ]
    };
};
