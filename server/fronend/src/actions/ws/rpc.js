import WSProvider from 'providers/ws-provider';

export function getTelemetry() {
    WSProvider.rpc('getTelemetry');
}

export function removeTelemetry() {
    WSProvider.rpc('removeTelemetry');
}

export function hiJack() {
    WSProvider.rpc('hiJack');
}
