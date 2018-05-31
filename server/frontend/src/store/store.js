import {createStore, applyMiddleware, compose} from 'redux';
import thunk from 'redux-thunk';
import {routerMiddleware} from 'react-router-redux'
import reducers from '../reducers';

let currentStore = null;

export function initStore(initState = {}, history) {
    const middleware = routerMiddleware(history);

    const composeEnhancers =
        typeof window === 'object' &&
        window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ ?
            window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({}) : compose;

    currentStore = createStore(
        reducers,
        initState,
        composeEnhancers(
            applyMiddleware(middleware, thunk)
        )
    );
    return currentStore;
}

export function getStore() {
    return currentStore;
}