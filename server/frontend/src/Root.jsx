import PropTypes from 'prop-types';
import React, {PureComponent} from 'react';
import {Provider} from 'react-redux';
import {ConnectedRouter} from 'react-router-redux'
import {LocaleProvider} from 'antd';
import enUs from 'antd/lib/locale-provider/en_US';
import moment from 'moment';
import 'moment/locale/es-us';
import Main from 'pages/Main';

moment.locale('es-us');

export default class Root extends PureComponent {

    static propTypes = {
        store: PropTypes.object.isRequired,
        history: PropTypes.object.isRequired
    };

    render() {
        const {store, history} = this.props;

        return (
            <Provider store={store}>
                <LocaleProvider locale={enUs}>
                    <ConnectedRouter history={history}>
                        <Main/>
                    </ConnectedRouter>
                </LocaleProvider>
            </Provider>
        );
    }
}
