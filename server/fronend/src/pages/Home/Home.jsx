import _ from 'lodash';
import React, {PureComponent, Fragment} from 'react';
import {connect} from 'react-redux';
import {Row, Col, Badge, Tooltip} from 'antd';
import Block from 'components/Block';
import {getTelemetry} from 'actions/ws/rpc';
import Styles from 'styles/constants';
import {root, vinContainer, rowContainer, itemContainer, itemBadge, noneData} from './main-home';

@connect(store => {
    const {ws} = store;
    return {
        telemetry: _.get(ws, 'telemetry', [])
    };
})
class Home extends PureComponent {

    state = {
        isLoading: true
    };

    async componentDidMount() {
        await getTelemetry();
    }

    componentWillReceiveProps(props) {
        if (props.telemetry !== this.props.telemetry) {
            this.setState({
                isLoading: false
            });
        }
    }

    render() {
        const {telemetry} = this.props,
            {isLoading} = this.state;

        return (
            <div className={root}>
                {!_.isEmpty(telemetry) ?
                    <Fragment>
                        <Block className={vinContainer}>
                            <Row type="flex" justify="space-between">
                                <Col>
                                    <span>VIN: </span>
                                    <Tooltip title="VIN of device">
                                        <span style={{color: Styles.colors.infoColor}}>
                                            {_.get(telemetry, '[0].vin', '')}
                                        </span>
                                    </Tooltip>
                                </Col>
                            </Row>
                        </Block>
                        {_.map(telemetry, (t, index) => (
                            <Row gutter={24} key={index} className={rowContainer}>
                                <Col span={5}>
                                    <Block title="Battery" isLoading={isLoading}>
                                        <div className={itemContainer}>
                                    <span className={itemBadge}>
                                        <Badge status="success"/>
                                    </span>
                                            <span style={{color: Styles.colors.successColor}}>{t.battery}</span>
                                        </div>
                                    </Block>
                                </Col>
                                <Col span={5}>
                                    <Block title="Total range" isLoading={isLoading}>
                                        <div className={itemContainer}>
                                    <span className={itemBadge}>
                                        <Badge status="processing"/>
                                    </span>
                                            <span style={{color: Styles.colors.infoColor}}>{t.totalRange}</span>
                                        </div>
                                    </Block>
                                </Col>
                                <Col span={5}>
                                    <Block title="Temperature" isLoading={isLoading}>
                                        <div className={itemContainer}>
                                    <span className={itemBadge}>
                                        <Badge status="warning"/>
                                    </span>
                                            <span style={{color: Styles.colors.warningColor}}>{t.temperature}</span>
                                        </div>
                                    </Block>
                                </Col>
                                <Col span={5}>
                                    <Block title="Mileage" isLoading={isLoading}>
                                        <div className={itemContainer}>
                                    <span className={itemBadge}>
                                        <Badge status="processing"/>
                                    </span>
                                            <span style={{color: Styles.colors.infoColor}}>{t.mileage}</span>
                                        </div>
                                    </Block>
                                </Col>
                                <Col span={4}>
                                    <Block title="Created" isLoading={isLoading}>
                                        <div className={itemContainer}>
                                    <span className={itemBadge}>
                                        <Badge status="default"/>
                                    </span>
                                            <span style={{color: Styles.colors.normalColor}}>{t.created}</span>
                                        </div>
                                    </Block>
                                </Col>
                            </Row>
                        ))}
                    </Fragment> :
                    <div className={noneData}>
                        <h1>None data</h1>
                    </div>
                }
            </div>
        );
    }
}

export default Home;
