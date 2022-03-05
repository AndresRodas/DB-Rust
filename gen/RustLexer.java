// Generated from C:/Users/Andres/Documents/Proyectos/DB-Rust/parser\RustLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class RustLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "CHAR", "STR1", "STR2", "USIZE", "VECTOR", "STRUCT", 
			"POW", "PRINT", "LET", "MUT", "FUNC", "ARROW1", "ARROW2", "ABS", "SQRT", 
			"TOSTR", "CLONE", "NEW", "LEN", "PUSH", "REMOVE", "CONTAINS", "INSERT", 
			"CAPACITY", "WCAPACITY", "MAIN", "IF", "MATCH", "LOOP", "WHILE", "FOR", 
			"IN", "BREAK", "CONTINUE", "RETURN", "MODULE", "PUB", "NUMBER", "STRING", 
			"ID", "PUNTO", "C_PTS", "D_PTS", "PYC", "COMA", "DIFERENTE", "IG_IG", 
			"NOT", "OR", "AND", "IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", 
			"MUL", "DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"CORIZQ", "CORDER", "WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'i64'", "'f64'", "'bool'", "'char'", "'String'", "'&str'", "'usize'", 
			"'vec'", "'struct'", "'pow'", "'println!'", "'let'", "'mut'", "'fn'", 
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


	public RustLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "RustLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2J\u01ef\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5"+
		"\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3"+
		"\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16"+
		"\3\16\3\16\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22"+
		"\3\22\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\27\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36"+
		"\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\""+
		"\3\"\3\"\3\"\3#\3#\3#\3#\3$\3$\3$\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3&"+
		"\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3)\3)\3)\3)\3*\6*\u0176"+
		"\n*\r*\16*\u0177\3*\3*\6*\u017c\n*\r*\16*\u017d\5*\u0180\n*\3+\3+\7+\u0184"+
		"\n+\f+\16+\u0187\13+\3+\3+\3,\3,\7,\u018d\n,\f,\16,\u0190\13,\3-\3-\3"+
		".\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\62\3\63\3\63\3\63\3\64\3"+
		"\64\3\65\3\65\3\65\3\66\3\66\3\66\3\67\3\67\38\38\38\39\39\39\3:\3:\3"+
		";\3;\3<\3<\3=\3=\3>\3>\3?\3?\3@\3@\3A\3A\3B\3B\3C\3C\3D\3D\3E\3E\3F\3"+
		"F\3G\6G\u01ce\nG\rG\16G\u01cf\3G\3G\3H\3H\3H\3H\7H\u01d8\nH\fH\16H\u01db"+
		"\13H\3H\3H\3H\3H\3H\3I\3I\3I\3I\7I\u01e6\nI\fI\16I\u01e9\13I\3I\3I\3J"+
		"\3J\3J\3\u01d9\2K\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65"+
		"\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64"+
		"g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089"+
		"F\u008bG\u008dH\u008fI\u0091J\u0093\2\3\2\t\3\2\62;\3\2$$\5\2C\\aac|\6"+
		"\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_"+
		"\2\u01f5\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2"+
		"\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27"+
		"\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2"+
		"\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2"+
		"\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2"+
		"\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2"+
		"\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S"+
		"\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2"+
		"\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2"+
		"\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y"+
		"\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3"+
		"\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2"+
		"\2\u008d\3\2\2\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\3\u0095\3\2\2\2\5\u0099"+
		"\3\2\2\2\7\u009d\3\2\2\2\t\u00a2\3\2\2\2\13\u00a7\3\2\2\2\r\u00ae\3\2"+
		"\2\2\17\u00b3\3\2\2\2\21\u00b9\3\2\2\2\23\u00bd\3\2\2\2\25\u00c4\3\2\2"+
		"\2\27\u00c8\3\2\2\2\31\u00d1\3\2\2\2\33\u00d5\3\2\2\2\35\u00d9\3\2\2\2"+
		"\37\u00dc\3\2\2\2!\u00df\3\2\2\2#\u00e2\3\2\2\2%\u00e6\3\2\2\2\'\u00eb"+
		"\3\2\2\2)\u00f5\3\2\2\2+\u00fb\3\2\2\2-\u00ff\3\2\2\2/\u0103\3\2\2\2\61"+
		"\u0108\3\2\2\2\63\u010f\3\2\2\2\65\u0118\3\2\2\2\67\u011f\3\2\2\29\u0128"+
		"\3\2\2\2;\u0136\3\2\2\2=\u013b\3\2\2\2?\u013e\3\2\2\2A\u0144\3\2\2\2C"+
		"\u0149\3\2\2\2E\u014f\3\2\2\2G\u0153\3\2\2\2I\u0156\3\2\2\2K\u015c\3\2"+
		"\2\2M\u0165\3\2\2\2O\u016c\3\2\2\2Q\u0170\3\2\2\2S\u0175\3\2\2\2U\u0181"+
		"\3\2\2\2W\u018a\3\2\2\2Y\u0191\3\2\2\2[\u0193\3\2\2\2]\u0196\3\2\2\2_"+
		"\u0198\3\2\2\2a\u019a\3\2\2\2c\u019c\3\2\2\2e\u019f\3\2\2\2g\u01a2\3\2"+
		"\2\2i\u01a4\3\2\2\2k\u01a7\3\2\2\2m\u01aa\3\2\2\2o\u01ac\3\2\2\2q\u01af"+
		"\3\2\2\2s\u01b2\3\2\2\2u\u01b4\3\2\2\2w\u01b6\3\2\2\2y\u01b8\3\2\2\2{"+
		"\u01ba\3\2\2\2}\u01bc\3\2\2\2\177\u01be\3\2\2\2\u0081\u01c0\3\2\2\2\u0083"+
		"\u01c2\3\2\2\2\u0085\u01c4\3\2\2\2\u0087\u01c6\3\2\2\2\u0089\u01c8\3\2"+
		"\2\2\u008b\u01ca\3\2\2\2\u008d\u01cd\3\2\2\2\u008f\u01d3\3\2\2\2\u0091"+
		"\u01e1\3\2\2\2\u0093\u01ec\3\2\2\2\u0095\u0096\7k\2\2\u0096\u0097\78\2"+
		"\2\u0097\u0098\7\66\2\2\u0098\4\3\2\2\2\u0099\u009a\7h\2\2\u009a\u009b"+
		"\78\2\2\u009b\u009c\7\66\2\2\u009c\6\3\2\2\2\u009d\u009e\7d\2\2\u009e"+
		"\u009f\7q\2\2\u009f\u00a0\7q\2\2\u00a0\u00a1\7n\2\2\u00a1\b\3\2\2\2\u00a2"+
		"\u00a3\7e\2\2\u00a3\u00a4\7j\2\2\u00a4\u00a5\7c\2\2\u00a5\u00a6\7t\2\2"+
		"\u00a6\n\3\2\2\2\u00a7\u00a8\7U\2\2\u00a8\u00a9\7v\2\2\u00a9\u00aa\7t"+
		"\2\2\u00aa\u00ab\7k\2\2\u00ab\u00ac\7p\2\2\u00ac\u00ad\7i\2\2\u00ad\f"+
		"\3\2\2\2\u00ae\u00af\7(\2\2\u00af\u00b0\7u\2\2\u00b0\u00b1\7v\2\2\u00b1"+
		"\u00b2\7t\2\2\u00b2\16\3\2\2\2\u00b3\u00b4\7w\2\2\u00b4\u00b5\7u\2\2\u00b5"+
		"\u00b6\7k\2\2\u00b6\u00b7\7|\2\2\u00b7\u00b8\7g\2\2\u00b8\20\3\2\2\2\u00b9"+
		"\u00ba\7x\2\2\u00ba\u00bb\7g\2\2\u00bb\u00bc\7e\2\2\u00bc\22\3\2\2\2\u00bd"+
		"\u00be\7u\2\2\u00be\u00bf\7v\2\2\u00bf\u00c0\7t\2\2\u00c0\u00c1\7w\2\2"+
		"\u00c1\u00c2\7e\2\2\u00c2\u00c3\7v\2\2\u00c3\24\3\2\2\2\u00c4\u00c5\7"+
		"r\2\2\u00c5\u00c6\7q\2\2\u00c6\u00c7\7y\2\2\u00c7\26\3\2\2\2\u00c8\u00c9"+
		"\7r\2\2\u00c9\u00ca\7t\2\2\u00ca\u00cb\7k\2\2\u00cb\u00cc\7p\2\2\u00cc"+
		"\u00cd\7v\2\2\u00cd\u00ce\7n\2\2\u00ce\u00cf\7p\2\2\u00cf\u00d0\7#\2\2"+
		"\u00d0\30\3\2\2\2\u00d1\u00d2\7n\2\2\u00d2\u00d3\7g\2\2\u00d3\u00d4\7"+
		"v\2\2\u00d4\32\3\2\2\2\u00d5\u00d6\7o\2\2\u00d6\u00d7\7w\2\2\u00d7\u00d8"+
		"\7v\2\2\u00d8\34\3\2\2\2\u00d9\u00da\7h\2\2\u00da\u00db\7p\2\2\u00db\36"+
		"\3\2\2\2\u00dc\u00dd\7/\2\2\u00dd\u00de\7@\2\2\u00de \3\2\2\2\u00df\u00e0"+
		"\7?\2\2\u00e0\u00e1\7@\2\2\u00e1\"\3\2\2\2\u00e2\u00e3\7c\2\2\u00e3\u00e4"+
		"\7d\2\2\u00e4\u00e5\7u\2\2\u00e5$\3\2\2\2\u00e6\u00e7\7u\2\2\u00e7\u00e8"+
		"\7s\2\2\u00e8\u00e9\7t\2\2\u00e9\u00ea\7v\2\2\u00ea&\3\2\2\2\u00eb\u00ec"+
		"\7v\2\2\u00ec\u00ed\7q\2\2\u00ed\u00ee\7a\2\2\u00ee\u00ef\7u\2\2\u00ef"+
		"\u00f0\7v\2\2\u00f0\u00f1\7t\2\2\u00f1\u00f2\7k\2\2\u00f2\u00f3\7p\2\2"+
		"\u00f3\u00f4\7i\2\2\u00f4(\3\2\2\2\u00f5\u00f6\7e\2\2\u00f6\u00f7\7n\2"+
		"\2\u00f7\u00f8\7q\2\2\u00f8\u00f9\7p\2\2\u00f9\u00fa\7g\2\2\u00fa*\3\2"+
		"\2\2\u00fb\u00fc\7p\2\2\u00fc\u00fd\7g\2\2\u00fd\u00fe\7y\2\2\u00fe,\3"+
		"\2\2\2\u00ff\u0100\7n\2\2\u0100\u0101\7g\2\2\u0101\u0102\7p\2\2\u0102"+
		".\3\2\2\2\u0103\u0104\7r\2\2\u0104\u0105\7w\2\2\u0105\u0106\7u\2\2\u0106"+
		"\u0107\7j\2\2\u0107\60\3\2\2\2\u0108\u0109\7t\2\2\u0109\u010a\7g\2\2\u010a"+
		"\u010b\7o\2\2\u010b\u010c\7q\2\2\u010c\u010d\7x\2\2\u010d\u010e\7g\2\2"+
		"\u010e\62\3\2\2\2\u010f\u0110\7e\2\2\u0110\u0111\7q\2\2\u0111\u0112\7"+
		"p\2\2\u0112\u0113\7v\2\2\u0113\u0114\7c\2\2\u0114\u0115\7k\2\2\u0115\u0116"+
		"\7p\2\2\u0116\u0117\7u\2\2\u0117\64\3\2\2\2\u0118\u0119\7k\2\2\u0119\u011a"+
		"\7p\2\2\u011a\u011b\7u\2\2\u011b\u011c\7g\2\2\u011c\u011d\7t\2\2\u011d"+
		"\u011e\7v\2\2\u011e\66\3\2\2\2\u011f\u0120\7e\2\2\u0120\u0121\7c\2\2\u0121"+
		"\u0122\7r\2\2\u0122\u0123\7c\2\2\u0123\u0124\7e\2\2\u0124\u0125\7k\2\2"+
		"\u0125\u0126\7v\2\2\u0126\u0127\7{\2\2\u01278\3\2\2\2\u0128\u0129\7y\2"+
		"\2\u0129\u012a\7k\2\2\u012a\u012b\7v\2\2\u012b\u012c\7j\2\2\u012c\u012d"+
		"\7a\2\2\u012d\u012e\7e\2\2\u012e\u012f\7c\2\2\u012f\u0130\7r\2\2\u0130"+
		"\u0131\7c\2\2\u0131\u0132\7e\2\2\u0132\u0133\7k\2\2\u0133\u0134\7v\2\2"+
		"\u0134\u0135\7{\2\2\u0135:\3\2\2\2\u0136\u0137\7o\2\2\u0137\u0138\7c\2"+
		"\2\u0138\u0139\7k\2\2\u0139\u013a\7p\2\2\u013a<\3\2\2\2\u013b\u013c\7"+
		"k\2\2\u013c\u013d\7h\2\2\u013d>\3\2\2\2\u013e\u013f\7o\2\2\u013f\u0140"+
		"\7c\2\2\u0140\u0141\7v\2\2\u0141\u0142\7e\2\2\u0142\u0143\7j\2\2\u0143"+
		"@\3\2\2\2\u0144\u0145\7n\2\2\u0145\u0146\7q\2\2\u0146\u0147\7q\2\2\u0147"+
		"\u0148\7r\2\2\u0148B\3\2\2\2\u0149\u014a\7y\2\2\u014a\u014b\7j\2\2\u014b"+
		"\u014c\7k\2\2\u014c\u014d\7n\2\2\u014d\u014e\7g\2\2\u014eD\3\2\2\2\u014f"+
		"\u0150\7h\2\2\u0150\u0151\7q\2\2\u0151\u0152\7t\2\2\u0152F\3\2\2\2\u0153"+
		"\u0154\7k\2\2\u0154\u0155\7p\2\2\u0155H\3\2\2\2\u0156\u0157\7d\2\2\u0157"+
		"\u0158\7t\2\2\u0158\u0159\7g\2\2\u0159\u015a\7c\2\2\u015a\u015b\7m\2\2"+
		"\u015bJ\3\2\2\2\u015c\u015d\7e\2\2\u015d\u015e\7q\2\2\u015e\u015f\7p\2"+
		"\2\u015f\u0160\7v\2\2\u0160\u0161\7k\2\2\u0161\u0162\7p\2\2\u0162\u0163"+
		"\7w\2\2\u0163\u0164\7g\2\2\u0164L\3\2\2\2\u0165\u0166\7t\2\2\u0166\u0167"+
		"\7g\2\2\u0167\u0168\7v\2\2\u0168\u0169\7w\2\2\u0169\u016a\7t\2\2\u016a"+
		"\u016b\7p\2\2\u016bN\3\2\2\2\u016c\u016d\7o\2\2\u016d\u016e\7q\2\2\u016e"+
		"\u016f\7f\2\2\u016fP\3\2\2\2\u0170\u0171\7r\2\2\u0171\u0172\7w\2\2\u0172"+
		"\u0173\7d\2\2\u0173R\3\2\2\2\u0174\u0176\t\2\2\2\u0175\u0174\3\2\2\2\u0176"+
		"\u0177\3\2\2\2\u0177\u0175\3\2\2\2\u0177\u0178\3\2\2\2\u0178\u017f\3\2"+
		"\2\2\u0179\u017b\7\60\2\2\u017a\u017c\t\2\2\2\u017b\u017a\3\2\2\2\u017c"+
		"\u017d\3\2\2\2\u017d\u017b\3\2\2\2\u017d\u017e\3\2\2\2\u017e\u0180\3\2"+
		"\2\2\u017f\u0179\3\2\2\2\u017f\u0180\3\2\2\2\u0180T\3\2\2\2\u0181\u0185"+
		"\7$\2\2\u0182\u0184\n\3\2\2\u0183\u0182\3\2\2\2\u0184\u0187\3\2\2\2\u0185"+
		"\u0183\3\2\2\2\u0185\u0186\3\2\2\2\u0186\u0188\3\2\2\2\u0187\u0185\3\2"+
		"\2\2\u0188\u0189\7$\2\2\u0189V\3\2\2\2\u018a\u018e\t\4\2\2\u018b\u018d"+
		"\t\5\2\2\u018c\u018b\3\2\2\2\u018d\u0190\3\2\2\2\u018e\u018c\3\2\2\2\u018e"+
		"\u018f\3\2\2\2\u018fX\3\2\2\2\u0190\u018e\3\2\2\2\u0191\u0192\7\60\2\2"+
		"\u0192Z\3\2\2\2\u0193\u0194\7<\2\2\u0194\u0195\7<\2\2\u0195\\\3\2\2\2"+
		"\u0196\u0197\7<\2\2\u0197^\3\2\2\2\u0198\u0199\7=\2\2\u0199`\3\2\2\2\u019a"+
		"\u019b\7.\2\2\u019bb\3\2\2\2\u019c\u019d\7#\2\2\u019d\u019e\7?\2\2\u019e"+
		"d\3\2\2\2\u019f\u01a0\7?\2\2\u01a0\u01a1\7?\2\2\u01a1f\3\2\2\2\u01a2\u01a3"+
		"\7#\2\2\u01a3h\3\2\2\2\u01a4\u01a5\7~\2\2\u01a5\u01a6\7~\2\2\u01a6j\3"+
		"\2\2\2\u01a7\u01a8\7(\2\2\u01a8\u01a9\7(\2\2\u01a9l\3\2\2\2\u01aa\u01ab"+
		"\7?\2\2\u01abn\3\2\2\2\u01ac\u01ad\7@\2\2\u01ad\u01ae\7?\2\2\u01aep\3"+
		"\2\2\2\u01af\u01b0\7>\2\2\u01b0\u01b1\7?\2\2\u01b1r\3\2\2\2\u01b2\u01b3"+
		"\7@\2\2\u01b3t\3\2\2\2\u01b4\u01b5\7>\2\2\u01b5v\3\2\2\2\u01b6\u01b7\7"+
		",\2\2\u01b7x\3\2\2\2\u01b8\u01b9\7\61\2\2\u01b9z\3\2\2\2\u01ba\u01bb\7"+
		"-\2\2\u01bb|\3\2\2\2\u01bc\u01bd\7/\2\2\u01bd~\3\2\2\2\u01be\u01bf\7\'"+
		"\2\2\u01bf\u0080\3\2\2\2\u01c0\u01c1\7*\2\2\u01c1\u0082\3\2\2\2\u01c2"+
		"\u01c3\7+\2\2\u01c3\u0084\3\2\2\2\u01c4\u01c5\7}\2\2\u01c5\u0086\3\2\2"+
		"\2\u01c6\u01c7\7\177\2\2\u01c7\u0088\3\2\2\2\u01c8\u01c9\7]\2\2\u01c9"+
		"\u008a\3\2\2\2\u01ca\u01cb\7_\2\2\u01cb\u008c\3\2\2\2\u01cc\u01ce\t\6"+
		"\2\2\u01cd\u01cc\3\2\2\2\u01ce\u01cf\3\2\2\2\u01cf\u01cd\3\2\2\2\u01cf"+
		"\u01d0\3\2\2\2\u01d0\u01d1\3\2\2\2\u01d1\u01d2\bG\2\2\u01d2\u008e\3\2"+
		"\2\2\u01d3\u01d4\7\61\2\2\u01d4\u01d5\7,\2\2\u01d5\u01d9\3\2\2\2\u01d6"+
		"\u01d8\13\2\2\2\u01d7\u01d6\3\2\2\2\u01d8\u01db\3\2\2\2\u01d9\u01da\3"+
		"\2\2\2\u01d9\u01d7\3\2\2\2\u01da\u01dc\3\2\2\2\u01db\u01d9\3\2\2\2\u01dc"+
		"\u01dd\7,\2\2\u01dd\u01de\7\61\2\2\u01de\u01df\3\2\2\2\u01df\u01e0\bH"+
		"\2\2\u01e0\u0090\3\2\2\2\u01e1\u01e2\7\61\2\2\u01e2\u01e3\7\61\2\2\u01e3"+
		"\u01e7\3\2\2\2\u01e4\u01e6\n\7\2\2\u01e5\u01e4\3\2\2\2\u01e6\u01e9\3\2"+
		"\2\2\u01e7\u01e5\3\2\2\2\u01e7\u01e8\3\2\2\2\u01e8\u01ea\3\2\2\2\u01e9"+
		"\u01e7\3\2\2\2\u01ea\u01eb\bI\2\2\u01eb\u0092\3\2\2\2\u01ec\u01ed\7^\2"+
		"\2\u01ed\u01ee\t\b\2\2\u01ee\u0094\3\2\2\2\13\2\u0177\u017d\u017f\u0185"+
		"\u018e\u01cf\u01d9\u01e7\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}