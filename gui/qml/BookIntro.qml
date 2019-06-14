import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0

Item {
    id: bookIntro    
    anchors.fill: parent
    property var bookData
    
    // property alias authorName: authorText.text
    // property alias title: titleText.text
    // property alias introduce: introText.text
    // property alias bookSource: bookSourceText.text
    // default property alias cover: coverImg.source
    // property string bookUrl: ""    
    Column{       
        anchors.fill: parent       
        Rectangle {
            id: bookInst
            width: parent.width
            height: parent.height/3
            border.width: 1
            Text{
                text:bookData.authorName
            }
        }

        Text {
            id: introduce
            text: bookData.introduce       
        }
        Row{
            width: parent.width                                  
            anchors.bottom: parent.bottom           
            Button{
                id: saveToShelf
                text:"放入书架"
            }
            Button{
                id: startRead
                text:"开始阅读"
            }
        }
    }
}