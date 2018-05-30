#include "nats_api.h"

QMap<QString, std::function<void(QJsonValue)>> getNatsApiMap() {
    QMap<QString, std::function<void(QJsonValue)>> methodsMap;
    methodsMap.insert("hiJack", methodsMap["hiJack"] = [](QJsonValue fields) {
        HmiProvider::intance()->setIsHiJack(true);
    });
    return methodsMap;
}
