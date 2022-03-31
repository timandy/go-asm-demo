#include "textflag.h"
#include "funcdata.h"

TEXT  ·getg0(SB), NOSPLIT, $32-16
    NO_LOCAL_POINTERS
    MOVW $0, ret_type+0(FP)
    MOVW $0, ret_data+8(FP)
    GO_RESULTS_INITIALIZED

    // get runtime.g type
    MOVW $type·runtime·g(SB), R8

    // get runtime·g0 variable
    MOVW runtime·g0(SB), R9

    // return interface{}
    MOVW R8, ret_type+0(FP)
    MOVW R9, ret_data+8(FP)
    RET