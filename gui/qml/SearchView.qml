import QtQuick 2.0
import CustomQmlTypes 1.0


ListView {
    id: searchResult        
    anchors.margins: 10
    anchors.fill: parent
    spacing: 8
    orientation:ListView.Vertical 
    delegate: BookItem{
        authorName: author
        title: name
        cover: cover_url
        introduce: intro
        bookSource: book_source
        bookUrl: book_url
    }
   
    model: SearchListModel{
        id: booksModel
    }

    Connections {
        target: search
        onAccepted: {
            // searchResult.model.clear()
            searchResult.model.doSearch(search.text)
        }
    }
}

    
    
    



