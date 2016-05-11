package testlib

/*
#cgo CPPFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra
#cgo CPPFLAGS: -DUNICODE -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_GUI_LIB -DQT_WIDGETS_LIB -DQT_TEST_LIB -DQT_TEST_LIB
#cgo CPPFLAGS: -IC:/Qt/Qt5.6.0/5.6/mingw49_32/include -IC:/Qt/Qt5.6.0/5.6/mingw49_32/mkspecs/win32-g++
#cgo CPPFLAGS: -IC:/Qt/Qt5.6.0/5.6/mingw49_32/include/QtCore -IC:/Qt/Qt5.6.0/5.6/mingw49_32/include/QtGui -IC:/Qt/Qt5.6.0/5.6/mingw49_32/include/QtWidgets -IC:/Qt/Qt5.6.0/5.6/mingw49_32/include/QtTest -IC:/Qt/Qt5.6.0/5.6/mingw49_32/include/QtTest

#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads -Wl,--allow-multiple-definition
#cgo LDFLAGS: -LC:/Qt/Qt5.6.0/5.6/mingw49_32/lib -lQt5Core -lQt5Gui -lQt5Widgets -lQt5Test -lQt5Test -lmingw32 -lqtmain -lshell32
*/
import "C"
