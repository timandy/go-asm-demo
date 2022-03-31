#include "textflag.h"
#include "funcdata.h"

TEXT  ·getg0(SB), NOSPLIT, $16-16
    NO_LOCAL_POINTERS
    MOVQ $0, ret_type+0(FP)
    MOVQ $0, ret_data+8(FP)
    GO_RESULTS_INITIALIZED

    // get runtime.g type
    MOVQ $type·runtime·g(SB), AX

    // get runtime·g0 variable
    MOVQ runtime·g0(SB), BX

    // return interface{}
    MOVQ AX, ret_type+0(FP)
    MOVQ BX, ret_data+8(FP)
    RET
