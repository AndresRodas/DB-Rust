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
		INSERT=26, CAPACITY=27, WCAPACITY=28, MAIN=29, IF=30, ELSE=31, MATCH=32, 
		LOOP=33, WHILE=34, FOR=35, IN=36, BREAK=37, CONTINUE=38, RETURN=39, MODULE=40, 
		PUB=41, NUMBER=42, STRING=43, ID=44, PUNTO=45, C_PTS=46, D_PTS=47, PYC=48, 
		COMA=49, DIFERENTE=50, IG_IG=51, NOT=52, OR=53, AND=54, IGUAL=55, MAYORIGUAL=56, 
		MENORIGUAL=57, MAYOR=58, MENOR=59, MUL=60, DIV=61, ADD=62, SUB=63, MOD=64, 
		PARIZQ=65, PARDER=66, LLAVEIZQ=67, LLAVEDER=68, CORIZQ=69, CORDER=70, 
		WHITESPACE=71, COMMENT=72, LINE_COMMENT=73;
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
			"CAPACITY", "WCAPACITY", "MAIN", "IF", "ELSE", "MATCH", "LOOP", "WHILE", 
			"FOR", "IN", "BREAK", "CONTINUE", "RETURN", "MODULE", "PUB", "NUMBER", 
			"STRING", "ID", "PUNTO", "C_PTS", "D_PTS", "PYC", "COMA", "DIFERENTE", 
			"IG_IG", "NOT", "OR", "AND", "IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", 
			"MENOR", "MUL", "DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", 
			"LLAVEDER", "CORIZQ", "CORDER", "WHITESPACE", "COMMENT", "LINE_COMMENT", 
			"ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'i64'", "'f64'", "'bool'", "'char'", "'String'", "'&str'", "'usize'", 
			"'vec'", "'struct'", "'pow'", "'println!'", "'let'", "'mut'", "'fn'", 
			"'->'", "'=>'", "'abs'", "'sqrt'", "'to_string'", "'clone'", "'new'", 
			"'len'", "'push'", "'remove'", "'contains'", "'insert'", "'capacity'", 
			"'with_capacity'", "'main'", "'if'", "'else'", "'match'", "'loop'", "'while'", 
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
			"INSERT", "CAPACITY", "WCAPACITY", "MAIN", "IF", "ELSE", "MATCH", "LOOP", 
			"WHILE", "FOR", "IN", "BREAK", "CONTINUE", "RETURN", "MODULE", "PUB", 
			"NUMBER", "STRING", "ID", "PUNTO", "C_PTS", "D_PTS", "PYC", "COMA", "DIFERENTE", 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2K\u01f6\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3"+
		"\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13"+
		"\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3"+
		"\16\3\16\3\16\3\16\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3"+
		"\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3"+
		"\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3"+
		"\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3"+
		"\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3"+
		"\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3"+
		"\36\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3!\3\""+
		"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3$\3$\3$\3%\3%\3%\3&\3&\3&\3&\3"+
		"&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3(\3(\3)\3)\3"+
		")\3)\3*\3*\3*\3*\3+\6+\u017d\n+\r+\16+\u017e\3+\3+\6+\u0183\n+\r+\16+"+
		"\u0184\5+\u0187\n+\3,\3,\7,\u018b\n,\f,\16,\u018e\13,\3,\3,\3-\3-\7-\u0194"+
		"\n-\f-\16-\u0197\13-\3.\3.\3/\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63"+
		"\3\63\3\63\3\64\3\64\3\64\3\65\3\65\3\66\3\66\3\66\3\67\3\67\3\67\38\3"+
		"8\39\39\39\3:\3:\3:\3;\3;\3<\3<\3=\3=\3>\3>\3?\3?\3@\3@\3A\3A\3B\3B\3"+
		"C\3C\3D\3D\3E\3E\3F\3F\3G\3G\3H\6H\u01d5\nH\rH\16H\u01d6\3H\3H\3I\3I\3"+
		"I\3I\7I\u01df\nI\fI\16I\u01e2\13I\3I\3I\3I\3I\3I\3J\3J\3J\3J\7J\u01ed"+
		"\nJ\fJ\16J\u01f0\13J\3J\3J\3K\3K\3K\3\u01e0\2L\3\3\5\4\7\5\t\6\13\7\r"+
		"\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25"+
		")\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O"+
		")Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081"+
		"B\u0083C\u0085D\u0087E\u0089F\u008bG\u008dH\u008fI\u0091J\u0093K\u0095"+
		"\2\3\2\t\3\2\62;\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\6\2\13\f\17\17\"\"^^"+
		"\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u01fc\2\3\3\2\2\2\2\5\3\2\2\2\2"+
		"\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2"+
		"\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2"+
		"\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2"+
		"\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2"+
		"\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2"+
		"\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2"+
		"M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3"+
		"\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2"+
		"\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2"+
		"s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177"+
		"\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2"+
		"\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2\2\2\u008f\3\2\2\2\2\u0091"+
		"\3\2\2\2\2\u0093\3\2\2\2\3\u0097\3\2\2\2\5\u009b\3\2\2\2\7\u009f\3\2\2"+
		"\2\t\u00a4\3\2\2\2\13\u00a9\3\2\2\2\r\u00b0\3\2\2\2\17\u00b5\3\2\2\2\21"+
		"\u00bb\3\2\2\2\23\u00bf\3\2\2\2\25\u00c6\3\2\2\2\27\u00ca\3\2\2\2\31\u00d3"+
		"\3\2\2\2\33\u00d7\3\2\2\2\35\u00db\3\2\2\2\37\u00de\3\2\2\2!\u00e1\3\2"+
		"\2\2#\u00e4\3\2\2\2%\u00e8\3\2\2\2\'\u00ed\3\2\2\2)\u00f7\3\2\2\2+\u00fd"+
		"\3\2\2\2-\u0101\3\2\2\2/\u0105\3\2\2\2\61\u010a\3\2\2\2\63\u0111\3\2\2"+
		"\2\65\u011a\3\2\2\2\67\u0121\3\2\2\29\u012a\3\2\2\2;\u0138\3\2\2\2=\u013d"+
		"\3\2\2\2?\u0140\3\2\2\2A\u0145\3\2\2\2C\u014b\3\2\2\2E\u0150\3\2\2\2G"+
		"\u0156\3\2\2\2I\u015a\3\2\2\2K\u015d\3\2\2\2M\u0163\3\2\2\2O\u016c\3\2"+
		"\2\2Q\u0173\3\2\2\2S\u0177\3\2\2\2U\u017c\3\2\2\2W\u0188\3\2\2\2Y\u0191"+
		"\3\2\2\2[\u0198\3\2\2\2]\u019a\3\2\2\2_\u019d\3\2\2\2a\u019f\3\2\2\2c"+
		"\u01a1\3\2\2\2e\u01a3\3\2\2\2g\u01a6\3\2\2\2i\u01a9\3\2\2\2k\u01ab\3\2"+
		"\2\2m\u01ae\3\2\2\2o\u01b1\3\2\2\2q\u01b3\3\2\2\2s\u01b6\3\2\2\2u\u01b9"+
		"\3\2\2\2w\u01bb\3\2\2\2y\u01bd\3\2\2\2{\u01bf\3\2\2\2}\u01c1\3\2\2\2\177"+
		"\u01c3\3\2\2\2\u0081\u01c5\3\2\2\2\u0083\u01c7\3\2\2\2\u0085\u01c9\3\2"+
		"\2\2\u0087\u01cb\3\2\2\2\u0089\u01cd\3\2\2\2\u008b\u01cf\3\2\2\2\u008d"+
		"\u01d1\3\2\2\2\u008f\u01d4\3\2\2\2\u0091\u01da\3\2\2\2\u0093\u01e8\3\2"+
		"\2\2\u0095\u01f3\3\2\2\2\u0097\u0098\7k\2\2\u0098\u0099\78\2\2\u0099\u009a"+
		"\7\66\2\2\u009a\4\3\2\2\2\u009b\u009c\7h\2\2\u009c\u009d\78\2\2\u009d"+
		"\u009e\7\66\2\2\u009e\6\3\2\2\2\u009f\u00a0\7d\2\2\u00a0\u00a1\7q\2\2"+
		"\u00a1\u00a2\7q\2\2\u00a2\u00a3\7n\2\2\u00a3\b\3\2\2\2\u00a4\u00a5\7e"+
		"\2\2\u00a5\u00a6\7j\2\2\u00a6\u00a7\7c\2\2\u00a7\u00a8\7t\2\2\u00a8\n"+
		"\3\2\2\2\u00a9\u00aa\7U\2\2\u00aa\u00ab\7v\2\2\u00ab\u00ac\7t\2\2\u00ac"+
		"\u00ad\7k\2\2\u00ad\u00ae\7p\2\2\u00ae\u00af\7i\2\2\u00af\f\3\2\2\2\u00b0"+
		"\u00b1\7(\2\2\u00b1\u00b2\7u\2\2\u00b2\u00b3\7v\2\2\u00b3\u00b4\7t\2\2"+
		"\u00b4\16\3\2\2\2\u00b5\u00b6\7w\2\2\u00b6\u00b7\7u\2\2\u00b7\u00b8\7"+
		"k\2\2\u00b8\u00b9\7|\2\2\u00b9\u00ba\7g\2\2\u00ba\20\3\2\2\2\u00bb\u00bc"+
		"\7x\2\2\u00bc\u00bd\7g\2\2\u00bd\u00be\7e\2\2\u00be\22\3\2\2\2\u00bf\u00c0"+
		"\7u\2\2\u00c0\u00c1\7v\2\2\u00c1\u00c2\7t\2\2\u00c2\u00c3\7w\2\2\u00c3"+
		"\u00c4\7e\2\2\u00c4\u00c5\7v\2\2\u00c5\24\3\2\2\2\u00c6\u00c7\7r\2\2\u00c7"+
		"\u00c8\7q\2\2\u00c8\u00c9\7y\2\2\u00c9\26\3\2\2\2\u00ca\u00cb\7r\2\2\u00cb"+
		"\u00cc\7t\2\2\u00cc\u00cd\7k\2\2\u00cd\u00ce\7p\2\2\u00ce\u00cf\7v\2\2"+
		"\u00cf\u00d0\7n\2\2\u00d0\u00d1\7p\2\2\u00d1\u00d2\7#\2\2\u00d2\30\3\2"+
		"\2\2\u00d3\u00d4\7n\2\2\u00d4\u00d5\7g\2\2\u00d5\u00d6\7v\2\2\u00d6\32"+
		"\3\2\2\2\u00d7\u00d8\7o\2\2\u00d8\u00d9\7w\2\2\u00d9\u00da\7v\2\2\u00da"+
		"\34\3\2\2\2\u00db\u00dc\7h\2\2\u00dc\u00dd\7p\2\2\u00dd\36\3\2\2\2\u00de"+
		"\u00df\7/\2\2\u00df\u00e0\7@\2\2\u00e0 \3\2\2\2\u00e1\u00e2\7?\2\2\u00e2"+
		"\u00e3\7@\2\2\u00e3\"\3\2\2\2\u00e4\u00e5\7c\2\2\u00e5\u00e6\7d\2\2\u00e6"+
		"\u00e7\7u\2\2\u00e7$\3\2\2\2\u00e8\u00e9\7u\2\2\u00e9\u00ea\7s\2\2\u00ea"+
		"\u00eb\7t\2\2\u00eb\u00ec\7v\2\2\u00ec&\3\2\2\2\u00ed\u00ee\7v\2\2\u00ee"+
		"\u00ef\7q\2\2\u00ef\u00f0\7a\2\2\u00f0\u00f1\7u\2\2\u00f1\u00f2\7v\2\2"+
		"\u00f2\u00f3\7t\2\2\u00f3\u00f4\7k\2\2\u00f4\u00f5\7p\2\2\u00f5\u00f6"+
		"\7i\2\2\u00f6(\3\2\2\2\u00f7\u00f8\7e\2\2\u00f8\u00f9\7n\2\2\u00f9\u00fa"+
		"\7q\2\2\u00fa\u00fb\7p\2\2\u00fb\u00fc\7g\2\2\u00fc*\3\2\2\2\u00fd\u00fe"+
		"\7p\2\2\u00fe\u00ff\7g\2\2\u00ff\u0100\7y\2\2\u0100,\3\2\2\2\u0101\u0102"+
		"\7n\2\2\u0102\u0103\7g\2\2\u0103\u0104\7p\2\2\u0104.\3\2\2\2\u0105\u0106"+
		"\7r\2\2\u0106\u0107\7w\2\2\u0107\u0108\7u\2\2\u0108\u0109\7j\2\2\u0109"+
		"\60\3\2\2\2\u010a\u010b\7t\2\2\u010b\u010c\7g\2\2\u010c\u010d\7o\2\2\u010d"+
		"\u010e\7q\2\2\u010e\u010f\7x\2\2\u010f\u0110\7g\2\2\u0110\62\3\2\2\2\u0111"+
		"\u0112\7e\2\2\u0112\u0113\7q\2\2\u0113\u0114\7p\2\2\u0114\u0115\7v\2\2"+
		"\u0115\u0116\7c\2\2\u0116\u0117\7k\2\2\u0117\u0118\7p\2\2\u0118\u0119"+
		"\7u\2\2\u0119\64\3\2\2\2\u011a\u011b\7k\2\2\u011b\u011c\7p\2\2\u011c\u011d"+
		"\7u\2\2\u011d\u011e\7g\2\2\u011e\u011f\7t\2\2\u011f\u0120\7v\2\2\u0120"+
		"\66\3\2\2\2\u0121\u0122\7e\2\2\u0122\u0123\7c\2\2\u0123\u0124\7r\2\2\u0124"+
		"\u0125\7c\2\2\u0125\u0126\7e\2\2\u0126\u0127\7k\2\2\u0127\u0128\7v\2\2"+
		"\u0128\u0129\7{\2\2\u01298\3\2\2\2\u012a\u012b\7y\2\2\u012b\u012c\7k\2"+
		"\2\u012c\u012d\7v\2\2\u012d\u012e\7j\2\2\u012e\u012f\7a\2\2\u012f\u0130"+
		"\7e\2\2\u0130\u0131\7c\2\2\u0131\u0132\7r\2\2\u0132\u0133\7c\2\2\u0133"+
		"\u0134\7e\2\2\u0134\u0135\7k\2\2\u0135\u0136\7v\2\2\u0136\u0137\7{\2\2"+
		"\u0137:\3\2\2\2\u0138\u0139\7o\2\2\u0139\u013a\7c\2\2\u013a\u013b\7k\2"+
		"\2\u013b\u013c\7p\2\2\u013c<\3\2\2\2\u013d\u013e\7k\2\2\u013e\u013f\7"+
		"h\2\2\u013f>\3\2\2\2\u0140\u0141\7g\2\2\u0141\u0142\7n\2\2\u0142\u0143"+
		"\7u\2\2\u0143\u0144\7g\2\2\u0144@\3\2\2\2\u0145\u0146\7o\2\2\u0146\u0147"+
		"\7c\2\2\u0147\u0148\7v\2\2\u0148\u0149\7e\2\2\u0149\u014a\7j\2\2\u014a"+
		"B\3\2\2\2\u014b\u014c\7n\2\2\u014c\u014d\7q\2\2\u014d\u014e\7q\2\2\u014e"+
		"\u014f\7r\2\2\u014fD\3\2\2\2\u0150\u0151\7y\2\2\u0151\u0152\7j\2\2\u0152"+
		"\u0153\7k\2\2\u0153\u0154\7n\2\2\u0154\u0155\7g\2\2\u0155F\3\2\2\2\u0156"+
		"\u0157\7h\2\2\u0157\u0158\7q\2\2\u0158\u0159\7t\2\2\u0159H\3\2\2\2\u015a"+
		"\u015b\7k\2\2\u015b\u015c\7p\2\2\u015cJ\3\2\2\2\u015d\u015e\7d\2\2\u015e"+
		"\u015f\7t\2\2\u015f\u0160\7g\2\2\u0160\u0161\7c\2\2\u0161\u0162\7m\2\2"+
		"\u0162L\3\2\2\2\u0163\u0164\7e\2\2\u0164\u0165\7q\2\2\u0165\u0166\7p\2"+
		"\2\u0166\u0167\7v\2\2\u0167\u0168\7k\2\2\u0168\u0169\7p\2\2\u0169\u016a"+
		"\7w\2\2\u016a\u016b\7g\2\2\u016bN\3\2\2\2\u016c\u016d\7t\2\2\u016d\u016e"+
		"\7g\2\2\u016e\u016f\7v\2\2\u016f\u0170\7w\2\2\u0170\u0171\7t\2\2\u0171"+
		"\u0172\7p\2\2\u0172P\3\2\2\2\u0173\u0174\7o\2\2\u0174\u0175\7q\2\2\u0175"+
		"\u0176\7f\2\2\u0176R\3\2\2\2\u0177\u0178\7r\2\2\u0178\u0179\7w\2\2\u0179"+
		"\u017a\7d\2\2\u017aT\3\2\2\2\u017b\u017d\t\2\2\2\u017c\u017b\3\2\2\2\u017d"+
		"\u017e\3\2\2\2\u017e\u017c\3\2\2\2\u017e\u017f\3\2\2\2\u017f\u0186\3\2"+
		"\2\2\u0180\u0182\7\60\2\2\u0181\u0183\t\2\2\2\u0182\u0181\3\2\2\2\u0183"+
		"\u0184\3\2\2\2\u0184\u0182\3\2\2\2\u0184\u0185\3\2\2\2\u0185\u0187\3\2"+
		"\2\2\u0186\u0180\3\2\2\2\u0186\u0187\3\2\2\2\u0187V\3\2\2\2\u0188\u018c"+
		"\7$\2\2\u0189\u018b\n\3\2\2\u018a\u0189\3\2\2\2\u018b\u018e\3\2\2\2\u018c"+
		"\u018a\3\2\2\2\u018c\u018d\3\2\2\2\u018d\u018f\3\2\2\2\u018e\u018c\3\2"+
		"\2\2\u018f\u0190\7$\2\2\u0190X\3\2\2\2\u0191\u0195\t\4\2\2\u0192\u0194"+
		"\t\5\2\2\u0193\u0192\3\2\2\2\u0194\u0197\3\2\2\2\u0195\u0193\3\2\2\2\u0195"+
		"\u0196\3\2\2\2\u0196Z\3\2\2\2\u0197\u0195\3\2\2\2\u0198\u0199\7\60\2\2"+
		"\u0199\\\3\2\2\2\u019a\u019b\7<\2\2\u019b\u019c\7<\2\2\u019c^\3\2\2\2"+
		"\u019d\u019e\7<\2\2\u019e`\3\2\2\2\u019f\u01a0\7=\2\2\u01a0b\3\2\2\2\u01a1"+
		"\u01a2\7.\2\2\u01a2d\3\2\2\2\u01a3\u01a4\7#\2\2\u01a4\u01a5\7?\2\2\u01a5"+
		"f\3\2\2\2\u01a6\u01a7\7?\2\2\u01a7\u01a8\7?\2\2\u01a8h\3\2\2\2\u01a9\u01aa"+
		"\7#\2\2\u01aaj\3\2\2\2\u01ab\u01ac\7~\2\2\u01ac\u01ad\7~\2\2\u01adl\3"+
		"\2\2\2\u01ae\u01af\7(\2\2\u01af\u01b0\7(\2\2\u01b0n\3\2\2\2\u01b1\u01b2"+
		"\7?\2\2\u01b2p\3\2\2\2\u01b3\u01b4\7@\2\2\u01b4\u01b5\7?\2\2\u01b5r\3"+
		"\2\2\2\u01b6\u01b7\7>\2\2\u01b7\u01b8\7?\2\2\u01b8t\3\2\2\2\u01b9\u01ba"+
		"\7@\2\2\u01bav\3\2\2\2\u01bb\u01bc\7>\2\2\u01bcx\3\2\2\2\u01bd\u01be\7"+
		",\2\2\u01bez\3\2\2\2\u01bf\u01c0\7\61\2\2\u01c0|\3\2\2\2\u01c1\u01c2\7"+
		"-\2\2\u01c2~\3\2\2\2\u01c3\u01c4\7/\2\2\u01c4\u0080\3\2\2\2\u01c5\u01c6"+
		"\7\'\2\2\u01c6\u0082\3\2\2\2\u01c7\u01c8\7*\2\2\u01c8\u0084\3\2\2\2\u01c9"+
		"\u01ca\7+\2\2\u01ca\u0086\3\2\2\2\u01cb\u01cc\7}\2\2\u01cc\u0088\3\2\2"+
		"\2\u01cd\u01ce\7\177\2\2\u01ce\u008a\3\2\2\2\u01cf\u01d0\7]\2\2\u01d0"+
		"\u008c\3\2\2\2\u01d1\u01d2\7_\2\2\u01d2\u008e\3\2\2\2\u01d3\u01d5\t\6"+
		"\2\2\u01d4\u01d3\3\2\2\2\u01d5\u01d6\3\2\2\2\u01d6\u01d4\3\2\2\2\u01d6"+
		"\u01d7\3\2\2\2\u01d7\u01d8\3\2\2\2\u01d8\u01d9\bH\2\2\u01d9\u0090\3\2"+
		"\2\2\u01da\u01db\7\61\2\2\u01db\u01dc\7,\2\2\u01dc\u01e0\3\2\2\2\u01dd"+
		"\u01df\13\2\2\2\u01de\u01dd\3\2\2\2\u01df\u01e2\3\2\2\2\u01e0\u01e1\3"+
		"\2\2\2\u01e0\u01de\3\2\2\2\u01e1\u01e3\3\2\2\2\u01e2\u01e0\3\2\2\2\u01e3"+
		"\u01e4\7,\2\2\u01e4\u01e5\7\61\2\2\u01e5\u01e6\3\2\2\2\u01e6\u01e7\bI"+
		"\2\2\u01e7\u0092\3\2\2\2\u01e8\u01e9\7\61\2\2\u01e9\u01ea\7\61\2\2\u01ea"+
		"\u01ee\3\2\2\2\u01eb\u01ed\n\7\2\2\u01ec\u01eb\3\2\2\2\u01ed\u01f0\3\2"+
		"\2\2\u01ee\u01ec\3\2\2\2\u01ee\u01ef\3\2\2\2\u01ef\u01f1\3\2\2\2\u01f0"+
		"\u01ee\3\2\2\2\u01f1\u01f2\bJ\2\2\u01f2\u0094\3\2\2\2\u01f3\u01f4\7^\2"+
		"\2\u01f4\u01f5\t\b\2\2\u01f5\u0096\3\2\2\2\13\2\u017e\u0184\u0186\u018c"+
		"\u0195\u01d6\u01e0\u01ee\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}