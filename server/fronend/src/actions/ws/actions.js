import {WS__TELEMETRY_GET} from './types';

export function getTelemetry(telemetry) {
    return {
        type: WS__TELEMETRY_GET,
        telemetry
    };
}
