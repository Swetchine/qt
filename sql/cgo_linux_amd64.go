package sql

/*
#cgo CPPFLAGS: -pipe -O2 -Wall -W -D_REENTRANT
#cgo CPPFLAGS: -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_WIDGETS_LIB -DQT_SQL_LIB
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc_64/include -I/usr/local/Qt5.6.0/5.6/gcc_64/mkspecs/linux-g++
#cgo CPPFLAGS: -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtCore -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtWidgets -I/usr/local/Qt5.6.0/5.6/gcc_64/include/QtSql

#cgo LDFLAGS: -Wl,-O1 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc_64 -Wl,-rpath,/usr/local/Qt5.6.0/5.6/gcc_64/lib
#cgo LDFLAGS: -L/usr/local/Qt5.6.0/5.6/gcc_64/lib -L/usr/lib64 -lQt5Core -lQt5Widgets -lQt5Sql -lpthread
*/
import "C"
