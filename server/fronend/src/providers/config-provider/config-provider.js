import config from '../../config.json';

class ConfigProvider {
    constructor() {
        this.config = _.get(config, 'data', {});
    };

    getConfig() {
        return this.config;
    }

    isDev() {
        return _.get(config, 'dev', false);
    }
}

export default new ConfigProvider();