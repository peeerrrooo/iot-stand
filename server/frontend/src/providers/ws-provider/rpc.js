import {getStore} from '../../store';
import {setTelemetry} from 'actions/ws/actions';
import {sendInfo} from 'utils/notification';
import dateFormat from 'dateformat';
import _ from "lodash";

export function getTelemetry(params) {
    const telemetry = _.map(_.get(params, 'telemetry', []), t => ({
            battery: _.get(t, 'battery', 0),
            totalRange: _.get(t, 'total_range', 0),
            temperature: _.get(t, 'temperature', 0),
            mileage: _.get(t, 'mileage', 0),
            vin: _.get(t, 'vin', 0),
            created: dateFormat(_.get(t, 'created', ''))
        })),
        {dispatch} = getStore();
    dispatch(setTelemetry(telemetry));
}

export function removeTelemetry() {
    const {dispatch} = getStore();
    dispatch(setTelemetry([]));
    sendInfo('Success clear telemetry');
}
