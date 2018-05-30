import {WS__TELEMETRY_GET} from './types';

export function setTelemetry(telemetry) {
    return {
        type: WS__TELEMETRY_GET,
        telemetry
    };
}