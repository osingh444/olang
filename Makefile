output: testlex.o lexer.o token.o
	gcc testlex.o lexer.o token.o -o output

testlex.o: testlex.c
	gcc -c testlex.c

token.o: token.c token.h
	gcc -c token.c

lexer.o: lexer.c lexer.h
	gcc -c lexer.c

clean:
	rm *.o output
