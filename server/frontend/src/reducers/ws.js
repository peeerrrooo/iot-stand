import _ from 'lodash';
import {WS__TELEMETRY_GET} from 'actions/ws/types';

const defaultState = {};

export default (state = defaultState, action = false) => {
    const type = _.get(action, 'type', false);
    switch (type) {
        case WS__TELEMETRY_GET:
            return {
                ...state,
                telemetry: action.telemetry
            };
        default:
            return state;
    }
}
