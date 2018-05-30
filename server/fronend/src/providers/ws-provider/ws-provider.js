import _ from 'lodash';
import ConfigProvider from 'providers/config-provider';
import UUIDv3 from 'uuid/v3';
import {getTelemetry} from './rpc';
import {updateTelemetry} from './notify';

class WSProvider {
    constructor() {
        if (_.get(ConfigProvider.getConfig(), 'ws.isRelative', false)) {
            this.host = `${location.host}${_.get(ConfigProvider.getConfig(), 'ws.uri', '')}`;
        } else {
            this.host = _.get(ConfigProvider.getConfig(), 'ws.uri', '');
        }
        this.socket = null;
    }

    methods = {
        'getTelemetry': getTelemetry
    };

    events = {
        'updateTelemetry': updateTelemetry
    };

    rpc(method = '', params = {}) {
        this.send({method, req_id: UUIDv3('ex.domain.io', UUIDv3.DNS), params});
    }

    run() {
        return new Promise((resolve) => {
            const socket = new WebSocket(`ws://${this.host}`);
            this.socket = socket;
            socket.onopen = () => {
                this.socket = socket;
                resolve();
            };
            socket.onclose = () => {
                this.socket = null;
                resolve();
            };
            socket.onerror = () => {
                this.socket = null;
                const self = this;
                resolve();
                setTimeout(() => {
                    self.run();
                }, 5000);
            };
            socket.onmessage = e => {
                this.onMessage(JSON.parse(e.data));
            };
        });
    }

    onMessage(data) {
        const method = _.get(data, 'method', ''),
            result = _.get(data, 'result', null),
            event = _.get(data, 'event', null),
            fields = _.get(data, 'fields', {});

        if (method && this.methods[method]) {
            this.methods[method](result);
        }
        else if (event && this.events[event]) {
            this.events[event](fields);
        }
    }

    send(data) {
        if (this.socket) {
            this.socket.send(JSON.stringify(data));
        }
    }
}

export default new WSProvider();
