#include "check_network.h"

CheckNetwork::CheckNetwork() {}

void CheckNetwork::run() {
    checkNetwork();
}

void CheckNetwork::checkNetwork() {
    if (system("ping -c 1 8.8.8.8")) {
        Logger::Error("NETWORK", "Failed check connect network");
        HmiProvider::intance()->setIsNetwork(false);
        sleep(1);
        this->checkNetwork();
    } else {
        Logger::Info("NETWORK", "Success check connect network");
        HmiProvider::intance()->setIsNetwork(true);
        sleep(1);
        this->checkNetwork();
    }
}
