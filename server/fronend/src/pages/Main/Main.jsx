import React, {PureComponent} from 'react';
import {Switch, Route, Redirect, withRouter} from 'react-router-dom';
import {Layout, Row, Col, Button, Tooltip} from 'antd';
import {hiJack} from 'actions/ws/rpc';
import {sendSuccess} from 'utils/notification';
import Home from 'pages/Home';
import {mainHeader, content, footer} from './main-main';

const {Header, Footer, Content} = Layout;

@withRouter
class Main extends PureComponent {
    render() {
        return (
            <Layout>
                <Header className={mainHeader}>
                    <Row type="flex" justify="space-between">
                        <Col>
                            <span>IoT Hub Service</span>
                        </Col>
                        <Col>
                            <Tooltip title="Send 'Hi Jack"/>
                            <Button
                                icon="cloud-upload"
                                onClick={this.handleClickHiJack}
                            >
                                Hi Jack!
                            </Button>
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
                    <Footer className={footer}>
                        © 2018 IoT Stand. All rights reserved.
                    </Footer>
                </Layout>
            </Layout>
        );
    }

    handleClickHiJack = async () => {
        await hiJack();
        await sendSuccess('Hi Jack', 'Hi Jack activated!')
    };
}

export default Main;
