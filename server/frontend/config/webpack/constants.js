const path = require('path');

const ROOT_PATH = path.resolve(__dirname, `./../../`);

module.exports = {
    ROOT_PATH: ROOT_PATH,
    SRC_PATH: path.resolve(ROOT_PATH, 'src')
};