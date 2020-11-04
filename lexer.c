#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "lexer.h"

static void allocation_failed(void) {
  printf("Out of memory.\n");
  exit(1);
}

static void file_fail(void) {
	printf("some file shit failed\n");
	exit(1);
}

static int isTokenBreak(char ch) {
	switch(ch) {
		case ' ': return 1;
		case '\n': return 1;
		default: return 0;
	}
}

static int isAlpha(char ch) {
	if((ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122) || ch == 45) {
		return 1;
	}
	return 0;
}

static int isNumeric(char ch) {
	if(ch <= 57 && ch >= 48) {
		return 1;
	}
	return 0;
}

static types getOperatorType(char ch) {
	switch(ch) {
		case '(': return LPAREN;
		case ')': return RPAREN;
		case '{': return LBRACK;
		case '}': return RBRACK;
		case '+': return ADD;
		case '-': return SUBTRACT;
		case '*': return MULTIPLY;
		case '/': return DIVIDE;
		case '\'': return SINGLEQUOTE;
		case '"': return DOUBLEQUOTE;
		case '>': return GREATERTHAN;
		case '<': return LESSTHAN;
		case ',': return COMMA;
		case '=': return EQUALS;
		case '.': return PERIOD;
		default: return NOTOP;
	}
}

static int isOperator(char ch) {
	return getOperatorType(ch) != NOTOP;
}

static int isSameType(types type, char ch) {
	switch (type) {
		case IDENTIFIER:
			if(isAlpha(ch) || isNumeric(ch)) {
				return 1;
			} else {
				return 0;
			}
		case NUMBER:
			if(isNumeric(ch)) {
				return 1;
			} else {
				return 0;
			}
		default: return 0;
	}
}

static types getType(char ch) {
	types t;
	if(isAlpha(ch)) {
		return IDENTIFIER;
	} else if(isNumeric(ch)) {
		return NUMBER;
	} else if(isOperator(ch)) {
		return OPERATOR;
	}
	return INVALID;
}

int getNumChars(FILE *file) {
	int numChars = 0;
	char ch;
	while(1) {
		ch = fgetc(file);
        if(ch == EOF) {
			break;
		}
        numChars += 1;
	}
	return numChars;
}

lexer_t* createLexer(char* filename) {
	lexer_t* lexer = malloc(sizeof(lexer_t));
	if(!lexer) {
		allocation_failed();
	}
	lexer->numTokens = 0;

	//read the file in
	FILE *fileptr = fopen(filename, "r");
	if(!fileptr) {
		file_fail();
	}

 	lexer->len = (unsigned long)getNumChars(fileptr);

	fseek(fileptr, 0, SEEK_SET);
	lexer->file = malloc((lexer->len + 1) * sizeof(char));

	fread(lexer->file, 1, lexer->len, fileptr);
	fclose(fileptr);
	lexer->file[lexer->len] = '\0';
	return lexer;
}

void lex(lexer_t* lexer) {
	token_t* currToken = createToken(BEGINNING, "", 0);
	lexer->tokens = currToken;
	unsigned int currpos = 0;

	while(currpos < lexer->len) {
		token_t* nextToken = advanceToken(&currpos, lexer->file);

		if(nextToken->len != 0) {
			currToken->next = nextToken;
			lexer->numTokens += 1;
			currToken = nextToken;
		}
	}
}

char* copyString(char* buf, int len) {
	char* newbuf = malloc((len + 1) * sizeof(char));
	if(!newbuf) {
		allocation_failed();
	}
	newbuf[len] = '\0';
	return strncpy(newbuf, buf, len);
}

token_t* advanceToken(int* currpos, char* file) {
	types currtype = getType(file[*currpos]);
	int maxlen = 20;
	int currlen = 0;
	char* token_val = malloc(maxlen * sizeof(char));
	if(!token_val) {
		allocation_failed();
	}

	while(!isTokenBreak(file[*currpos])) {
		if(currtype != getType(file[*currpos])) {
			char* short_token = copyString(token_val, currlen);
			free(token_val);
			return createToken(currtype, short_token, currlen);
		}

		if(file[*currpos] == '/' && file[*currpos + 1] == '/') {
			while(file[*currpos] != '\n') {
				*currpos += 1;
			}
			break;
		}

		if(currlen == maxlen) {
			token_val = realloc(token_val, maxlen * 2);
			if(!token_val) {
				allocation_failed();
			}
			maxlen *= 2;
		}

		token_val[currlen] = file[*currpos];
		currlen += 1;
		*currpos += 1;
	}

	*currpos += 1;
	char* short_token = copyString(token_val, currlen);
	free(token_val);
	return createToken(currtype, short_token, currlen);
}

void deleteLexer(lexer_t* lexer) {
	deleteTokens(lexer->tokens);
	free(lexer->file);
	free(lexer);
}
