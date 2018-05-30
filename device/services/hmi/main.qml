import QtQuick 2.9
import QtQuick.Window 2.2
import QtQuick.VirtualKeyboard 2.2
import iot.hmi.data 1.0

Window {
    id: window
    visible: true
    visibility: Window.Maximized
    title: qsTr("HMI")

    Rectangle {
        anchors.fill: parent
        id: root
        color: "#37474F"

        Rectangle {
            id: telemetryContainer
            visible: !timerJack.running
            width: parent.width * 0.5
            height: parent.height * 0.3
            color: telemetryButtonArea.pressed ? "#00796B" : "#009688"
            radius: root.width * 0.1
            border {
                width: 0
            }

            anchors {
                verticalCenter: parent.verticalCenter
                horizontalCenter: parent.horizontalCenter
            }

            Text {
                font.pointSize: root.width * 0.02
                text: "Send Telemetry"
                color: "#fff"
                anchors {
                    verticalCenter: parent.verticalCenter
                    horizontalCenter: parent.horizontalCenter
                }
            }

            MouseArea {
                id: telemetryButtonArea
                anchors.fill: parent
                cursorShape: Qt.PointingHandCursor
                onClicked: {
                    Api.updateTelemetry()
                }
            }
        }

        Item {
            id: hiJackContainer
            width: parent.width * 0.5
            height: parent.height * 0.3
            visible: timerJack.running

            anchors {
                verticalCenter: parent.verticalCenter
                horizontalCenter: parent.horizontalCenter
            }

            Text {
                id: hiJackText
                font.pointSize: root.width * 0.08
                text: "I'm here"
                color: "#fff"
                anchors {
                    verticalCenter: parent.verticalCenter
                    horizontalCenter: parent.horizontalCenter
                }
            }
        }

        Timer {
            id: timerJack
            interval: 6000
            onTriggered: {
                Api.isHiJack = false
            }
            running: Api.isHiJack
        }

        SequentialAnimation {
            running: timerJack.running

            ColorAnimation {
                target: root
                property: "color"
                duration: 1500
                from: "#FBC02D"
                to: "#FFA000"
                easing.type: Easing.InOutQuad
            }

            ColorAnimation {
                target: root
                property: "color"
                duration: 1500
                from: "#FFA000"
                to: "#F57C00"
                easing.type: Easing.InOutQuad
            }

            ColorAnimation {
                target: root
                property: "color"
                duration: 1500
                from: "#F57C00"
                to: "#E64A19"
                easing.type: Easing.InOutQuad
            }

            ColorAnimation {
                target: root
                property: "color"
                duration: 1500
                from: "#E64A19"
                to: "#37474F"
                easing.type: Easing.InOutQuad
            }
        }

        SequentialAnimation {
            running: timerJack.running

            NumberAnimation {
                target: hiJackText
                property: "scale"
                duration: 3000
                from: 1
                to: 2
                easing.type: Easing.InOutQuad
            }

            NumberAnimation {
                target: hiJackText
                property: "scale"
                duration: 3000
                from: 2
                to: 1
                easing.type: Easing.InOutQuad
            }
        }
    }
}
