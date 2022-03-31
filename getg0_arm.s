#include "textflag.h"
#include "funcdata.h"

TEXT  ·getg0(SB), NOSPLIT, $8-8
    NO_LOCAL_POINTERS

    // get runtime.g type
    MOVW $type·runtime·g(SB), R0
    RET
