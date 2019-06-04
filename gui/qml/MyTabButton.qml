import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
 TabButton {
    id: allBooks
    text: qsTr("所有书籍")
    contentItem: Text {
        text: allBooks.text
        font: allBooks.font
        opacity: enabled ? 1.0 : 0.3
        color: "#000000"
        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
        elide: Text.ElideRight
    }
    background: Rectangle {
        implicitWidth: 100
        implicitHeight: 40
        // opacity: enabled ? 1 : 0.3

        Rectangle {
            width: parent.width
            height: 1
            color: allBooks.checked ? "#000000" : "#FFFFFF"
            anchors.bottom: parent.bottom
        }
    }
}