import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Window 2.2


Rectangle{
    id:bookItem
    width: parent.width
    height:48*root.dpi
    // anchors.fill: parent    
    anchors.margins: 10
    property alias authorName: authorText.text
    property alias title: titleText.text
    property alias introduce: introText.text
    property alias bookSource: bookSourceText.text
    default property alias cover: coverImg.source
    property string bookUrl: ""
    
    signal clicked

    Row {
        width: parent.width
        Rectangle{
            width: parent.width/4
            height: width/4*5
            radius: 20
            Image{
                id: coverImg      
                asynchronous: true                    
                anchors.fill: parent                
                anchors.margins: 10
                source: source?source:"qrc:/images/drawer.png"
            }
        }
        
        Column {                                  
            leftPadding: 10  
            spacing: 10         
            Text{
                id: titleText
                color: "lightsteelblue"
                font.bold: true
                // font.pixelSize: font.pixelSize*1.2
                text: "书名"
            }
            Text{
                id: authorText              
                text: "作者"
            }
            Text{
                id: introText
                text: "简介"
            }
            Text{
                id: bookSourceText
                text: "书源"
            }
        }
        
    }
    
    MouseArea {
        anchors.fill: parent
        onClicked: {
            console.log(title+" clicked!"+" root.y:"+root.y+" root.height:"+root.height+" popup.y:"+bookPopup.y);
            bookPopup.open()
            
            // bookItem.clicked()
        }
    }
    ListView.onAdd:{
        console.log("new item added. title:"+title+" url:"+bookUrl+"cover:\'"+cover+"\'");
        if (cover!=""){
            coverImg.source=cover
        }else{
            coverImg.source="qrc:/images/drawer.png"
        }
    } 
    Rectangle {
        color: "black"      
        width: parent.width
        height:1
        anchors.top: parent.bottom        
    } 
    Popup {
        id:bookPopup
        width: root.width/3*2
        height: root.height/2
        // anchors.centerIn: Overlay.overlay
        x: Math.round((parent.width - width) / 2)
        y: Math.round((Screen.desktopAvailableHeight - height) / 2)
        // Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
        // x:(root.width-width)/2
        // y:0
        modal: true
        focus: true
        closePolicy: Popup.CloseOnPressOutside | Popup.CloseOnPressOutsideParent

        Rectangle
        {
//            width: 400
//            height: 300
            anchors.fill: parent
            Text {
                id: mytext
                font.pixelSize: 24
                text: qsTr("Popup 内容显示模块")
            }
        }
    }
   
}