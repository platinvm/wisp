// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab true
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar WispLexer;

COMMA	: ',';
COLON	: ':';
LSQUAR	: '[';
RSQUAR	: ']';
LCURLY	: '{';
RCURLY	: '}';

BOOLEAN: 'true' | 'false';

fragment SIGN		: [+-];
fragment BIN		: [01];
fragment DEC		: [0-9];
fragment HEX_UPPER	: [0-9A-F];
fragment HEX_LOWER	: [0-9a-f];
fragment HEX		: HEX_UPPER | HEX_LOWER;

INTEGER		: SIGN? DEC+;
FLOAT		: INTEGER '.' DEC+;
BINARY		: SIGN? '0b' BIN+;
HEXADECIMAL	: SIGN? '0x' HEX+;
PERCENTAGE	: (INTEGER | FLOAT) '%';

fragment ESC	: '\\' ('"' | '\\' | 'n' | 'r' | 't');
STRING			: '"' ( ESC | ~('"' | '\\' | '\r' | '\n'))* '"';

MULTILINE_STRING: '"""' (.)*? '"""';

fragment IP_OCTET	: DEC | [1-9] DEC | '1' DEC DEC | '2' [0-4] DEC | '25' [0-5];
fragment MAC_OCTET	: HEX HEX;
fragment H16		: HEX HEX? HEX? HEX?;
fragment LS32		: H16 ':' H16 | IPV4;

MAC		: MAC_OCTET ':' MAC_OCTET ':' MAC_OCTET ':' MAC_OCTET ':' MAC_OCTET ':' MAC_OCTET;
IPV4	: IP_OCTET '.' IP_OCTET '.' IP_OCTET '.' IP_OCTET | 'localhost';

IPV6:
	H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16
	| '::' H16 (':' H16)? (':' H16)? (':' H16)? (':' H16)? (':' H16)? (':' H16)?
	| H16 '::' H16 (':' H16)? (':' H16)? (':' H16)? (':' H16)? (':' H16)?
	| H16 ':' H16 '::' H16 (':' H16)? (':' H16)? (':' H16)? (':' H16)?
	| H16 ':' H16 ':' H16 '::' H16 (':' H16)? (':' H16)? (':' H16)?
	| H16 ':' H16 ':' H16 ':' H16 '::' H16 (':' H16)? (':' H16)?
	| H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16 '::'
	| H16 ':' H16 ':' H16 ':' H16 ':' H16 '::' H16 (':' H16)?
	| H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16 '::' H16
	| H16 ':' H16 ':' H16 ':' H16 ':' H16 ':' H16 '::'
	| H16 ':' H16 ':' H16 ':' H16 ':' H16 '::' LS32
	| H16 ':' H16 ':' H16 ':' H16 ':' H16 '::'
	| H16 ':' H16 ':' H16 ':' H16 '::' LS32
	| H16 ':' H16 ':' H16 ':' H16 '::'
	| H16 ':' H16 ':' H16 '::' LS32
	| H16 ':' H16 ':' H16 '::'
	| H16 ':' H16 '::' LS32
	| H16 ':' H16 '::'
	| H16 '::' LS32
	| '::' LS32
	| H16 '::'
;

fragment COLOR_SHORT		: HEX HEX HEX;
fragment COLOR_SHORT_ALPHA	: COLOR_SHORT HEX;
fragment COLOR_FULL			: COLOR_SHORT COLOR_SHORT;
fragment COLOR_FULL_ALPHA	: COLOR_FULL HEX HEX;

COLOR: '#' ( COLOR_SHORT | COLOR_SHORT_ALPHA | COLOR_FULL | COLOR_FULL_ALPHA);

fragment ALPHA_NUMERIC		: [a-zA-Z0-9]+;
fragment ALPHA_NUM_HYPHEN	: [a-zA-Z0-9-]+;
fragment NORMAL				: DEC '.' DEC '.' DEC;
fragment PRE_RELEASE		: NORMAL '-' PRE_RELEASE_PART ('.' PRE_RELEASE_PART)*;
fragment BUILD				: NORMAL '+' BUILD_PART ('.' BUILD_PART)*;
fragment PRE_RELEASE_BUILD:
	NORMAL '-' PRE_RELEASE_PART ('.' PRE_RELEASE_PART)* '+' BUILD_PART ('.' BUILD_PART)*
;
fragment PRE_RELEASE_PART	: ALPHA_NUMERIC | DEC;
fragment BUILD_PART			: ALPHA_NUM_HYPHEN | DEC;

VERSION: NORMAL | PRE_RELEASE | BUILD | PRE_RELEASE_BUILD;

DURATION:
	(INTEGER | FLOAT) (
		's'		// second
		| 'ms'	// millisecond
		| 'us'	// microsecond
		| 'ns'	// nanosecond
		| 'm'	// minute
		| 'h'	// hour
		| 'd'	// day
		| 'w'	// week
		| 'mo'	// month
		| 'y'	// year
	) DURATION?
;

fragment PARTIAL_TIME	: DEC DEC ':' DEC DEC ':' DEC DEC ('.' DEC+)?;
fragment TIME_OFFSET	: 'Z' | ('+' | '-') DEC DEC ':' DEC DEC;
fragment FULL_DATE		: DEC DEC DEC DEC '-' DEC DEC '-' DEC DEC;
fragment FULL_TIME		: PARTIAL_TIME TIME_OFFSET;

TIMESTAMP: FULL_DATE 'T' FULL_TIME;

ID: [a-zA-Z_][a-zA-Z_0-9]*;

WS					: [ \t\n\r\f]+	-> skip;
COMMENT				: '//' ~[\n\r]*	-> skip;
MULTILINE_COMMENT	: '/*' .*? '*/'	-> skip;