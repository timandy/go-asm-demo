#include "textflag.h"
#include "funcdata.h"

TEXT  ·getg0(SB), NOSPLIT, $8-8
    NO_LOCAL_POINTERS

    // get runtime.g type
    MOVW $type·runtime·g(SB), R8

    // get runtime·g0 variable
    MOVW $0, R9

    // return interface{}
    MOVW R8, ret_lo+0(FP)
    MOVW R9, ret_hi+4(FP)
    RET
