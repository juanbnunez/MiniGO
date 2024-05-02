// Generated from D:/Tec/Compi/MiniGo1/MiniGO/compiler/goParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class goParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		SEMI=1, SVD=2, LP=3, RP=4, TIL=5, COL=6, PL=7, MIN=8, MUL=9, DIV=10, CM=11, 
		LSB=12, RSB=13, LCB=14, RCB=15, INC=16, DEC=17, MOD=18, LSH=19, RSH=20, 
		AMPER=21, BC=22, VB=23, CARET=24, EQ=25, NEQ=26, LT=27, GT=28, LTE=29, 
		GTE=30, LAND=31, LOR=32, NEG=33, ASOP=34, AAOP=35, BAOP=36, SAOP=37, BOAOP=38, 
		MAOP=39, BXOOP=40, LSAOP=41, RSAOP=42, BCAOP=43, RAOP=44, DAOP=45, DOT=46, 
		DBS=47, BS=48, PACKAGE=49, IF=50, WHILE=51, LET=52, THEN=53, ELSE=54, 
		IN=55, DO=56, BEGIN=57, END=58, CONST=59, VAR=60, TYPE=61, FUNC=62, STRUCT=63, 
		APPEND=64, LEN=65, CAP=66, PRINT=67, PRINTLN=68, RETURN=69, BREAK=70, 
		CONTINUE=71, FOR=72, SWITCH=73, CASE=74, DEFAULT=75, ID=76, INTLITERAL=77, 
		FLOATLITERAL=78, RUNELITERAL=79, RAWSTRINGLITERAL=80, INTERPRETEDSTRINGLITERAL=81, 
		LETTER=82, DIGIT=83, LINE_COMMENT=84, SPACES=85;
	public static final int
		RULE_root = 0, RULE_topDeclarationList = 1, RULE_variableDecl = 2, RULE_innerVarDecls = 3, 
		RULE_singleVarDecl = 4, RULE_singleVarDeclNoExps = 5, RULE_typeDecl = 6, 
		RULE_innerTypeDecls = 7, RULE_singleTypeDecl = 8, RULE_funcDecl = 9, RULE_funcFrontDecl = 10, 
		RULE_funcArgDecls = 11, RULE_declType = 12, RULE_sliceDeclType = 13, RULE_arrayDeclType = 14, 
		RULE_structDeclType = 15, RULE_structMemDecls = 16, RULE_identifierList = 17, 
		RULE_expression = 18, RULE_expressionList = 19, RULE_primaryExpression = 20, 
		RULE_operand = 21, RULE_literal = 22, RULE_index = 23, RULE_arguments = 24, 
		RULE_selector = 25, RULE_appendExpression = 26, RULE_lengthExpression = 27, 
		RULE_capExpression = 28, RULE_statementList = 29, RULE_block = 30, RULE_statement = 31, 
		RULE_simpleStatement = 32, RULE_assignmentStatement = 33, RULE_ifStatement = 34, 
		RULE_loop = 35, RULE_switch = 36, RULE_expressionCaseClauseList = 37, 
		RULE_expressionCaseClause = 38, RULE_expressionSwitchCase = 39, RULE_operator = 40, 
		RULE_epsilon = 41;
	private static String[] makeRuleNames() {
		return new String[] {
			"root", "topDeclarationList", "variableDecl", "innerVarDecls", "singleVarDecl", 
			"singleVarDeclNoExps", "typeDecl", "innerTypeDecls", "singleTypeDecl", 
			"funcDecl", "funcFrontDecl", "funcArgDecls", "declType", "sliceDeclType", 
			"arrayDeclType", "structDeclType", "structMemDecls", "identifierList", 
			"expression", "expressionList", "primaryExpression", "operand", "literal", 
			"index", "arguments", "selector", "appendExpression", "lengthExpression", 
			"capExpression", "statementList", "block", "statement", "simpleStatement", 
			"assignmentStatement", "ifStatement", "loop", "switch", "expressionCaseClauseList", 
			"expressionCaseClause", "expressionSwitchCase", "operator", "epsilon"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "';'", "':='", "'('", "')'", "'~'", "':'", "'+'", "'-'", "'*'", 
			"'/'", "','", "'['", "']'", "'{'", "'}'", "'++'", "'--'", "'%'", "'<<'", 
			"'>>'", "'&'", "'&^'", "'|'", "'^'", "'=='", "'!='", "'<'", "'>'", "'<='", 
			"'>='", "'&&'", "'||'", "'!'", "'='", "'+='", "'&='", "'-='", "'|='", 
			"'*='", "'^='", "'<<='", "'>>='", "'&^='", "'%='", "'/='", "'.'", "'\\'", 
			"'''", "'package'", "'if'", "'while'", "'let'", "'then'", "'else'", "'in'", 
			"'do'", "'begin'", "'end'", "'const'", "'var'", "'type'", "'func'", "'struct'", 
			"'append'", "'len'", "'cap'", "'print'", "'println'", "'return'", "'break'", 
			"'continue'", "'for'", "'switch'", "'case'", "'default'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "SEMI", "SVD", "LP", "RP", "TIL", "COL", "PL", "MIN", "MUL", "DIV", 
			"CM", "LSB", "RSB", "LCB", "RCB", "INC", "DEC", "MOD", "LSH", "RSH", 
			"AMPER", "BC", "VB", "CARET", "EQ", "NEQ", "LT", "GT", "LTE", "GTE", 
			"LAND", "LOR", "NEG", "ASOP", "AAOP", "BAOP", "SAOP", "BOAOP", "MAOP", 
			"BXOOP", "LSAOP", "RSAOP", "BCAOP", "RAOP", "DAOP", "DOT", "DBS", "BS", 
			"PACKAGE", "IF", "WHILE", "LET", "THEN", "ELSE", "IN", "DO", "BEGIN", 
			"END", "CONST", "VAR", "TYPE", "FUNC", "STRUCT", "APPEND", "LEN", "CAP", 
			"PRINT", "PRINTLN", "RETURN", "BREAK", "CONTINUE", "FOR", "SWITCH", "CASE", 
			"DEFAULT", "ID", "INTLITERAL", "FLOATLITERAL", "RUNELITERAL", "RAWSTRINGLITERAL", 
			"INTERPRETEDSTRINGLITERAL", "LETTER", "DIGIT", "LINE_COMMENT", "SPACES"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "goParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public goParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RootContext extends ParserRuleContext {
		public RootContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_root; }
	 
		public RootContext() { }
		public void copyFrom(RootContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RootASTContext extends RootContext {
		public TerminalNode PACKAGE() { return getToken(goParser.PACKAGE, 0); }
		public TerminalNode ID() { return getToken(goParser.ID, 0); }
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public TopDeclarationListContext topDeclarationList() {
			return getRuleContext(TopDeclarationListContext.class,0);
		}
		public RootASTContext(RootContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterRootAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitRootAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitRootAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final RootContext root() throws RecognitionException {
		RootContext _localctx = new RootContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_root);
		try {
			_localctx = new RootASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(84);
			match(PACKAGE);
			setState(85);
			match(ID);
			setState(86);
			match(SEMI);
			setState(87);
			topDeclarationList();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TopDeclarationListContext extends ParserRuleContext {
		public TopDeclarationListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_topDeclarationList; }
	 
		public TopDeclarationListContext() { }
		public void copyFrom(TopDeclarationListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TopDeclarationListASTContext extends TopDeclarationListContext {
		public List<VariableDeclContext> variableDecl() {
			return getRuleContexts(VariableDeclContext.class);
		}
		public VariableDeclContext variableDecl(int i) {
			return getRuleContext(VariableDeclContext.class,i);
		}
		public List<TypeDeclContext> typeDecl() {
			return getRuleContexts(TypeDeclContext.class);
		}
		public TypeDeclContext typeDecl(int i) {
			return getRuleContext(TypeDeclContext.class,i);
		}
		public List<FuncDeclContext> funcDecl() {
			return getRuleContexts(FuncDeclContext.class);
		}
		public FuncDeclContext funcDecl(int i) {
			return getRuleContext(FuncDeclContext.class,i);
		}
		public TopDeclarationListASTContext(TopDeclarationListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterTopDeclarationListAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitTopDeclarationListAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitTopDeclarationListAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TopDeclarationListContext topDeclarationList() throws RecognitionException {
		TopDeclarationListContext _localctx = new TopDeclarationListContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_topDeclarationList);
		int _la;
		try {
			_localctx = new TopDeclarationListASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8070450532247928840L) != 0)) {
				{
				setState(92);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case LP:
				case VAR:
					{
					setState(89);
					variableDecl();
					}
					break;
				case TYPE:
					{
					setState(90);
					typeDecl();
					}
					break;
				case FUNC:
					{
					setState(91);
					funcDecl();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(96);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclContext extends ParserRuleContext {
		public VariableDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDecl; }
	 
		public VariableDeclContext() { }
		public void copyFrom(VariableDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InnerVarVDASTContext extends VariableDeclContext {
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public InnerVarDeclsContext innerVarDecls() {
			return getRuleContext(InnerVarDeclsContext.class,0);
		}
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public InnerVarVDASTContext(VariableDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterInnerVarVDAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitInnerVarVDAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitInnerVarVDAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LpVDASTContext extends VariableDeclContext {
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public LpVDASTContext(VariableDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterLpVDAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitLpVDAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitLpVDAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarVDASTContext extends VariableDeclContext {
		public TerminalNode VAR() { return getToken(goParser.VAR, 0); }
		public SingleVarDeclContext singleVarDecl() {
			return getRuleContext(SingleVarDeclContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public SingleVarVDASTContext(VariableDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSingleVarVDAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSingleVarVDAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSingleVarVDAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final VariableDeclContext variableDecl() throws RecognitionException {
		VariableDeclContext _localctx = new VariableDeclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_variableDecl);
		try {
			setState(109);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				_localctx = new SingleVarVDASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(97);
				match(VAR);
				setState(98);
				singleVarDecl();
				setState(99);
				match(SEMI);
				}
				break;
			case 2:
				_localctx = new InnerVarVDASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(101);
				match(LP);
				setState(102);
				innerVarDecls();
				setState(103);
				match(RP);
				setState(104);
				match(SEMI);
				}
				break;
			case 3:
				_localctx = new LpVDASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(106);
				match(LP);
				setState(107);
				match(RP);
				setState(108);
				match(SEMI);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InnerVarDeclsContext extends ParserRuleContext {
		public InnerVarDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_innerVarDecls; }
	 
		public InnerVarDeclsContext() { }
		public void copyFrom(InnerVarDeclsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InnerVarDeclsASTContext extends InnerVarDeclsContext {
		public List<SingleVarDeclContext> singleVarDecl() {
			return getRuleContexts(SingleVarDeclContext.class);
		}
		public SingleVarDeclContext singleVarDecl(int i) {
			return getRuleContext(SingleVarDeclContext.class,i);
		}
		public List<TerminalNode> SEMI() { return getTokens(goParser.SEMI); }
		public TerminalNode SEMI(int i) {
			return getToken(goParser.SEMI, i);
		}
		public InnerVarDeclsASTContext(InnerVarDeclsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterInnerVarDeclsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitInnerVarDeclsAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitInnerVarDeclsAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InnerVarDeclsContext innerVarDecls() throws RecognitionException {
		InnerVarDeclsContext _localctx = new InnerVarDeclsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_innerVarDecls);
		int _la;
		try {
			_localctx = new InnerVarDeclsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			singleVarDecl();
			setState(112);
			match(SEMI);
			setState(118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ID) {
				{
				{
				setState(113);
				singleVarDecl();
				setState(114);
				match(SEMI);
				}
				}
				setState(120);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclContext extends ParserRuleContext {
		public SingleVarDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleVarDecl; }
	 
		public SingleVarDeclContext() { }
		public void copyFrom(SingleVarDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdListDTSVDASTContext extends SingleVarDeclContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public TerminalNode ASOP() { return getToken(goParser.ASOP, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public IdListDTSVDASTContext(SingleVarDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIdListDTSVDAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIdListDTSVDAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIdListDTSVDAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdListSVDASTContext extends SingleVarDeclContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public TerminalNode ASOP() { return getToken(goParser.ASOP, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public IdListSVDASTContext(SingleVarDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIdListSVDAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIdListSVDAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIdListSVDAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDASTContext extends SingleVarDeclContext {
		public SingleVarDeclNoExpsContext singleVarDeclNoExps() {
			return getRuleContext(SingleVarDeclNoExpsContext.class,0);
		}
		public SingleVarDASTContext(SingleVarDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSingleVarDAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSingleVarDAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSingleVarDAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SingleVarDeclContext singleVarDecl() throws RecognitionException {
		SingleVarDeclContext _localctx = new SingleVarDeclContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_singleVarDecl);
		try {
			setState(131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new IdListDTSVDASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(121);
				identifierList();
				setState(122);
				declType();
				setState(123);
				match(ASOP);
				setState(124);
				expressionList();
				}
				break;
			case 2:
				_localctx = new IdListSVDASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(126);
				identifierList();
				setState(127);
				match(ASOP);
				setState(128);
				expressionList();
				}
				break;
			case 3:
				_localctx = new SingleVarDASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(130);
				singleVarDeclNoExps();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclNoExpsContext extends ParserRuleContext {
		public SingleVarDeclNoExpsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleVarDeclNoExps; }
	 
		public SingleVarDeclNoExpsContext() { }
		public void copyFrom(SingleVarDeclNoExpsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SingleVarDeclNoExpsASTContext extends SingleVarDeclNoExpsContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public SingleVarDeclNoExpsASTContext(SingleVarDeclNoExpsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSingleVarDeclNoExpsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSingleVarDeclNoExpsAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSingleVarDeclNoExpsAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SingleVarDeclNoExpsContext singleVarDeclNoExps() throws RecognitionException {
		SingleVarDeclNoExpsContext _localctx = new SingleVarDeclNoExpsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_singleVarDeclNoExps);
		try {
			_localctx = new SingleVarDeclNoExpsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			identifierList();
			setState(134);
			declType();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeDeclContext extends ParserRuleContext {
		public TypeDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDecl; }
	 
		public TypeDeclContext() { }
		public void copyFrom(TypeDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeTDASTContext extends TypeDeclContext {
		public TerminalNode TYPE() { return getToken(goParser.TYPE, 0); }
		public SingleTypeDeclContext singleTypeDecl() {
			return getRuleContext(SingleTypeDeclContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public InnerTypeDeclsContext innerTypeDecls() {
			return getRuleContext(InnerTypeDeclsContext.class,0);
		}
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public TypeTDASTContext(TypeDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterTypeTDAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitTypeTDAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitTypeTDAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final TypeDeclContext typeDecl() throws RecognitionException {
		TypeDeclContext _localctx = new TypeDeclContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_typeDecl);
		try {
			_localctx = new TypeTDASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			match(TYPE);
			setState(148);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(137);
				singleTypeDecl();
				setState(138);
				match(SEMI);
				}
				break;
			case 2:
				{
				setState(140);
				match(LP);
				setState(141);
				innerTypeDecls();
				setState(142);
				match(RP);
				setState(143);
				match(SEMI);
				}
				break;
			case 3:
				{
				setState(145);
				match(LP);
				setState(146);
				match(RP);
				setState(147);
				match(SEMI);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InnerTypeDeclsContext extends ParserRuleContext {
		public InnerTypeDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_innerTypeDecls; }
	 
		public InnerTypeDeclsContext() { }
		public void copyFrom(InnerTypeDeclsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InnerTypeDeclsASTContext extends InnerTypeDeclsContext {
		public List<SingleTypeDeclContext> singleTypeDecl() {
			return getRuleContexts(SingleTypeDeclContext.class);
		}
		public SingleTypeDeclContext singleTypeDecl(int i) {
			return getRuleContext(SingleTypeDeclContext.class,i);
		}
		public List<TerminalNode> SEMI() { return getTokens(goParser.SEMI); }
		public TerminalNode SEMI(int i) {
			return getToken(goParser.SEMI, i);
		}
		public InnerTypeDeclsASTContext(InnerTypeDeclsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterInnerTypeDeclsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitInnerTypeDeclsAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitInnerTypeDeclsAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final InnerTypeDeclsContext innerTypeDecls() throws RecognitionException {
		InnerTypeDeclsContext _localctx = new InnerTypeDeclsContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_innerTypeDecls);
		int _la;
		try {
			_localctx = new InnerTypeDeclsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			singleTypeDecl();
			setState(151);
			match(SEMI);
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ID) {
				{
				{
				setState(152);
				singleTypeDecl();
				setState(153);
				match(SEMI);
				}
				}
				setState(159);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SingleTypeDeclContext extends ParserRuleContext {
		public SingleTypeDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleTypeDecl; }
	 
		public SingleTypeDeclContext() { }
		public void copyFrom(SingleTypeDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdSYDASTContext extends SingleTypeDeclContext {
		public TerminalNode ID() { return getToken(goParser.ID, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public IdSYDASTContext(SingleTypeDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIdSYDAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIdSYDAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIdSYDAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SingleTypeDeclContext singleTypeDecl() throws RecognitionException {
		SingleTypeDeclContext _localctx = new SingleTypeDeclContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_singleTypeDecl);
		try {
			_localctx = new IdSYDASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(160);
			match(ID);
			setState(161);
			declType();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncDeclContext extends ParserRuleContext {
		public FuncDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcDecl; }
	 
		public FuncDeclContext() { }
		public void copyFrom(FuncDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncDeclASTContext extends FuncDeclContext {
		public FuncFrontDeclContext funcFrontDecl() {
			return getRuleContext(FuncFrontDeclContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public FuncDeclASTContext(FuncDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterFuncDeclAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitFuncDeclAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitFuncDeclAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FuncDeclContext funcDecl() throws RecognitionException {
		FuncDeclContext _localctx = new FuncDeclContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_funcDecl);
		try {
			_localctx = new FuncDeclASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(163);
			funcFrontDecl();
			setState(164);
			block();
			setState(165);
			match(SEMI);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncFrontDeclContext extends ParserRuleContext {
		public FuncFrontDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcFrontDecl; }
	 
		public FuncFrontDeclContext() { }
		public void copyFrom(FuncFrontDeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncFFDASTContext extends FuncFrontDeclContext {
		public TerminalNode FUNC() { return getToken(goParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(goParser.ID, 0); }
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public FuncArgDeclsContext funcArgDecls() {
			return getRuleContext(FuncArgDeclsContext.class,0);
		}
		public List<EpsilonContext> epsilon() {
			return getRuleContexts(EpsilonContext.class);
		}
		public EpsilonContext epsilon(int i) {
			return getRuleContext(EpsilonContext.class,i);
		}
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public FuncFFDASTContext(FuncFrontDeclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterFuncFFDAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitFuncFFDAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitFuncFFDAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FuncFrontDeclContext funcFrontDecl() throws RecognitionException {
		FuncFrontDeclContext _localctx = new FuncFrontDeclContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_funcFrontDecl);
		try {
			_localctx = new FuncFFDASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			match(FUNC);
			setState(168);
			match(ID);
			setState(169);
			match(LP);
			setState(172);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				{
				setState(170);
				funcArgDecls();
				}
				break;
			case RP:
				{
				setState(171);
				epsilon();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(174);
			match(RP);
			setState(177);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LP:
			case LSB:
			case STRUCT:
			case ID:
				{
				setState(175);
				declType();
				}
				break;
			case LCB:
				{
				setState(176);
				epsilon();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncArgDeclsContext extends ParserRuleContext {
		public FuncArgDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcArgDecls; }
	 
		public FuncArgDeclsContext() { }
		public void copyFrom(FuncArgDeclsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncArgDeclsASTContext extends FuncArgDeclsContext {
		public List<SingleVarDeclNoExpsContext> singleVarDeclNoExps() {
			return getRuleContexts(SingleVarDeclNoExpsContext.class);
		}
		public SingleVarDeclNoExpsContext singleVarDeclNoExps(int i) {
			return getRuleContext(SingleVarDeclNoExpsContext.class,i);
		}
		public List<TerminalNode> CM() { return getTokens(goParser.CM); }
		public TerminalNode CM(int i) {
			return getToken(goParser.CM, i);
		}
		public FuncArgDeclsASTContext(FuncArgDeclsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterFuncArgDeclsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitFuncArgDeclsAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitFuncArgDeclsAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FuncArgDeclsContext funcArgDecls() throws RecognitionException {
		FuncArgDeclsContext _localctx = new FuncArgDeclsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_funcArgDecls);
		int _la;
		try {
			_localctx = new FuncArgDeclsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			singleVarDeclNoExps();
			setState(184);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CM) {
				{
				{
				setState(180);
				match(CM);
				setState(181);
				singleVarDeclNoExps();
				}
				}
				setState(186);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclTypeContext extends ParserRuleContext {
		public DeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declType; }
	 
		public DeclTypeContext() { }
		public void copyFrom(DeclTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclTypeDTASTContext extends DeclTypeContext {
		public StructDeclTypeContext structDeclType() {
			return getRuleContext(StructDeclTypeContext.class,0);
		}
		public StructDeclTypeDTASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterStructDeclTypeDTAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitStructDeclTypeDTAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitStructDeclTypeDTAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceDeclTypeDTASTContext extends DeclTypeContext {
		public SliceDeclTypeContext sliceDeclType() {
			return getRuleContext(SliceDeclTypeContext.class,0);
		}
		public SliceDeclTypeDTASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSliceDeclTypeDTAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSliceDeclTypeDTAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSliceDeclTypeDTAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArrayDeclTypeDTASTContext extends DeclTypeContext {
		public ArrayDeclTypeContext arrayDeclType() {
			return getRuleContext(ArrayDeclTypeContext.class,0);
		}
		public ArrayDeclTypeDTASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterArrayDeclTypeDTAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitArrayDeclTypeDTAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitArrayDeclTypeDTAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LpDTASTContext extends DeclTypeContext {
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public LpDTASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterLpDTAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitLpDTAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitLpDTAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdDTASTContext extends DeclTypeContext {
		public TerminalNode ID() { return getToken(goParser.ID, 0); }
		public IdDTASTContext(DeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIdDTAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIdDTAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIdDTAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final DeclTypeContext declType() throws RecognitionException {
		DeclTypeContext _localctx = new DeclTypeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_declType);
		try {
			setState(195);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				_localctx = new LpDTASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(187);
				match(LP);
				setState(188);
				declType();
				setState(189);
				match(RP);
				}
				break;
			case 2:
				_localctx = new IdDTASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(191);
				match(ID);
				}
				break;
			case 3:
				_localctx = new SliceDeclTypeDTASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(192);
				sliceDeclType();
				}
				break;
			case 4:
				_localctx = new ArrayDeclTypeDTASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(193);
				arrayDeclType();
				}
				break;
			case 5:
				_localctx = new StructDeclTypeDTASTContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(194);
				structDeclType();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SliceDeclTypeContext extends ParserRuleContext {
		public SliceDeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceDeclType; }
	 
		public SliceDeclTypeContext() { }
		public void copyFrom(SliceDeclTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LsbSDTASTContext extends SliceDeclTypeContext {
		public TerminalNode LSB() { return getToken(goParser.LSB, 0); }
		public TerminalNode RSB() { return getToken(goParser.RSB, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public LsbSDTASTContext(SliceDeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterLsbSDTAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitLsbSDTAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitLsbSDTAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SliceDeclTypeContext sliceDeclType() throws RecognitionException {
		SliceDeclTypeContext _localctx = new SliceDeclTypeContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_sliceDeclType);
		try {
			_localctx = new LsbSDTASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(197);
			match(LSB);
			setState(198);
			match(RSB);
			setState(199);
			declType();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayDeclTypeContext extends ParserRuleContext {
		public ArrayDeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayDeclType; }
	 
		public ArrayDeclTypeContext() { }
		public void copyFrom(ArrayDeclTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LsbADTASTContext extends ArrayDeclTypeContext {
		public TerminalNode LSB() { return getToken(goParser.LSB, 0); }
		public TerminalNode INTLITERAL() { return getToken(goParser.INTLITERAL, 0); }
		public TerminalNode RSB() { return getToken(goParser.RSB, 0); }
		public DeclTypeContext declType() {
			return getRuleContext(DeclTypeContext.class,0);
		}
		public LsbADTASTContext(ArrayDeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterLsbADTAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitLsbADTAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitLsbADTAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArrayDeclTypeContext arrayDeclType() throws RecognitionException {
		ArrayDeclTypeContext _localctx = new ArrayDeclTypeContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_arrayDeclType);
		try {
			_localctx = new LsbADTASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			match(LSB);
			setState(202);
			match(INTLITERAL);
			setState(203);
			match(RSB);
			setState(204);
			declType();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclTypeContext extends ParserRuleContext {
		public StructDeclTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDeclType; }
	 
		public StructDeclTypeContext() { }
		public void copyFrom(StructDeclTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructSDTASTContext extends StructDeclTypeContext {
		public TerminalNode STRUCT() { return getToken(goParser.STRUCT, 0); }
		public TerminalNode LCB() { return getToken(goParser.LCB, 0); }
		public TerminalNode RCB() { return getToken(goParser.RCB, 0); }
		public StructMemDeclsContext structMemDecls() {
			return getRuleContext(StructMemDeclsContext.class,0);
		}
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public StructSDTASTContext(StructDeclTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterStructSDTAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitStructSDTAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitStructSDTAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructDeclTypeContext structDeclType() throws RecognitionException {
		StructDeclTypeContext _localctx = new StructDeclTypeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_structDeclType);
		try {
			_localctx = new StructSDTASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
			match(STRUCT);
			setState(207);
			match(LCB);
			setState(210);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				{
				setState(208);
				structMemDecls();
				}
				break;
			case RCB:
				{
				setState(209);
				epsilon();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(212);
			match(RCB);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructMemDeclsContext extends ParserRuleContext {
		public StructMemDeclsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structMemDecls; }
	 
		public StructMemDeclsContext() { }
		public void copyFrom(StructMemDeclsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructMemDeclsASTContext extends StructMemDeclsContext {
		public List<SingleVarDeclNoExpsContext> singleVarDeclNoExps() {
			return getRuleContexts(SingleVarDeclNoExpsContext.class);
		}
		public SingleVarDeclNoExpsContext singleVarDeclNoExps(int i) {
			return getRuleContext(SingleVarDeclNoExpsContext.class,i);
		}
		public List<TerminalNode> SEMI() { return getTokens(goParser.SEMI); }
		public TerminalNode SEMI(int i) {
			return getToken(goParser.SEMI, i);
		}
		public StructMemDeclsASTContext(StructMemDeclsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterStructMemDeclsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitStructMemDeclsAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitStructMemDeclsAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StructMemDeclsContext structMemDecls() throws RecognitionException {
		StructMemDeclsContext _localctx = new StructMemDeclsContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_structMemDecls);
		int _la;
		try {
			_localctx = new StructMemDeclsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			singleVarDeclNoExps();
			setState(215);
			match(SEMI);
			setState(221);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ID) {
				{
				{
				setState(216);
				singleVarDeclNoExps();
				setState(217);
				match(SEMI);
				}
				}
				setState(223);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierListContext extends ParserRuleContext {
		public IdentifierListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierList; }
	 
		public IdentifierListContext() { }
		public void copyFrom(IdentifierListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierListASTContext extends IdentifierListContext {
		public List<TerminalNode> ID() { return getTokens(goParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(goParser.ID, i);
		}
		public List<TerminalNode> CM() { return getTokens(goParser.CM); }
		public TerminalNode CM(int i) {
			return getToken(goParser.CM, i);
		}
		public IdentifierListASTContext(IdentifierListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIdentifierListAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIdentifierListAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIdentifierListAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IdentifierListContext identifierList() throws RecognitionException {
		IdentifierListContext _localctx = new IdentifierListContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_identifierList);
		int _la;
		try {
			_localctx = new IdentifierListASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
			match(ID);
			setState(229);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CM) {
				{
				{
				setState(225);
				match(CM);
				setState(226);
				match(ID);
				}
				}
				setState(231);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	 
		public ExpressionContext() { }
		public void copyFrom(ExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionEASTContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public ExpressionEASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterExpressionEAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitExpressionEAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitExpressionEAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionEASTContext extends ExpressionContext {
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public PrimaryExpressionEASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterPrimaryExpressionEAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitPrimaryExpressionEAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitPrimaryExpressionEAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OperatorEASTContext extends ExpressionContext {
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public OperatorEASTContext(ExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterOperatorEAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitOperatorEAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitOperatorEAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(237);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LP:
			case APPEND:
			case LEN:
			case CAP:
			case ID:
			case INTLITERAL:
			case FLOATLITERAL:
			case RUNELITERAL:
			case RAWSTRINGLITERAL:
			case INTERPRETEDSTRINGLITERAL:
				{
				_localctx = new PrimaryExpressionEASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(233);
				primaryExpression(0);
				}
				break;
			case PL:
			case MIN:
			case MUL:
			case DIV:
			case MOD:
			case LSH:
			case RSH:
			case AMPER:
			case BC:
			case VB:
			case CARET:
			case EQ:
			case NEQ:
			case LT:
			case GT:
			case LTE:
			case GTE:
			case LAND:
			case LOR:
			case NEG:
			case ASOP:
			case AAOP:
			case BAOP:
			case SAOP:
			case BOAOP:
			case MAOP:
			case BXOOP:
			case LSAOP:
			case RSAOP:
			case BCAOP:
			case RAOP:
			case DAOP:
				{
				_localctx = new OperatorEASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(234);
				operator();
				setState(235);
				expression(1);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(245);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpressionEASTContext(new ExpressionContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_expression);
					setState(239);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(240);
					operator();
					setState(241);
					expression(3);
					}
					} 
				}
				setState(247);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionListContext extends ParserRuleContext {
		public ExpressionListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionList; }
	 
		public ExpressionListContext() { }
		public void copyFrom(ExpressionListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionListASTContext extends ExpressionListContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> CM() { return getTokens(goParser.CM); }
		public TerminalNode CM(int i) {
			return getToken(goParser.CM, i);
		}
		public ExpressionListASTContext(ExpressionListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterExpressionListAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitExpressionListAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitExpressionListAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionListContext expressionList() throws RecognitionException {
		ExpressionListContext _localctx = new ExpressionListContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_expressionList);
		int _la;
		try {
			_localctx = new ExpressionListASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(248);
			expression(0);
			setState(253);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CM) {
				{
				{
				setState(249);
				match(CM);
				setState(250);
				expression(0);
				}
				}
				setState(255);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionContext extends ParserRuleContext {
		public PrimaryExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryExpression; }
	 
		public PrimaryExpressionContext() { }
		public void copyFrom(PrimaryExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LengthExpPEASTContext extends PrimaryExpressionContext {
		public LengthExpressionContext lengthExpression() {
			return getRuleContext(LengthExpressionContext.class,0);
		}
		public LengthExpPEASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterLengthExpPEAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitLengthExpPEAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitLengthExpPEAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AppendExpPEASTContext extends PrimaryExpressionContext {
		public AppendExpressionContext appendExpression() {
			return getRuleContext(AppendExpressionContext.class,0);
		}
		public AppendExpPEASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterAppendExpPEAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitAppendExpPEAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitAppendExpPEAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OperandPEASTContext extends PrimaryExpressionContext {
		public OperandContext operand() {
			return getRuleContext(OperandContext.class,0);
		}
		public OperandPEASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterOperandPEAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitOperandPEAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitOperandPEAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpPEASTContext extends PrimaryExpressionContext {
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public SelectorContext selector() {
			return getRuleContext(SelectorContext.class,0);
		}
		public IndexContext index() {
			return getRuleContext(IndexContext.class,0);
		}
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public PrimaryExpPEASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterPrimaryExpPEAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitPrimaryExpPEAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitPrimaryExpPEAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CapExpPEASTContext extends PrimaryExpressionContext {
		public CapExpressionContext capExpression() {
			return getRuleContext(CapExpressionContext.class,0);
		}
		public CapExpPEASTContext(PrimaryExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterCapExpPEAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitCapExpPEAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitCapExpPEAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final PrimaryExpressionContext primaryExpression() throws RecognitionException {
		return primaryExpression(0);
	}

	private PrimaryExpressionContext primaryExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PrimaryExpressionContext _localctx = new PrimaryExpressionContext(_ctx, _parentState);
		PrimaryExpressionContext _prevctx = _localctx;
		int _startState = 40;
		enterRecursionRule(_localctx, 40, RULE_primaryExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(261);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LP:
			case ID:
			case INTLITERAL:
			case FLOATLITERAL:
			case RUNELITERAL:
			case RAWSTRINGLITERAL:
			case INTERPRETEDSTRINGLITERAL:
				{
				_localctx = new OperandPEASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(257);
				operand();
				}
				break;
			case APPEND:
				{
				_localctx = new AppendExpPEASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(258);
				appendExpression();
				}
				break;
			case LEN:
				{
				_localctx = new LengthExpPEASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(259);
				lengthExpression();
				}
				break;
			case CAP:
				{
				_localctx = new CapExpPEASTContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(260);
				capExpression();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(271);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PrimaryExpPEASTContext(new PrimaryExpressionContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_primaryExpression);
					setState(263);
					if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
					setState(267);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case DOT:
						{
						setState(264);
						selector();
						}
						break;
					case LSB:
						{
						setState(265);
						index();
						}
						break;
					case LP:
						{
						setState(266);
						arguments();
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
					} 
				}
				setState(273);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperandContext extends ParserRuleContext {
		public OperandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operand; }
	 
		public OperandContext() { }
		public void copyFrom(OperandContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralOASTContext extends OperandContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public LiteralOASTContext(OperandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterLiteralOAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitLiteralOAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitLiteralOAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdOASTContext extends OperandContext {
		public TerminalNode ID() { return getToken(goParser.ID, 0); }
		public IdOASTContext(OperandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIdOAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIdOAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIdOAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionOASTContext extends OperandContext {
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public ExpressionOASTContext(OperandContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterExpressionOAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitExpressionOAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitExpressionOAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final OperandContext operand() throws RecognitionException {
		OperandContext _localctx = new OperandContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_operand);
		try {
			setState(280);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTLITERAL:
			case FLOATLITERAL:
			case RUNELITERAL:
			case RAWSTRINGLITERAL:
			case INTERPRETEDSTRINGLITERAL:
				_localctx = new LiteralOASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(274);
				literal();
				}
				break;
			case ID:
				_localctx = new IdOASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(275);
				match(ID);
				}
				break;
			case LP:
				_localctx = new ExpressionOASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(276);
				match(LP);
				setState(277);
				expression(0);
				setState(278);
				match(RP);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	 
		public LiteralContext() { }
		public void copyFrom(LiteralContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntliteralASTContext extends LiteralContext {
		public TerminalNode INTLITERAL() { return getToken(goParser.INTLITERAL, 0); }
		public IntliteralASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIntliteralAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIntliteralAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIntliteralAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InterpretedliteralASTContext extends LiteralContext {
		public TerminalNode INTERPRETEDSTRINGLITERAL() { return getToken(goParser.INTERPRETEDSTRINGLITERAL, 0); }
		public InterpretedliteralASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterInterpretedliteralAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitInterpretedliteralAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitInterpretedliteralAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RunliteralASTContext extends LiteralContext {
		public TerminalNode RUNELITERAL() { return getToken(goParser.RUNELITERAL, 0); }
		public RunliteralASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterRunliteralAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitRunliteralAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitRunliteralAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatliteralASTContext extends LiteralContext {
		public TerminalNode FLOATLITERAL() { return getToken(goParser.FLOATLITERAL, 0); }
		public FloatliteralASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterFloatliteralAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitFloatliteralAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitFloatliteralAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RawsliteralASTContext extends LiteralContext {
		public TerminalNode RAWSTRINGLITERAL() { return getToken(goParser.RAWSTRINGLITERAL, 0); }
		public RawsliteralASTContext(LiteralContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterRawsliteralAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitRawsliteralAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitRawsliteralAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_literal);
		try {
			setState(287);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTLITERAL:
				_localctx = new IntliteralASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(282);
				match(INTLITERAL);
				}
				break;
			case FLOATLITERAL:
				_localctx = new FloatliteralASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(283);
				match(FLOATLITERAL);
				}
				break;
			case RUNELITERAL:
				_localctx = new RunliteralASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(284);
				match(RUNELITERAL);
				}
				break;
			case RAWSTRINGLITERAL:
				_localctx = new RawsliteralASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(285);
				match(RAWSTRINGLITERAL);
				}
				break;
			case INTERPRETEDSTRINGLITERAL:
				_localctx = new InterpretedliteralASTContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(286);
				match(INTERPRETEDSTRINGLITERAL);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IndexContext extends ParserRuleContext {
		public IndexContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_index; }
	 
		public IndexContext() { }
		public void copyFrom(IndexContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdexASTContext extends IndexContext {
		public TerminalNode LSB() { return getToken(goParser.LSB, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RSB() { return getToken(goParser.RSB, 0); }
		public IdexASTContext(IndexContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIdexAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIdexAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIdexAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IndexContext index() throws RecognitionException {
		IndexContext _localctx = new IndexContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_index);
		try {
			_localctx = new IdexASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(289);
			match(LSB);
			setState(290);
			expression(0);
			setState(291);
			match(RSB);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentsContext extends ParserRuleContext {
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	 
		public ArgumentsContext() { }
		public void copyFrom(ArgumentsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentsASTContext extends ArgumentsContext {
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public ArgumentsASTContext(ArgumentsContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterArgumentsAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitArgumentsAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitArgumentsAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_arguments);
		try {
			_localctx = new ArgumentsASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(293);
			match(LP);
			setState(296);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LP:
			case PL:
			case MIN:
			case MUL:
			case DIV:
			case MOD:
			case LSH:
			case RSH:
			case AMPER:
			case BC:
			case VB:
			case CARET:
			case EQ:
			case NEQ:
			case LT:
			case GT:
			case LTE:
			case GTE:
			case LAND:
			case LOR:
			case NEG:
			case ASOP:
			case AAOP:
			case BAOP:
			case SAOP:
			case BOAOP:
			case MAOP:
			case BXOOP:
			case LSAOP:
			case RSAOP:
			case BCAOP:
			case RAOP:
			case DAOP:
			case APPEND:
			case LEN:
			case CAP:
			case ID:
			case INTLITERAL:
			case FLOATLITERAL:
			case RUNELITERAL:
			case RAWSTRINGLITERAL:
			case INTERPRETEDSTRINGLITERAL:
				{
				setState(294);
				expressionList();
				}
				break;
			case RP:
				{
				setState(295);
				epsilon();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(298);
			match(RP);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SelectorContext extends ParserRuleContext {
		public SelectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selector; }
	 
		public SelectorContext() { }
		public void copyFrom(SelectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SelectorASTContext extends SelectorContext {
		public TerminalNode DOT() { return getToken(goParser.DOT, 0); }
		public TerminalNode ID() { return getToken(goParser.ID, 0); }
		public SelectorASTContext(SelectorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSelectorAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSelectorAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSelectorAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SelectorContext selector() throws RecognitionException {
		SelectorContext _localctx = new SelectorContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_selector);
		try {
			_localctx = new SelectorASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(300);
			match(DOT);
			setState(301);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AppendExpressionContext extends ParserRuleContext {
		public AppendExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_appendExpression; }
	 
		public AppendExpressionContext() { }
		public void copyFrom(AppendExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AppendExpressionASTContext extends AppendExpressionContext {
		public TerminalNode APPEND() { return getToken(goParser.APPEND, 0); }
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CM() { return getToken(goParser.CM, 0); }
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public AppendExpressionASTContext(AppendExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterAppendExpressionAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitAppendExpressionAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitAppendExpressionAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AppendExpressionContext appendExpression() throws RecognitionException {
		AppendExpressionContext _localctx = new AppendExpressionContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_appendExpression);
		try {
			_localctx = new AppendExpressionASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(303);
			match(APPEND);
			setState(304);
			match(LP);
			setState(305);
			expression(0);
			setState(306);
			match(CM);
			setState(307);
			expression(0);
			setState(308);
			match(RP);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LengthExpressionContext extends ParserRuleContext {
		public LengthExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lengthExpression; }
	 
		public LengthExpressionContext() { }
		public void copyFrom(LengthExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LengthExpressionASTContext extends LengthExpressionContext {
		public TerminalNode LEN() { return getToken(goParser.LEN, 0); }
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public LengthExpressionASTContext(LengthExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterLengthExpressionAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitLengthExpressionAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitLengthExpressionAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LengthExpressionContext lengthExpression() throws RecognitionException {
		LengthExpressionContext _localctx = new LengthExpressionContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_lengthExpression);
		try {
			_localctx = new LengthExpressionASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(310);
			match(LEN);
			setState(311);
			match(LP);
			setState(312);
			expression(0);
			setState(313);
			match(RP);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CapExpressionContext extends ParserRuleContext {
		public CapExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_capExpression; }
	 
		public CapExpressionContext() { }
		public void copyFrom(CapExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CapExpressionASTContext extends CapExpressionContext {
		public TerminalNode CAP() { return getToken(goParser.CAP, 0); }
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public CapExpressionASTContext(CapExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterCapExpressionAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitCapExpressionAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitCapExpressionAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final CapExpressionContext capExpression() throws RecognitionException {
		CapExpressionContext _localctx = new CapExpressionContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_capExpression);
		try {
			_localctx = new CapExpressionASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(315);
			match(CAP);
			setState(316);
			match(LP);
			setState(317);
			expression(0);
			setState(318);
			match(RP);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementListContext extends ParserRuleContext {
		public StatementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementList; }
	 
		public StatementListContext() { }
		public void copyFrom(StatementListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementListASTContext extends StatementListContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public StatementListASTContext(StatementListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterStatementListAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitStatementListAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitStatementListAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StatementListContext statementList() throws RecognitionException {
		StatementListContext _localctx = new StatementListContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_statementList);
		int _la;
		try {
			_localctx = new StatementListASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(323);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3459960782471317386L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 259071L) != 0)) {
				{
				{
				setState(320);
				statement();
				}
				}
				setState(325);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	 
		public BlockContext() { }
		public void copyFrom(BlockContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BlockASTContext extends BlockContext {
		public TerminalNode LCB() { return getToken(goParser.LCB, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public TerminalNode RCB() { return getToken(goParser.RCB, 0); }
		public BlockASTContext(BlockContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitBlockAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitBlockAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_block);
		try {
			_localctx = new BlockASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(326);
			match(LCB);
			setState(327);
			statementList();
			setState(328);
			match(RCB);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	 
		public StatementContext() { }
		public void copyFrom(StatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakSASTContext extends StatementContext {
		public TerminalNode BREAK() { return getToken(goParser.BREAK, 0); }
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public BreakSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterBreakSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitBreakSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitBreakSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BlockSASTContext extends StatementContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public BlockSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterBlockSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitBlockSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitBlockSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclSASTContext extends StatementContext {
		public VariableDeclContext variableDecl() {
			return getRuleContext(VariableDeclContext.class,0);
		}
		public VariableDeclSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterVariableDeclSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitVariableDeclSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitVariableDeclSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueSASTContext extends StatementContext {
		public TerminalNode CONTINUE() { return getToken(goParser.CONTINUE, 0); }
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public ContinueSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterContinueSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitContinueSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitContinueSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeDeclSASTContext extends StatementContext {
		public TypeDeclContext typeDecl() {
			return getRuleContext(TypeDeclContext.class,0);
		}
		public TypeDeclSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterTypeDeclSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitTypeDeclSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitTypeDeclSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintlnSASTContext extends StatementContext {
		public TerminalNode PRINTLN() { return getToken(goParser.PRINTLN, 0); }
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public PrintlnSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterPrintlnSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitPrintlnSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitPrintlnSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchSASTContext extends StatementContext {
		public SwitchContext switch_() {
			return getRuleContext(SwitchContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public SwitchSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSwitchSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSwitchSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSwitchSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementSASTContext extends StatementContext {
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public IfStatementSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIfStatementSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIfStatementSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIfStatementSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnSASTContext extends StatementContext {
		public TerminalNode RETURN() { return getToken(goParser.RETURN, 0); }
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public ReturnSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterReturnSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitReturnSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitReturnSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SimpleStatementSASTContext extends StatementContext {
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public SimpleStatementSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSimpleStatementSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSimpleStatementSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSimpleStatementSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LoopSASTContext extends StatementContext {
		public LoopContext loop() {
			return getRuleContext(LoopContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public LoopSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterLoopSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitLoopSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitLoopSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintSASTContext extends StatementContext {
		public TerminalNode PRINT() { return getToken(goParser.PRINT, 0); }
		public TerminalNode LP() { return getToken(goParser.LP, 0); }
		public TerminalNode RP() { return getToken(goParser.RP, 0); }
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public PrintSASTContext(StatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterPrintSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitPrintSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitPrintSAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_statement);
		try {
			setState(376);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				_localctx = new PrintSASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(330);
				match(PRINT);
				setState(331);
				match(LP);
				setState(334);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case LP:
				case PL:
				case MIN:
				case MUL:
				case DIV:
				case MOD:
				case LSH:
				case RSH:
				case AMPER:
				case BC:
				case VB:
				case CARET:
				case EQ:
				case NEQ:
				case LT:
				case GT:
				case LTE:
				case GTE:
				case LAND:
				case LOR:
				case NEG:
				case ASOP:
				case AAOP:
				case BAOP:
				case SAOP:
				case BOAOP:
				case MAOP:
				case BXOOP:
				case LSAOP:
				case RSAOP:
				case BCAOP:
				case RAOP:
				case DAOP:
				case APPEND:
				case LEN:
				case CAP:
				case ID:
				case INTLITERAL:
				case FLOATLITERAL:
				case RUNELITERAL:
				case RAWSTRINGLITERAL:
				case INTERPRETEDSTRINGLITERAL:
					{
					setState(332);
					expressionList();
					}
					break;
				case RP:
					{
					setState(333);
					epsilon();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(336);
				match(RP);
				setState(337);
				match(SEMI);
				}
				break;
			case 2:
				_localctx = new PrintlnSASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(339);
				match(PRINTLN);
				setState(340);
				match(LP);
				setState(343);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case LP:
				case PL:
				case MIN:
				case MUL:
				case DIV:
				case MOD:
				case LSH:
				case RSH:
				case AMPER:
				case BC:
				case VB:
				case CARET:
				case EQ:
				case NEQ:
				case LT:
				case GT:
				case LTE:
				case GTE:
				case LAND:
				case LOR:
				case NEG:
				case ASOP:
				case AAOP:
				case BAOP:
				case SAOP:
				case BOAOP:
				case MAOP:
				case BXOOP:
				case LSAOP:
				case RSAOP:
				case BCAOP:
				case RAOP:
				case DAOP:
				case APPEND:
				case LEN:
				case CAP:
				case ID:
				case INTLITERAL:
				case FLOATLITERAL:
				case RUNELITERAL:
				case RAWSTRINGLITERAL:
				case INTERPRETEDSTRINGLITERAL:
					{
					setState(341);
					expressionList();
					}
					break;
				case RP:
					{
					setState(342);
					epsilon();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(345);
				match(RP);
				setState(346);
				match(SEMI);
				}
				break;
			case 3:
				_localctx = new ReturnSASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(348);
				match(RETURN);
				setState(351);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case LP:
				case PL:
				case MIN:
				case MUL:
				case DIV:
				case MOD:
				case LSH:
				case RSH:
				case AMPER:
				case BC:
				case VB:
				case CARET:
				case EQ:
				case NEQ:
				case LT:
				case GT:
				case LTE:
				case GTE:
				case LAND:
				case LOR:
				case NEG:
				case ASOP:
				case AAOP:
				case BAOP:
				case SAOP:
				case BOAOP:
				case MAOP:
				case BXOOP:
				case LSAOP:
				case RSAOP:
				case BCAOP:
				case RAOP:
				case DAOP:
				case APPEND:
				case LEN:
				case CAP:
				case ID:
				case INTLITERAL:
				case FLOATLITERAL:
				case RUNELITERAL:
				case RAWSTRINGLITERAL:
				case INTERPRETEDSTRINGLITERAL:
					{
					setState(349);
					expression(0);
					}
					break;
				case SEMI:
					{
					setState(350);
					epsilon();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(353);
				match(SEMI);
				}
				break;
			case 4:
				_localctx = new BreakSASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(355);
				match(BREAK);
				setState(356);
				match(SEMI);
				}
				break;
			case 5:
				_localctx = new ContinueSASTContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(357);
				match(CONTINUE);
				setState(358);
				match(SEMI);
				}
				break;
			case 6:
				_localctx = new SimpleStatementSASTContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(359);
				simpleStatement();
				setState(360);
				match(SEMI);
				}
				break;
			case 7:
				_localctx = new BlockSASTContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(362);
				block();
				setState(363);
				match(SEMI);
				}
				break;
			case 8:
				_localctx = new SwitchSASTContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(365);
				switch_();
				setState(366);
				match(SEMI);
				}
				break;
			case 9:
				_localctx = new IfStatementSASTContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(368);
				ifStatement();
				setState(369);
				match(SEMI);
				}
				break;
			case 10:
				_localctx = new LoopSASTContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(371);
				loop();
				setState(372);
				match(SEMI);
				}
				break;
			case 11:
				_localctx = new TypeDeclSASTContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(374);
				typeDecl();
				}
				break;
			case 12:
				_localctx = new VariableDeclSASTContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(375);
				variableDecl();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SimpleStatementContext extends ParserRuleContext {
		public SimpleStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleStatement; }
	 
		public SimpleStatementContext() { }
		public void copyFrom(SimpleStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssigmentStatementSSASTContext extends SimpleStatementContext {
		public AssignmentStatementContext assignmentStatement() {
			return getRuleContext(AssignmentStatementContext.class,0);
		}
		public AssigmentStatementSSASTContext(SimpleStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterAssigmentStatementSSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitAssigmentStatementSSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitAssigmentStatementSSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionSSASTContext extends SimpleStatementContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode INC() { return getToken(goParser.INC, 0); }
		public TerminalNode DEC() { return getToken(goParser.DEC, 0); }
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public ExpressionSSASTContext(SimpleStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterExpressionSSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitExpressionSSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitExpressionSSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpListSSASTContext extends SimpleStatementContext {
		public List<ExpressionListContext> expressionList() {
			return getRuleContexts(ExpressionListContext.class);
		}
		public ExpressionListContext expressionList(int i) {
			return getRuleContext(ExpressionListContext.class,i);
		}
		public TerminalNode SVD() { return getToken(goParser.SVD, 0); }
		public ExpListSSASTContext(SimpleStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterExpListSSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitExpListSSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitExpListSSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class EpsilonSSASTContext extends SimpleStatementContext {
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public EpsilonSSASTContext(SimpleStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterEpsilonSSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitEpsilonSSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitEpsilonSSAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SimpleStatementContext simpleStatement() throws RecognitionException {
		SimpleStatementContext _localctx = new SimpleStatementContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_simpleStatement);
		try {
			setState(390);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				_localctx = new EpsilonSSASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(378);
				epsilon();
				}
				break;
			case 2:
				_localctx = new ExpressionSSASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(379);
				expression(0);
				setState(383);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case INC:
					{
					setState(380);
					match(INC);
					}
					break;
				case DEC:
					{
					setState(381);
					match(DEC);
					}
					break;
				case SEMI:
				case LCB:
					{
					setState(382);
					epsilon();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				break;
			case 3:
				_localctx = new AssigmentStatementSSASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(385);
				assignmentStatement();
				}
				break;
			case 4:
				_localctx = new ExpListSSASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(386);
				expressionList();
				setState(387);
				match(SVD);
				setState(388);
				expressionList();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentStatementContext extends ParserRuleContext {
		public AssignmentStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentStatement; }
	 
		public AssignmentStatementContext() { }
		public void copyFrom(AssignmentStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpListASASTContext extends AssignmentStatementContext {
		public List<ExpressionListContext> expressionList() {
			return getRuleContexts(ExpressionListContext.class);
		}
		public ExpressionListContext expressionList(int i) {
			return getRuleContext(ExpressionListContext.class,i);
		}
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public ExpListASASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterExpListASAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitExpListASAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitExpListASAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpASASTContext extends AssignmentStatementContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public OperatorContext operator() {
			return getRuleContext(OperatorContext.class,0);
		}
		public ExpASASTContext(AssignmentStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterExpASAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitExpASAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitExpASAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AssignmentStatementContext assignmentStatement() throws RecognitionException {
		AssignmentStatementContext _localctx = new AssignmentStatementContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_assignmentStatement);
		try {
			setState(400);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				_localctx = new ExpListASASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(392);
				expressionList();
				setState(393);
				operator();
				setState(394);
				expressionList();
				}
				break;
			case 2:
				_localctx = new ExpASASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(396);
				expression(0);
				setState(397);
				operator();
				setState(398);
				expression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	 
		public IfStatementContext() { }
		public void copyFrom(IfStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IsExpBlockIsISASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(goParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(goParser.ELSE, 0); }
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public IsExpBlockIsISASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIsExpBlockIsISAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIsExpBlockIsISAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIsExpBlockIsISAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IsSSExpBlockISASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(goParser.IF, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IsSSExpBlockISASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIsSSExpBlockISAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIsSSExpBlockISAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIsSSExpBlockISAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IsExpBlockISASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(goParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(goParser.ELSE, 0); }
		public IsExpBlockISASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIsExpBlockISAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIsExpBlockISAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIsExpBlockISAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IsSSExpBlockBlockASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(goParser.IF, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(goParser.ELSE, 0); }
		public IsSSExpBlockBlockASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIsSSExpBlockBlockAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIsSSExpBlockBlockAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIsSSExpBlockBlockAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IsExpressionBlockISASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(goParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IsExpressionBlockISASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIsExpressionBlockISAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIsExpressionBlockISAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIsExpressionBlockISAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IsSSExpBlockifSASTContext extends IfStatementContext {
		public TerminalNode IF() { return getToken(goParser.IF, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(goParser.ELSE, 0); }
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public IsSSExpBlockifSASTContext(IfStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterIsSSExpBlockifSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitIsSSExpBlockifSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitIsSSExpBlockifSAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_ifStatement);
		try {
			setState(440);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				_localctx = new IsExpressionBlockISASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(402);
				match(IF);
				setState(403);
				expression(0);
				setState(404);
				block();
				}
				break;
			case 2:
				_localctx = new IsExpBlockIsISASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(406);
				match(IF);
				setState(407);
				expression(0);
				setState(408);
				block();
				setState(409);
				match(ELSE);
				setState(410);
				ifStatement();
				}
				break;
			case 3:
				_localctx = new IsExpBlockISASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(412);
				match(IF);
				setState(413);
				expression(0);
				setState(414);
				block();
				setState(415);
				match(ELSE);
				setState(416);
				block();
				}
				break;
			case 4:
				_localctx = new IsSSExpBlockISASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(418);
				match(IF);
				setState(419);
				simpleStatement();
				setState(420);
				match(SEMI);
				setState(421);
				expression(0);
				setState(422);
				block();
				}
				break;
			case 5:
				_localctx = new IsSSExpBlockifSASTContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(424);
				match(IF);
				setState(425);
				simpleStatement();
				setState(426);
				match(SEMI);
				setState(427);
				expression(0);
				setState(428);
				block();
				setState(429);
				match(ELSE);
				setState(430);
				ifStatement();
				}
				break;
			case 6:
				_localctx = new IsSSExpBlockBlockASTContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(432);
				match(IF);
				setState(433);
				simpleStatement();
				setState(434);
				match(SEMI);
				setState(435);
				expression(0);
				setState(436);
				block();
				setState(437);
				match(ELSE);
				setState(438);
				block();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LoopContext extends ParserRuleContext {
		public LoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loop; }
	 
		public LoopContext() { }
		public void copyFrom(LoopContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FExpBlockLASTContext extends LoopContext {
		public TerminalNode FOR() { return getToken(goParser.FOR, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FExpBlockLASTContext(LoopContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterFExpBlockLAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitFExpBlockLAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitFExpBlockLAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FSimpStateSimpStateBlockLASTContext extends LoopContext {
		public TerminalNode FOR() { return getToken(goParser.FOR, 0); }
		public List<SimpleStatementContext> simpleStatement() {
			return getRuleContexts(SimpleStatementContext.class);
		}
		public SimpleStatementContext simpleStatement(int i) {
			return getRuleContext(SimpleStatementContext.class,i);
		}
		public List<TerminalNode> SEMI() { return getTokens(goParser.SEMI); }
		public TerminalNode SEMI(int i) {
			return getToken(goParser.SEMI, i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FSimpStateSimpStateBlockLASTContext(LoopContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterFSimpStateSimpStateBlockLAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitFSimpStateSimpStateBlockLAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitFSimpStateSimpStateBlockLAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FBlockLASTContext extends LoopContext {
		public TerminalNode FOR() { return getToken(goParser.FOR, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FBlockLASTContext(LoopContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterFBlockLAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitFBlockLAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitFBlockLAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FSimpStateExpSBlockLASTContext extends LoopContext {
		public TerminalNode FOR() { return getToken(goParser.FOR, 0); }
		public List<SimpleStatementContext> simpleStatement() {
			return getRuleContexts(SimpleStatementContext.class);
		}
		public SimpleStatementContext simpleStatement(int i) {
			return getRuleContext(SimpleStatementContext.class,i);
		}
		public List<TerminalNode> SEMI() { return getTokens(goParser.SEMI); }
		public TerminalNode SEMI(int i) {
			return getToken(goParser.SEMI, i);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FSimpStateExpSBlockLASTContext(LoopContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterFSimpStateExpSBlockLAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitFSimpStateExpSBlockLAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitFSimpStateExpSBlockLAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LoopContext loop() throws RecognitionException {
		LoopContext _localctx = new LoopContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_loop);
		try {
			setState(463);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				_localctx = new FBlockLASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(442);
				match(FOR);
				setState(443);
				block();
				}
				break;
			case 2:
				_localctx = new FExpBlockLASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(444);
				match(FOR);
				setState(445);
				expression(0);
				setState(446);
				block();
				}
				break;
			case 3:
				_localctx = new FSimpStateExpSBlockLASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(448);
				match(FOR);
				setState(449);
				simpleStatement();
				setState(450);
				match(SEMI);
				setState(451);
				expression(0);
				setState(452);
				match(SEMI);
				setState(453);
				simpleStatement();
				setState(454);
				block();
				}
				break;
			case 4:
				_localctx = new FSimpStateSimpStateBlockLASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(456);
				match(FOR);
				setState(457);
				simpleStatement();
				setState(458);
				match(SEMI);
				setState(459);
				match(SEMI);
				setState(460);
				simpleStatement();
				setState(461);
				block();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchContext extends ParserRuleContext {
		public SwitchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch; }
	 
		public SwitchContext() { }
		public void copyFrom(SwitchContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SSimpleStatExpCCListSASTContext extends SwitchContext {
		public TerminalNode SWITCH() { return getToken(goParser.SWITCH, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public TerminalNode LCB() { return getToken(goParser.LCB, 0); }
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public TerminalNode RCB() { return getToken(goParser.RCB, 0); }
		public SSimpleStatExpCCListSASTContext(SwitchContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSSimpleStatExpCCListSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSSimpleStatExpCCListSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSSimpleStatExpCCListSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SExpressionSASTContext extends SwitchContext {
		public TerminalNode SWITCH() { return getToken(goParser.SWITCH, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LCB() { return getToken(goParser.LCB, 0); }
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public TerminalNode RCB() { return getToken(goParser.RCB, 0); }
		public SExpressionSASTContext(SwitchContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSExpressionSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSExpressionSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSExpressionSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SBlockSASTContext extends SwitchContext {
		public TerminalNode SWITCH() { return getToken(goParser.SWITCH, 0); }
		public TerminalNode LCB() { return getToken(goParser.LCB, 0); }
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public TerminalNode RCB() { return getToken(goParser.RCB, 0); }
		public SBlockSASTContext(SwitchContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSBlockSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSBlockSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSBlockSAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SSimpleStatSASTContext extends SwitchContext {
		public TerminalNode SWITCH() { return getToken(goParser.SWITCH, 0); }
		public SimpleStatementContext simpleStatement() {
			return getRuleContext(SimpleStatementContext.class,0);
		}
		public TerminalNode SEMI() { return getToken(goParser.SEMI, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LCB() { return getToken(goParser.LCB, 0); }
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public TerminalNode RCB() { return getToken(goParser.RCB, 0); }
		public SSimpleStatSASTContext(SwitchContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterSSimpleStatSAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitSSimpleStatSAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitSSimpleStatSAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SwitchContext switch_() throws RecognitionException {
		SwitchContext _localctx = new SwitchContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_switch);
		try {
			setState(491);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				_localctx = new SSimpleStatSASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(465);
				match(SWITCH);
				setState(466);
				simpleStatement();
				setState(467);
				match(SEMI);
				setState(468);
				expression(0);
				setState(469);
				match(LCB);
				setState(470);
				expressionCaseClauseList();
				setState(471);
				match(RCB);
				}
				break;
			case 2:
				_localctx = new SExpressionSASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(473);
				match(SWITCH);
				setState(474);
				expression(0);
				setState(475);
				match(LCB);
				setState(476);
				expressionCaseClauseList();
				setState(477);
				match(RCB);
				}
				break;
			case 3:
				_localctx = new SSimpleStatExpCCListSASTContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(479);
				match(SWITCH);
				setState(480);
				simpleStatement();
				setState(481);
				match(SEMI);
				setState(482);
				match(LCB);
				setState(483);
				expressionCaseClauseList();
				setState(484);
				match(RCB);
				}
				break;
			case 4:
				_localctx = new SBlockSASTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(486);
				match(SWITCH);
				setState(487);
				match(LCB);
				setState(488);
				expressionCaseClauseList();
				setState(489);
				match(RCB);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionCaseClauseListContext extends ParserRuleContext {
		public ExpressionCaseClauseListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionCaseClauseList; }
	 
		public ExpressionCaseClauseListContext() { }
		public void copyFrom(ExpressionCaseClauseListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpCaseClauseECCLASTContext extends ExpressionCaseClauseListContext {
		public ExpressionCaseClauseContext expressionCaseClause() {
			return getRuleContext(ExpressionCaseClauseContext.class,0);
		}
		public ExpressionCaseClauseListContext expressionCaseClauseList() {
			return getRuleContext(ExpressionCaseClauseListContext.class,0);
		}
		public ExpCaseClauseECCLASTContext(ExpressionCaseClauseListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterExpCaseClauseECCLAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitExpCaseClauseECCLAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitExpCaseClauseECCLAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class EpsilonECCLASTContext extends ExpressionCaseClauseListContext {
		public EpsilonContext epsilon() {
			return getRuleContext(EpsilonContext.class,0);
		}
		public EpsilonECCLASTContext(ExpressionCaseClauseListContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterEpsilonECCLAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitEpsilonECCLAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitEpsilonECCLAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionCaseClauseListContext expressionCaseClauseList() throws RecognitionException {
		ExpressionCaseClauseListContext _localctx = new ExpressionCaseClauseListContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_expressionCaseClauseList);
		try {
			setState(497);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RCB:
				_localctx = new EpsilonECCLASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(493);
				epsilon();
				}
				break;
			case CASE:
			case DEFAULT:
				_localctx = new ExpCaseClauseECCLASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(494);
				expressionCaseClause();
				setState(495);
				expressionCaseClauseList();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionCaseClauseContext extends ParserRuleContext {
		public ExpressionCaseClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionCaseClause; }
	 
		public ExpressionCaseClauseContext() { }
		public void copyFrom(ExpressionCaseClauseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpSwitchCaseECCASTContext extends ExpressionCaseClauseContext {
		public ExpressionSwitchCaseContext expressionSwitchCase() {
			return getRuleContext(ExpressionSwitchCaseContext.class,0);
		}
		public TerminalNode COL() { return getToken(goParser.COL, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public ExpSwitchCaseECCASTContext(ExpressionCaseClauseContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterExpSwitchCaseECCAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitExpSwitchCaseECCAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitExpSwitchCaseECCAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionCaseClauseContext expressionCaseClause() throws RecognitionException {
		ExpressionCaseClauseContext _localctx = new ExpressionCaseClauseContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_expressionCaseClause);
		try {
			_localctx = new ExpSwitchCaseECCASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(499);
			expressionSwitchCase();
			setState(500);
			match(COL);
			setState(501);
			statementList();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionSwitchCaseContext extends ParserRuleContext {
		public ExpressionSwitchCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionSwitchCase; }
	 
		public ExpressionSwitchCaseContext() { }
		public void copyFrom(ExpressionSwitchCaseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DefaultESCASTContext extends ExpressionSwitchCaseContext {
		public TerminalNode DEFAULT() { return getToken(goParser.DEFAULT, 0); }
		public DefaultESCASTContext(ExpressionSwitchCaseContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterDefaultESCAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitDefaultESCAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitDefaultESCAST(this);
			else return visitor.visitChildren(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CaseESCASTContext extends ExpressionSwitchCaseContext {
		public TerminalNode CASE() { return getToken(goParser.CASE, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public CaseESCASTContext(ExpressionSwitchCaseContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterCaseESCAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitCaseESCAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitCaseESCAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ExpressionSwitchCaseContext expressionSwitchCase() throws RecognitionException {
		ExpressionSwitchCaseContext _localctx = new ExpressionSwitchCaseContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_expressionSwitchCase);
		try {
			setState(506);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CASE:
				_localctx = new CaseESCASTContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(503);
				match(CASE);
				setState(504);
				expressionList();
				}
				break;
			case DEFAULT:
				_localctx = new DefaultESCASTContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(505);
				match(DEFAULT);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperatorContext extends ParserRuleContext {
		public TerminalNode MUL() { return getToken(goParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(goParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(goParser.MOD, 0); }
		public TerminalNode LSH() { return getToken(goParser.LSH, 0); }
		public TerminalNode RSH() { return getToken(goParser.RSH, 0); }
		public TerminalNode AMPER() { return getToken(goParser.AMPER, 0); }
		public TerminalNode BC() { return getToken(goParser.BC, 0); }
		public TerminalNode PL() { return getToken(goParser.PL, 0); }
		public TerminalNode MIN() { return getToken(goParser.MIN, 0); }
		public TerminalNode VB() { return getToken(goParser.VB, 0); }
		public TerminalNode CARET() { return getToken(goParser.CARET, 0); }
		public TerminalNode EQ() { return getToken(goParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(goParser.NEQ, 0); }
		public TerminalNode LT() { return getToken(goParser.LT, 0); }
		public TerminalNode LTE() { return getToken(goParser.LTE, 0); }
		public TerminalNode GT() { return getToken(goParser.GT, 0); }
		public TerminalNode GTE() { return getToken(goParser.GTE, 0); }
		public TerminalNode LAND() { return getToken(goParser.LAND, 0); }
		public TerminalNode LOR() { return getToken(goParser.LOR, 0); }
		public TerminalNode NEG() { return getToken(goParser.NEG, 0); }
		public TerminalNode ASOP() { return getToken(goParser.ASOP, 0); }
		public TerminalNode AAOP() { return getToken(goParser.AAOP, 0); }
		public TerminalNode BAOP() { return getToken(goParser.BAOP, 0); }
		public TerminalNode SAOP() { return getToken(goParser.SAOP, 0); }
		public TerminalNode BOAOP() { return getToken(goParser.BOAOP, 0); }
		public TerminalNode MAOP() { return getToken(goParser.MAOP, 0); }
		public TerminalNode BXOOP() { return getToken(goParser.BXOOP, 0); }
		public TerminalNode LSAOP() { return getToken(goParser.LSAOP, 0); }
		public TerminalNode RSAOP() { return getToken(goParser.RSAOP, 0); }
		public TerminalNode BCAOP() { return getToken(goParser.BCAOP, 0); }
		public TerminalNode RAOP() { return getToken(goParser.RAOP, 0); }
		public TerminalNode DAOP() { return getToken(goParser.DAOP, 0); }
		public OperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterOperator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitOperator(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitOperator(this);
			else return visitor.visitChildren(this);
		}
	}

	public final OperatorContext operator() throws RecognitionException {
		OperatorContext _localctx = new OperatorContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(508);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 70368743917440L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EpsilonContext extends ParserRuleContext {
		public EpsilonContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_epsilon; }
	 
		public EpsilonContext() { }
		public void copyFrom(EpsilonContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class EpsilonASTContext extends EpsilonContext {
		public EpsilonASTContext(EpsilonContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).enterEpsilonAST(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof goParserListener ) ((goParserListener)listener).exitEpsilonAST(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof goParserVisitor ) return ((goParserVisitor<? extends T>)visitor).visitEpsilonAST(this);
			else return visitor.visitChildren(this);
		}
	}

	public final EpsilonContext epsilon() throws RecognitionException {
		EpsilonContext _localctx = new EpsilonContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_epsilon);
		try {
			_localctx = new EpsilonASTContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 18:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 20:
			return primaryExpression_sempred((PrimaryExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean primaryExpression_sempred(PrimaryExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001U\u0201\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0005\u0001]\b\u0001"+
		"\n\u0001\f\u0001`\t\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u0002n\b\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003u\b\u0003\n\u0003\f\u0003"+
		"x\t\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004"+
		"\u0084\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u0095\b\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007"+
		"\u009c\b\u0007\n\u0007\f\u0007\u009f\t\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003"+
		"\n\u00ad\b\n\u0001\n\u0001\n\u0001\n\u0003\n\u00b2\b\n\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0005\u000b\u00b7\b\u000b\n\u000b\f\u000b\u00ba\t\u000b"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003"+
		"\f\u00c4\b\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0003\u000f\u00d3\b\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u00dc\b\u0010\n"+
		"\u0010\f\u0010\u00df\t\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0005"+
		"\u0011\u00e4\b\u0011\n\u0011\f\u0011\u00e7\t\u0011\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u00ee\b\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0005\u0012\u00f4\b\u0012\n\u0012"+
		"\f\u0012\u00f7\t\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013"+
		"\u00fc\b\u0013\n\u0013\f\u0013\u00ff\t\u0013\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u0106\b\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u010c\b\u0014\u0005\u0014\u010e"+
		"\b\u0014\n\u0014\f\u0014\u0111\t\u0014\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u0119\b\u0015\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u0120\b\u0016"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0003\u0018\u0129\b\u0018\u0001\u0018\u0001\u0018\u0001\u0019"+
		"\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001c\u0001\u001d\u0005\u001d\u0142\b\u001d\n\u001d\f\u001d\u0145"+
		"\t\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u014f\b\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0003"+
		"\u001f\u0158\b\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0003\u001f\u0160\b\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u0179\b\u001f\u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0003 \u0180\b \u0001 \u0001 \u0001 \u0001 \u0001"+
		" \u0003 \u0187\b \u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0003!\u0191\b!\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0003\"\u01b9\b\"\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0003#\u01d0"+
		"\b#\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0003$\u01ec\b$\u0001%\u0001"+
		"%\u0001%\u0001%\u0003%\u01f2\b%\u0001&\u0001&\u0001&\u0001&\u0001\'\u0001"+
		"\'\u0001\'\u0003\'\u01fb\b\'\u0001(\u0001(\u0001)\u0001)\u0001)\u0000"+
		"\u0002$(*\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016"+
		"\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLNPR\u0000\u0001\u0002"+
		"\u0000\u0007\n\u0012-\u021d\u0000T\u0001\u0000\u0000\u0000\u0002^\u0001"+
		"\u0000\u0000\u0000\u0004m\u0001\u0000\u0000\u0000\u0006o\u0001\u0000\u0000"+
		"\u0000\b\u0083\u0001\u0000\u0000\u0000\n\u0085\u0001\u0000\u0000\u0000"+
		"\f\u0088\u0001\u0000\u0000\u0000\u000e\u0096\u0001\u0000\u0000\u0000\u0010"+
		"\u00a0\u0001\u0000\u0000\u0000\u0012\u00a3\u0001\u0000\u0000\u0000\u0014"+
		"\u00a7\u0001\u0000\u0000\u0000\u0016\u00b3\u0001\u0000\u0000\u0000\u0018"+
		"\u00c3\u0001\u0000\u0000\u0000\u001a\u00c5\u0001\u0000\u0000\u0000\u001c"+
		"\u00c9\u0001\u0000\u0000\u0000\u001e\u00ce\u0001\u0000\u0000\u0000 \u00d6"+
		"\u0001\u0000\u0000\u0000\"\u00e0\u0001\u0000\u0000\u0000$\u00ed\u0001"+
		"\u0000\u0000\u0000&\u00f8\u0001\u0000\u0000\u0000(\u0105\u0001\u0000\u0000"+
		"\u0000*\u0118\u0001\u0000\u0000\u0000,\u011f\u0001\u0000\u0000\u0000."+
		"\u0121\u0001\u0000\u0000\u00000\u0125\u0001\u0000\u0000\u00002\u012c\u0001"+
		"\u0000\u0000\u00004\u012f\u0001\u0000\u0000\u00006\u0136\u0001\u0000\u0000"+
		"\u00008\u013b\u0001\u0000\u0000\u0000:\u0143\u0001\u0000\u0000\u0000<"+
		"\u0146\u0001\u0000\u0000\u0000>\u0178\u0001\u0000\u0000\u0000@\u0186\u0001"+
		"\u0000\u0000\u0000B\u0190\u0001\u0000\u0000\u0000D\u01b8\u0001\u0000\u0000"+
		"\u0000F\u01cf\u0001\u0000\u0000\u0000H\u01eb\u0001\u0000\u0000\u0000J"+
		"\u01f1\u0001\u0000\u0000\u0000L\u01f3\u0001\u0000\u0000\u0000N\u01fa\u0001"+
		"\u0000\u0000\u0000P\u01fc\u0001\u0000\u0000\u0000R\u01fe\u0001\u0000\u0000"+
		"\u0000TU\u00051\u0000\u0000UV\u0005L\u0000\u0000VW\u0005\u0001\u0000\u0000"+
		"WX\u0003\u0002\u0001\u0000X\u0001\u0001\u0000\u0000\u0000Y]\u0003\u0004"+
		"\u0002\u0000Z]\u0003\f\u0006\u0000[]\u0003\u0012\t\u0000\\Y\u0001\u0000"+
		"\u0000\u0000\\Z\u0001\u0000\u0000\u0000\\[\u0001\u0000\u0000\u0000]`\u0001"+
		"\u0000\u0000\u0000^\\\u0001\u0000\u0000\u0000^_\u0001\u0000\u0000\u0000"+
		"_\u0003\u0001\u0000\u0000\u0000`^\u0001\u0000\u0000\u0000ab\u0005<\u0000"+
		"\u0000bc\u0003\b\u0004\u0000cd\u0005\u0001\u0000\u0000dn\u0001\u0000\u0000"+
		"\u0000ef\u0005\u0003\u0000\u0000fg\u0003\u0006\u0003\u0000gh\u0005\u0004"+
		"\u0000\u0000hi\u0005\u0001\u0000\u0000in\u0001\u0000\u0000\u0000jk\u0005"+
		"\u0003\u0000\u0000kl\u0005\u0004\u0000\u0000ln\u0005\u0001\u0000\u0000"+
		"ma\u0001\u0000\u0000\u0000me\u0001\u0000\u0000\u0000mj\u0001\u0000\u0000"+
		"\u0000n\u0005\u0001\u0000\u0000\u0000op\u0003\b\u0004\u0000pv\u0005\u0001"+
		"\u0000\u0000qr\u0003\b\u0004\u0000rs\u0005\u0001\u0000\u0000su\u0001\u0000"+
		"\u0000\u0000tq\u0001\u0000\u0000\u0000ux\u0001\u0000\u0000\u0000vt\u0001"+
		"\u0000\u0000\u0000vw\u0001\u0000\u0000\u0000w\u0007\u0001\u0000\u0000"+
		"\u0000xv\u0001\u0000\u0000\u0000yz\u0003\"\u0011\u0000z{\u0003\u0018\f"+
		"\u0000{|\u0005\"\u0000\u0000|}\u0003&\u0013\u0000}\u0084\u0001\u0000\u0000"+
		"\u0000~\u007f\u0003\"\u0011\u0000\u007f\u0080\u0005\"\u0000\u0000\u0080"+
		"\u0081\u0003&\u0013\u0000\u0081\u0084\u0001\u0000\u0000\u0000\u0082\u0084"+
		"\u0003\n\u0005\u0000\u0083y\u0001\u0000\u0000\u0000\u0083~\u0001\u0000"+
		"\u0000\u0000\u0083\u0082\u0001\u0000\u0000\u0000\u0084\t\u0001\u0000\u0000"+
		"\u0000\u0085\u0086\u0003\"\u0011\u0000\u0086\u0087\u0003\u0018\f\u0000"+
		"\u0087\u000b\u0001\u0000\u0000\u0000\u0088\u0094\u0005=\u0000\u0000\u0089"+
		"\u008a\u0003\u0010\b\u0000\u008a\u008b\u0005\u0001\u0000\u0000\u008b\u0095"+
		"\u0001\u0000\u0000\u0000\u008c\u008d\u0005\u0003\u0000\u0000\u008d\u008e"+
		"\u0003\u000e\u0007\u0000\u008e\u008f\u0005\u0004\u0000\u0000\u008f\u0090"+
		"\u0005\u0001\u0000\u0000\u0090\u0095\u0001\u0000\u0000\u0000\u0091\u0092"+
		"\u0005\u0003\u0000\u0000\u0092\u0093\u0005\u0004\u0000\u0000\u0093\u0095"+
		"\u0005\u0001\u0000\u0000\u0094\u0089\u0001\u0000\u0000\u0000\u0094\u008c"+
		"\u0001\u0000\u0000\u0000\u0094\u0091\u0001\u0000\u0000\u0000\u0095\r\u0001"+
		"\u0000\u0000\u0000\u0096\u0097\u0003\u0010\b\u0000\u0097\u009d\u0005\u0001"+
		"\u0000\u0000\u0098\u0099\u0003\u0010\b\u0000\u0099\u009a\u0005\u0001\u0000"+
		"\u0000\u009a\u009c\u0001\u0000\u0000\u0000\u009b\u0098\u0001\u0000\u0000"+
		"\u0000\u009c\u009f\u0001\u0000\u0000\u0000\u009d\u009b\u0001\u0000\u0000"+
		"\u0000\u009d\u009e\u0001\u0000\u0000\u0000\u009e\u000f\u0001\u0000\u0000"+
		"\u0000\u009f\u009d\u0001\u0000\u0000\u0000\u00a0\u00a1\u0005L\u0000\u0000"+
		"\u00a1\u00a2\u0003\u0018\f\u0000\u00a2\u0011\u0001\u0000\u0000\u0000\u00a3"+
		"\u00a4\u0003\u0014\n\u0000\u00a4\u00a5\u0003<\u001e\u0000\u00a5\u00a6"+
		"\u0005\u0001\u0000\u0000\u00a6\u0013\u0001\u0000\u0000\u0000\u00a7\u00a8"+
		"\u0005>\u0000\u0000\u00a8\u00a9\u0005L\u0000\u0000\u00a9\u00ac\u0005\u0003"+
		"\u0000\u0000\u00aa\u00ad\u0003\u0016\u000b\u0000\u00ab\u00ad\u0003R)\u0000"+
		"\u00ac\u00aa\u0001\u0000\u0000\u0000\u00ac\u00ab\u0001\u0000\u0000\u0000"+
		"\u00ad\u00ae\u0001\u0000\u0000\u0000\u00ae\u00b1\u0005\u0004\u0000\u0000"+
		"\u00af\u00b2\u0003\u0018\f\u0000\u00b0\u00b2\u0003R)\u0000\u00b1\u00af"+
		"\u0001\u0000\u0000\u0000\u00b1\u00b0\u0001\u0000\u0000\u0000\u00b2\u0015"+
		"\u0001\u0000\u0000\u0000\u00b3\u00b8\u0003\n\u0005\u0000\u00b4\u00b5\u0005"+
		"\u000b\u0000\u0000\u00b5\u00b7\u0003\n\u0005\u0000\u00b6\u00b4\u0001\u0000"+
		"\u0000\u0000\u00b7\u00ba\u0001\u0000\u0000\u0000\u00b8\u00b6\u0001\u0000"+
		"\u0000\u0000\u00b8\u00b9\u0001\u0000\u0000\u0000\u00b9\u0017\u0001\u0000"+
		"\u0000\u0000\u00ba\u00b8\u0001\u0000\u0000\u0000\u00bb\u00bc\u0005\u0003"+
		"\u0000\u0000\u00bc\u00bd\u0003\u0018\f\u0000\u00bd\u00be\u0005\u0004\u0000"+
		"\u0000\u00be\u00c4\u0001\u0000\u0000\u0000\u00bf\u00c4\u0005L\u0000\u0000"+
		"\u00c0\u00c4\u0003\u001a\r\u0000\u00c1\u00c4\u0003\u001c\u000e\u0000\u00c2"+
		"\u00c4\u0003\u001e\u000f\u0000\u00c3\u00bb\u0001\u0000\u0000\u0000\u00c3"+
		"\u00bf\u0001\u0000\u0000\u0000\u00c3\u00c0\u0001\u0000\u0000\u0000\u00c3"+
		"\u00c1\u0001\u0000\u0000\u0000\u00c3\u00c2\u0001\u0000\u0000\u0000\u00c4"+
		"\u0019\u0001\u0000\u0000\u0000\u00c5\u00c6\u0005\f\u0000\u0000\u00c6\u00c7"+
		"\u0005\r\u0000\u0000\u00c7\u00c8\u0003\u0018\f\u0000\u00c8\u001b\u0001"+
		"\u0000\u0000\u0000\u00c9\u00ca\u0005\f\u0000\u0000\u00ca\u00cb\u0005M"+
		"\u0000\u0000\u00cb\u00cc\u0005\r\u0000\u0000\u00cc\u00cd\u0003\u0018\f"+
		"\u0000\u00cd\u001d\u0001\u0000\u0000\u0000\u00ce\u00cf\u0005?\u0000\u0000"+
		"\u00cf\u00d2\u0005\u000e\u0000\u0000\u00d0\u00d3\u0003 \u0010\u0000\u00d1"+
		"\u00d3\u0003R)\u0000\u00d2\u00d0\u0001\u0000\u0000\u0000\u00d2\u00d1\u0001"+
		"\u0000\u0000\u0000\u00d3\u00d4\u0001\u0000\u0000\u0000\u00d4\u00d5\u0005"+
		"\u000f\u0000\u0000\u00d5\u001f\u0001\u0000\u0000\u0000\u00d6\u00d7\u0003"+
		"\n\u0005\u0000\u00d7\u00dd\u0005\u0001\u0000\u0000\u00d8\u00d9\u0003\n"+
		"\u0005\u0000\u00d9\u00da\u0005\u0001\u0000\u0000\u00da\u00dc\u0001\u0000"+
		"\u0000\u0000\u00db\u00d8\u0001\u0000\u0000\u0000\u00dc\u00df\u0001\u0000"+
		"\u0000\u0000\u00dd\u00db\u0001\u0000\u0000\u0000\u00dd\u00de\u0001\u0000"+
		"\u0000\u0000\u00de!\u0001\u0000\u0000\u0000\u00df\u00dd\u0001\u0000\u0000"+
		"\u0000\u00e0\u00e5\u0005L\u0000\u0000\u00e1\u00e2\u0005\u000b\u0000\u0000"+
		"\u00e2\u00e4\u0005L\u0000\u0000\u00e3\u00e1\u0001\u0000\u0000\u0000\u00e4"+
		"\u00e7\u0001\u0000\u0000\u0000\u00e5\u00e3\u0001\u0000\u0000\u0000\u00e5"+
		"\u00e6\u0001\u0000\u0000\u0000\u00e6#\u0001\u0000\u0000\u0000\u00e7\u00e5"+
		"\u0001\u0000\u0000\u0000\u00e8\u00e9\u0006\u0012\uffff\uffff\u0000\u00e9"+
		"\u00ee\u0003(\u0014\u0000\u00ea\u00eb\u0003P(\u0000\u00eb\u00ec\u0003"+
		"$\u0012\u0001\u00ec\u00ee\u0001\u0000\u0000\u0000\u00ed\u00e8\u0001\u0000"+
		"\u0000\u0000\u00ed\u00ea\u0001\u0000\u0000\u0000\u00ee\u00f5\u0001\u0000"+
		"\u0000\u0000\u00ef\u00f0\n\u0002\u0000\u0000\u00f0\u00f1\u0003P(\u0000"+
		"\u00f1\u00f2\u0003$\u0012\u0003\u00f2\u00f4\u0001\u0000\u0000\u0000\u00f3"+
		"\u00ef\u0001\u0000\u0000\u0000\u00f4\u00f7\u0001\u0000\u0000\u0000\u00f5"+
		"\u00f3\u0001\u0000\u0000\u0000\u00f5\u00f6\u0001\u0000\u0000\u0000\u00f6"+
		"%\u0001\u0000\u0000\u0000\u00f7\u00f5\u0001\u0000\u0000\u0000\u00f8\u00fd"+
		"\u0003$\u0012\u0000\u00f9\u00fa\u0005\u000b\u0000\u0000\u00fa\u00fc\u0003"+
		"$\u0012\u0000\u00fb\u00f9\u0001\u0000\u0000\u0000\u00fc\u00ff\u0001\u0000"+
		"\u0000\u0000\u00fd\u00fb\u0001\u0000\u0000\u0000\u00fd\u00fe\u0001\u0000"+
		"\u0000\u0000\u00fe\'\u0001\u0000\u0000\u0000\u00ff\u00fd\u0001\u0000\u0000"+
		"\u0000\u0100\u0101\u0006\u0014\uffff\uffff\u0000\u0101\u0106\u0003*\u0015"+
		"\u0000\u0102\u0106\u00034\u001a\u0000\u0103\u0106\u00036\u001b\u0000\u0104"+
		"\u0106\u00038\u001c\u0000\u0105\u0100\u0001\u0000\u0000\u0000\u0105\u0102"+
		"\u0001\u0000\u0000\u0000\u0105\u0103\u0001\u0000\u0000\u0000\u0105\u0104"+
		"\u0001\u0000\u0000\u0000\u0106\u010f\u0001\u0000\u0000\u0000\u0107\u010b"+
		"\n\u0004\u0000\u0000\u0108\u010c\u00032\u0019\u0000\u0109\u010c\u0003"+
		".\u0017\u0000\u010a\u010c\u00030\u0018\u0000\u010b\u0108\u0001\u0000\u0000"+
		"\u0000\u010b\u0109\u0001\u0000\u0000\u0000\u010b\u010a\u0001\u0000\u0000"+
		"\u0000\u010c\u010e\u0001\u0000\u0000\u0000\u010d\u0107\u0001\u0000\u0000"+
		"\u0000\u010e\u0111\u0001\u0000\u0000\u0000\u010f\u010d\u0001\u0000\u0000"+
		"\u0000\u010f\u0110\u0001\u0000\u0000\u0000\u0110)\u0001\u0000\u0000\u0000"+
		"\u0111\u010f\u0001\u0000\u0000\u0000\u0112\u0119\u0003,\u0016\u0000\u0113"+
		"\u0119\u0005L\u0000\u0000\u0114\u0115\u0005\u0003\u0000\u0000\u0115\u0116"+
		"\u0003$\u0012\u0000\u0116\u0117\u0005\u0004\u0000\u0000\u0117\u0119\u0001"+
		"\u0000\u0000\u0000\u0118\u0112\u0001\u0000\u0000\u0000\u0118\u0113\u0001"+
		"\u0000\u0000\u0000\u0118\u0114\u0001\u0000\u0000\u0000\u0119+\u0001\u0000"+
		"\u0000\u0000\u011a\u0120\u0005M\u0000\u0000\u011b\u0120\u0005N\u0000\u0000"+
		"\u011c\u0120\u0005O\u0000\u0000\u011d\u0120\u0005P\u0000\u0000\u011e\u0120"+
		"\u0005Q\u0000\u0000\u011f\u011a\u0001\u0000\u0000\u0000\u011f\u011b\u0001"+
		"\u0000\u0000\u0000\u011f\u011c\u0001\u0000\u0000\u0000\u011f\u011d\u0001"+
		"\u0000\u0000\u0000\u011f\u011e\u0001\u0000\u0000\u0000\u0120-\u0001\u0000"+
		"\u0000\u0000\u0121\u0122\u0005\f\u0000\u0000\u0122\u0123\u0003$\u0012"+
		"\u0000\u0123\u0124\u0005\r\u0000\u0000\u0124/\u0001\u0000\u0000\u0000"+
		"\u0125\u0128\u0005\u0003\u0000\u0000\u0126\u0129\u0003&\u0013\u0000\u0127"+
		"\u0129\u0003R)\u0000\u0128\u0126\u0001\u0000\u0000\u0000\u0128\u0127\u0001"+
		"\u0000\u0000\u0000\u0129\u012a\u0001\u0000\u0000\u0000\u012a\u012b\u0005"+
		"\u0004\u0000\u0000\u012b1\u0001\u0000\u0000\u0000\u012c\u012d\u0005.\u0000"+
		"\u0000\u012d\u012e\u0005L\u0000\u0000\u012e3\u0001\u0000\u0000\u0000\u012f"+
		"\u0130\u0005@\u0000\u0000\u0130\u0131\u0005\u0003\u0000\u0000\u0131\u0132"+
		"\u0003$\u0012\u0000\u0132\u0133\u0005\u000b\u0000\u0000\u0133\u0134\u0003"+
		"$\u0012\u0000\u0134\u0135\u0005\u0004\u0000\u0000\u01355\u0001\u0000\u0000"+
		"\u0000\u0136\u0137\u0005A\u0000\u0000\u0137\u0138\u0005\u0003\u0000\u0000"+
		"\u0138\u0139\u0003$\u0012\u0000\u0139\u013a\u0005\u0004\u0000\u0000\u013a"+
		"7\u0001\u0000\u0000\u0000\u013b\u013c\u0005B\u0000\u0000\u013c\u013d\u0005"+
		"\u0003\u0000\u0000\u013d\u013e\u0003$\u0012\u0000\u013e\u013f\u0005\u0004"+
		"\u0000\u0000\u013f9\u0001\u0000\u0000\u0000\u0140\u0142\u0003>\u001f\u0000"+
		"\u0141\u0140\u0001\u0000\u0000\u0000\u0142\u0145\u0001\u0000\u0000\u0000"+
		"\u0143\u0141\u0001\u0000\u0000\u0000\u0143\u0144\u0001\u0000\u0000\u0000"+
		"\u0144;\u0001\u0000\u0000\u0000\u0145\u0143\u0001\u0000\u0000\u0000\u0146"+
		"\u0147\u0005\u000e\u0000\u0000\u0147\u0148\u0003:\u001d\u0000\u0148\u0149"+
		"\u0005\u000f\u0000\u0000\u0149=\u0001\u0000\u0000\u0000\u014a\u014b\u0005"+
		"C\u0000\u0000\u014b\u014e\u0005\u0003\u0000\u0000\u014c\u014f\u0003&\u0013"+
		"\u0000\u014d\u014f\u0003R)\u0000\u014e\u014c\u0001\u0000\u0000\u0000\u014e"+
		"\u014d\u0001\u0000\u0000\u0000\u014f\u0150\u0001\u0000\u0000\u0000\u0150"+
		"\u0151\u0005\u0004\u0000\u0000\u0151\u0152\u0005\u0001\u0000\u0000\u0152"+
		"\u0179\u0001\u0000\u0000\u0000\u0153\u0154\u0005D\u0000\u0000\u0154\u0157"+
		"\u0005\u0003\u0000\u0000\u0155\u0158\u0003&\u0013\u0000\u0156\u0158\u0003"+
		"R)\u0000\u0157\u0155\u0001\u0000\u0000\u0000\u0157\u0156\u0001\u0000\u0000"+
		"\u0000\u0158\u0159\u0001\u0000\u0000\u0000\u0159\u015a\u0005\u0004\u0000"+
		"\u0000\u015a\u015b\u0005\u0001\u0000\u0000\u015b\u0179\u0001\u0000\u0000"+
		"\u0000\u015c\u015f\u0005E\u0000\u0000\u015d\u0160\u0003$\u0012\u0000\u015e"+
		"\u0160\u0003R)\u0000\u015f\u015d\u0001\u0000\u0000\u0000\u015f\u015e\u0001"+
		"\u0000\u0000\u0000\u0160\u0161\u0001\u0000\u0000\u0000\u0161\u0162\u0005"+
		"\u0001\u0000\u0000\u0162\u0179\u0001\u0000\u0000\u0000\u0163\u0164\u0005"+
		"F\u0000\u0000\u0164\u0179\u0005\u0001\u0000\u0000\u0165\u0166\u0005G\u0000"+
		"\u0000\u0166\u0179\u0005\u0001\u0000\u0000\u0167\u0168\u0003@ \u0000\u0168"+
		"\u0169\u0005\u0001\u0000\u0000\u0169\u0179\u0001\u0000\u0000\u0000\u016a"+
		"\u016b\u0003<\u001e\u0000\u016b\u016c\u0005\u0001\u0000\u0000\u016c\u0179"+
		"\u0001\u0000\u0000\u0000\u016d\u016e\u0003H$\u0000\u016e\u016f\u0005\u0001"+
		"\u0000\u0000\u016f\u0179\u0001\u0000\u0000\u0000\u0170\u0171\u0003D\""+
		"\u0000\u0171\u0172\u0005\u0001\u0000\u0000\u0172\u0179\u0001\u0000\u0000"+
		"\u0000\u0173\u0174\u0003F#\u0000\u0174\u0175\u0005\u0001\u0000\u0000\u0175"+
		"\u0179\u0001\u0000\u0000\u0000\u0176\u0179\u0003\f\u0006\u0000\u0177\u0179"+
		"\u0003\u0004\u0002\u0000\u0178\u014a\u0001\u0000\u0000\u0000\u0178\u0153"+
		"\u0001\u0000\u0000\u0000\u0178\u015c\u0001\u0000\u0000\u0000\u0178\u0163"+
		"\u0001\u0000\u0000\u0000\u0178\u0165\u0001\u0000\u0000\u0000\u0178\u0167"+
		"\u0001\u0000\u0000\u0000\u0178\u016a\u0001\u0000\u0000\u0000\u0178\u016d"+
		"\u0001\u0000\u0000\u0000\u0178\u0170\u0001\u0000\u0000\u0000\u0178\u0173"+
		"\u0001\u0000\u0000\u0000\u0178\u0176\u0001\u0000\u0000\u0000\u0178\u0177"+
		"\u0001\u0000\u0000\u0000\u0179?\u0001\u0000\u0000\u0000\u017a\u0187\u0003"+
		"R)\u0000\u017b\u017f\u0003$\u0012\u0000\u017c\u0180\u0005\u0010\u0000"+
		"\u0000\u017d\u0180\u0005\u0011\u0000\u0000\u017e\u0180\u0003R)\u0000\u017f"+
		"\u017c\u0001\u0000\u0000\u0000\u017f\u017d\u0001\u0000\u0000\u0000\u017f"+
		"\u017e\u0001\u0000\u0000\u0000\u0180\u0187\u0001\u0000\u0000\u0000\u0181"+
		"\u0187\u0003B!\u0000\u0182\u0183\u0003&\u0013\u0000\u0183\u0184\u0005"+
		"\u0002\u0000\u0000\u0184\u0185\u0003&\u0013\u0000\u0185\u0187\u0001\u0000"+
		"\u0000\u0000\u0186\u017a\u0001\u0000\u0000\u0000\u0186\u017b\u0001\u0000"+
		"\u0000\u0000\u0186\u0181\u0001\u0000\u0000\u0000\u0186\u0182\u0001\u0000"+
		"\u0000\u0000\u0187A\u0001\u0000\u0000\u0000\u0188\u0189\u0003&\u0013\u0000"+
		"\u0189\u018a\u0003P(\u0000\u018a\u018b\u0003&\u0013\u0000\u018b\u0191"+
		"\u0001\u0000\u0000\u0000\u018c\u018d\u0003$\u0012\u0000\u018d\u018e\u0003"+
		"P(\u0000\u018e\u018f\u0003$\u0012\u0000\u018f\u0191\u0001\u0000\u0000"+
		"\u0000\u0190\u0188\u0001\u0000\u0000\u0000\u0190\u018c\u0001\u0000\u0000"+
		"\u0000\u0191C\u0001\u0000\u0000\u0000\u0192\u0193\u00052\u0000\u0000\u0193"+
		"\u0194\u0003$\u0012\u0000\u0194\u0195\u0003<\u001e\u0000\u0195\u01b9\u0001"+
		"\u0000\u0000\u0000\u0196\u0197\u00052\u0000\u0000\u0197\u0198\u0003$\u0012"+
		"\u0000\u0198\u0199\u0003<\u001e\u0000\u0199\u019a\u00056\u0000\u0000\u019a"+
		"\u019b\u0003D\"\u0000\u019b\u01b9\u0001\u0000\u0000\u0000\u019c\u019d"+
		"\u00052\u0000\u0000\u019d\u019e\u0003$\u0012\u0000\u019e\u019f\u0003<"+
		"\u001e\u0000\u019f\u01a0\u00056\u0000\u0000\u01a0\u01a1\u0003<\u001e\u0000"+
		"\u01a1\u01b9\u0001\u0000\u0000\u0000\u01a2\u01a3\u00052\u0000\u0000\u01a3"+
		"\u01a4\u0003@ \u0000\u01a4\u01a5\u0005\u0001\u0000\u0000\u01a5\u01a6\u0003"+
		"$\u0012\u0000\u01a6\u01a7\u0003<\u001e\u0000\u01a7\u01b9\u0001\u0000\u0000"+
		"\u0000\u01a8\u01a9\u00052\u0000\u0000\u01a9\u01aa\u0003@ \u0000\u01aa"+
		"\u01ab\u0005\u0001\u0000\u0000\u01ab\u01ac\u0003$\u0012\u0000\u01ac\u01ad"+
		"\u0003<\u001e\u0000\u01ad\u01ae\u00056\u0000\u0000\u01ae\u01af\u0003D"+
		"\"\u0000\u01af\u01b9\u0001\u0000\u0000\u0000\u01b0\u01b1\u00052\u0000"+
		"\u0000\u01b1\u01b2\u0003@ \u0000\u01b2\u01b3\u0005\u0001\u0000\u0000\u01b3"+
		"\u01b4\u0003$\u0012\u0000\u01b4\u01b5\u0003<\u001e\u0000\u01b5\u01b6\u0005"+
		"6\u0000\u0000\u01b6\u01b7\u0003<\u001e\u0000\u01b7\u01b9\u0001\u0000\u0000"+
		"\u0000\u01b8\u0192\u0001\u0000\u0000\u0000\u01b8\u0196\u0001\u0000\u0000"+
		"\u0000\u01b8\u019c\u0001\u0000\u0000\u0000\u01b8\u01a2\u0001\u0000\u0000"+
		"\u0000\u01b8\u01a8\u0001\u0000\u0000\u0000\u01b8\u01b0\u0001\u0000\u0000"+
		"\u0000\u01b9E\u0001\u0000\u0000\u0000\u01ba\u01bb\u0005H\u0000\u0000\u01bb"+
		"\u01d0\u0003<\u001e\u0000\u01bc\u01bd\u0005H\u0000\u0000\u01bd\u01be\u0003"+
		"$\u0012\u0000\u01be\u01bf\u0003<\u001e\u0000\u01bf\u01d0\u0001\u0000\u0000"+
		"\u0000\u01c0\u01c1\u0005H\u0000\u0000\u01c1\u01c2\u0003@ \u0000\u01c2"+
		"\u01c3\u0005\u0001\u0000\u0000\u01c3\u01c4\u0003$\u0012\u0000\u01c4\u01c5"+
		"\u0005\u0001\u0000\u0000\u01c5\u01c6\u0003@ \u0000\u01c6\u01c7\u0003<"+
		"\u001e\u0000\u01c7\u01d0\u0001\u0000\u0000\u0000\u01c8\u01c9\u0005H\u0000"+
		"\u0000\u01c9\u01ca\u0003@ \u0000\u01ca\u01cb\u0005\u0001\u0000\u0000\u01cb"+
		"\u01cc\u0005\u0001\u0000\u0000\u01cc\u01cd\u0003@ \u0000\u01cd\u01ce\u0003"+
		"<\u001e\u0000\u01ce\u01d0\u0001\u0000\u0000\u0000\u01cf\u01ba\u0001\u0000"+
		"\u0000\u0000\u01cf\u01bc\u0001\u0000\u0000\u0000\u01cf\u01c0\u0001\u0000"+
		"\u0000\u0000\u01cf\u01c8\u0001\u0000\u0000\u0000\u01d0G\u0001\u0000\u0000"+
		"\u0000\u01d1\u01d2\u0005I\u0000\u0000\u01d2\u01d3\u0003@ \u0000\u01d3"+
		"\u01d4\u0005\u0001\u0000\u0000\u01d4\u01d5\u0003$\u0012\u0000\u01d5\u01d6"+
		"\u0005\u000e\u0000\u0000\u01d6\u01d7\u0003J%\u0000\u01d7\u01d8\u0005\u000f"+
		"\u0000\u0000\u01d8\u01ec\u0001\u0000\u0000\u0000\u01d9\u01da\u0005I\u0000"+
		"\u0000\u01da\u01db\u0003$\u0012\u0000\u01db\u01dc\u0005\u000e\u0000\u0000"+
		"\u01dc\u01dd\u0003J%\u0000\u01dd\u01de\u0005\u000f\u0000\u0000\u01de\u01ec"+
		"\u0001\u0000\u0000\u0000\u01df\u01e0\u0005I\u0000\u0000\u01e0\u01e1\u0003"+
		"@ \u0000\u01e1\u01e2\u0005\u0001\u0000\u0000\u01e2\u01e3\u0005\u000e\u0000"+
		"\u0000\u01e3\u01e4\u0003J%\u0000\u01e4\u01e5\u0005\u000f\u0000\u0000\u01e5"+
		"\u01ec\u0001\u0000\u0000\u0000\u01e6\u01e7\u0005I\u0000\u0000\u01e7\u01e8"+
		"\u0005\u000e\u0000\u0000\u01e8\u01e9\u0003J%\u0000\u01e9\u01ea\u0005\u000f"+
		"\u0000\u0000\u01ea\u01ec\u0001\u0000\u0000\u0000\u01eb\u01d1\u0001\u0000"+
		"\u0000\u0000\u01eb\u01d9\u0001\u0000\u0000\u0000\u01eb\u01df\u0001\u0000"+
		"\u0000\u0000\u01eb\u01e6\u0001\u0000\u0000\u0000\u01ecI\u0001\u0000\u0000"+
		"\u0000\u01ed\u01f2\u0003R)\u0000\u01ee\u01ef\u0003L&\u0000\u01ef\u01f0"+
		"\u0003J%\u0000\u01f0\u01f2\u0001\u0000\u0000\u0000\u01f1\u01ed\u0001\u0000"+
		"\u0000\u0000\u01f1\u01ee\u0001\u0000\u0000\u0000\u01f2K\u0001\u0000\u0000"+
		"\u0000\u01f3\u01f4\u0003N\'\u0000\u01f4\u01f5\u0005\u0006\u0000\u0000"+
		"\u01f5\u01f6\u0003:\u001d\u0000\u01f6M\u0001\u0000\u0000\u0000\u01f7\u01f8"+
		"\u0005J\u0000\u0000\u01f8\u01fb\u0003&\u0013\u0000\u01f9\u01fb\u0005K"+
		"\u0000\u0000\u01fa\u01f7\u0001\u0000\u0000\u0000\u01fa\u01f9\u0001\u0000"+
		"\u0000\u0000\u01fbO\u0001\u0000\u0000\u0000\u01fc\u01fd\u0007\u0000\u0000"+
		"\u0000\u01fdQ\u0001\u0000\u0000\u0000\u01fe\u01ff\u0001\u0000\u0000\u0000"+
		"\u01ffS\u0001\u0000\u0000\u0000$\\^mv\u0083\u0094\u009d\u00ac\u00b1\u00b8"+
		"\u00c3\u00d2\u00dd\u00e5\u00ed\u00f5\u00fd\u0105\u010b\u010f\u0118\u011f"+
		"\u0128\u0143\u014e\u0157\u015f\u0178\u017f\u0186\u0190\u01b8\u01cf\u01eb"+
		"\u01f1\u01fa";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}