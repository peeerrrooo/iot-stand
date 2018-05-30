import {notification} from 'antd';

export function sendSuccess(message, description = '') {
    return notification['success']({
        message,
        description
    });
}

export function sendError(message, description = '') {
    return notification['error']({
        message,
        description
    });
}

export function sendWarning(message, description = '') {
    return notification['warning']({
        message,
        description
    });
}

export function sendInfo(message, description = '') {
    return notification['info']({
        message,
        description
    });
}