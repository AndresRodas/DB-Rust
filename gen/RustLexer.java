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
		TRU=10, FAL=11, POW=12, PRINT=13, LET=14, MUT=15, FUNC=16, ARROW1=17, 
		ARROW2=18, ABS=19, SQRT=20, TOSTR=21, TOOWN=22, CLONE=23, NEW=24, LEN=25, 
		PUSH=26, REMOVE=27, CONTAINS=28, INSERT=29, CAPACITY=30, WCAPACITY=31, 
		MAIN=32, IF=33, ELSE=34, MATCH=35, LOOP=36, WHILE=37, FOR=38, IN=39, BREAK=40, 
		CONTINUE=41, RETURN=42, MODULE=43, PUB=44, NUMBER=45, STRING=46, ID=47, 
		CHARACTER=48, PUNTO=49, C_PTS=50, D_PTS=51, PYC=52, COMA=53, DIFERENTE=54, 
		IG_IG=55, NOT=56, OR=57, PLEC=58, UNDERSCORE=59, AND=60, IGUAL=61, MAYORIGUAL=62, 
		MENORIGUAL=63, MAYOR=64, MENOR=65, MUL=66, DIV=67, ADD=68, SUB=69, MOD=70, 
		PARIZQ=71, PARDER=72, LLAVEIZQ=73, LLAVEDER=74, CORIZQ=75, CORDER=76, 
		AMP=77, WHITESPACE=78, COMMENT=79, LINE_COMMENT=80;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "CHAR", "STR1", "STR2", "USIZE", "VECTOR", "STRUCT", 
			"TRU", "FAL", "POW", "PRINT", "LET", "MUT", "FUNC", "ARROW1", "ARROW2", 
			"ABS", "SQRT", "TOSTR", "TOOWN", "CLONE", "NEW", "LEN", "PUSH", "REMOVE", 
			"CONTAINS", "INSERT", "CAPACITY", "WCAPACITY", "MAIN", "IF", "ELSE", 
			"MATCH", "LOOP", "WHILE", "FOR", "IN", "BREAK", "CONTINUE", "RETURN", 
			"MODULE", "PUB", "NUMBER", "STRING", "ID", "CHARACTER", "PUNTO", "C_PTS", 
			"D_PTS", "PYC", "COMA", "DIFERENTE", "IG_IG", "NOT", "OR", "PLEC", "UNDERSCORE", 
			"AND", "IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"CORIZQ", "CORDER", "AMP", "WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'i64'", "'f64'", "'bool'", "'char'", "'String'", "'&str'", "'usize'", 
			"'vec'", "'struct'", "'true'", "'false'", "'pow'", "'println!'", "'let'", 
			"'mut'", "'fn'", "'->'", "'=>'", "'abs'", "'sqrt'", "'to_string()'", 
			"'to_owned()'", "'clone'", "'new'", "'len'", "'push'", "'remove'", "'contains'", 
			"'insert'", "'capacity'", "'with_capacity'", "'main'", "'if'", "'else'", 
			"'match'", "'loop'", "'while'", "'for'", "'in'", "'break'", "'continue'", 
			"'return'", "'mod'", "'pub'", null, null, null, null, "'.'", "'::'", 
			"':'", "';'", "','", "'!='", "'=='", "'!'", "'||'", "'|'", "'_'", "'&&'", 
			"'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", 
			"'('", "')'", "'{'", "'}'", "'['", "']'", "'&'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "CHAR", "STR1", "STR2", "USIZE", "VECTOR", 
			"STRUCT", "TRU", "FAL", "POW", "PRINT", "LET", "MUT", "FUNC", "ARROW1", 
			"ARROW2", "ABS", "SQRT", "TOSTR", "TOOWN", "CLONE", "NEW", "LEN", "PUSH", 
			"REMOVE", "CONTAINS", "INSERT", "CAPACITY", "WCAPACITY", "MAIN", "IF", 
			"ELSE", "MATCH", "LOOP", "WHILE", "FOR", "IN", "BREAK", "CONTINUE", "RETURN", 
			"MODULE", "PUB", "NUMBER", "STRING", "ID", "CHARACTER", "PUNTO", "C_PTS", 
			"D_PTS", "PYC", "COMA", "DIFERENTE", "IG_IG", "NOT", "OR", "PLEC", "UNDERSCORE", 
			"AND", "IGUAL", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"CORIZQ", "CORDER", "AMP", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2R\u0226\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\4P\tP\4Q\tQ\4R\tR\3\2\3\2\3\2"+
		"\3\2\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t"+
		"\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3"+
		"\16\3\16\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3"+
		"\22\3\22\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3"+
		"\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3"+
		"\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3"+
		"\30\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3"+
		"\34\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3"+
		"\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3"+
		"\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3!\3!\3!"+
		"\3!\3!\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3&"+
		"\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3(\3(\3(\3)\3)\3)\3)\3)\3)\3*\3*\3*\3"+
		"*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3+\3,\3,\3,\3,\3-\3-\3-\3-\3.\6.\u01a3"+
		"\n.\r.\16.\u01a4\3.\3.\6.\u01a9\n.\r.\16.\u01aa\5.\u01ad\n.\3/\3/\7/\u01b1"+
		"\n/\f/\16/\u01b4\13/\3/\3/\3\60\3\60\7\60\u01ba\n\60\f\60\16\60\u01bd"+
		"\13\60\3\61\3\61\3\61\3\61\3\62\3\62\3\63\3\63\3\63\3\64\3\64\3\65\3\65"+
		"\3\66\3\66\3\67\3\67\3\67\38\38\38\39\39\3:\3:\3:\3;\3;\3<\3<\3=\3=\3"+
		"=\3>\3>\3?\3?\3?\3@\3@\3@\3A\3A\3B\3B\3C\3C\3D\3D\3E\3E\3F\3F\3G\3G\3"+
		"H\3H\3I\3I\3J\3J\3K\3K\3L\3L\3M\3M\3N\3N\3O\6O\u0205\nO\rO\16O\u0206\3"+
		"O\3O\3P\3P\3P\3P\7P\u020f\nP\fP\16P\u0212\13P\3P\3P\3P\3P\3P\3Q\3Q\3Q"+
		"\3Q\7Q\u021d\nQ\fQ\16Q\u0220\13Q\3Q\3Q\3R\3R\3R\3\u0210\2S\3\3\5\4\7\5"+
		"\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23"+
		"%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G"+
		"%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{"+
		"?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008dH\u008fI\u0091"+
		"J\u0093K\u0095L\u0097M\u0099N\u009bO\u009dP\u009fQ\u00a1R\u00a3\2\3\2"+
		"\n\3\2\62;\3\2$$\4\2C\\c|\6\2\62;C\\aac|\3\2))\6\2\13\f\17\17\"\"^^\4"+
		"\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u022c\2\3\3\2\2\2\2\5\3\2\2\2\2\7"+
		"\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2"+
		"\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2"+
		"\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2"+
		"\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2"+
		"\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2"+
		"\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M"+
		"\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2"+
		"\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2"+
		"\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s"+
		"\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177"+
		"\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2"+
		"\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2\2\2\u008f\3\2\2\2\2\u0091"+
		"\3\2\2\2\2\u0093\3\2\2\2\2\u0095\3\2\2\2\2\u0097\3\2\2\2\2\u0099\3\2\2"+
		"\2\2\u009b\3\2\2\2\2\u009d\3\2\2\2\2\u009f\3\2\2\2\2\u00a1\3\2\2\2\3\u00a5"+
		"\3\2\2\2\5\u00a9\3\2\2\2\7\u00ad\3\2\2\2\t\u00b2\3\2\2\2\13\u00b7\3\2"+
		"\2\2\r\u00be\3\2\2\2\17\u00c3\3\2\2\2\21\u00c9\3\2\2\2\23\u00cd\3\2\2"+
		"\2\25\u00d4\3\2\2\2\27\u00d9\3\2\2\2\31\u00df\3\2\2\2\33\u00e3\3\2\2\2"+
		"\35\u00ec\3\2\2\2\37\u00f0\3\2\2\2!\u00f4\3\2\2\2#\u00f7\3\2\2\2%\u00fa"+
		"\3\2\2\2\'\u00fd\3\2\2\2)\u0101\3\2\2\2+\u0106\3\2\2\2-\u0112\3\2\2\2"+
		"/\u011d\3\2\2\2\61\u0123\3\2\2\2\63\u0127\3\2\2\2\65\u012b\3\2\2\2\67"+
		"\u0130\3\2\2\29\u0137\3\2\2\2;\u0140\3\2\2\2=\u0147\3\2\2\2?\u0150\3\2"+
		"\2\2A\u015e\3\2\2\2C\u0163\3\2\2\2E\u0166\3\2\2\2G\u016b\3\2\2\2I\u0171"+
		"\3\2\2\2K\u0176\3\2\2\2M\u017c\3\2\2\2O\u0180\3\2\2\2Q\u0183\3\2\2\2S"+
		"\u0189\3\2\2\2U\u0192\3\2\2\2W\u0199\3\2\2\2Y\u019d\3\2\2\2[\u01a2\3\2"+
		"\2\2]\u01ae\3\2\2\2_\u01b7\3\2\2\2a\u01be\3\2\2\2c\u01c2\3\2\2\2e\u01c4"+
		"\3\2\2\2g\u01c7\3\2\2\2i\u01c9\3\2\2\2k\u01cb\3\2\2\2m\u01cd\3\2\2\2o"+
		"\u01d0\3\2\2\2q\u01d3\3\2\2\2s\u01d5\3\2\2\2u\u01d8\3\2\2\2w\u01da\3\2"+
		"\2\2y\u01dc\3\2\2\2{\u01df\3\2\2\2}\u01e1\3\2\2\2\177\u01e4\3\2\2\2\u0081"+
		"\u01e7\3\2\2\2\u0083\u01e9\3\2\2\2\u0085\u01eb\3\2\2\2\u0087\u01ed\3\2"+
		"\2\2\u0089\u01ef\3\2\2\2\u008b\u01f1\3\2\2\2\u008d\u01f3\3\2\2\2\u008f"+
		"\u01f5\3\2\2\2\u0091\u01f7\3\2\2\2\u0093\u01f9\3\2\2\2\u0095\u01fb\3\2"+
		"\2\2\u0097\u01fd\3\2\2\2\u0099\u01ff\3\2\2\2\u009b\u0201\3\2\2\2\u009d"+
		"\u0204\3\2\2\2\u009f\u020a\3\2\2\2\u00a1\u0218\3\2\2\2\u00a3\u0223\3\2"+
		"\2\2\u00a5\u00a6\7k\2\2\u00a6\u00a7\78\2\2\u00a7\u00a8\7\66\2\2\u00a8"+
		"\4\3\2\2\2\u00a9\u00aa\7h\2\2\u00aa\u00ab\78\2\2\u00ab\u00ac\7\66\2\2"+
		"\u00ac\6\3\2\2\2\u00ad\u00ae\7d\2\2\u00ae\u00af\7q\2\2\u00af\u00b0\7q"+
		"\2\2\u00b0\u00b1\7n\2\2\u00b1\b\3\2\2\2\u00b2\u00b3\7e\2\2\u00b3\u00b4"+
		"\7j\2\2\u00b4\u00b5\7c\2\2\u00b5\u00b6\7t\2\2\u00b6\n\3\2\2\2\u00b7\u00b8"+
		"\7U\2\2\u00b8\u00b9\7v\2\2\u00b9\u00ba\7t\2\2\u00ba\u00bb\7k\2\2\u00bb"+
		"\u00bc\7p\2\2\u00bc\u00bd\7i\2\2\u00bd\f\3\2\2\2\u00be\u00bf\7(\2\2\u00bf"+
		"\u00c0\7u\2\2\u00c0\u00c1\7v\2\2\u00c1\u00c2\7t\2\2\u00c2\16\3\2\2\2\u00c3"+
		"\u00c4\7w\2\2\u00c4\u00c5\7u\2\2\u00c5\u00c6\7k\2\2\u00c6\u00c7\7|\2\2"+
		"\u00c7\u00c8\7g\2\2\u00c8\20\3\2\2\2\u00c9\u00ca\7x\2\2\u00ca\u00cb\7"+
		"g\2\2\u00cb\u00cc\7e\2\2\u00cc\22\3\2\2\2\u00cd\u00ce\7u\2\2\u00ce\u00cf"+
		"\7v\2\2\u00cf\u00d0\7t\2\2\u00d0\u00d1\7w\2\2\u00d1\u00d2\7e\2\2\u00d2"+
		"\u00d3\7v\2\2\u00d3\24\3\2\2\2\u00d4\u00d5\7v\2\2\u00d5\u00d6\7t\2\2\u00d6"+
		"\u00d7\7w\2\2\u00d7\u00d8\7g\2\2\u00d8\26\3\2\2\2\u00d9\u00da\7h\2\2\u00da"+
		"\u00db\7c\2\2\u00db\u00dc\7n\2\2\u00dc\u00dd\7u\2\2\u00dd\u00de\7g\2\2"+
		"\u00de\30\3\2\2\2\u00df\u00e0\7r\2\2\u00e0\u00e1\7q\2\2\u00e1\u00e2\7"+
		"y\2\2\u00e2\32\3\2\2\2\u00e3\u00e4\7r\2\2\u00e4\u00e5\7t\2\2\u00e5\u00e6"+
		"\7k\2\2\u00e6\u00e7\7p\2\2\u00e7\u00e8\7v\2\2\u00e8\u00e9\7n\2\2\u00e9"+
		"\u00ea\7p\2\2\u00ea\u00eb\7#\2\2\u00eb\34\3\2\2\2\u00ec\u00ed\7n\2\2\u00ed"+
		"\u00ee\7g\2\2\u00ee\u00ef\7v\2\2\u00ef\36\3\2\2\2\u00f0\u00f1\7o\2\2\u00f1"+
		"\u00f2\7w\2\2\u00f2\u00f3\7v\2\2\u00f3 \3\2\2\2\u00f4\u00f5\7h\2\2\u00f5"+
		"\u00f6\7p\2\2\u00f6\"\3\2\2\2\u00f7\u00f8\7/\2\2\u00f8\u00f9\7@\2\2\u00f9"+
		"$\3\2\2\2\u00fa\u00fb\7?\2\2\u00fb\u00fc\7@\2\2\u00fc&\3\2\2\2\u00fd\u00fe"+
		"\7c\2\2\u00fe\u00ff\7d\2\2\u00ff\u0100\7u\2\2\u0100(\3\2\2\2\u0101\u0102"+
		"\7u\2\2\u0102\u0103\7s\2\2\u0103\u0104\7t\2\2\u0104\u0105\7v\2\2\u0105"+
		"*\3\2\2\2\u0106\u0107\7v\2\2\u0107\u0108\7q\2\2\u0108\u0109\7a\2\2\u0109"+
		"\u010a\7u\2\2\u010a\u010b\7v\2\2\u010b\u010c\7t\2\2\u010c\u010d\7k\2\2"+
		"\u010d\u010e\7p\2\2\u010e\u010f\7i\2\2\u010f\u0110\7*\2\2\u0110\u0111"+
		"\7+\2\2\u0111,\3\2\2\2\u0112\u0113\7v\2\2\u0113\u0114\7q\2\2\u0114\u0115"+
		"\7a\2\2\u0115\u0116\7q\2\2\u0116\u0117\7y\2\2\u0117\u0118\7p\2\2\u0118"+
		"\u0119\7g\2\2\u0119\u011a\7f\2\2\u011a\u011b\7*\2\2\u011b\u011c\7+\2\2"+
		"\u011c.\3\2\2\2\u011d\u011e\7e\2\2\u011e\u011f\7n\2\2\u011f\u0120\7q\2"+
		"\2\u0120\u0121\7p\2\2\u0121\u0122\7g\2\2\u0122\60\3\2\2\2\u0123\u0124"+
		"\7p\2\2\u0124\u0125\7g\2\2\u0125\u0126\7y\2\2\u0126\62\3\2\2\2\u0127\u0128"+
		"\7n\2\2\u0128\u0129\7g\2\2\u0129\u012a\7p\2\2\u012a\64\3\2\2\2\u012b\u012c"+
		"\7r\2\2\u012c\u012d\7w\2\2\u012d\u012e\7u\2\2\u012e\u012f\7j\2\2\u012f"+
		"\66\3\2\2\2\u0130\u0131\7t\2\2\u0131\u0132\7g\2\2\u0132\u0133\7o\2\2\u0133"+
		"\u0134\7q\2\2\u0134\u0135\7x\2\2\u0135\u0136\7g\2\2\u01368\3\2\2\2\u0137"+
		"\u0138\7e\2\2\u0138\u0139\7q\2\2\u0139\u013a\7p\2\2\u013a\u013b\7v\2\2"+
		"\u013b\u013c\7c\2\2\u013c\u013d\7k\2\2\u013d\u013e\7p\2\2\u013e\u013f"+
		"\7u\2\2\u013f:\3\2\2\2\u0140\u0141\7k\2\2\u0141\u0142\7p\2\2\u0142\u0143"+
		"\7u\2\2\u0143\u0144\7g\2\2\u0144\u0145\7t\2\2\u0145\u0146\7v\2\2\u0146"+
		"<\3\2\2\2\u0147\u0148\7e\2\2\u0148\u0149\7c\2\2\u0149\u014a\7r\2\2\u014a"+
		"\u014b\7c\2\2\u014b\u014c\7e\2\2\u014c\u014d\7k\2\2\u014d\u014e\7v\2\2"+
		"\u014e\u014f\7{\2\2\u014f>\3\2\2\2\u0150\u0151\7y\2\2\u0151\u0152\7k\2"+
		"\2\u0152\u0153\7v\2\2\u0153\u0154\7j\2\2\u0154\u0155\7a\2\2\u0155\u0156"+
		"\7e\2\2\u0156\u0157\7c\2\2\u0157\u0158\7r\2\2\u0158\u0159\7c\2\2\u0159"+
		"\u015a\7e\2\2\u015a\u015b\7k\2\2\u015b\u015c\7v\2\2\u015c\u015d\7{\2\2"+
		"\u015d@\3\2\2\2\u015e\u015f\7o\2\2\u015f\u0160\7c\2\2\u0160\u0161\7k\2"+
		"\2\u0161\u0162\7p\2\2\u0162B\3\2\2\2\u0163\u0164\7k\2\2\u0164\u0165\7"+
		"h\2\2\u0165D\3\2\2\2\u0166\u0167\7g\2\2\u0167\u0168\7n\2\2\u0168\u0169"+
		"\7u\2\2\u0169\u016a\7g\2\2\u016aF\3\2\2\2\u016b\u016c\7o\2\2\u016c\u016d"+
		"\7c\2\2\u016d\u016e\7v\2\2\u016e\u016f\7e\2\2\u016f\u0170\7j\2\2\u0170"+
		"H\3\2\2\2\u0171\u0172\7n\2\2\u0172\u0173\7q\2\2\u0173\u0174\7q\2\2\u0174"+
		"\u0175\7r\2\2\u0175J\3\2\2\2\u0176\u0177\7y\2\2\u0177\u0178\7j\2\2\u0178"+
		"\u0179\7k\2\2\u0179\u017a\7n\2\2\u017a\u017b\7g\2\2\u017bL\3\2\2\2\u017c"+
		"\u017d\7h\2\2\u017d\u017e\7q\2\2\u017e\u017f\7t\2\2\u017fN\3\2\2\2\u0180"+
		"\u0181\7k\2\2\u0181\u0182\7p\2\2\u0182P\3\2\2\2\u0183\u0184\7d\2\2\u0184"+
		"\u0185\7t\2\2\u0185\u0186\7g\2\2\u0186\u0187\7c\2\2\u0187\u0188\7m\2\2"+
		"\u0188R\3\2\2\2\u0189\u018a\7e\2\2\u018a\u018b\7q\2\2\u018b\u018c\7p\2"+
		"\2\u018c\u018d\7v\2\2\u018d\u018e\7k\2\2\u018e\u018f\7p\2\2\u018f\u0190"+
		"\7w\2\2\u0190\u0191\7g\2\2\u0191T\3\2\2\2\u0192\u0193\7t\2\2\u0193\u0194"+
		"\7g\2\2\u0194\u0195\7v\2\2\u0195\u0196\7w\2\2\u0196\u0197\7t\2\2\u0197"+
		"\u0198\7p\2\2\u0198V\3\2\2\2\u0199\u019a\7o\2\2\u019a\u019b\7q\2\2\u019b"+
		"\u019c\7f\2\2\u019cX\3\2\2\2\u019d\u019e\7r\2\2\u019e\u019f\7w\2\2\u019f"+
		"\u01a0\7d\2\2\u01a0Z\3\2\2\2\u01a1\u01a3\t\2\2\2\u01a2\u01a1\3\2\2\2\u01a3"+
		"\u01a4\3\2\2\2\u01a4\u01a2\3\2\2\2\u01a4\u01a5\3\2\2\2\u01a5\u01ac\3\2"+
		"\2\2\u01a6\u01a8\7\60\2\2\u01a7\u01a9\t\2\2\2\u01a8\u01a7\3\2\2\2\u01a9"+
		"\u01aa\3\2\2\2\u01aa\u01a8\3\2\2\2\u01aa\u01ab\3\2\2\2\u01ab\u01ad\3\2"+
		"\2\2\u01ac\u01a6\3\2\2\2\u01ac\u01ad\3\2\2\2\u01ad\\\3\2\2\2\u01ae\u01b2"+
		"\7$\2\2\u01af\u01b1\n\3\2\2\u01b0\u01af\3\2\2\2\u01b1\u01b4\3\2\2\2\u01b2"+
		"\u01b0\3\2\2\2\u01b2\u01b3\3\2\2\2\u01b3\u01b5\3\2\2\2\u01b4\u01b2\3\2"+
		"\2\2\u01b5\u01b6\7$\2\2\u01b6^\3\2\2\2\u01b7\u01bb\t\4\2\2\u01b8\u01ba"+
		"\t\5\2\2\u01b9\u01b8\3\2\2\2\u01ba\u01bd\3\2\2\2\u01bb\u01b9\3\2\2\2\u01bb"+
		"\u01bc\3\2\2\2\u01bc`\3\2\2\2\u01bd\u01bb\3\2\2\2\u01be\u01bf\t\6\2\2"+
		"\u01bf\u01c0\t\4\2\2\u01c0\u01c1\t\6\2\2\u01c1b\3\2\2\2\u01c2\u01c3\7"+
		"\60\2\2\u01c3d\3\2\2\2\u01c4\u01c5\7<\2\2\u01c5\u01c6\7<\2\2\u01c6f\3"+
		"\2\2\2\u01c7\u01c8\7<\2\2\u01c8h\3\2\2\2\u01c9\u01ca\7=\2\2\u01caj\3\2"+
		"\2\2\u01cb\u01cc\7.\2\2\u01ccl\3\2\2\2\u01cd\u01ce\7#\2\2\u01ce\u01cf"+
		"\7?\2\2\u01cfn\3\2\2\2\u01d0\u01d1\7?\2\2\u01d1\u01d2\7?\2\2\u01d2p\3"+
		"\2\2\2\u01d3\u01d4\7#\2\2\u01d4r\3\2\2\2\u01d5\u01d6\7~\2\2\u01d6\u01d7"+
		"\7~\2\2\u01d7t\3\2\2\2\u01d8\u01d9\7~\2\2\u01d9v\3\2\2\2\u01da\u01db\7"+
		"a\2\2\u01dbx\3\2\2\2\u01dc\u01dd\7(\2\2\u01dd\u01de\7(\2\2\u01dez\3\2"+
		"\2\2\u01df\u01e0\7?\2\2\u01e0|\3\2\2\2\u01e1\u01e2\7@\2\2\u01e2\u01e3"+
		"\7?\2\2\u01e3~\3\2\2\2\u01e4\u01e5\7>\2\2\u01e5\u01e6\7?\2\2\u01e6\u0080"+
		"\3\2\2\2\u01e7\u01e8\7@\2\2\u01e8\u0082\3\2\2\2\u01e9\u01ea\7>\2\2\u01ea"+
		"\u0084\3\2\2\2\u01eb\u01ec\7,\2\2\u01ec\u0086\3\2\2\2\u01ed\u01ee\7\61"+
		"\2\2\u01ee\u0088\3\2\2\2\u01ef\u01f0\7-\2\2\u01f0\u008a\3\2\2\2\u01f1"+
		"\u01f2\7/\2\2\u01f2\u008c\3\2\2\2\u01f3\u01f4\7\'\2\2\u01f4\u008e\3\2"+
		"\2\2\u01f5\u01f6\7*\2\2\u01f6\u0090\3\2\2\2\u01f7\u01f8\7+\2\2\u01f8\u0092"+
		"\3\2\2\2\u01f9\u01fa\7}\2\2\u01fa\u0094\3\2\2\2\u01fb\u01fc\7\177\2\2"+
		"\u01fc\u0096\3\2\2\2\u01fd\u01fe\7]\2\2\u01fe\u0098\3\2\2\2\u01ff\u0200"+
		"\7_\2\2\u0200\u009a\3\2\2\2\u0201\u0202\7(\2\2\u0202\u009c\3\2\2\2\u0203"+
		"\u0205\t\7\2\2\u0204\u0203\3\2\2\2\u0205\u0206\3\2\2\2\u0206\u0204\3\2"+
		"\2\2\u0206\u0207\3\2\2\2\u0207\u0208\3\2\2\2\u0208\u0209\bO\2\2\u0209"+
		"\u009e\3\2\2\2\u020a\u020b\7\61\2\2\u020b\u020c\7,\2\2\u020c\u0210\3\2"+
		"\2\2\u020d\u020f\13\2\2\2\u020e\u020d\3\2\2\2\u020f\u0212\3\2\2\2\u0210"+
		"\u0211\3\2\2\2\u0210\u020e\3\2\2\2\u0211\u0213\3\2\2\2\u0212\u0210\3\2"+
		"\2\2\u0213\u0214\7,\2\2\u0214\u0215\7\61\2\2\u0215\u0216\3\2\2\2\u0216"+
		"\u0217\bP\2\2\u0217\u00a0\3\2\2\2\u0218\u0219\7\61\2\2\u0219\u021a\7\61"+
		"\2\2\u021a\u021e\3\2\2\2\u021b\u021d\n\b\2\2\u021c\u021b\3\2\2\2\u021d"+
		"\u0220\3\2\2\2\u021e\u021c\3\2\2\2\u021e\u021f\3\2\2\2\u021f\u0221\3\2"+
		"\2\2\u0220\u021e\3\2\2\2\u0221\u0222\bQ\2\2\u0222\u00a2\3\2\2\2\u0223"+
		"\u0224\7^\2\2\u0224\u0225\t\t\2\2\u0225\u00a4\3\2\2\2\13\2\u01a4\u01aa"+
		"\u01ac\u01b2\u01bb\u0206\u0210\u021e\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}