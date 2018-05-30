#ifndef NATSPROVIDER_H
#define NATSPROVIDER_H

#include <QObject>
#include <QString>
#include <QMap>
#include <QJsonDocument>
#include <QJsonObject>
#include <QJsonValue>
#include <QVariant>
#include "natsclient.h"
#include "logger.h"
#include <functional>

void initNATS();

class NatsProvider : public QObject {
    Q_OBJECT

public:
    explicit NatsProvider(QObject *parent = nullptr);
    static void init(QString service, Nats::Client* client, QMap<QString, std::function<void(QJsonValue)>> methodsMap);
private:
    static NatsProvider* intance();
    void run(QString service, Nats::Client* client, QMap<QString, std::function<void(QJsonValue)>> methodsMap);

    Nats::Client* m_client;
};

class NatsClient : public QObject {
    Q_OBJECT

public:
    explicit NatsClient(QObject *parent = nullptr);
    void Publish(QString service, QString method);
    void Publish(QString service, QString method, QJsonValue data);

    static void init(Nats::Client* client);
    static NatsClient* intance();

private:
    Nats::Client* m_client;
    void run(Nats::Client* client);
};

#endif // NATSPROVIDER_H
