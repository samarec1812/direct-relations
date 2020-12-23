grammar Relation;

SPACE: '\n' |  ' ' | '\t';

ID_COMBINER
        : '---'
        | '-->'
        | '>--'
        | '<--'
        | '--<'
        | '><'
        | '<>'
        | '-/-'
        ;

OPERATOR_PARALLEL : '#';
OPERATOR_SUM : '+' ;
OPERATOR_POINT : '.';
L_PAR : '(';
R_PAR : ')';


operator
    :  OPERATOR_PARALLEL
    | OPERATOR_SUM
    | OPERATOR_POINT
    | SPACE
    ;

combiner
    :  ID_COMBINER
    |  combiner (operator combiner)+
    |  L_PAR combiner R_PAR
    ;



program
       : combiner
       ;

