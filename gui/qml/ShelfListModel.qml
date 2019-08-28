import QtQuick 2.6
import QtQuick.Layouts 1.3
import QtQuick.Controls 2.0
import QMLFileIo 1.0


Item {
    property string source: ""
    property ListModel model : ListModel {
        id:jsonModel
        onDataChanged: {
            sortModel();
            shelfFile.write(serialize());
        }
    }
    property alias count: jsonModel.count


    QMLFileIo{
        id:shelfFile
        path: parent.source
    }

    onSourceChanged:{
        console.log("source changed to "+source);
        jsonModel.clear();
        if(source !=""){
            shelfFile.path = source
            var objectArray = JSON.parse(shelfFile.readAll())
        }
        for ( var key in objectArray ) {
            var jo = objectArray[key];
            jsonModel.append( jo );
        }
        // console.log(serialize());
    }

    function serialize() {
        var shelves =[];
        for(var i = 0; i < model.count; i++) {
            // console.log(JSON.stringify(model.get(i),null,4));
            shelves.push(model.get(i));
        }
        return JSON.stringify(shelves,null,4);
    }

    function sortModel(){
        var n;
        var i;
        for (n=0; n < model.count; n++){
            for (i=n+1; i < model.count; i++){
                if (model.get(n).finalDate < model.get(i).finalDate){
                    model.move(i, n, 1);
                    n=0;
                }
            }
        }
    }
}
