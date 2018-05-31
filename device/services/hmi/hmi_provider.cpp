#include "hmi_provider.h"

HmiProvider::HmiProvider(QObject *parent) :
    QObject(parent) {
}

HmiProvider* HmiProvider::intance() {
    static auto item = new HmiProvider();
    return item;
}

void HmiProvider::init() {
    HmiProvider::intance();
}

void HmiProvider::updateTelemetry() {
    NatsClient::intance()->Publish("telemetry", "updateTelemetry");
}

bool HmiProvider::getIsHiJack() {
    return m_hiJack;
}

void HmiProvider::setIsHiJack(const bool hiJack) {
    m_hiJack = hiJack;
    emit hiJackChanged();
}

bool HmiProvider::getIsNetwork() {
    return m_isNetwork;
}

void HmiProvider::setIsNetwork(const bool network) {
    m_isNetwork = network;
    emit networkChanged();
}
