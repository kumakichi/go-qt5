#include <QtGui>
#include <QLabel>
#include <QDialog>
#include <QApplication>

#define FUNC_NULL   0
#define FUNC_TWO    2

typedef void (*NULL_FUNC) ();
typedef int (*TWO_FUNC) (int, int);
NULL_FUNC oooo;
TWO_FUNC xxxx;

extern "C" void qtDebug(const char *typeName)
{
	qDebug() << "Debug:" << typeName;
}

extern "C" int start(const char *typeName)
{
	int ret, argc = 0;
	char **argv = 0;

	oooo();
	ret = xxxx(7, 8);
	qDebug() << "Calculate 7+8 using func written in golang, equals" << ret;

	QApplication a(argc, argv);
	QDialog w;
	QLabel l(&w);
	l.setText(typeName);
	w.show();
	return a.exec();
}

extern "C" void bind_go_export_funcs_for_fucking_ms(int fn_type, void *fn)
{
	switch (fn_type) {
	case FUNC_NULL:
		oooo = (NULL_FUNC) fn;
		break;
	case FUNC_TWO:
		xxxx = (TWO_FUNC) fn;
		break;
	}
}
