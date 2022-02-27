// Generated from c:\Users\Andres\Documents\Proyectos\DB-Rust\Rust.g4 by ANTLR 4.8

    import arrayList "github.com/colegno/arraylist"
    

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Rust extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, CHAR=4, STR1=5, STR2=6, USIZE=7, VECTOR=8, STRUCT=9, 
		POW=10, PRINT=11, LET=12, MUT=13, FUNC=14, ARROW1=15, ARROW2=16, ABS=17, 
		SQRT=18, TOSTR=19, CLONE=20, NEW=21, LEN=22, PUSH=23, REMOVE=24, CONTAINS=25, 
		INSERT=26, CAPACITY=27, WCAPACITY=28, MAIN=29, IF=30, MATCH=31, LOOP=32, 
		WHILE=33, FOR=34, IN=35, BREAK=36, CONTINUE=37, RETURN=38, MODULE=39, 
		PUB=40, NUMBER=41, STRING=42, ID=43, PUNTO=44, C_PTS=45, D_PTS=46, PYC=47, 
		COMA=48, DIFERENTE=49, IG_IG=50, NOT=51, OR=52, AND=53, IGUAL=54, MAYORIGUAL=55, 
		MENORIGUAL=56, MAYOR=57, MENOR=58, MUL=59, DIV=60, ADD=61, SUB=62, MOD=63, 
		PARIZQ=64, PARDER=65, LLAVEIZQ=66, LLAVEDER=67, CORIZQ=68, CORDER=69, 
		WHITESPACE=70, COMMENT=71, LINE_COMMENT=72;
	public static final int
		RULE_start = 0, RULE_global_env = 1, RULE_main = 2, RULE_instructions = 3, 
		RULE_instruction = 4, RULE_declaration = 5, RULE_function = 6, RULE_module = 7, 
		RULE_types = 8, RULE_listParams = 9, RULE_expression = 10, RULE_expr_arit = 11, 
		RULE_primitive = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "global_env", "main", "instructions", "instruction", "declaration", 
			"function", "module", "types", "listParams", "expression", "expr_arit", 
			"primitive"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'i64'", "'f64'", "'bool'", "'char'", "'String'", "'&str'", "'usize'", 
			"'vec'", "'struct'", "'pow'", "'println'", "'let'", "'mut'", "'fn'", 
			"'->'", "'=>'", "'abs'", "'sqrt'", "'to_string'", "'clone'", "'new'", 
			"'len'", "'push'", "'remove'", "'contains'", "'insert'", "'capacity'", 
			"'with_capacity'", "'main'", "'if'", "'match'", "'loop'", "'while'", 
			"'for'", "'in'", "'break'", "'continue'", "'return'", "'mod'", "'pub'", 
			null, null, null, "'.'", "'::'", "':'", "';'", "','", "'!='", "'=='", 
			"'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", 
			"'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHAR", "STR1", "STR2", "USIZE", "VECTOR", 
			"STRUCT", "POW", "PRINT", "LET", "MUT", "FUNC", "ARROW1", "ARROW2", "ABS", 
			"SQRT", "TOSTR", "CLONE", "NEW", "LEN", "PUSH", "REMOVE", "CONTAINS", 
			"INSERT", "CAPACITY", "WCAPACITY", "MAIN", "IF", "MATCH", "LOOP", "WHILE", 
			"FOR", "IN", "BREAK", "CONTINUE", "RETURN", "MODULE", "PUB", "NUMBER", 
			"STRING", "ID", "PUNTO", "C_PTS", "D_PTS", "PYC", "COMA", "DIFERENTE", 
			"IG_IG", "NOT", "OR", "AND", "IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", 
			"MENOR", "MUL", "DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", 
			"LLAVEDER", "CORIZQ", "CORDER", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
	public String getGrammarFileName() { return "Rust.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Rust(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public string inst;
		public List<Global_envContext> global_env() {
			return getRuleContexts(Global_envContext.class);
		}
		public Global_envContext global_env(int i) {
			return getRuleContext(Global_envContext.class,i);
		}
		public MainContext main() {
			return getRuleContext(MainContext.class,0);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			setState(42);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(26);
				global_env(0);
				setState(27);
				main();
				setState(28);
				global_env(0);
				 _localctx.inst = "mundo" 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(31);
				global_env(0);
				setState(32);
				main();
				 _localctx.inst = "mundo" 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(35);
				main();
				setState(36);
				global_env(0);
				 _localctx.inst = "hola" 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(39);
				main();
				 _localctx.inst = "hola" 
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

	public static class Global_envContext extends ParserRuleContext {
		public DeclarationContext declaration() {
			return getRuleContext(DeclarationContext.class,0);
		}
		public FunctionContext function() {
			return getRuleContext(FunctionContext.class,0);
		}
		public ModuleContext module() {
			return getRuleContext(ModuleContext.class,0);
		}
		public Global_envContext global_env() {
			return getRuleContext(Global_envContext.class,0);
		}
		public Global_envContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_global_env; }
	}

	public final Global_envContext global_env() throws RecognitionException {
		return global_env(0);
	}

	private Global_envContext global_env(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Global_envContext _localctx = new Global_envContext(_ctx, _parentState);
		Global_envContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_global_env, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LET:
				{
				setState(45);
				declaration();
				}
				break;
			case FUNC:
				{
				setState(46);
				function();
				}
				break;
			case MODULE:
				{
				setState(47);
				module();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(58);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(56);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
					case 1:
						{
						_localctx = new Global_envContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_global_env);
						setState(50);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(51);
						declaration();
						}
						break;
					case 2:
						{
						_localctx = new Global_envContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_global_env);
						setState(52);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(53);
						function();
						}
						break;
					case 3:
						{
						_localctx = new Global_envContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_global_env);
						setState(54);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(55);
						module();
						}
						break;
					}
					} 
				}
				setState(60);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
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

	public static class MainContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(Rust.FUNC, 0); }
		public TerminalNode MAIN() { return getToken(Rust.MAIN, 0); }
		public TerminalNode PARIZQ() { return getToken(Rust.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Rust.PARDER, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Rust.LLAVEIZQ, 0); }
		public InstructionsContext instructions() {
			return getRuleContext(InstructionsContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Rust.LLAVEDER, 0); }
		public MainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main; }
	}

	public final MainContext main() throws RecognitionException {
		MainContext _localctx = new MainContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_main);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(61);
			match(FUNC);
			setState(62);
			match(MAIN);
			setState(63);
			match(PARIZQ);
			setState(64);
			match(PARDER);
			setState(65);
			match(LLAVEIZQ);
			setState(66);
			instructions();
			setState(67);
			match(LLAVEDER);
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

	public static class InstructionsContext extends ParserRuleContext {
		public List<InstructionContext> instruction() {
			return getRuleContexts(InstructionContext.class);
		}
		public InstructionContext instruction(int i) {
			return getRuleContext(InstructionContext.class,i);
		}
		public InstructionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instructions; }
	}

	public final InstructionsContext instructions() throws RecognitionException {
		InstructionsContext _localctx = new InstructionsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_instructions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PRINT) {
				{
				{
				setState(69);
				instruction();
				}
				}
				setState(74);
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

	public static class InstructionContext extends ParserRuleContext {
		public TerminalNode PRINT() { return getToken(Rust.PRINT, 0); }
		public TerminalNode PYC() { return getToken(Rust.PYC, 0); }
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_instruction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			match(PRINT);
			setState(76);
			match(PYC);
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

	public static class DeclarationContext extends ParserRuleContext {
		public TerminalNode LET() { return getToken(Rust.LET, 0); }
		public TerminalNode MUT() { return getToken(Rust.MUT, 0); }
		public TerminalNode ID() { return getToken(Rust.ID, 0); }
		public TerminalNode D_PTS() { return getToken(Rust.D_PTS, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Rust.IGUAL, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PYC() { return getToken(Rust.PYC, 0); }
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_declaration);
		try {
			setState(108);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(78);
				match(LET);
				setState(79);
				match(MUT);
				setState(80);
				match(ID);
				setState(81);
				match(D_PTS);
				setState(82);
				types();
				setState(83);
				match(IGUAL);
				setState(84);
				expression();
				setState(85);
				match(PYC);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(87);
				match(LET);
				setState(88);
				match(MUT);
				setState(89);
				match(ID);
				setState(90);
				match(IGUAL);
				setState(91);
				expression();
				setState(92);
				match(PYC);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(94);
				match(LET);
				setState(95);
				match(ID);
				setState(96);
				match(D_PTS);
				setState(97);
				types();
				setState(98);
				match(IGUAL);
				setState(99);
				expression();
				setState(100);
				match(PYC);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(102);
				match(LET);
				setState(103);
				match(ID);
				setState(104);
				match(IGUAL);
				setState(105);
				expression();
				setState(106);
				match(PYC);
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

	public static class FunctionContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(Rust.FUNC, 0); }
		public TerminalNode ID() { return getToken(Rust.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(Rust.PARIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Rust.PARDER, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Rust.LLAVEIZQ, 0); }
		public InstructionsContext instructions() {
			return getRuleContext(InstructionsContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Rust.LLAVEDER, 0); }
		public FunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function; }
	}

	public final FunctionContext function() throws RecognitionException {
		FunctionContext _localctx = new FunctionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_function);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			match(FUNC);
			setState(111);
			match(ID);
			setState(112);
			match(PARIZQ);
			setState(113);
			listParams(0);
			setState(114);
			match(PARDER);
			setState(115);
			match(LLAVEIZQ);
			setState(116);
			instructions();
			setState(117);
			match(LLAVEDER);
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

	public static class ModuleContext extends ParserRuleContext {
		public TerminalNode MODULE() { return getToken(Rust.MODULE, 0); }
		public TerminalNode ID() { return getToken(Rust.ID, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Rust.LLAVEIZQ, 0); }
		public TerminalNode LLAVEDER() { return getToken(Rust.LLAVEDER, 0); }
		public ModuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_module; }
	}

	public final ModuleContext module() throws RecognitionException {
		ModuleContext _localctx = new ModuleContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_module);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			match(MODULE);
			setState(120);
			match(ID);
			setState(121);
			match(LLAVEIZQ);
			setState(122);
			match(LLAVEDER);
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

	public static class TypesContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(Rust.INT, 0); }
		public TerminalNode FLOAT() { return getToken(Rust.FLOAT, 0); }
		public TerminalNode BOOL() { return getToken(Rust.BOOL, 0); }
		public TerminalNode CHAR() { return getToken(Rust.CHAR, 0); }
		public TerminalNode STR1() { return getToken(Rust.STR1, 0); }
		public TerminalNode VECTOR() { return getToken(Rust.VECTOR, 0); }
		public TerminalNode STRUCT() { return getToken(Rust.STRUCT, 0); }
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_types);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << BOOL) | (1L << CHAR) | (1L << STR1) | (1L << VECTOR) | (1L << STRUCT))) != 0)) ) {
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

	public static class ListParamsContext extends ParserRuleContext {
		public *arrayList.List l;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode COMA() { return getToken(Rust.COMA, 0); }
		public ListParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParams; }
	}

	public final ListParamsContext listParams() throws RecognitionException {
		return listParams(0);
	}

	private ListParamsContext listParams(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsContext _localctx = new ListParamsContext(_ctx, _parentState);
		ListParamsContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(127);
			expression();
			 
			}
			_ctx.stop = _input.LT(-1);
			setState(137);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_listParams);
					setState(130);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(131);
					match(COMA);
					setState(132);
					expression();
					 
					}
					} 
				}
				setState(139);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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

	public static class ExpressionContext extends ParserRuleContext {
		public Expr_aritContext expr_arit() {
			return getRuleContext(Expr_aritContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			expr_arit(0);
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

	public static class Expr_aritContext extends ParserRuleContext {
		public Expr_aritContext opIz;
		public Token op;
		public Expr_aritContext opDe;
		public TerminalNode CORIZQ() { return getToken(Rust.CORIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(Rust.CORDER, 0); }
		public TerminalNode PARIZQ() { return getToken(Rust.PARIZQ, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Rust.PARDER, 0); }
		public PrimitiveContext primitive() {
			return getRuleContext(PrimitiveContext.class,0);
		}
		public List<Expr_aritContext> expr_arit() {
			return getRuleContexts(Expr_aritContext.class);
		}
		public Expr_aritContext expr_arit(int i) {
			return getRuleContext(Expr_aritContext.class,i);
		}
		public TerminalNode MUL() { return getToken(Rust.MUL, 0); }
		public TerminalNode DIV() { return getToken(Rust.DIV, 0); }
		public TerminalNode ADD() { return getToken(Rust.ADD, 0); }
		public TerminalNode SUB() { return getToken(Rust.SUB, 0); }
		public TerminalNode MENOR() { return getToken(Rust.MENOR, 0); }
		public TerminalNode MENORIGUAL() { return getToken(Rust.MENORIGUAL, 0); }
		public TerminalNode MAYORIGUAL() { return getToken(Rust.MAYORIGUAL, 0); }
		public TerminalNode MAYOR() { return getToken(Rust.MAYOR, 0); }
		public Expr_aritContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_arit; }
	}

	public final Expr_aritContext expr_arit() throws RecognitionException {
		return expr_arit(0);
	}

	private Expr_aritContext expr_arit(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_aritContext _localctx = new Expr_aritContext(_ctx, _parentState);
		Expr_aritContext _prevctx = _localctx;
		int _startState = 22;
		enterRecursionRule(_localctx, 22, RULE_expr_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CORIZQ:
				{
				setState(143);
				match(CORIZQ);
				setState(144);
				listParams(0);
				setState(145);
				match(CORDER);
				}
				break;
			case PARIZQ:
				{
				setState(147);
				match(PARIZQ);
				setState(148);
				expression();
				setState(149);
				match(PARDER);
				}
				break;
			case NUMBER:
			case STRING:
				{
				setState(151);
				primitive();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(165);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(163);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(154);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(155);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MUL || _la==DIV) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(156);
						((Expr_aritContext)_localctx).opDe = expr_arit(7);
						}
						break;
					case 2:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(157);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(158);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(159);
						((Expr_aritContext)_localctx).opDe = expr_arit(6);
						}
						break;
					case 3:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(160);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(161);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MAYORIGUAL) | (1L << MENORIGUAL) | (1L << MAYOR) | (1L << MENOR))) != 0)) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(162);
						((Expr_aritContext)_localctx).opDe = expr_arit(5);
						}
						break;
					}
					} 
				}
				setState(167);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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

	public static class PrimitiveContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(Rust.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(Rust.STRING, 0); }
		public PrimitiveContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitive; }
	}

	public final PrimitiveContext primitive() throws RecognitionException {
		PrimitiveContext _localctx = new PrimitiveContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_primitive);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(168);
			_la = _input.LA(1);
			if ( !(_la==NUMBER || _la==STRING) ) {
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return global_env_sempred((Global_envContext)_localctx, predIndex);
		case 9:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 11:
			return expr_arit_sempred((Expr_aritContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean global_env_sempred(Global_envContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		case 2:
			return precpred(_ctx, 4);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_arit_sempred(Expr_aritContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 6);
		case 5:
			return precpred(_ctx, 5);
		case 6:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3J\u00ad\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\5\2-\n\2\3\3\3\3\3\3\3\3\5\3\63\n\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\7\3;\n\3\f\3\16\3>\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\5\7\5I\n\5\f\5\16\5L\13\5\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\5\7o\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t"+
		"\3\t\3\t\3\t\3\t\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\7\13\u008a\n\13\f\13\16\13\u008d\13\13\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\5\r\u009b\n\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\7"+
		"\r\u00a6\n\r\f\r\16\r\u00a9\13\r\3\16\3\16\3\16\2\5\4\24\30\17\2\4\6\b"+
		"\n\f\16\20\22\24\26\30\32\2\7\4\2\3\7\n\13\3\2=>\3\2?@\3\29<\3\2+,\2\u00b1"+
		"\2,\3\2\2\2\4\62\3\2\2\2\6?\3\2\2\2\bJ\3\2\2\2\nM\3\2\2\2\fn\3\2\2\2\16"+
		"p\3\2\2\2\20y\3\2\2\2\22~\3\2\2\2\24\u0080\3\2\2\2\26\u008e\3\2\2\2\30"+
		"\u009a\3\2\2\2\32\u00aa\3\2\2\2\34\35\5\4\3\2\35\36\5\6\4\2\36\37\5\4"+
		"\3\2\37 \b\2\1\2 -\3\2\2\2!\"\5\4\3\2\"#\5\6\4\2#$\b\2\1\2$-\3\2\2\2%"+
		"&\5\6\4\2&\'\5\4\3\2\'(\b\2\1\2(-\3\2\2\2)*\5\6\4\2*+\b\2\1\2+-\3\2\2"+
		"\2,\34\3\2\2\2,!\3\2\2\2,%\3\2\2\2,)\3\2\2\2-\3\3\2\2\2./\b\3\1\2/\63"+
		"\5\f\7\2\60\63\5\16\b\2\61\63\5\20\t\2\62.\3\2\2\2\62\60\3\2\2\2\62\61"+
		"\3\2\2\2\63<\3\2\2\2\64\65\f\b\2\2\65;\5\f\7\2\66\67\f\7\2\2\67;\5\16"+
		"\b\289\f\6\2\29;\5\20\t\2:\64\3\2\2\2:\66\3\2\2\2:8\3\2\2\2;>\3\2\2\2"+
		"<:\3\2\2\2<=\3\2\2\2=\5\3\2\2\2><\3\2\2\2?@\7\20\2\2@A\7\37\2\2AB\7B\2"+
		"\2BC\7C\2\2CD\7D\2\2DE\5\b\5\2EF\7E\2\2F\7\3\2\2\2GI\5\n\6\2HG\3\2\2\2"+
		"IL\3\2\2\2JH\3\2\2\2JK\3\2\2\2K\t\3\2\2\2LJ\3\2\2\2MN\7\r\2\2NO\7\61\2"+
		"\2O\13\3\2\2\2PQ\7\16\2\2QR\7\17\2\2RS\7-\2\2ST\7\60\2\2TU\5\22\n\2UV"+
		"\78\2\2VW\5\26\f\2WX\7\61\2\2Xo\3\2\2\2YZ\7\16\2\2Z[\7\17\2\2[\\\7-\2"+
		"\2\\]\78\2\2]^\5\26\f\2^_\7\61\2\2_o\3\2\2\2`a\7\16\2\2ab\7-\2\2bc\7\60"+
		"\2\2cd\5\22\n\2de\78\2\2ef\5\26\f\2fg\7\61\2\2go\3\2\2\2hi\7\16\2\2ij"+
		"\7-\2\2jk\78\2\2kl\5\26\f\2lm\7\61\2\2mo\3\2\2\2nP\3\2\2\2nY\3\2\2\2n"+
		"`\3\2\2\2nh\3\2\2\2o\r\3\2\2\2pq\7\20\2\2qr\7-\2\2rs\7B\2\2st\5\24\13"+
		"\2tu\7C\2\2uv\7D\2\2vw\5\b\5\2wx\7E\2\2x\17\3\2\2\2yz\7)\2\2z{\7-\2\2"+
		"{|\7D\2\2|}\7E\2\2}\21\3\2\2\2~\177\t\2\2\2\177\23\3\2\2\2\u0080\u0081"+
		"\b\13\1\2\u0081\u0082\5\26\f\2\u0082\u0083\b\13\1\2\u0083\u008b\3\2\2"+
		"\2\u0084\u0085\f\4\2\2\u0085\u0086\7\62\2\2\u0086\u0087\5\26\f\2\u0087"+
		"\u0088\b\13\1\2\u0088\u008a\3\2\2\2\u0089\u0084\3\2\2\2\u008a\u008d\3"+
		"\2\2\2\u008b\u0089\3\2\2\2\u008b\u008c\3\2\2\2\u008c\25\3\2\2\2\u008d"+
		"\u008b\3\2\2\2\u008e\u008f\5\30\r\2\u008f\27\3\2\2\2\u0090\u0091\b\r\1"+
		"\2\u0091\u0092\7F\2\2\u0092\u0093\5\24\13\2\u0093\u0094\7G\2\2\u0094\u009b"+
		"\3\2\2\2\u0095\u0096\7B\2\2\u0096\u0097\5\26\f\2\u0097\u0098\7C\2\2\u0098"+
		"\u009b\3\2\2\2\u0099\u009b\5\32\16\2\u009a\u0090\3\2\2\2\u009a\u0095\3"+
		"\2\2\2\u009a\u0099\3\2\2\2\u009b\u00a7\3\2\2\2\u009c\u009d\f\b\2\2\u009d"+
		"\u009e\t\3\2\2\u009e\u00a6\5\30\r\t\u009f\u00a0\f\7\2\2\u00a0\u00a1\t"+
		"\4\2\2\u00a1\u00a6\5\30\r\b\u00a2\u00a3\f\6\2\2\u00a3\u00a4\t\5\2\2\u00a4"+
		"\u00a6\5\30\r\7\u00a5\u009c\3\2\2\2\u00a5\u009f\3\2\2\2\u00a5\u00a2\3"+
		"\2\2\2\u00a6\u00a9\3\2\2\2\u00a7\u00a5\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8"+
		"\31\3\2\2\2\u00a9\u00a7\3\2\2\2\u00aa\u00ab\t\6\2\2\u00ab\33\3\2\2\2\f"+
		",\62:<Jn\u008b\u009a\u00a5\u00a7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}