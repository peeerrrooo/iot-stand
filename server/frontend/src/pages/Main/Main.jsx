import _ from 'lodash';
import React, {PureComponent} from 'react';
import {Switch, Route, Redirect, withRouter} from 'react-router-dom';
import {Layout, Row, Col, Button, Tooltip, Modal} from 'antd';
import {Element} from 'react-scroll';
import {hiJack, removeTelemetry} from 'actions/ws/rpc';
import Home from 'pages/Home';
import {mainHeader, buttonRemove, content, footer} from './main-main';
import {connect} from "react-redux";

const {Header, Footer, Content} = Layout,
    AntDConfirm = Modal.confirm;

@withRouter
@connect(store => {
    const {ws} = store;
    return {
        telemetry: _.get(ws, 'telemetry', [])
    };
})
class Main extends PureComponent {
    render() {
        const {telemetry} = this.props;

        return (
            <Layout>
                <Header className={mainHeader}>
                    <Row type="flex" justify="space-between">
                        <Col>
                            <span>IoT Hub Service</span>
                        </Col>
                        <Col>
                            {!_.isEmpty(telemetry) &&
                            <Button
                                className={buttonRemove}
                                icon="delete"
                                onClick={this.handleClickPurge}
                            >
                                Clear all telemetry
                            </Button>
                            }
                            <Tooltip title="Send 'Hi Jack">
                                <Button
                                    icon="cloud-upload"
                                    onClick={this.handleClickHiJack}
                                >
                                    Hi Jack!
                                </Button>
                            </Tooltip>
                        </Col>
                    </Row>
                </Header>
                <Layout>
                    <Content className={content}>
                        <Switch>
                            <Route path="/home" component={Home} exact/>
                            <Redirect to="/home"/>
                        </Switch>
                    </Content>
                    <Element name="end"/>
                    <Footer className={footer}>
                        © 2018 IoT Stand. All rights reserved.
                    </Footer>
                </Layout>
            </Layout>
        );
    }

    handleClickHiJack = async () => {
        AntDConfirm({
            title: `Are you sure active 'Hi Jack'?`,
            okText: 'Yes',
            okType: 'danger',
            cancelText: 'Cancel',
            async onOk() {
                await hiJack();
            },
            onCancel() {

            }
        });
    };

    handleClickPurge = () => {
        AntDConfirm({
            title: `Are you sure delete all telemetry?`,
            okText: 'Yes',
            okType: 'danger',
            cancelText: 'Cancel',
            async onOk() {
                await removeTelemetry();
            },
            onCancel() {

            }
        });
    };
}

export default Main;
