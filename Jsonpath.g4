grammar Jsonpath;

//
// Derived from https://github.com/stevenalexander/antlr4-jsonpath-grammar
//

jsonpath: '$' path EOF
        ;

path : nodeExpr+
     ;

nodeExpr : dotExpr
         | selectExpr
         ;

dotExpr : DOTS 
        | DOTS (Identifier | STAR | Length )
        ;

selectExpr : '[' (starExpr | rangeExpr | itemsExpr | namesExpr | filterExpr | scriptExpr)? ']'
           ;

rangeExpr : (startIndex=INT)? ':' (endIndex=INT (':' step=INT)?)?
          ;

itemsExpr : INT (',' INT)*
          ;

namesExpr: QUOTED (',' QUOTED)*
        ;

starExpr : STAR
         ;

filterExpr : '?(' queryExpr ')'
           ;

scriptExpr : '(' valueExpr ')'
           ;

queryExpr : '@.' exists=Identifier
          | '@.' name=Identifier op=COMP value=(INT | NUMBER | QUOTED | 'true' | 'false' | 'null')
          | '@.' name=Identifier op='=~' value=REGEX
          ;

valueExpr : '@.' name=Identifier op=OP value=(INT | NUMBER )
          | '@.' name=Length op=OP value=(INT | NUMBER )
          ;

OP : '+' | '-' | '/' | '*'
   ;

COMP : '>' | '>=' | '<' | '<=' | '==' | '!='
     ;

Identifier : [a-zA-Z][a-zA-Z0-9_]*
           ;

Length : 'length()'
       ;

INT : '-'? ('0'..'9')+
    ;

NUMBER : INT ('.' ('0'..'9')+)?
       ;

QUOTED : '\'' (~'\'')* '\''
       ;

REGEX : '/' (~'/')* '/' ('i'|'u'|'m'|'S')*
      ;

DOTS : '.'
     | '..'
     ;

STAR : '*'
     ;

WS : [ \t\n\r]+ -> skip
   ;
