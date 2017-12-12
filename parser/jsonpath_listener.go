// Generated from Jsonpath.g4 by ANTLR 4.7.

package parser // Jsonpath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// JsonpathListener is a complete listener for a parse tree produced by JsonpathParser.
type JsonpathListener interface {
	antlr.ParseTreeListener

	// EnterJsonpath is called when entering the jsonpath production.
	EnterJsonpath(c *JsonpathContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterNodeExpr is called when entering the nodeExpr production.
	EnterNodeExpr(c *NodeExprContext)

	// EnterDotExpr is called when entering the dotExpr production.
	EnterDotExpr(c *DotExprContext)

	// EnterSelectExpr is called when entering the selectExpr production.
	EnterSelectExpr(c *SelectExprContext)

	// EnterRangeExpr is called when entering the rangeExpr production.
	EnterRangeExpr(c *RangeExprContext)

	// EnterItemsExpr is called when entering the itemsExpr production.
	EnterItemsExpr(c *ItemsExprContext)

	// EnterNamesExpr is called when entering the namesExpr production.
	EnterNamesExpr(c *NamesExprContext)

	// EnterStarExpr is called when entering the starExpr production.
	EnterStarExpr(c *StarExprContext)

	// EnterFilterExpr is called when entering the filterExpr production.
	EnterFilterExpr(c *FilterExprContext)

	// EnterScriptExpr is called when entering the scriptExpr production.
	EnterScriptExpr(c *ScriptExprContext)

	// EnterQueryExpr is called when entering the queryExpr production.
	EnterQueryExpr(c *QueryExprContext)

	// EnterValueExpr is called when entering the valueExpr production.
	EnterValueExpr(c *ValueExprContext)

	// ExitJsonpath is called when exiting the jsonpath production.
	ExitJsonpath(c *JsonpathContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitNodeExpr is called when exiting the nodeExpr production.
	ExitNodeExpr(c *NodeExprContext)

	// ExitDotExpr is called when exiting the dotExpr production.
	ExitDotExpr(c *DotExprContext)

	// ExitSelectExpr is called when exiting the selectExpr production.
	ExitSelectExpr(c *SelectExprContext)

	// ExitRangeExpr is called when exiting the rangeExpr production.
	ExitRangeExpr(c *RangeExprContext)

	// ExitItemsExpr is called when exiting the itemsExpr production.
	ExitItemsExpr(c *ItemsExprContext)

	// ExitNamesExpr is called when exiting the namesExpr production.
	ExitNamesExpr(c *NamesExprContext)

	// ExitStarExpr is called when exiting the starExpr production.
	ExitStarExpr(c *StarExprContext)

	// ExitFilterExpr is called when exiting the filterExpr production.
	ExitFilterExpr(c *FilterExprContext)

	// ExitScriptExpr is called when exiting the scriptExpr production.
	ExitScriptExpr(c *ScriptExprContext)

	// ExitQueryExpr is called when exiting the queryExpr production.
	ExitQueryExpr(c *QueryExprContext)

	// ExitValueExpr is called when exiting the valueExpr production.
	ExitValueExpr(c *ValueExprContext)
}
