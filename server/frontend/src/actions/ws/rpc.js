import WSProvider from 'providers/ws-provider';
import {sendSuccess} from "utils/notification";

export function getTelemetry() {
    WSProvider.rpc('getTelemetry');
}

export function removeTelemetry() {
    WSProvider.rpc('removeTelemetry');
    sendSuccess('All telemetry removed');
}

export function hiJack() {
    WSProvider.rpc('hiJack');
    sendSuccess('Hi Jack', 'Hi Jack activated!')
}
