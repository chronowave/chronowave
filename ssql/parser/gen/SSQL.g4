/*
 * Copyright 2020 Rock Lei Wang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Package parser declares an expression parser with support for macro
 * expansion.
 */

grammar SSQL;

// Grammar Rules
// =============

start
    : FIND selection WHERE expression orderBy? limit? EOF
    ;

selection
    : attribute (',' attribute)*
    | GROUP_BY '(' groupBy (',' groupBy)* ')' (',' aggregate)*
    ;

attribute
    : IDENTIFIER | aggregate
    ;

aggregate
    : (AVG | MAX | MIN | SUM | COUNT) '(' IDENTIFIER ')' | percentile
    ;

percentile
    : PERCENTILE '(' IDENTIFIER ',' REAL_NUMBER ')'
    ;

groupBy
    : IDENTIFIER | partition
    ;

partition
    : PARTITION '(' IDENTIFIER ',' INTEGER ')'
    ;

expression
    : tuple tuple*
    ;

tuple
    : vector | or | and
    ;

vector
    : '[' IDENTIFIER? PATH (predicate | vector+)? ']'
    ;

or
    : '{' tuple+ '}'
    ;

and
    : '{&' tuple+ '}'
    ;

predicate
    : eq | neq | gt | ge | lt | le | in | between | contain | exist
    | timeframe | key
    ;

eq
    : EQ '(' scalar ')'
    ;

neq
    : NEQ '(' scalar ')'
    ;

gt
    : GT '(' scalar ')'
    ;

ge
    : GE '(' scalar ')'
    ;

lt
    : LT '(' scalar ')'
    ;

le
    : LE '(' scalar ')'
    ;

in
    : IN '(' list ')'
    ;

between
    : BETWEEN '(' INTEGER ',' INTEGER ')'
    | BETWEEN '(' REAL_NUMBER ',' REAL_NUMBER ')'
    ;

contain
    : CONTAIN '(' STRING ')'
    ;

exist
    : EXIST ('(' ')')?
    ;

timeframe
    : TIMEFRAME '(' INTEGER ',' INTEGER ')'
    ;

key
    : KEY '(' (INTEGER | STRING) ')'
    ;

scalar
    : REAL_NUMBER | INTEGER+
    ;

list
    : stringList | doubleList | intList
    ;

stringList
    : STRING (',' STRING)*
    ;

doubleList
    : REAL_NUMBER (',' REAL_NUMBER)*
    ;

intList
    : INTEGER (',' INTEGER)*
    ;

orderBy
    : ORDER_BY order (',' order)*
    ;

order
    : IDENTIFIER dir=(ASC | DESC)?
    ;

limit
    : LIMIT INTEGER+
    ;

// Lexer Rules
// ===========
AVG : 'AVG';
MAX : 'MAX';
MIN : 'MIN';
SUM : 'SUM';
COUNT : 'COUNT';
PERCENTILE : 'PCTL';
PARTITION : 'PART';

EQ : 'EQ';
NEQ : 'NEQ';
IN: 'IN';
LT : 'LT';
LE : 'LE';
GE : 'GE';
GT : 'GT';
BETWEEN : 'BETWEEN';
CONTAIN : 'CONTAIN';
EXIST : 'EXIST';
TIMEFRAME : 'TIMEFRAME';
KEY : 'KEY';

FIND : 'FIND';
WHERE : 'WHERE';
ORDER_BY : 'ORDER-BY';
GROUP_BY : 'GROUP-BY';
LIMIT : 'LIMIT';
ASC : 'ASC';
DESC : 'DESC';
NAME : (LETTER | '_') (LETTER | DIGIT | '_' | '.' | '-')*;
PATH : '/' | '/' NAME ('/' NAME)*;

STRING : DQUOTA_STRING | SQUOTA_STRING |  BQUOTA_STRING;
INTEGER : '0'+ | NON_ZERO_DIGIT DIGIT*;
REAL_NUMBER
    : (DIGIT+)? '.' DIGIT+
    | DIGIT+ '.' EXPONENT
    | (DIGIT+)? '.' (DIGIT+ EXPONENT)
    | DIGIT+ EXPONENT
    ;

fragment LETTER : 'A'..'Z' | 'a'..'z';
fragment NON_ZERO_DIGIT  : '1'..'9';
fragment DIGIT  : '0'..'9';
fragment EXPONENT: 'E' [-+]? DIGIT+;
fragment DQUOTA_STRING : '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
fragment SQUOTA_STRING : '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\'';
fragment BQUOTA_STRING : '`' ( '\\'. | '``' | ~('`'|'\\'))* '`';

IDENTIFIER : '$' NAME;
WS : [ \n\t\r] -> skip;