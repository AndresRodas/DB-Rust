// Generated from c:\Users\Andres\Documents\Proyectos\DB-Rust\RustLexer.g4 by ANTLR 4.8
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2J\u01ee\b\1\4\2\t"+
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
		"\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16"+
		"\3\16\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22\3\22"+
		"\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\27\3\27\3\27"+
		"\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36"+
		"\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3\"\3\"\3\"\3\""+
		"\3\"\3\"\3#\3#\3#\3#\3$\3$\3$\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3&\3&\3"+
		"&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3)\3)\3)\3)\3*\6*\u0175\n"+
		"*\r*\16*\u0176\3*\3*\6*\u017b\n*\r*\16*\u017c\5*\u017f\n*\3+\3+\7+\u0183"+
		"\n+\f+\16+\u0186\13+\3+\3+\3,\3,\7,\u018c\n,\f,\16,\u018f\13,\3-\3-\3"+
		".\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\62\3\63\3\63\3\63\3\64\3"+
		"\64\3\65\3\65\3\65\3\66\3\66\3\66\3\67\3\67\38\38\38\39\39\39\3:\3:\3"+
		";\3;\3<\3<\3=\3=\3>\3>\3?\3?\3@\3@\3A\3A\3B\3B\3C\3C\3D\3D\3E\3E\3F\3"+
		"F\3G\6G\u01cd\nG\rG\16G\u01ce\3G\3G\3H\3H\3H\3H\7H\u01d7\nH\fH\16H\u01da"+
		"\13H\3H\3H\3H\3H\3H\3I\3I\3I\3I\7I\u01e5\nI\fI\16I\u01e8\13I\3I\3I\3J"+
		"\3J\3J\3\u01d8\2K\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65"+
		"\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64"+
		"g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089"+
		"F\u008bG\u008dH\u008fI\u0091J\u0093\2\3\2\t\3\2\62;\3\2$$\5\2C\\aac|\6"+
		"\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_"+
		"\2\u01f4\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2"+
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
		"\2\27\u00c8\3\2\2\2\31\u00d0\3\2\2\2\33\u00d4\3\2\2\2\35\u00d8\3\2\2\2"+
		"\37\u00db\3\2\2\2!\u00de\3\2\2\2#\u00e1\3\2\2\2%\u00e5\3\2\2\2\'\u00ea"+
		"\3\2\2\2)\u00f4\3\2\2\2+\u00fa\3\2\2\2-\u00fe\3\2\2\2/\u0102\3\2\2\2\61"+
		"\u0107\3\2\2\2\63\u010e\3\2\2\2\65\u0117\3\2\2\2\67\u011e\3\2\2\29\u0127"+
		"\3\2\2\2;\u0135\3\2\2\2=\u013a\3\2\2\2?\u013d\3\2\2\2A\u0143\3\2\2\2C"+
		"\u0148\3\2\2\2E\u014e\3\2\2\2G\u0152\3\2\2\2I\u0155\3\2\2\2K\u015b\3\2"+
		"\2\2M\u0164\3\2\2\2O\u016b\3\2\2\2Q\u016f\3\2\2\2S\u0174\3\2\2\2U\u0180"+
		"\3\2\2\2W\u0189\3\2\2\2Y\u0190\3\2\2\2[\u0192\3\2\2\2]\u0195\3\2\2\2_"+
		"\u0197\3\2\2\2a\u0199\3\2\2\2c\u019b\3\2\2\2e\u019e\3\2\2\2g\u01a1\3\2"+
		"\2\2i\u01a3\3\2\2\2k\u01a6\3\2\2\2m\u01a9\3\2\2\2o\u01ab\3\2\2\2q\u01ae"+
		"\3\2\2\2s\u01b1\3\2\2\2u\u01b3\3\2\2\2w\u01b5\3\2\2\2y\u01b7\3\2\2\2{"+
		"\u01b9\3\2\2\2}\u01bb\3\2\2\2\177\u01bd\3\2\2\2\u0081\u01bf\3\2\2\2\u0083"+
		"\u01c1\3\2\2\2\u0085\u01c3\3\2\2\2\u0087\u01c5\3\2\2\2\u0089\u01c7\3\2"+
		"\2\2\u008b\u01c9\3\2\2\2\u008d\u01cc\3\2\2\2\u008f\u01d2\3\2\2\2\u0091"+
		"\u01e0\3\2\2\2\u0093\u01eb\3\2\2\2\u0095\u0096\7k\2\2\u0096\u0097\78\2"+
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
		"\u00cd\7v\2\2\u00cd\u00ce\7n\2\2\u00ce\u00cf\7p\2\2\u00cf\30\3\2\2\2\u00d0"+
		"\u00d1\7n\2\2\u00d1\u00d2\7g\2\2\u00d2\u00d3\7v\2\2\u00d3\32\3\2\2\2\u00d4"+
		"\u00d5\7o\2\2\u00d5\u00d6\7w\2\2\u00d6\u00d7\7v\2\2\u00d7\34\3\2\2\2\u00d8"+
		"\u00d9\7h\2\2\u00d9\u00da\7p\2\2\u00da\36\3\2\2\2\u00db\u00dc\7/\2\2\u00dc"+
		"\u00dd\7@\2\2\u00dd \3\2\2\2\u00de\u00df\7?\2\2\u00df\u00e0\7@\2\2\u00e0"+
		"\"\3\2\2\2\u00e1\u00e2\7c\2\2\u00e2\u00e3\7d\2\2\u00e3\u00e4\7u\2\2\u00e4"+
		"$\3\2\2\2\u00e5\u00e6\7u\2\2\u00e6\u00e7\7s\2\2\u00e7\u00e8\7t\2\2\u00e8"+
		"\u00e9\7v\2\2\u00e9&\3\2\2\2\u00ea\u00eb\7v\2\2\u00eb\u00ec\7q\2\2\u00ec"+
		"\u00ed\7a\2\2\u00ed\u00ee\7u\2\2\u00ee\u00ef\7v\2\2\u00ef\u00f0\7t\2\2"+
		"\u00f0\u00f1\7k\2\2\u00f1\u00f2\7p\2\2\u00f2\u00f3\7i\2\2\u00f3(\3\2\2"+
		"\2\u00f4\u00f5\7e\2\2\u00f5\u00f6\7n\2\2\u00f6\u00f7\7q\2\2\u00f7\u00f8"+
		"\7p\2\2\u00f8\u00f9\7g\2\2\u00f9*\3\2\2\2\u00fa\u00fb\7p\2\2\u00fb\u00fc"+
		"\7g\2\2\u00fc\u00fd\7y\2\2\u00fd,\3\2\2\2\u00fe\u00ff\7n\2\2\u00ff\u0100"+
		"\7g\2\2\u0100\u0101\7p\2\2\u0101.\3\2\2\2\u0102\u0103\7r\2\2\u0103\u0104"+
		"\7w\2\2\u0104\u0105\7u\2\2\u0105\u0106\7j\2\2\u0106\60\3\2\2\2\u0107\u0108"+
		"\7t\2\2\u0108\u0109\7g\2\2\u0109\u010a\7o\2\2\u010a\u010b\7q\2\2\u010b"+
		"\u010c\7x\2\2\u010c\u010d\7g\2\2\u010d\62\3\2\2\2\u010e\u010f\7e\2\2\u010f"+
		"\u0110\7q\2\2\u0110\u0111\7p\2\2\u0111\u0112\7v\2\2\u0112\u0113\7c\2\2"+
		"\u0113\u0114\7k\2\2\u0114\u0115\7p\2\2\u0115\u0116\7u\2\2\u0116\64\3\2"+
		"\2\2\u0117\u0118\7k\2\2\u0118\u0119\7p\2\2\u0119\u011a\7u\2\2\u011a\u011b"+
		"\7g\2\2\u011b\u011c\7t\2\2\u011c\u011d\7v\2\2\u011d\66\3\2\2\2\u011e\u011f"+
		"\7e\2\2\u011f\u0120\7c\2\2\u0120\u0121\7r\2\2\u0121\u0122\7c\2\2\u0122"+
		"\u0123\7e\2\2\u0123\u0124\7k\2\2\u0124\u0125\7v\2\2\u0125\u0126\7{\2\2"+
		"\u01268\3\2\2\2\u0127\u0128\7y\2\2\u0128\u0129\7k\2\2\u0129\u012a\7v\2"+
		"\2\u012a\u012b\7j\2\2\u012b\u012c\7a\2\2\u012c\u012d\7e\2\2\u012d\u012e"+
		"\7c\2\2\u012e\u012f\7r\2\2\u012f\u0130\7c\2\2\u0130\u0131\7e\2\2\u0131"+
		"\u0132\7k\2\2\u0132\u0133\7v\2\2\u0133\u0134\7{\2\2\u0134:\3\2\2\2\u0135"+
		"\u0136\7o\2\2\u0136\u0137\7c\2\2\u0137\u0138\7k\2\2\u0138\u0139\7p\2\2"+
		"\u0139<\3\2\2\2\u013a\u013b\7k\2\2\u013b\u013c\7h\2\2\u013c>\3\2\2\2\u013d"+
		"\u013e\7o\2\2\u013e\u013f\7c\2\2\u013f\u0140\7v\2\2\u0140\u0141\7e\2\2"+
		"\u0141\u0142\7j\2\2\u0142@\3\2\2\2\u0143\u0144\7n\2\2\u0144\u0145\7q\2"+
		"\2\u0145\u0146\7q\2\2\u0146\u0147\7r\2\2\u0147B\3\2\2\2\u0148\u0149\7"+
		"y\2\2\u0149\u014a\7j\2\2\u014a\u014b\7k\2\2\u014b\u014c\7n\2\2\u014c\u014d"+
		"\7g\2\2\u014dD\3\2\2\2\u014e\u014f\7h\2\2\u014f\u0150\7q\2\2\u0150\u0151"+
		"\7t\2\2\u0151F\3\2\2\2\u0152\u0153\7k\2\2\u0153\u0154\7p\2\2\u0154H\3"+
		"\2\2\2\u0155\u0156\7d\2\2\u0156\u0157\7t\2\2\u0157\u0158\7g\2\2\u0158"+
		"\u0159\7c\2\2\u0159\u015a\7m\2\2\u015aJ\3\2\2\2\u015b\u015c\7e\2\2\u015c"+
		"\u015d\7q\2\2\u015d\u015e\7p\2\2\u015e\u015f\7v\2\2\u015f\u0160\7k\2\2"+
		"\u0160\u0161\7p\2\2\u0161\u0162\7w\2\2\u0162\u0163\7g\2\2\u0163L\3\2\2"+
		"\2\u0164\u0165\7t\2\2\u0165\u0166\7g\2\2\u0166\u0167\7v\2\2\u0167\u0168"+
		"\7w\2\2\u0168\u0169\7t\2\2\u0169\u016a\7p\2\2\u016aN\3\2\2\2\u016b\u016c"+
		"\7o\2\2\u016c\u016d\7q\2\2\u016d\u016e\7f\2\2\u016eP\3\2\2\2\u016f\u0170"+
		"\7r\2\2\u0170\u0171\7w\2\2\u0171\u0172\7d\2\2\u0172R\3\2\2\2\u0173\u0175"+
		"\t\2\2\2\u0174\u0173\3\2\2\2\u0175\u0176\3\2\2\2\u0176\u0174\3\2\2\2\u0176"+
		"\u0177\3\2\2\2\u0177\u017e\3\2\2\2\u0178\u017a\7\60\2\2\u0179\u017b\t"+
		"\2\2\2\u017a\u0179\3\2\2\2\u017b\u017c\3\2\2\2\u017c\u017a\3\2\2\2\u017c"+
		"\u017d\3\2\2\2\u017d\u017f\3\2\2\2\u017e\u0178\3\2\2\2\u017e\u017f\3\2"+
		"\2\2\u017fT\3\2\2\2\u0180\u0184\7$\2\2\u0181\u0183\n\3\2\2\u0182\u0181"+
		"\3\2\2\2\u0183\u0186\3\2\2\2\u0184\u0182\3\2\2\2\u0184\u0185\3\2\2\2\u0185"+
		"\u0187\3\2\2\2\u0186\u0184\3\2\2\2\u0187\u0188\7$\2\2\u0188V\3\2\2\2\u0189"+
		"\u018d\t\4\2\2\u018a\u018c\t\5\2\2\u018b\u018a\3\2\2\2\u018c\u018f\3\2"+
		"\2\2\u018d\u018b\3\2\2\2\u018d\u018e\3\2\2\2\u018eX\3\2\2\2\u018f\u018d"+
		"\3\2\2\2\u0190\u0191\7\60\2\2\u0191Z\3\2\2\2\u0192\u0193\7<\2\2\u0193"+
		"\u0194\7<\2\2\u0194\\\3\2\2\2\u0195\u0196\7<\2\2\u0196^\3\2\2\2\u0197"+
		"\u0198\7=\2\2\u0198`\3\2\2\2\u0199\u019a\7.\2\2\u019ab\3\2\2\2\u019b\u019c"+
		"\7#\2\2\u019c\u019d\7?\2\2\u019dd\3\2\2\2\u019e\u019f\7?\2\2\u019f\u01a0"+
		"\7?\2\2\u01a0f\3\2\2\2\u01a1\u01a2\7#\2\2\u01a2h\3\2\2\2\u01a3\u01a4\7"+
		"~\2\2\u01a4\u01a5\7~\2\2\u01a5j\3\2\2\2\u01a6\u01a7\7(\2\2\u01a7\u01a8"+
		"\7(\2\2\u01a8l\3\2\2\2\u01a9\u01aa\7?\2\2\u01aan\3\2\2\2\u01ab\u01ac\7"+
		"@\2\2\u01ac\u01ad\7?\2\2\u01adp\3\2\2\2\u01ae\u01af\7>\2\2\u01af\u01b0"+
		"\7?\2\2\u01b0r\3\2\2\2\u01b1\u01b2\7@\2\2\u01b2t\3\2\2\2\u01b3\u01b4\7"+
		">\2\2\u01b4v\3\2\2\2\u01b5\u01b6\7,\2\2\u01b6x\3\2\2\2\u01b7\u01b8\7\61"+
		"\2\2\u01b8z\3\2\2\2\u01b9\u01ba\7-\2\2\u01ba|\3\2\2\2\u01bb\u01bc\7/\2"+
		"\2\u01bc~\3\2\2\2\u01bd\u01be\7\'\2\2\u01be\u0080\3\2\2\2\u01bf\u01c0"+
		"\7*\2\2\u01c0\u0082\3\2\2\2\u01c1\u01c2\7+\2\2\u01c2\u0084\3\2\2\2\u01c3"+
		"\u01c4\7}\2\2\u01c4\u0086\3\2\2\2\u01c5\u01c6\7\177\2\2\u01c6\u0088\3"+
		"\2\2\2\u01c7\u01c8\7]\2\2\u01c8\u008a\3\2\2\2\u01c9\u01ca\7_\2\2\u01ca"+
		"\u008c\3\2\2\2\u01cb\u01cd\t\6\2\2\u01cc\u01cb\3\2\2\2\u01cd\u01ce\3\2"+
		"\2\2\u01ce\u01cc\3\2\2\2\u01ce\u01cf\3\2\2\2\u01cf\u01d0\3\2\2\2\u01d0"+
		"\u01d1\bG\2\2\u01d1\u008e\3\2\2\2\u01d2\u01d3\7\61\2\2\u01d3\u01d4\7,"+
		"\2\2\u01d4\u01d8\3\2\2\2\u01d5\u01d7\13\2\2\2\u01d6\u01d5\3\2\2\2\u01d7"+
		"\u01da\3\2\2\2\u01d8\u01d9\3\2\2\2\u01d8\u01d6\3\2\2\2\u01d9\u01db\3\2"+
		"\2\2\u01da\u01d8\3\2\2\2\u01db\u01dc\7,\2\2\u01dc\u01dd\7\61\2\2\u01dd"+
		"\u01de\3\2\2\2\u01de\u01df\bH\2\2\u01df\u0090\3\2\2\2\u01e0\u01e1\7\61"+
		"\2\2\u01e1\u01e2\7\61\2\2\u01e2\u01e6\3\2\2\2\u01e3\u01e5\n\7\2\2\u01e4"+
		"\u01e3\3\2\2\2\u01e5\u01e8\3\2\2\2\u01e6\u01e4\3\2\2\2\u01e6\u01e7\3\2"+
		"\2\2\u01e7\u01e9\3\2\2\2\u01e8\u01e6\3\2\2\2\u01e9\u01ea\bI\2\2\u01ea"+
		"\u0092\3\2\2\2\u01eb\u01ec\7^\2\2\u01ec\u01ed\t\b\2\2\u01ed\u0094\3\2"+
		"\2\2\13\2\u0176\u017c\u017e\u0184\u018d\u01ce\u01d8\u01e6\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}