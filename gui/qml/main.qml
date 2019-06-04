import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Dialogs 1.0
import QtQuick.Controls.Styles 1.4
import CustomQmlTypes 1.0

ApplicationWindow {
    id: root
    visible: true
    title: "L:A_N:application_ID:Kindle小说"
    minimumWidth: 600
    minimumHeight: 800

    // 头部工具栏
    header: ToolBar{
        // style: ToolBarStyle{
        //     background:Rectangle{
        //         color: "#FFFFFF"
        //     }
        // }
        RowLayout{
            id: toolbar
            anchors.fill: parent                              
            ToolButton {
                id: settings                
                Layout.alignment: Qt.AlignLeft            
                property alias source: settingsIcon.source
                contentItem: Image {
                    id: settingsIcon
                    fillMode: Image.Pad
                    horizontalAlignment: Image.AlignHCenter
                    verticalAlignment: Image.AlignVCenter
                    source: stackView.depth > 1 ? "qrc:/images/back.png" : "qrc:/images/drawer.png"
                }
                background: Rectangle {
                    color: "#FFFFFF"
                }
                onClicked: {
                    if (stackView.depth > 1) {
                        stackView.pop()
                        listView.currentIndex = -1
                    } else {
                        drawer.open()
                    }
                }
            }

            TextField{
                id: search
                // anchors.fill: parent
                Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter            
                Layout.fillWidth: true             
                placeholderText: qsTr("搜索书名、作者")
                onFocusChanged: function changeText() {
                    // stackView.push("./search.qml")
                    tabbar.currentIndex = 1
                    pageLoader.source ="SearchView.qml"

                }
                background: Rectangle {
                    color: "#FFFFFF"
                }
                // onAccepted: { 
                    // console.log("Enter key word:",search.text)
                    // stackView.push("./search.qml")
                    // tabbar.currentIndex = 1
                    // pageLoader.source ="SearchView.qml"
                    // searchWS.sendTextMessage(search.text)
                // }
            }

            ToolButton {
                id: menu
                // anchors.fill:parent.TopRight
                // anchors.right: parent.right
                Layout.alignment: Qt.AlignRight
                property  alias source: menuIcon.source
                    
                contentItem: Image {
                    id: menuIcon
                    fillMode: Image.Pad
                    horizontalAlignment: Image.AlignHCenter
                    verticalAlignment: Image.AlignVCenter
                    source: "qrc:/images/menu@4x.png"
                }
                onClicked: optionsMenu.open()
                background: Rectangle {
                    color: "#FFFFFF"
                }
                Menu {
                    id: optionsMenu
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
        }  // end of toolbar 
    }

    StackView {
        id: stackView
        width: parent.width
        initialItem: mainView
        anchors.top: header.bottom
    }

    
    Item {
        id: mainView       
        TabBar {
            id: tabbar
            width: root.width
            anchors.topMargin: 2
            anchors.top: parent.top
            background: Rectangle {
                color: "#FFFFFF"
            }
            MyTabButton {
                id: allBooks
                text: qsTr("所有书籍")
            }
            MyTabButton {
                id: searchList
                text: qsTr("搜索列表")
            }
            MyTabButton {
                id: discover
                text: qsTr("发现")
            }
        }
        StackLayout {
            width: parent.width
            height: root.height - header.height-tabbar.height            
            currentIndex: tabbar.currentIndex
            anchors.top: tabbar.bottom
            anchors.topMargin: 10          
            Item {
                id: allBooksTab
                Text{
                    text: qsTr("暂未实现")
                    font: allBooks.font
                    opacity: enabled ? 1.0 : 0.3
                    color: "#000000"
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                    elide: Text.ElideRight
                }
            }
            // Item {
            //     id: searchListTab
            //     Text{
            //         text: searchList.text
            //         font: searchList.font
            //         opacity: enabled ? 1.0 : 0.3
            //         color: "#000000"
            //         horizontalAlignment: Text.AlignHCenter
            //         verticalAlignment: Text.AlignVCenter
            //         elide: Text.ElideRight
            //     }
            // }
            // SearchView{}
            Loader { id: pageLoader }
            Item {
                id: discoverTab
                Text{
                    // text: discover.text
                    text: qsTr("暂未实现")
                    font: discover.font
                    opacity: enabled ? 1.0 : 0.3
                    color: "#000000"
                    horizontalAlignment: Text.AlignHCenter
                    verticalAlignment: Text.AlignVCenter
                    elide: Text.ElideRight
                }
            }
        }
    }
    // InputPanel {
    //     id: vkb
    //     visible: true
    //     active: true
    //     anchors.fill:parent
    // }
    // FileDialog {
    //     id: fileDialog
    //     title: "选择书源文件(json格式)"
    //     folder: shortcuts.home
    //     onAccepted: {
    //         console.log("You chose: " + fileDialog.fileUrls)
    //         // Qt.quit()
    //         fileDialog.visible = false
    //     }
    //     onRejected: {
    //         console.log("Canceled")
    //         fileDialog.visible = false
    //         // Qt.quit()
    //     }
    //     Component.onCompleted: visible = false
    // }
}