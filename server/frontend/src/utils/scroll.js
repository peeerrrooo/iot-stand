import {scroller} from 'react-scroll';

export function scrollTo(name) {
    scroller.scrollTo(name, {
        duration: 0,
        smooth: true
    });
}
