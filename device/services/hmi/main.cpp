#include <QObject>
#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include <QQmlEngine>
#include <QJSEngine>
#include <QJSValue>
#include "natsprovider.h"
#include "hmi_provider.h"
#include "nats_api.h"

static QObject *hmi_provider(QQmlEngine *engine, QJSEngine *scriptEngine) {
    Q_UNUSED(engine)
    Q_UNUSED(scriptEngine)
    auto instance = HmiProvider::intance();
    return instance;
}

int main(int argc, char *argv[])
{
    qputenv("QT_IM_MODULE", QByteArray("qtvirtualkeyboard"));

    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);

    QGuiApplication app(argc, argv);

    HmiProvider::init();

    // Init NATS.
    Nats::Client client;
    QObject::connect(&client, &Nats::Client::connected, [&client] {
        Logger::Info("NATS", QString("Success connect to NATS"));
        NatsClient::init(&client);
        NatsProvider::init("hmi", &client, getNatsApiMap());
    });
    QObject::connect(&client, &Nats::Client::error, [](const QString &error) {
        Logger::Info("NATS", QString("Error connect to NATS: '%1'").arg(error));
    });
    client.connect("127.0.0.1", 4222);

    // Init HMI provider.
    qmlRegisterSingletonType<HmiProvider>("iot.hmi.data", 1, 0, "Api", hmi_provider);

    QQmlApplicationEngine engine;
    engine.load(QUrl(QStringLiteral("qrc:/main.qml")));
    if (engine.rootObjects().isEmpty())
        return -1;

    return app.exec();
}
