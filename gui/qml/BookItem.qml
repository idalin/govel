import QtQuick 2.6
import QtQuick.Layouts 1.3

Rectangle{
    id:bookItem
    width: parent.width
    height:190
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
        Image{
            id: coverImg      
            asynchronous: true     
            width: parent.width/4
            height: width/4*5
            anchors.margins: 10
            source: source?source:"qrc:/images/drawer.png"
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
    Rectangle{
        height:1
        width: parent.width
        color: "black"
        
    }
    MouseArea {
        anchors.fill: parent
        onClicked: {
            console.log(title+" clicked!");
            
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
}