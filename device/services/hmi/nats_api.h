#ifndef NATS_API_H
#define NATS_API_H

#include <QMap>
#include <QJsonValue>
#include <functional>
#include "hmi_provider.h"

QMap<QString, std::function<void(QJsonValue)>> getNatsApiMap();

#endif // NATS_API_H
