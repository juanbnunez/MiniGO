// Generated from D:/Tec/Compi/MiniGo1/MiniGO/compiler/goParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link goParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface goParserVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by the {@code rootAST}
	 * labeled alternative in {@link goParser#root}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRootAST(goParser.RootASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code topDeclarationListAST}
	 * labeled alternative in {@link goParser#topDeclarationList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTopDeclarationListAST(goParser.TopDeclarationListASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code singleVarVDAST}
	 * labeled alternative in {@link goParser#variableDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSingleVarVDAST(goParser.SingleVarVDASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code innerVarVDAST}
	 * labeled alternative in {@link goParser#variableDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInnerVarVDAST(goParser.InnerVarVDASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code lpVDAST}
	 * labeled alternative in {@link goParser#variableDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLpVDAST(goParser.LpVDASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code innerVarDeclsAST}
	 * labeled alternative in {@link goParser#innerVarDecls}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInnerVarDeclsAST(goParser.InnerVarDeclsASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code idListDTSVDAST}
	 * labeled alternative in {@link goParser#singleVarDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdListDTSVDAST(goParser.IdListDTSVDASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code idListSVDAST}
	 * labeled alternative in {@link goParser#singleVarDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdListSVDAST(goParser.IdListSVDASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code singleVarDAST}
	 * labeled alternative in {@link goParser#singleVarDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSingleVarDAST(goParser.SingleVarDASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code singleVarDeclNoExpsAST}
	 * labeled alternative in {@link goParser#singleVarDeclNoExps}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSingleVarDeclNoExpsAST(goParser.SingleVarDeclNoExpsASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code typeTDAST}
	 * labeled alternative in {@link goParser#typeDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeTDAST(goParser.TypeTDASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code innerTypeDeclsAST}
	 * labeled alternative in {@link goParser#innerTypeDecls}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInnerTypeDeclsAST(goParser.InnerTypeDeclsASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code idSYDAST}
	 * labeled alternative in {@link goParser#singleTypeDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdSYDAST(goParser.IdSYDASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code funcDeclAST}
	 * labeled alternative in {@link goParser#funcDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFuncDeclAST(goParser.FuncDeclASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code funcFFDAST}
	 * labeled alternative in {@link goParser#funcFrontDecl}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFuncFFDAST(goParser.FuncFFDASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code funcArgDeclsAST}
	 * labeled alternative in {@link goParser#funcArgDecls}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFuncArgDeclsAST(goParser.FuncArgDeclsASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code lpDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLpDTAST(goParser.LpDTASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code idDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdDTAST(goParser.IdDTASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code sliceDeclTypeDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSliceDeclTypeDTAST(goParser.SliceDeclTypeDTASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code arrayDeclTypeDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArrayDeclTypeDTAST(goParser.ArrayDeclTypeDTASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code structDeclTypeDTAST}
	 * labeled alternative in {@link goParser#declType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructDeclTypeDTAST(goParser.StructDeclTypeDTASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code lsbSDTAST}
	 * labeled alternative in {@link goParser#sliceDeclType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLsbSDTAST(goParser.LsbSDTASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code lsbADTAST}
	 * labeled alternative in {@link goParser#arrayDeclType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLsbADTAST(goParser.LsbADTASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code structSDTAST}
	 * labeled alternative in {@link goParser#structDeclType}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructSDTAST(goParser.StructSDTASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code structMemDeclsAST}
	 * labeled alternative in {@link goParser#structMemDecls}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStructMemDeclsAST(goParser.StructMemDeclsASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code identifierListAST}
	 * labeled alternative in {@link goParser#identifierList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdentifierListAST(goParser.IdentifierListASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code expressionEAST}
	 * labeled alternative in {@link goParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressionEAST(goParser.ExpressionEASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code primaryExpressionEAST}
	 * labeled alternative in {@link goParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimaryExpressionEAST(goParser.PrimaryExpressionEASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code operatorEAST}
	 * labeled alternative in {@link goParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperatorEAST(goParser.OperatorEASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code expressionListAST}
	 * labeled alternative in {@link goParser#expressionList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressionListAST(goParser.ExpressionListASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code lengthExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLengthExpPEAST(goParser.LengthExpPEASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code appendExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAppendExpPEAST(goParser.AppendExpPEASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code operandPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperandPEAST(goParser.OperandPEASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code primaryExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrimaryExpPEAST(goParser.PrimaryExpPEASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code capExpPEAST}
	 * labeled alternative in {@link goParser#primaryExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCapExpPEAST(goParser.CapExpPEASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code literalOAST}
	 * labeled alternative in {@link goParser#operand}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLiteralOAST(goParser.LiteralOASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code idOAST}
	 * labeled alternative in {@link goParser#operand}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdOAST(goParser.IdOASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code expressionOAST}
	 * labeled alternative in {@link goParser#operand}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressionOAST(goParser.ExpressionOASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code intliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIntliteralAST(goParser.IntliteralASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code floatliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFloatliteralAST(goParser.FloatliteralASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code runliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRunliteralAST(goParser.RunliteralASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code rawsliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRawsliteralAST(goParser.RawsliteralASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code interpretedliteralAST}
	 * labeled alternative in {@link goParser#literal}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitInterpretedliteralAST(goParser.InterpretedliteralASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code idexAST}
	 * labeled alternative in {@link goParser#index}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIdexAST(goParser.IdexASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code argumentsAST}
	 * labeled alternative in {@link goParser#arguments}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArgumentsAST(goParser.ArgumentsASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code selectorAST}
	 * labeled alternative in {@link goParser#selector}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSelectorAST(goParser.SelectorASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code appendExpressionAST}
	 * labeled alternative in {@link goParser#appendExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAppendExpressionAST(goParser.AppendExpressionASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code lengthExpressionAST}
	 * labeled alternative in {@link goParser#lengthExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLengthExpressionAST(goParser.LengthExpressionASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code capExpressionAST}
	 * labeled alternative in {@link goParser#capExpression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCapExpressionAST(goParser.CapExpressionASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code statementListAST}
	 * labeled alternative in {@link goParser#statementList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStatementListAST(goParser.StatementListASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code blockAST}
	 * labeled alternative in {@link goParser#block}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBlockAST(goParser.BlockASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code printSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrintSAST(goParser.PrintSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code printlnSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitPrintlnSAST(goParser.PrintlnSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code returnSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitReturnSAST(goParser.ReturnSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code breakSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBreakSAST(goParser.BreakSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code continueSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitContinueSAST(goParser.ContinueSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code simpleStatementSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSimpleStatementSAST(goParser.SimpleStatementSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code blockSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBlockSAST(goParser.BlockSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code switchSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSwitchSAST(goParser.SwitchSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ifStatementSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIfStatementSAST(goParser.IfStatementSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code loopSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLoopSAST(goParser.LoopSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code typeDeclSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitTypeDeclSAST(goParser.TypeDeclSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code variableDeclSAST}
	 * labeled alternative in {@link goParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitVariableDeclSAST(goParser.VariableDeclSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code epsilonSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEpsilonSSAST(goParser.EpsilonSSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code expressionSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpressionSSAST(goParser.ExpressionSSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code assigmentStatementSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAssigmentStatementSSAST(goParser.AssigmentStatementSSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code expListSSAST}
	 * labeled alternative in {@link goParser#simpleStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpListSSAST(goParser.ExpListSSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code expListASAST}
	 * labeled alternative in {@link goParser#assignmentStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpListASAST(goParser.ExpListASASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code expASAST}
	 * labeled alternative in {@link goParser#assignmentStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpASAST(goParser.ExpASASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code isExpressionBlockISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIsExpressionBlockISAST(goParser.IsExpressionBlockISASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code isExpBlockIsISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIsExpBlockIsISAST(goParser.IsExpBlockIsISASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code isExpBlockISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIsExpBlockISAST(goParser.IsExpBlockISASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code isSSExpBlockISAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIsSSExpBlockISAST(goParser.IsSSExpBlockISASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code isSSExpBlockifSAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIsSSExpBlockifSAST(goParser.IsSSExpBlockifSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code isSSExpBlockBlockAST}
	 * labeled alternative in {@link goParser#ifStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIsSSExpBlockBlockAST(goParser.IsSSExpBlockBlockASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code fBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFBlockLAST(goParser.FBlockLASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code fExpBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFExpBlockLAST(goParser.FExpBlockLASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code fSimpStateExpSBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFSimpStateExpSBlockLAST(goParser.FSimpStateExpSBlockLASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code fSimpStateSimpStateBlockLAST}
	 * labeled alternative in {@link goParser#loop}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFSimpStateSimpStateBlockLAST(goParser.FSimpStateSimpStateBlockLASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code sSimpleStatSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSSimpleStatSAST(goParser.SSimpleStatSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code sExpressionSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSExpressionSAST(goParser.SExpressionSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code sSimpleStatExpCCListSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSSimpleStatExpCCListSAST(goParser.SSimpleStatExpCCListSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code sBlockSAST}
	 * labeled alternative in {@link goParser#switch}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSBlockSAST(goParser.SBlockSASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code epsilonECCLAST}
	 * labeled alternative in {@link goParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEpsilonECCLAST(goParser.EpsilonECCLASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code expCaseClauseECCLAST}
	 * labeled alternative in {@link goParser#expressionCaseClauseList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpCaseClauseECCLAST(goParser.ExpCaseClauseECCLASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code expSwitchCaseECCAST}
	 * labeled alternative in {@link goParser#expressionCaseClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpSwitchCaseECCAST(goParser.ExpSwitchCaseECCASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code caseESCAST}
	 * labeled alternative in {@link goParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCaseESCAST(goParser.CaseESCASTContext ctx);
	/**
	 * Visit a parse tree produced by the {@code defaultESCAST}
	 * labeled alternative in {@link goParser#expressionSwitchCase}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitDefaultESCAST(goParser.DefaultESCASTContext ctx);
	/**
	 * Visit a parse tree produced by {@link goParser#operator}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOperator(goParser.OperatorContext ctx);
	/**
	 * Visit a parse tree produced by the {@code epsilonAST}
	 * labeled alternative in {@link goParser#epsilon}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEpsilonAST(goParser.EpsilonASTContext ctx);
}