#ifndef LOGGER_H
#define LOGGER_H

#include <QString>

class Logger {
public:
    Logger();
    static void Info(QString service, QString message);
    static void Error(QString service, QString message);
};

#endif // LOGGER_H
