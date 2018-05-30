#include "logger.h"

Logger::Logger() {}

void Logger::Info(QString service, QString message) {
    qInfo("%s", qUtf8Printable(QString("%1: %2").arg(service, message)));
}

void Logger::Error(QString service, QString message) {
    qWarning("%s", qUtf8Printable(QString("%1: %2").arg(service, message)));
}
