#!/bin/sh

export GOARCH=arm
export GOOS=linux
export CGO_ENABLED=1
export GOROOT=/usr/lib/go
export GOPATH=/home/dalin/go
export PATH=$GOROOT/bin:/home/dalin/go/bin:$PATH
export CROSS=/home/${USER}/x-tools/arm-kobo-linux-gnueabihf/bin/arm-kobo-linux-gnueabihf
#export SYSROOT=/home/${USER}/x-tools/arm-kobo-linux-gnueabihf/arm-kobo-linux-gnueabihf/sysroot
export AR=${CROSS}-ar
export AS=${CROSS}-as
export CC=${CROSS}-gcc
export CXX=${CROSS}-g++
export LD=${CROSS}-ld
export RANLIB=${CROSS}-ranlib
export CGO_CFLAGS="-O3 -march=armv7-a -mfpu=neon -mfloat-abi=hard -D__arm__ -D__ARM_NEON__ -fPIC -fno-omit-frame-pointer -funwind-tables -Wl,--no-merge-exidx-entries"
export QT_VERSION="5.13.2"
export QT_API="5.13.0"
export QT_DIR="/home/${USER}/qt-bin/qt-linux-${QT_VERSION}-kobo"
# export QT_INC="-I/opt/kindle-qt5/target/include/QtAccessibilitySupport -I/opt/kindle-qt5/target/include/QtConcurrent -I/opt/kindle-qt5/target/include/QtCore -I/opt/kindle-qt5/target/include/QtDBus -I/opt/kindle-qt5/target/include/QtDeviceDiscoverySupport -I/opt/kindle-qt5/target/include/QtEventDispatcherSupport -I/opt/kindle-qt5/target/include/QtFbSupport -I/opt/kindle-qt5/target/include/QtFontDatabaseSupport -I/opt/kindle-qt5/target/include/QtGlxSupport -I/opt/kindle-qt5/target/include/QtGui -I/opt/kindle-qt5/target/include/QtHelp -I/opt/kindle-qt5/target/include/QtInputSupport -I/opt/kindle-qt5/target/include/QtNetwork -I/opt/kindle-qt5/target/include/QtOpenGL -I/opt/kindle-qt5/target/include/QtOpenGLExtensions -I/opt/kindle-qt5/target/include/QtPacketProtocol -I/opt/kindle-qt5/target/include/QtPlatformCompositorSupport -I/opt/kindle-qt5/target/include/QtPlatformHeaders -I/opt/kindle-qt5/target/include/QtPrintSupport -I/opt/kindle-qt5/target/include/QtQml -I/opt/kindle-qt5/target/include/QtQmlDebug -I/opt/kindle-qt5/target/include/QtQuick -I/opt/kindle-qt5/target/include/QtQuickControls2 -I/opt/kindle-qt5/target/include/QtQuickParticles -I/opt/kindle-qt5/target/include/QtQuickTemplates2 -I/opt/kindle-qt5/target/include/QtQuickTest -I/opt/kindle-qt5/target/include/QtQuickWidgets -I/opt/kindle-qt5/target/include/QtRemoteObjects -I/opt/kindle-qt5/target/include/QtRepParser -I/opt/kindle-qt5/target/include/QtScript -I/opt/kindle-qt5/target/include/QtScriptTools -I/opt/kindle-qt5/target/include/QtServiceSupport -I/opt/kindle-qt5/target/include/QtSql -I/opt/kindle-qt5/target/include/QtTest -I/opt/kindle-qt5/target/include/QtThemeSupport -I/opt/kindle-qt5/target/include/QtUiPlugin -I/opt/kindle-qt5/target/include/QtUiTools -I/opt/kindle-qt5/target/include/QtWebChannel -I/opt/kindle-qt5/target/include/QtWidgets -I/opt/kindle-qt5/target/include/QtX11Extras -I/opt/kindle-qt5/target/include/QtXml -I/opt/kindle-qt5/target/include/QtXmlPatterns -I/opt/kindle-qt5/target/include/QtZlib"
export QT_QMAKE_DIR=${QT_DIR}/bin
# export QT_DIR=/opt/kindle-qt5/host/target
# export CGO_CFLAGS="-g -O2 -I/opt/kindle-qt5/target/include $QT_INC -L/home/dalin/opt/lib -L/home/dalin/opt/cross-gcc-linaro/arm-linux-gnueabi/lib -L/opt/kindle-qt5/target/lib --sysroot $SYSROOT"
# export CGO_CXXFLAGS="-g -O2 -I/opt/kindle-qt5/target/include $QT_INC -L/home/dalin/opt/lib -L/home/dalin/opt/cross-gcc-linaro/arm-linux-gnueabi/lib -L/opt/kindle-qt5/target/lib --sysroot $SYSROOT -std=c++11"
# export CGO_LDFLAGS="-g -O2 -I/opt/kindle-qt5/target/include $QT_INC -L/home/dalin/opt/lib -L/home/dalin/opt/cross-gcc-linaro/arm-linux-gnueabi/lib -L/opt/kindle-qt5/target/lib --sysroot $SYSROOT"
