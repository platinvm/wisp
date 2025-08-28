// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab true
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

parser grammar WispParser;

options {
	tokenVocab = WispLexer;
}

program: expression EOF;

expression: literal | array | object;

array: '[' (expression (',' expression)* ','?)? ']';

object: '{' (pair (',' pair)* ','?)? '}';

pair: (ID | STRING) ':' expression;

literal:
	BOOLEAN
	| INTEGER
	| FLOAT
	| BINARY
	| HEXADECIMAL
	| STRING
	| MULTILINE_STRING
	| IPV4
	| IPV6
	| MAC
	| COLOR
	| VERSION
	| DURATION
	| TIMESTAMP
	| PERCENTAGE
;