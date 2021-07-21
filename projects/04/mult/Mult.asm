// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// Put your code here.
@sum
M=0

// 残りカウントをR1で初期化
@R1
D=M
@count 
M=D

(LOOP)
    // チェック
    @count
    D=M
    @END
    D;JEQ

    // 合計値計算
    @R0
    D=M
    @sum
    M=M+D

    // カウントを減らす
    @count
    M=M-1

    // ループ繰り返す
    @LOOP
    0;JMP
(END)

// R2に記録
@sum
D=M
@R2
M=D

@END
0;JMP