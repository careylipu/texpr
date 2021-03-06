// Code generated from TExprParser.g4 by ANTLR 4.8. DO NOT EDIT.

package ast // TExprParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TExprParser.
type TExprParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TExprParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by TExprParser#binaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#matchExpression.
	VisitMatchExpression(ctx *MatchExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#inExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#isTypeExpression.
	VisitIsTypeExpression(ctx *IsTypeExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#calcExpression.
	VisitCalcExpression(ctx *CalcExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#isNotTypeExpression.
	VisitIsNotTypeExpression(ctx *IsNotTypeExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#notInExpression.
	VisitNotInExpression(ctx *NotInExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#comparatorExpression.
	VisitComparatorExpression(ctx *ComparatorExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#variableExpression.
	VisitVariableExpression(ctx *VariableExpressionContext) interface{}

	// Visit a parse tree produced by TExprParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by TExprParser#comparator.
	VisitComparator(ctx *ComparatorContext) interface{}

	// Visit a parse tree produced by TExprParser#binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by TExprParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by TExprParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by TExprParser#container.
	VisitContainer(ctx *ContainerContext) interface{}

	// Visit a parse tree produced by TExprParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by TExprParser#calc.
	VisitCalc(ctx *CalcContext) interface{}

	// Visit a parse tree produced by TExprParser#bit.
	VisitBit(ctx *BitContext) interface{}

	// Visit a parse tree produced by TExprParser#shift.
	VisitShift(ctx *ShiftContext) interface{}

	// Visit a parse tree produced by TExprParser#plus.
	VisitPlus(ctx *PlusContext) interface{}

	// Visit a parse tree produced by TExprParser#multiplying.
	VisitMultiplying(ctx *MultiplyingContext) interface{}

	// Visit a parse tree produced by TExprParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by TExprParser#scientific.
	VisitScientific(ctx *ScientificContext) interface{}

	// Visit a parse tree produced by TExprParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by TExprParser#funcname.
	VisitFuncname(ctx *FuncnameContext) interface{}

	// Visit a parse tree produced by TExprParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by TExprParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by TExprParser#regex.
	VisitRegex(ctx *RegexContext) interface{}

	// Visit a parse tree produced by TExprParser#kind.
	VisitKind(ctx *KindContext) interface{}

	// Visit a parse tree produced by TExprParser#strings.
	VisitStrings(ctx *StringsContext) interface{}

	// Visit a parse tree produced by TExprParser#integers.
	VisitIntegers(ctx *IntegersContext) interface{}

	// Visit a parse tree produced by TExprParser#floats.
	VisitFloats(ctx *FloatsContext) interface{}

	// Visit a parse tree produced by TExprParser#booleans.
	VisitBooleans(ctx *BooleansContext) interface{}
}
