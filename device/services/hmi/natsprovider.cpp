#include "natsprovider.h"

NatsProvider::NatsProvider(QObject *parent) :
    QObject(parent) {
}

void NatsProvider::run(QString service, Nats::Client* client, QMap<QString, std::function<void(QJsonValue)>> methodsMap) {
    m_client = client;
    QString serviceName = QString("%1-SERVICE").arg(service.toUpper());
    Nats::Subscription *s = m_client->subscribe(QString("%1-service").arg(service));
    Logger::Info(serviceName, QString("Connect subsribe: '%1'").arg(serviceName));

    QObject::connect(s, &Nats::Subscription::received, [serviceName, s, methodsMap] {
        Logger::Info(serviceName, QString("Get message: '%1'").arg(s->message));

        // Parse event map.
        auto doc = QJsonDocument::fromJson(s->message.toUtf8());
        auto ob = doc.object();
        auto method = ob["method"].toString();
        auto data = ob["data"];
        if (methodsMap.contains(method)) {
            Logger::Info(serviceName, QString("Call method: %1").arg(method));
            methodsMap[method](data);
        }
    });
}

NatsProvider* NatsProvider::intance() {
    static auto provider = new NatsProvider();
    return provider;
}

void NatsProvider::init(QString service, Nats::Client* client, QMap<QString, std::function<void(QJsonValue)>> methodsMap) {
    auto provider = NatsProvider::intance();
    provider->run(service, client, methodsMap);
}

NatsClient::NatsClient(QObject *parent) :
    QObject(parent) {
}

void NatsClient::init(Nats::Client *client) {
    auto instance = NatsClient::intance();
    instance->run(client);
}

NatsClient* NatsClient::intance() {
    static auto provider = new NatsClient();
    return provider;
}

void NatsClient::run(Nats::Client* client) {
    m_client = client;
}

void NatsClient::Publish(QString service, QString method) {
    if (m_client != nullptr) {
        auto ob = QJsonObject();
        ob["method"] = method;
        auto doc = QJsonDocument(ob);
        m_client->publish(QString("%1-service").arg(service), QString(doc.toJson()));
        Logger::Info("HMI", QString("Send NATS to method: %1").arg(method));
    }
}

void NatsClient::Publish(QString service, QString method, QJsonValue data) {
    if (m_client != nullptr) {
        auto ob = QJsonObject();
        ob["method"] = method;
        ob["data"] = data;
        auto doc = QJsonDocument(ob);
        m_client->publish(QString("%1-service").arg(service), QString(doc.toJson()));
        Logger::Info("HMI", QString("Send NATS to method: %1").arg(method));
    }
}
