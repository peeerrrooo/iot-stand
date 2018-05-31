#ifndef HMI_PROVIDER_H
#define HMI_PROVIDER_H

#include <QObject>
#include "natsprovider.h"

class HmiProvider : public QObject
{
    Q_OBJECT
    Q_PROPERTY(bool isNetwork READ getIsNetwork WRITE setIsNetwork NOTIFY networkChanged)
    Q_PROPERTY(bool isHiJack READ getIsHiJack WRITE setIsHiJack NOTIFY hiJackChanged)

public:
    explicit HmiProvider(QObject *parent = nullptr);
    static HmiProvider* intance();
    static void init();

    Q_INVOKABLE void updateTelemetry();

    bool getIsNetwork();
    void setIsNetwork(const bool network);

    bool getIsHiJack();
    void setIsHiJack(const bool hiJack);

signals:
    void networkChanged();
    void hiJackChanged();

private:
    bool m_isNetwork = false;
    bool m_hiJack = false;
};


#endif // HMI_PROVIDER_H
