// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.
// 初期化
// 最大スクリーンアドレスの計算
@8192 //256*32
D=A
@SCREEN
D=D+A
@MAXADDRESS //最大スクリーンアドレス
M=D

// 以下、無限ループ
(KEY)

@SCREEN
D=A
@address // スクリーンアドレス初期化
M=D

@KBD //キーボード取得
D=M

// 振り分け
@WHITE
D;JEQ
@BLACK
0;JMP

(WHITE)
@color
M=0
@LOOP
0;JMP

(BLACK)
@color
M=-1
@LOOP
0;JMP

// スクリーン書き換え
(LOOP)
@color
D=M

@address
A=M // 値のアドレスに移動
M=D //アドレスの値をcolorに

D=A+1 //アドレスに1加えたところに移動
@address // 必要！！！
M=D // 次に移動するために新たなアドレスを値として保存

@MAXADDRESS
D=M-D //Dが0かどうか

@LOOP
D;JNE

@KEY
0;JMP
(END)