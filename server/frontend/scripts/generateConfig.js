const config = require('../config/config'),
    fs = require('fs'),
    path = require('path'),
    filePath = path.resolve(__dirname, `./../src/config.json`),
    NODE_ENV = config.get('NODE_ENV'),
    data = NODE_ENV == 'production' ?
        {
            dev: false,
            data: config.get('config:prod')
        } :
        {
            dev: true,
            data: config.get('config:dev')
        };

let outStream = fs.createWriteStream(filePath);
outStream.write(JSON.stringify(data));
outStream.on('close', err => {
    if (err) {
        console.error(`Error: ${err}`);
        return;
    }

    console.log('Resolve generate config.json');
});