//go:build amd64
#include "textflag.h"

// func Rdtsc() uint64
TEXT ·Rdtsc(SB), NOSPLIT, $0-8
    RDTSC
    MOVL DX, R8
    SHLQ $32, R8
    MOVL AX, R9
    ORQ R9, R8
    MOVQ R8, ret+0(FP)
    RET
