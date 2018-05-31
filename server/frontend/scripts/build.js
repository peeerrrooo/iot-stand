const webpack = require(`webpack`),
    webpackConfig = require(`../config/webpack/webpack.config`);

webpack(webpackConfig).run((err, stats) => {
    if (err) {
        console.log(err);
    }

    console.log(stats.toString({
        assets: true,
        colors: true,
        version: true,
        hash: true,
        timings: true,
        chunks: false,
        chunkModules: false
    }));
});