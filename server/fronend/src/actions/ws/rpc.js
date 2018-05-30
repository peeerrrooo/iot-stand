import WSProvider from 'providers/ws-provider';

export function getTelemetry() {
    WSProvider.rpc('getTelemetry');
}

export function hiJack() {
    WSProvider.rpc('hiJack');
}
