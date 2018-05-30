import {notification} from 'antd';

export function sendSuccess(message, description = '') {
    return notification['success']({
        message,
        description,
        placement: 'bottomLeft'
    });
}

export function sendError(message, description = '') {
    return notification['error']({
        message,
        description,
        placement: 'bottomLeft'
    });
}

export function sendWarning(message, description = '') {
    return notification['warning']({
        message,
        description,
        placement: 'bottomLeft'
    });
}

export function sendInfo(message, description = '') {
    return notification['info']({
        message,
        description,
        placement: 'bottomLeft'
    });
}
