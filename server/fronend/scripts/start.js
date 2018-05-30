const webpack = require('webpack'),
    WebpackDevServer = require('webpack-dev-server'),
    webpackConfig = require('../config/webpack/webpack.config'),
    config = require('../config/config');

new WebpackDevServer(webpack(webpackConfig), {
    publicPath: webpackConfig.output.publicPath,
    hot: false,
    historyApiFallback: true,
    quiet: false,
    noInfo: false,
    stats: {
        assets: true,
        colors: true,
        version: true,
        hash: true,
        timings: true,
        chunks: false,
        chunkModules: false
    }
}).listen(config.get('server:port'), '0.0.0.0', err => {
    if (err) {
        console.error(err);
    }
});
