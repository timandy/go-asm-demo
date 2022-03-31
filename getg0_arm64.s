#include "textflag.h"
#include "funcdata.h"

TEXT  ·getg0(SB), NOSPLIT, $32-16
    NO_LOCAL_POINTERS
    MOVD $0, ret_type+0(FP)
    MOVD $0, ret_data+8(FP)
    GO_RESULTS_INITIALIZED

    // get runtime.g type
    MOVD $type·runtime·g(SB), R8

    // get runtime·g0 variable
    MOVD runtime·g0(SB), R9

    // return interface{}
    MOVD R8, ret_type+0(FP)
    MOVD R9, ret_data+8(FP)
    RET
