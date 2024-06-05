@aquiStr = global [3 x i8] c"%d\00"

declare i32 @printf(i8* %0, i32 %1)

declare i32 @puts(i8* %0)

define void @main() {
0:
	%1 = alloca i32
	store i32 0, i32* %1
	%2 = alloca [2 x i32]
	store [2 x i32] zeroinitializer, [2 x i32]* %2
	%3 = alloca [22 x i8]
	%4 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 0
	store i8 91, i8* %4
	%5 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 1
	store i8 54, i8* %5
	%6 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 2
	store i8 32, i8* %6
	%7 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 3
	store i8 120, i8* %7
	%8 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 4
	store i8 32, i8* %8
	%9 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 5
	store i8 105, i8* %9
	%10 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 6
	store i8 56, i8* %10
	%11 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 7
	store i8 93, i8* %11
	%12 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 8
	store i8 32, i8* %12
	%13 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 9
	store i8 99, i8* %13
	%14 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 10
	store i8 34, i8* %14
	%15 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 11
	store i8 92, i8* %15
	%16 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 12
	store i8 50, i8* %16
	%17 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 13
	store i8 50, i8* %17
	%18 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 14
	store i8 104, i8* %18
	%19 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 15
	store i8 111, i8* %19
	%20 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 16
	store i8 108, i8* %20
	%21 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 17
	store i8 97, i8* %21
	%22 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 18
	store i8 92, i8* %22
	%23 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 19
	store i8 50, i8* %23
	%24 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 20
	store i8 50, i8* %24
	%25 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 21
	store i8 34, i8* %25
	%26 = getelementptr [22 x i8], [22 x i8]* %3, i32 0, i32 0
	ret void
}
