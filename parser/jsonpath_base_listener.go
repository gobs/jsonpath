// Generated from Jsonpath.g4 by ANTLR 4.7.

package parser // Jsonpath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJsonpathListener is a complete listener for a parse tree produced by JsonpathParser.
type BaseJsonpathListener struct{}

var _ JsonpathListener = &BaseJsonpathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonpathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonpathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonpathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonpathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJsonpath is called when production jsonpath is entered.
func (s *BaseJsonpathListener) EnterJsonpath(ctx *JsonpathContext) {}

// ExitJsonpath is called when production jsonpath is exited.
func (s *BaseJsonpathListener) ExitJsonpath(ctx *JsonpathContext) {}

// EnterPath is called when production path is entered.
func (s *BaseJsonpathListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseJsonpathListener) ExitPath(ctx *PathContext) {}

// EnterNodeExpr is called when production nodeExpr is entered.
func (s *BaseJsonpathListener) EnterNodeExpr(ctx *NodeExprContext) {}

// ExitNodeExpr is called when production nodeExpr is exited.
func (s *BaseJsonpathListener) ExitNodeExpr(ctx *NodeExprContext) {}

// EnterDotExpr is called when production dotExpr is entered.
func (s *BaseJsonpathListener) EnterDotExpr(ctx *DotExprContext) {}

// ExitDotExpr is called when production dotExpr is exited.
func (s *BaseJsonpathListener) ExitDotExpr(ctx *DotExprContext) {}

// EnterSelectExpr is called when production selectExpr is entered.
func (s *BaseJsonpathListener) EnterSelectExpr(ctx *SelectExprContext) {}

// ExitSelectExpr is called when production selectExpr is exited.
func (s *BaseJsonpathListener) ExitSelectExpr(ctx *SelectExprContext) {}

// EnterRangeExpr is called when production rangeExpr is entered.
func (s *BaseJsonpathListener) EnterRangeExpr(ctx *RangeExprContext) {}

// ExitRangeExpr is called when production rangeExpr is exited.
func (s *BaseJsonpathListener) ExitRangeExpr(ctx *RangeExprContext) {}

// EnterItemsExpr is called when production itemsExpr is entered.
func (s *BaseJsonpathListener) EnterItemsExpr(ctx *ItemsExprContext) {}

// ExitItemsExpr is called when production itemsExpr is exited.
func (s *BaseJsonpathListener) ExitItemsExpr(ctx *ItemsExprContext) {}

// EnterNamesExpr is called when production namesExpr is entered.
func (s *BaseJsonpathListener) EnterNamesExpr(ctx *NamesExprContext) {}

// ExitNamesExpr is called when production namesExpr is exited.
func (s *BaseJsonpathListener) ExitNamesExpr(ctx *NamesExprContext) {}

// EnterStarExpr is called when production starExpr is entered.
func (s *BaseJsonpathListener) EnterStarExpr(ctx *StarExprContext) {}

// ExitStarExpr is called when production starExpr is exited.
func (s *BaseJsonpathListener) ExitStarExpr(ctx *StarExprContext) {}

// EnterFilterExpr is called when production filterExpr is entered.
func (s *BaseJsonpathListener) EnterFilterExpr(ctx *FilterExprContext) {}

// ExitFilterExpr is called when production filterExpr is exited.
func (s *BaseJsonpathListener) ExitFilterExpr(ctx *FilterExprContext) {}

// EnterScriptExpr is called when production scriptExpr is entered.
func (s *BaseJsonpathListener) EnterScriptExpr(ctx *ScriptExprContext) {}

// ExitScriptExpr is called when production scriptExpr is exited.
func (s *BaseJsonpathListener) ExitScriptExpr(ctx *ScriptExprContext) {}

// EnterQueryExpr is called when production queryExpr is entered.
func (s *BaseJsonpathListener) EnterQueryExpr(ctx *QueryExprContext) {}

// ExitQueryExpr is called when production queryExpr is exited.
func (s *BaseJsonpathListener) ExitQueryExpr(ctx *QueryExprContext) {}

// EnterValueExpr is called when production valueExpr is entered.
func (s *BaseJsonpathListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production valueExpr is exited.
func (s *BaseJsonpathListener) ExitValueExpr(ctx *ValueExprContext) {}
