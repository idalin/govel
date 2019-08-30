import QtQuick 2.6
import QtQuick.Window 2.2
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QtQuick.Dialogs 1.0
import QtQuick.Controls.Styles 1.4
import QtQuick.VirtualKeyboard 2.2
import QtQuick.VirtualKeyboard.Settings 2.2

import CustomQmlTypes 1.0

ApplicationWindow {
    id: root
    visible: true
    title: "L:A_N:application_ID:Kindle小说"
    property real dpi: Screen.pixelDensity.toFixed(2)
    minimumWidth: 600
    minimumHeight: 800

    Setting{
        id: settings
        source:"../config.json"
        Component.onCompleted: {
            for(var c in settings.config){
                console.log("key: " + c+" value: " + settings.config[c])
            }
        }
    }

    // 头部工具栏
    header: ToolBar{
        RowLayout{
            id: toolbar
            anchors.fill: parent
            ToolButton {
                id: settingBottun
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
                        // listView.currentIndex = -1
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
                    tabbar.currentIndex = 1
                    // if(focus){
                        vkb.visible = true
                    // }
                    // pageLoader.source ="SearchView.qml"

                }
                background: Rectangle {
                    color: "#FFFFFF"
                }
            }

            ToolButton {
                id: menu
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
        onDepthChanged:{
            if(depth<=1){
                header.visible=true;
                stackView.anchors.top=header.bottom;

            }
        }
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

            // 书架
            ListView {
                id: myShelf
                // anchors.margins: 10
                // anchors.fill: parent
                spacing: 8
                orientation:ListView.Vertical
                delegate: ShelfItem{}
                ShelfListModel{
                   id: shelfModel
                   source: "./myBookShelf.json"
                }
                model: shelfModel.model

            }

            // 搜索页
            ListView {
                id: searchResult
                spacing: 8
                orientation:ListView.Vertical
                delegate: BookItem{
                    authorName: author
                    title: name
                    cover: coverUrl
                    intro: introduce
                    bookUrl: noteUrl
                    bookSource: tag
                }
                model: SearchListModel{
                    id: booksModel
                }

                Connections {
                    target: search
                    onAccepted: {
                        searchResult.model.doSearch(search.text)
                    }
                }
            }

            // 发现页，尚未实现
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
    InputPanel {
        id: vkb
        visible: false
        height: root.height/3
        anchors.right: parent.right
        anchors.left: parent.left
        anchors.bottom: parent.bottom
        onActiveChanged: {
            if(!active) { visible = false; }
        }
        Component.onCompleted: {
            VirtualKeyboardSettings.locale = "zh_CN.utf8";
            VirtualKeyboardSettings.styleName = "eink";
        }
    }
    
    function showBook(book) {
        console.log(book.title);       
    }

}
