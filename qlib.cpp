#include <QtGui>
#include <QLabel>
#include <QDialog>
#include <QApplication>

extern "C" void qtDebug(const char *typeName)
{
    qDebug() << "Debug:" << typeName;
}

extern "C" void func_written_in_go() {}

extern "C" int start(const char *typeName) {
    int argc =0;
    char **argv = 0;
    QApplication a(argc, argv);
    func_written_in_go();
    QDialog w;
    QLabel l(&w);
    l.setText(typeName);
    w.show();
    return a.exec();
}
