import QtQuick 2.7
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Window 2.2
import QtGraphicalEffects 1.0

Rectangle{
    id:shelfItem
    width: parent.width
    height:root.width/5
    // anchors.fill: parent    
    // anchors.margins: height/15
    property alias authorName: authorText.text
    property alias title: titleText.text
    property string intro
    property alias bookSource: bookSourceText.text
    default property alias cover: coverImg.source
    property string bookUrl: ""
    property bool allowUpdate   
	property int chapter_list_size
	property int dur_chapter 
    property int unreadChapter: chapter_list_size-dur_chapter-1
	property string dur_chapter_name
	property int durChapterPage
	property string finalDate  
	property string finalRefreshDate
	property int group      
	property bool hasUpdate
	property bool isLoading
	property string last_chapter_name
	property int new_chapters
	property string noteUrl
	property int serialNumber
	property string tag
	property bool useReplaceRule
    
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
                source: source?source:"qrc:/images/drawer.png"
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
            padding:height/10
            spacing: height/15     
            width: parent.width-coverImgItem.width-badge.width    
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
                id: durChapterName              
                text: dur_chapter_name
            }
            // TextMetrics {
            //     id: introTextMetrics
            //     // width: parent.width
            //     // font.family: "Arial"
            //     elide: Text.ElideRight
            //     elideWidth: parent.width*2
            //     text: "介绍"
            // }
            Text{
                id: lastChapter
                // width:parent.width
                // wrapMode: Text.WordWrap
                text: last_chapter_name
            }
            Text{
                id: bookSourceText
                wrapMode: Text.WordWrap
                text: "书源"
            }
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
                if (new_chapters>0){
                    unreadBadge.color="red"
                }else if(unreadChapter>0){
                    unreadBadge.color="black"
                }else{
                    unreadBadge.visible=false
                }
                if(new_chapters>99||unreadChapter>99){
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