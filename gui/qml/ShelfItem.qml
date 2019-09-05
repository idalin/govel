import QtQuick 2.7
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Window 2.2
import QtGraphicalEffects 1.0

Rectangle{
    id:shelfItem
    width: parent.width
    height:root.width/5 
    property int unreadChapter: book.chapterListSize-book.durChapter-1
    property var book: myShelf.model.get(index)
    
    signal clicked

    Row {
        width: parent.width
        Item{
            // width: parent.width/4
            id:coverImgItem
            height: shelfItem.height
            width: height/5*4           
            // anchors.margins: height/20           
            // radius: 20
            Image{
                id: coverImg      
                asynchronous: true                    
                anchors.fill: parent                
                // anchors.margins: height/20
                source: book.bookInfoBean.coverUrl?book.bookInfoBean.coverUrl:"qrc:/images/drawer.png"
            }
            Rectangle{
                id: mask
                anchors.fill: parent
                visible: false
                radius: 5
            }
            OpacityMask{    
                anchors.fill: parent
                source: coverImg
                maskSource: mask
            }
        }
        
        Column {                                  
            // leftPadding: height/10  
            // padding:height/10
            // spacing: height/15     
            width: parent.width-coverImgItem.width-badge.width    
            Text{
                id: titleText
                color: "lightsteelblue"
                font.bold: true
                // font.pixelSize: font.pixelSize*1.2
                text: book.bookInfoBean.name
            }
            Text{
                id: authorText              
                text: book.bookInfoBean.author
            }
            Text{
                id: durChapterName              
                text: book.durChapterName
            }

            Text{
                id: lastChapter
                text: book.lastChapterName
            }
            // Text{
            //     id: bookSourceText
            //     wrapMode: Text.WordWrap
            //     text: book.bookInfoBean.origin
            // }
        }
        
        Item{
            id:badge
            height:parent.height/5
            width: unreadChapter>99?height*1.5:height
            anchors.rightMargin: height/2
            Rectangle{
                id:unreadBadge
                anchors.fill: parent  
                smooth: true
                radius:height/2
                color:"black"
                Label{
                    id: unread
                    anchors.fill: parent
                    verticalAlignment: Text.AlignHCenter
                    horizontalAlignment: Text.AlignHCenter
                    text:unreadChapter
                    color: "white"                   
                }
            }
            
            Component.onCompleted: {
                if (book.newChapters>0){
                    unreadBadge.color="red"
                }else if(unreadChapter>0){
                    unreadBadge.color="black"
                }else{
                    unreadBadge.visible=false
                }
                if(book.newChapters>99||unreadChapter>99){
                    badge.width=badge.height*2
                }
            }
            
        }

    }
    
    MouseArea {
        anchors.fill: parent
        onClicked: {
            // console.log(title+" clicked!"+" root.y:"+root.y+" root.height:"+root.height+" popup.y:"+bookPopup.y+" popup.x:"+bookPopup.x+" item.y:"+bookItem.y);
            // console.log("model is "+ title+" .index is:"+index);            
            myShelf.currentIndex=index;
            myShelf.model.setProperty(index,"finalDate",( new Date()).valueOf());
            // console.log(myShelf.model.get(index).finalDate);
            stackView.push("Reader.qml")
            
            // bookPopup.open()
            
            // bookItem.clicked()
        }
    }
    Rectangle {
        color: "black"      
        width: parent.width
        height:1
        anchors.top: parent.bottom        
    }    
    
}