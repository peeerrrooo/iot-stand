const nconf = require('nconf'),
    path = require('path'),
    fs = require('fs');

const configFile = path.resolve(__dirname, './../', 'etc/config.yaml');

try {
    const stat = fs.statSync(configFile);
    if (!(stat && stat.isFile()))
        throw new Error('Configuration file ' + configFile + ' not found');
} catch (e) {
    console.log('Configuration file ' + configFile + ' not found');
    process.exit(1);
}

nconf.formats.yaml = require('nconf-yaml');
nconf.argv()
    .env()
    .file({file: configFile, format: nconf.formats.yaml});

module.exports = nconf;