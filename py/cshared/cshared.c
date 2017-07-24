#include "py.h"
#include <stdio.h>

int main() {
    char* c = "Hello cgo.";
    GoCall(c);
}