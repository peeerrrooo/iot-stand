import React from 'react';
import PropTypes from 'prop-types';
import {Card, Tooltip, Badge} from 'antd';
import {Spin} from 'antd';
import classNames from 'classnames';
import {root, titleCount, loader, spinLoader, cardContainer, cardLoader} from './main-block';

const Block = props => {
    const {className, children, title, onClick, isLoading} = props;
    let cardProps = {};
    _.each(props, (value, key) => {
        if (key !== 'isLoading' &&
            key !== children) {
            cardProps[key] = value;
        }
    });
    cardProps = _.merge({}, cardProps, {
        className: classNames(cardContainer, className, {
            [cardLoader]: isLoading
        }),
        children: null,
        onClick
    });
    return (
        <div className={root}>
            <Card
                hoverable
                {...cardProps}
            >
                {children}
            </Card>
            {isLoading &&
            <div className={loader}>
                <Spin className={spinLoader}/>
            </div>
            }
        </div>
    );
};

Block.propTypes = {
    className: PropTypes.string,
    title: PropTypes.any,
    isLoading: PropTypes.bool,
    onClick: PropTypes.func,
    isHover: PropTypes.bool
};

Block.defaultProps = {
    className: '',
    isLoading: false,
    isHover: false
};

export default Block;

