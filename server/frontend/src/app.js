import React from 'react';
import {render} from 'react-dom';
import {initStore} from './store';
import WSProvider from 'providers/ws-provider';
import history from 'utils/history';
import Root from './Root';
import './styles/typography';
import './styles/antd';

function renderApp(Component, store, history) {
    render(
        <Component store={store} history={history}/>,
        document.getElementById('app')
    );
}

const initState = initStore({}, history);

(async function () {
    await WSProvider.run();
    renderApp(Root, initState, history);
})();

