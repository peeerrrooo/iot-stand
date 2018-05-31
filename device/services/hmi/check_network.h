#ifndef CHECK_NETWORK_H
#define CHECK_NETWORK_H
#include <iostream>
#include <QObject>
#include <QThread>
#include <QTimer>
#include "logger.h"
#include "hmi_provider.h"

class CheckNetwork : public QThread {
public:
    explicit CheckNetwork();
    void run();
    void checkNetwork();
};

#endif // CHECK_NETWORK_H
