#include "textflag.h"
#include "funcdata.h"

TEXT ·getg0(SB), NOSPLIT, $8-8
    NO_LOCAL_POINTERS
    MOVL $0, ret_type+0(FP)
    MOVL $0, ret_data+4(FP)
    GO_RESULTS_INITIALIZED

    // get runtime.g type
    MOVL $type·runtime·g(SB), AX

    // get runtime·g0 variable
    MOVL $runtime·g0(SB), BX

    // return interface{}
    MOVL AX, ret_type+0(FP)
    MOVL BX, ret_data+4(FP)
    RET
