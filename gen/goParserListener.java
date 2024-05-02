// Generated from D:/Tec/Compi/MiniGo1/MiniGO/compiler/goParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link goParser}.
 */
public interface goParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by the {@code rootAST}
	 * labeled alternative in {@link goParser#root}.
	 * @param ctx the parse tree
	 */
	void enterRootAST(goParser.RootASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code rootAST}
	 * labeled alternative in {@link goParser#root}.
	 * @param ctx the parse tree
	 */
	void exitRootAST(goParser.RootASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code topDeclarationListAST}
	 * labeled alternative in {@link goParser#topDeclarationList}.
	 * @param ctx the parse tree
	 */
	void enterTopDeclarationListAST(goParser.TopDeclarationListASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code topDeclarationListAST}
	 * labeled alternative in {@link goParser#topDeclarationList}.
	 * @param ctx the parse tree
	 */
	void exitTopDeclarationListAST(goParser.TopDeclarationListASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code singleVarVDAST}
	 * labeled alternative in {@link goParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void enterSingleVarVDAST(goParser.SingleVarVDASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code singleVarVDAST}
	 * labeled alternative in {@link goParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void exitSingleVarVDAST(goParser.SingleVarVDASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code innerVarVDAST}
	 * labeled alternative in {@link goParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void enterInnerVarVDAST(goParser.InnerVarVDASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code innerVarVDAST}
	 * labeled alternative in {@link goParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void exitInnerVarVDAST(goParser.InnerVarVDASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code lpVDAST}
	 * labeled alternative in {@link goParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void enterLpVDAST(goParser.LpVDASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code lpVDAST}
	 * labeled alternative in {@link goParser#variableDecl}.
	 * @param ctx the parse tree
	 */
	void exitLpVDAST(goParser.LpVDASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code innerVarDeclsAST}
	 * labeled alternative in {@link goParser#innerVarDecls}.
	 * @param ctx the parse tree
	 */
	void enterInnerVarDeclsAST(goParser.InnerVarDeclsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code innerVarDeclsAST}
	 * labeled alternative in {@link goParser#innerVarDecls}.
	 * @param ctx the parse tree
	 */
	void exitInnerVarDeclsAST(goParser.InnerVarDeclsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code idListDTSVDAST}
	 * labeled alternative in {@link goParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void enterIdListDTSVDAST(goParser.IdListDTSVDASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code idListDTSVDAST}
	 * labeled alternative in {@link goParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void exitIdListDTSVDAST(goParser.IdListDTSVDASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code idListSVDAST}
	 * labeled alternative in {@link goParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void enterIdListSVDAST(goParser.IdListSVDASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code idListSVDAST}
	 * labeled alternative in {@link goParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void exitIdListSVDAST(goParser.IdListSVDASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code singleVarDAST}
	 * labeled alternative in {@link goParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void enterSingleVarDAST(goParser.SingleVarDASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code singleVarDAST}
	 * labeled alternative in {@link goParser#singleVarDecl}.
	 * @param ctx the parse tree
	 */
	void exitSingleVarDAST(goParser.SingleVarDASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code singleVarDeclNoExpsAST}
	 * labeled alternative in {@link goParser#singleVarDeclNoExps}.
	 * @param ctx the parse tree
	 */
	void enterSingleVarDeclNoExpsAST(goParser.SingleVarDeclNoExpsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code singleVarDeclNoExpsAST}
	 * labeled alternative in {@link goParser#singleVarDeclNoExps}.
	 * @param ctx the parse tree
	 */
	void exitSingleVarDeclNoExpsAST(goParser.SingleVarDeclNoExpsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code typeTDAST}
	 * labeled alternative in {@link goParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void enterTypeTDAST(goParser.TypeTDASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code typeTDAST}
	 * labeled alternative in {@link goParser#typeDecl}.
	 * @param ctx the parse tree
	 */
	void exitTypeTDAST(goParser.TypeTDASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code innerTypeDeclsAST}
	 * labeled alternative in {@link goParser#innerTypeDecls}.
	 * @param ctx the parse tree
	 */
	void enterInnerTypeDeclsAST(goParser.InnerTypeDeclsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code innerTypeDeclsAST}
	 * labeled alternative in {@link goParser#innerTypeDecls}.
	 * @param ctx the parse tree
	 */
	void exitInnerTypeDeclsAST(goParser.InnerTypeDeclsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code idSYDAST}
	 * labeled alternative in {@link goParser#singleTypeDecl}.
	 * @param ctx the parse tree
	 */
	void enterIdSYDAST(goParser.IdSYDASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code idSYDAST}
	 * labeled alternative in {@link goParser#singleTypeDecl}.
	 * @param ctx the parse tree
	 */
	void exitIdSYDAST(goParser.IdSYDASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code funcDeclAST}
	 * labeled alternative in {@link goParser#funcDecl}.
	 * @param ctx the parse tree
	 */
	void enterFuncDeclAST(goParser.FuncDeclASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code funcDeclAST}
	 * labeled alternative in {@link goParser#funcDecl}.
	 * @param ctx the parse tree
	 */
	void exitFuncDeclAST(goParser.FuncDeclASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code funcFFDAST}
	 * labeled alternative in {@link goParser#funcFrontDecl}.
	 * @param ctx the parse tree
	 */
	void enterFuncFFDAST(goParser.FuncFFDASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code funcFFDAST}
	 * labeled alternative in {@link goParser#funcFrontDecl}.
	 * @param ctx the parse tree
	 */
	void exitFuncFFDAST(goParser.FuncFFDASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code funcArgDeclsAST}
	 * labeled alternative in {@link goParser#funcArgDecls}.
	 * @param ctx the parse tree
	 */
	void enterFuncArgDeclsAST(goParser.FuncArgDeclsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code funcArgDeclsAST}
	 * labeled alternative in {@link goParser#funcArgDecls}.
	 * @param ctx the parse tree
	 */
	void exitFuncArgDeclsAST(goParser.FuncArgDeclsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code lpDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterLpDTAST(goParser.LpDTASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code lpDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitLpDTAST(goParser.LpDTASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code idDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterIdDTAST(goParser.IdDTASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code idDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitIdDTAST(goParser.IdDTASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sliceDeclTypeDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterSliceDeclTypeDTAST(goParser.SliceDeclTypeDTASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sliceDeclTypeDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitSliceDeclTypeDTAST(goParser.SliceDeclTypeDTASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code arrayDeclTypeDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterArrayDeclTypeDTAST(goParser.ArrayDeclTypeDTASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code arrayDeclTypeDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitArrayDeclTypeDTAST(goParser.ArrayDeclTypeDTASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structDeclTypeDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void enterStructDeclTypeDTAST(goParser.StructDeclTypeDTASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structDeclTypeDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 */
	void exitStructDeclTypeDTAST(goParser.StructDeclTypeDTASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code lsbSDTAST}
	 * labeled alternative in {@link goParser#sliceDeclType}.
	 * @param ctx the parse tree
	 */
	void enterLsbSDTAST(goParser.LsbSDTASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code lsbSDTAST}
	 * labeled alternative in {@link goParser#sliceDeclType}.
	 * @param ctx the parse tree
	 */
	void exitLsbSDTAST(goParser.LsbSDTASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code lsbADTAST}
	 * labeled alternative in {@link goParser#arrayDeclType}.
	 * @param ctx the parse tree
	 */
	void enterLsbADTAST(goParser.LsbADTASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code lsbADTAST}
	 * labeled alternative in {@link goParser#arrayDeclType}.
	 * @param ctx the parse tree
	 */
	void exitLsbADTAST(goParser.LsbADTASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structSDTAST}
	 * labeled alternative in {@link goParser#structDeclType}.
	 * @param ctx the parse tree
	 */
	void enterStructSDTAST(goParser.StructSDTASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structSDTAST}
	 * labeled alternative in {@link goParser#structDeclType}.
	 * @param ctx the parse tree
	 */
	void exitStructSDTAST(goParser.StructSDTASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structMemDeclsAST}
	 * labeled alternative in {@link goParser#structMemDecls}.
	 * @param ctx the parse tree
	 */
	void enterStructMemDeclsAST(goParser.StructMemDeclsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structMemDeclsAST}
	 * labeled alternative in {@link goParser#structMemDecls}.
	 * @param ctx the parse tree
	 */
	void exitStructMemDeclsAST(goParser.StructMemDeclsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code identifierListAST}
	 * labeled alternative in {@link goParser#identifierList}.
	 * @param ctx the parse tree
	 */
	void enterIdentifierListAST(goParser.IdentifierListASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code identifierListAST}
	 * labeled alternative in {@link goParser#identifierList}.
	 * @param ctx the parse tree
	 */
	void exitIdentifierListAST(goParser.IdentifierListASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionEAST}
	 * labeled alternative in {@link goParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpressionEAST(goParser.ExpressionEASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionEAST}
	 * labeled alternative in {@link goParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpressionEAST(goParser.ExpressionEASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primaryExpressionEAST}
	 * labeled alternative in {@link goParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpressionEAST(goParser.PrimaryExpressionEASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primaryExpressionEAST}
	 * labeled alternative in {@link goParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpressionEAST(goParser.PrimaryExpressionEASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code operatorEAST}
	 * labeled alternative in {@link goParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterOperatorEAST(goParser.OperatorEASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code operatorEAST}
	 * labeled alternative in {@link goParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitOperatorEAST(goParser.OperatorEASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionListAST}
	 * labeled alternative in {@link goParser#expressionList}.
	 * @param ctx the parse tree
	 */
	void enterExpressionListAST(goParser.ExpressionListASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionListAST}
	 * labeled alternative in {@link goParser#expressionList}.
	 * @param ctx the parse tree
	 */
	void exitExpressionListAST(goParser.ExpressionListASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code lengthExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterLengthExpPEAST(goParser.LengthExpPEASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code lengthExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitLengthExpPEAST(goParser.LengthExpPEASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code appendExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterAppendExpPEAST(goParser.AppendExpPEASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code appendExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitAppendExpPEAST(goParser.AppendExpPEASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code operandPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterOperandPEAST(goParser.OperandPEASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code operandPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitOperandPEAST(goParser.OperandPEASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code primaryExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpPEAST(goParser.PrimaryExpPEASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code primaryExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpPEAST(goParser.PrimaryExpPEASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code capExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterCapExpPEAST(goParser.CapExpPEASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code capExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitCapExpPEAST(goParser.CapExpPEASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code literalOAST}
	 * labeled alternative in {@link goParser#operand}.
	 * @param ctx the parse tree
	 */
	void enterLiteralOAST(goParser.LiteralOASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code literalOAST}
	 * labeled alternative in {@link goParser#operand}.
	 * @param ctx the parse tree
	 */
	void exitLiteralOAST(goParser.LiteralOASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code idOAST}
	 * labeled alternative in {@link goParser#operand}.
	 * @param ctx the parse tree
	 */
	void enterIdOAST(goParser.IdOASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code idOAST}
	 * labeled alternative in {@link goParser#operand}.
	 * @param ctx the parse tree
	 */
	void exitIdOAST(goParser.IdOASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionOAST}
	 * labeled alternative in {@link goParser#operand}.
	 * @param ctx the parse tree
	 */
	void enterExpressionOAST(goParser.ExpressionOASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionOAST}
	 * labeled alternative in {@link goParser#operand}.
	 * @param ctx the parse tree
	 */
	void exitExpressionOAST(goParser.ExpressionOASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code intliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterIntliteralAST(goParser.IntliteralASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code intliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitIntliteralAST(goParser.IntliteralASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code floatliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterFloatliteralAST(goParser.FloatliteralASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code floatliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitFloatliteralAST(goParser.FloatliteralASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code runliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterRunliteralAST(goParser.RunliteralASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code runliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitRunliteralAST(goParser.RunliteralASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code rawsliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterRawsliteralAST(goParser.RawsliteralASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code rawsliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitRawsliteralAST(goParser.RawsliteralASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code interpretedliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterInterpretedliteralAST(goParser.InterpretedliteralASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code interpretedliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitInterpretedliteralAST(goParser.InterpretedliteralASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code idexAST}
	 * labeled alternative in {@link goParser#index}.
	 * @param ctx the parse tree
	 */
	void enterIdexAST(goParser.IdexASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code idexAST}
	 * labeled alternative in {@link goParser#index}.
	 * @param ctx the parse tree
	 */
	void exitIdexAST(goParser.IdexASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code argumentsAST}
	 * labeled alternative in {@link goParser#arguments}.
	 * @param ctx the parse tree
	 */
	void enterArgumentsAST(goParser.ArgumentsASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code argumentsAST}
	 * labeled alternative in {@link goParser#arguments}.
	 * @param ctx the parse tree
	 */
	void exitArgumentsAST(goParser.ArgumentsASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code selectorAST}
	 * labeled alternative in {@link goParser#selector}.
	 * @param ctx the parse tree
	 */
	void enterSelectorAST(goParser.SelectorASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code selectorAST}
	 * labeled alternative in {@link goParser#selector}.
	 * @param ctx the parse tree
	 */
	void exitSelectorAST(goParser.SelectorASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code appendExpressionAST}
	 * labeled alternative in {@link goParser#appendExpression}.
	 * @param ctx the parse tree
	 */
	void enterAppendExpressionAST(goParser.AppendExpressionASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code appendExpressionAST}
	 * labeled alternative in {@link goParser#appendExpression}.
	 * @param ctx the parse tree
	 */
	void exitAppendExpressionAST(goParser.AppendExpressionASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code lengthExpressionAST}
	 * labeled alternative in {@link goParser#lengthExpression}.
	 * @param ctx the parse tree
	 */
	void enterLengthExpressionAST(goParser.LengthExpressionASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code lengthExpressionAST}
	 * labeled alternative in {@link goParser#lengthExpression}.
	 * @param ctx the parse tree
	 */
	void exitLengthExpressionAST(goParser.LengthExpressionASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code capExpressionAST}
	 * labeled alternative in {@link goParser#capExpression}.
	 * @param ctx the parse tree
	 */
	void enterCapExpressionAST(goParser.CapExpressionASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code capExpressionAST}
	 * labeled alternative in {@link goParser#capExpression}.
	 * @param ctx the parse tree
	 */
	void exitCapExpressionAST(goParser.CapExpressionASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code statementListAST}
	 * labeled alternative in {@link goParser#statementList}.
	 * @param ctx the parse tree
	 */
	void enterStatementListAST(goParser.StatementListASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code statementListAST}
	 * labeled alternative in {@link goParser#statementList}.
	 * @param ctx the parse tree
	 */
	void exitStatementListAST(goParser.StatementListASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code blockAST}
	 * labeled alternative in {@link goParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlockAST(goParser.BlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code blockAST}
	 * labeled alternative in {@link goParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlockAST(goParser.BlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code printSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterPrintSAST(goParser.PrintSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code printSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitPrintSAST(goParser.PrintSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code printlnSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterPrintlnSAST(goParser.PrintlnSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code printlnSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitPrintlnSAST(goParser.PrintlnSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code returnSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterReturnSAST(goParser.ReturnSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code returnSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitReturnSAST(goParser.ReturnSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code breakSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterBreakSAST(goParser.BreakSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code breakSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitBreakSAST(goParser.BreakSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code continueSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterContinueSAST(goParser.ContinueSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code continueSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitContinueSAST(goParser.ContinueSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code simpleStatementSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterSimpleStatementSAST(goParser.SimpleStatementSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code simpleStatementSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitSimpleStatementSAST(goParser.SimpleStatementSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code blockSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterBlockSAST(goParser.BlockSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code blockSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitBlockSAST(goParser.BlockSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code switchSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterSwitchSAST(goParser.SwitchSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code switchSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitSwitchSAST(goParser.SwitchSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ifStatementSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatementSAST(goParser.IfStatementSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ifStatementSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatementSAST(goParser.IfStatementSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code loopSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterLoopSAST(goParser.LoopSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code loopSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitLoopSAST(goParser.LoopSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code typeDeclSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterTypeDeclSAST(goParser.TypeDeclSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code typeDeclSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitTypeDeclSAST(goParser.TypeDeclSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code variableDeclSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclSAST(goParser.VariableDeclSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code variableDeclSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclSAST(goParser.VariableDeclSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code epsilonSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void enterEpsilonSSAST(goParser.EpsilonSSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code epsilonSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void exitEpsilonSSAST(goParser.EpsilonSSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expressionSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void enterExpressionSSAST(goParser.ExpressionSSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expressionSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void exitExpressionSSAST(goParser.ExpressionSSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code assigmentStatementSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssigmentStatementSSAST(goParser.AssigmentStatementSSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code assigmentStatementSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssigmentStatementSSAST(goParser.AssigmentStatementSSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expListSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void enterExpListSSAST(goParser.ExpListSSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expListSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 */
	void exitExpListSSAST(goParser.ExpListSSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expListASAST}
	 * labeled alternative in {@link goParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterExpListASAST(goParser.ExpListASASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expListASAST}
	 * labeled alternative in {@link goParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitExpListASAST(goParser.ExpListASASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expASAST}
	 * labeled alternative in {@link goParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void enterExpASAST(goParser.ExpASASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expASAST}
	 * labeled alternative in {@link goParser#assignmentStatement}.
	 * @param ctx the parse tree
	 */
	void exitExpASAST(goParser.ExpASASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code isExpressionBlockISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIsExpressionBlockISAST(goParser.IsExpressionBlockISASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code isExpressionBlockISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIsExpressionBlockISAST(goParser.IsExpressionBlockISASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code isExpBlockIsISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIsExpBlockIsISAST(goParser.IsExpBlockIsISASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code isExpBlockIsISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIsExpBlockIsISAST(goParser.IsExpBlockIsISASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code isExpBlockISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIsExpBlockISAST(goParser.IsExpBlockISASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code isExpBlockISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIsExpBlockISAST(goParser.IsExpBlockISASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code isSSExpBlockISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIsSSExpBlockISAST(goParser.IsSSExpBlockISASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code isSSExpBlockISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIsSSExpBlockISAST(goParser.IsSSExpBlockISASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code isSSExpBlockifSAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIsSSExpBlockifSAST(goParser.IsSSExpBlockifSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code isSSExpBlockifSAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIsSSExpBlockifSAST(goParser.IsSSExpBlockifSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code isSSExpBlockBlockAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIsSSExpBlockBlockAST(goParser.IsSSExpBlockBlockASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code isSSExpBlockBlockAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIsSSExpBlockBlockAST(goParser.IsSSExpBlockBlockASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code fBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterFBlockLAST(goParser.FBlockLASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitFBlockLAST(goParser.FBlockLASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code fExpBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterFExpBlockLAST(goParser.FExpBlockLASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fExpBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitFExpBlockLAST(goParser.FExpBlockLASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code fSimpStateExpSBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterFSimpStateExpSBlockLAST(goParser.FSimpStateExpSBlockLASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fSimpStateExpSBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitFSimpStateExpSBlockLAST(goParser.FSimpStateExpSBlockLASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code fSimpStateSimpStateBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterFSimpStateSimpStateBlockLAST(goParser.FSimpStateSimpStateBlockLASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fSimpStateSimpStateBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitFSimpStateSimpStateBlockLAST(goParser.FSimpStateSimpStateBlockLASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sSimpleStatSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 */
	void enterSSimpleStatSAST(goParser.SSimpleStatSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sSimpleStatSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 */
	void exitSSimpleStatSAST(goParser.SSimpleStatSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sExpressionSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 */
	void enterSExpressionSAST(goParser.SExpressionSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sExpressionSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 */
	void exitSExpressionSAST(goParser.SExpressionSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sSimpleStatExpCCListSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 */
	void enterSSimpleStatExpCCListSAST(goParser.SSimpleStatExpCCListSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sSimpleStatExpCCListSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 */
	void exitSSimpleStatExpCCListSAST(goParser.SSimpleStatExpCCListSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sBlockSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 */
	void enterSBlockSAST(goParser.SBlockSASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sBlockSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 */
	void exitSBlockSAST(goParser.SBlockSASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code epsilonECCLAST}
	 * labeled alternative in {@link goParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 */
	void enterEpsilonECCLAST(goParser.EpsilonECCLASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code epsilonECCLAST}
	 * labeled alternative in {@link goParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 */
	void exitEpsilonECCLAST(goParser.EpsilonECCLASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expCaseClauseECCLAST}
	 * labeled alternative in {@link goParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 */
	void enterExpCaseClauseECCLAST(goParser.ExpCaseClauseECCLASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expCaseClauseECCLAST}
	 * labeled alternative in {@link goParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 */
	void exitExpCaseClauseECCLAST(goParser.ExpCaseClauseECCLASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expSwitchCaseECCAST}
	 * labeled alternative in {@link goParser#expressionCaseClause}.
	 * @param ctx the parse tree
	 */
	void enterExpSwitchCaseECCAST(goParser.ExpSwitchCaseECCASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expSwitchCaseECCAST}
	 * labeled alternative in {@link goParser#expressionCaseClause}.
	 * @param ctx the parse tree
	 */
	void exitExpSwitchCaseECCAST(goParser.ExpSwitchCaseECCASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code caseESCAST}
	 * labeled alternative in {@link goParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 */
	void enterCaseESCAST(goParser.CaseESCASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code caseESCAST}
	 * labeled alternative in {@link goParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 */
	void exitCaseESCAST(goParser.CaseESCASTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code defaultESCAST}
	 * labeled alternative in {@link goParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 */
	void enterDefaultESCAST(goParser.DefaultESCASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code defaultESCAST}
	 * labeled alternative in {@link goParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 */
	void exitDefaultESCAST(goParser.DefaultESCASTContext ctx);
	/**
	 * Enter a parse tree produced by {@link goParser#operator}.
	 * @param ctx the parse tree
	 */
	void enterOperator(goParser.OperatorContext ctx);
	/**
	 * Exit a parse tree produced by {@link goParser#operator}.
	 * @param ctx the parse tree
	 */
	void exitOperator(goParser.OperatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code epsilonAST}
	 * labeled alternative in {@link goParser#epsilon}.
	 * @param ctx the parse tree
	 */
	void enterEpsilonAST(goParser.EpsilonASTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code epsilonAST}
	 * labeled alternative in {@link goParser#epsilon}.
	 * @param ctx the parse tree
	 */
	void exitEpsilonAST(goParser.EpsilonASTContext ctx);
}