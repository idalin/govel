import QtQuick 2.6
import QtQuick.Controls 2.0
import QtQuick.Layouts 1.3
import QMLFileIo 1.0

Item{
    property string source
    property var config

    QMLFileIo {
        id:configFile
        path: source
        onPathChanged: loadConfig();
    }
    onConfigChanged: saveConfig();

    function loadConfig() {
        config = JSON.parse(configFile.readAll());
    }

    function saveConfig() {
        console.log("setting changed,write to file.")
        configFile.write(JSON.stringify(config,null,4));
    }
}