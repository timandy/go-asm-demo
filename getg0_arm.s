#include "textflag.h"
#include "funcdata.h"

TEXT  ·getg0(SB), NOSPLIT, $8-8
    NO_LOCAL_POINTERS

    // get runtime.g type
    MOVL $type·runtime·g(SB), R0

    // get runtime·g0 variable
    MOVL $0, R1

    // return interface{}
    MOVL R0, ret_type+0(FP)
    MOVL R1, ret_data+4(FP)
    RTS
