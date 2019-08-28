import QtQuick 2.6
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3
import QMLFileIo 1.0

Rectangle {
    id: reader
    width: 600
    height: 800
    property int pageNumber
    property int currentPage: 1
    property alias maxLineCount: content.maximumLineCount
    property string chapterName
    property string bookName
    property string contentText   
    property string chapterURL 
    
    color: "transparent"
    QMLFileIo{
        id: fileReader
        path: "/home/dalin/go/src/github.com/idalin/govel/storage/cache/明朝败家子-httpswwwzwducom/01177-第一千零三章：诛之.nb"
    }

    Rectangle{
        id: readerHeader
        width: parent.width
        height:80
        color: "white"
        z:100
        RowLayout{
            anchors.fill: parent
            ToolButton {
                id: back
                Layout.alignment: Qt.AlignLeft
                contentItem: Image {
                    id: backIcon
                    fillMode: Image.Pad
                    horizontalAlignment: Image.AlignHCenter
                    verticalAlignment: Image.AlignVCenter
                    source: "qrc:/images/back.png"
                }
                background: Rectangle {
                    color: "#FFFFFF"
                }
                onClicked: {
                    stackView.pop()                 
                }
            }
            Label{
                text:myShelf.currentItem.title
                font.pixelSize: 22
            }
            ToolButton {
                // id: readerMenu
                Layout.alignment: Qt.AlignRight
                contentItem: Image {
                    id: menuIcon
                    fillMode: Image.Pad
                    horizontalAlignment: Image.AlignHCenter
                    verticalAlignment: Image.AlignVCenter
                    source: "qrc:/images/menu@4x.png"
                }
                onClicked: readerMenu.open()
                background: Rectangle {
                    color: "#FFFFFF"
                }
                Menu {
                    id: readerMenu
                    x: parent.width - width
                    transformOrigin: Menu.TopRight
                    MenuItem {
                        text: "添加本地"
                        onTriggered: fileDialog.open()
                    }
                    MenuItem {
                        text: "添加网址"
                    }
                    MenuItem {
                        text: "退出"
                        onClicked:{
                            Qt.quit()
                        }
                    }
                }
            }
        }
    }

    Rectangle{
        id:left
        height: parent.height
        width: parent.width/3
        // border.color:"black"
        // border.width:1
        MouseArea {
            anchors.fill: parent
            onClicked:{
                content.topPadding+=reader.height;
                console.log("pre page clicked.")
                reader.currentPage-=1;
                if(content.topPadding>0){
                    content.topPadding=0;
                    reader.currentPage=1;
                }
            }
        }
    }
    Rectangle{
        id:leftMiddle
        height: parent.height/5
        x: parent.width/3
        width: parent.width/3
        // border.color:"black"
        // border.width:1
        MouseArea {
            anchors.fill: parent
            onClicked:{
                console.log("pre page clicked.");
                content.topPadding+=reader.height;
                reader.currentPage-=1;
                if(content.topPadding>0){
                    content.topPadding=0;
                    reader.currentPage=1;
                }
            }
        }
    }
    Rectangle{
        id:menu
        height: parent.height/5*3
        width: parent.width/3
        x: parent.width/3
        y: parent.height/5
        border.color:"blue"
        border.width:1
        MouseArea {
            anchors.fill: parent
            onClicked:{
                readerHeader.visible=readerHeader.visible?false:true;
                console.log("switch header.")
            }
        }
    }
    Rectangle{
        id:right
        height: parent.height
        width: parent.width/3
        x: parent.width/3*2
        // border.color:"red"
        // border.width:1
        MouseArea {
            anchors.fill: parent
            onClicked:{               
                console.log("next page clicked.");
                // console.log(content.text);
                if(reader.currentPage<reader.pageNumber){
                    content.topPadding-=reader.height;
                    reader.currentPage+=1;
                }
            }
        }
    }
    Rectangle{
        id:rightMiddle
        height: parent.height/5
        width: parent.width/3
        x: parent.width/3
        y: parent.height/5*4
        // border.color:"red"
        // border.width:1
        MouseArea {
            anchors.fill: parent
            onClicked:{
                console.log("next page clicked.")
                if(reader.currentPage<reader.pageNumber){
                    content.topPadding-=reader.height;
                    reader.currentPage+=1;
                }
            }
        }
    }

    Text {
        id:content
        height: parent.height
        width: parent.width
        wrapMode: Text.WordWrap
        color: "black"
        font.pixelSize:20
        lineHeight: font.pixelSize+5
        lineHeightMode: Text.FixedHeight
        verticalAlignment:Text.AlignTop
        textFormat: Text.PlainText
        padding: 5
        clip: true
        text:"fasdfadsfasdfffffffffffffffffd"       
        Component.onCompleted:{
            console.log("height:"+content.height);
            console.log("line height: "+content.lineHeight);           
            var ml=Math.ceil(content.height/content.lineHeight-1);
            // console.log("maximumLineCount: "+ content.maximumLineCount)
            reader.pageNumber=Math.ceil(content.lineCount/ml);
            console.log("line count:"+ content.lineCount);
            console.log("reader.pageNumber:" + reader.pageNumber);
            console.log("reading form file...");
            console.log("read done.")
        }
    }
    
    
    Component.onCompleted: {
        if(stackView){
            stackView.anchors.top=header.top;
        }
        readerHeader.visible=false;           
        // console.log( fileReader.readAll());
        content.text = fileReader.readAll();
        if (root){
            console.log("root height is:"+root.height);
            reader.height = root.height;
            reader.width = root.width;
        }
        if(header){
            header.visible = false;
        }
    }
    
}
