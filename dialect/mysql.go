package dialect

import (
	"regexp"
	"strings"
)

type MySQLDialect struct {
	dialect int
	name    string
}

func NewMySQLDialect() *MySQLDialect {
	var d MySQLDialect

	d.dialect = MySQL
	d.name = "MySQL"

	return &d
}

func (d MySQLDialect) Dialect() int {
	return d.dialect
}
func (d MySQLDialect) DialectName() string {
	return d.name
}
func (d MySQLDialect) CaseFolding() int {
	return NoFolding
}
func (d MySQLDialect) IdentQuoteChar() string {
	return "\""
}
func (d MySQLDialect) StringQuoteChar() string {
	return "'"
}

// MaxOperatorLength returns the length of the longest operator
// supported by the dialect
func (d MySQLDialect) MaxOperatorLength() int {
	return 3
}

// IsDatatype returns a boolean indicating if the supplied string
// (or string slice) is considered to be a datatype in MySQL
func (d MySQLDialect) IsDatatype(s ...string) bool {

	var mysqlDatatypes = map[string]bool{
		"bigint":                 true, // [(n)]
		"bigint (n)":             true, // [(n)]
		"binary":                 true, // [(n)]
		"binary (n)":             true, // [(n)]
		"bit":                    true,
		"blob":                   true,
		"boolean":                true,
		"bool":                   true,
		"char":                   true, // (n)
		"char (n)":               true, // (n)
		"character":              true, // (n)
		"character (n)":          true, // (n)
		"datetime":               true,
		"date":                   true,
		"decimal":                true, // [(p[,s])]
		"decimal (n)":            true, // [(p[,s])]
		"decimal (n,n)":          true, // [(p[,s])]
		"dec":                    true, // [(p[,s])]
		"dec (n)":                true, // [(p[,s])]
		"dec (n,n)":              true, // [(p[,s])]
		"double precision":       true, // [(p[,s])]
		"double precision (n)":   true, // [(p[,s])]
		"double precision (n,n)": true, // [(p[,s])]
		"double":                 true, // [(p[,s])]
		"double (n)":             true, // [(p[,s])]
		"double (n,n)":           true, // [(p[,s])]
		"enum":                   true,
		"float":                  true, // [(p[,s])]
		"float (n)":              true, // [(p[,s])]
		"float (n,n)":            true, // [(p[,s])]
		"integer":                true, // [(n)]
		"integer (n)":            true, // [(n)]
		"int":                    true, // [(n)]
		"int (n)":                true, // [(n)]
		"longblob":               true,
		"longtext":               true,
		"mediumblob":             true,
		"mediumint":              true, // (n)
		"mediumint (n)":          true, // (n)
		"mediumtext":             true,
		"nchar":                  true, // (n)
		"nchar (n)":              true, // (n)
		"nvarchar":               true, // (n)
		"nvarchar (n)":           true, // (n)
		"numeric":                true, // (p,s)
		"numeric (n)":            true, // (p,s)
		"numeric (n,n)":          true, // (p,s)
		"real":                   true, // (p,s)
		"real (n)":               true, // (p,s)
		"real (n,n)":             true, // (p,s)
		"set":                    true,
		"smallint":               true, // [(n)]
		"smallint (n)":           true, // [(n)]
		"text":                   true,
		"timestamp":              true,
		"time":                   true,
		"tinyblob":               true,
		"tinyint":                true,
		"tinytext":               true,
		"varbinary":              true, // (n)
		"varbinary (n)":          true, // (n)
		"varchar":                true,
		"year":                   true,
		"geometry":               true, //GIS extension
		"geometrycollection":     true, //GIS extension
		"linestring":             true, //GIS extension
		"multilinestring":        true, //GIS extension
		"multipoint":             true, //GIS extension
		"multipolygon":           true, //GIS extension
		"point":                  true, //GIS extension
		"polygon":                true, //GIS extension
	}

	var z []string
	rn := regexp.MustCompile(`^[0-9]+$`)

	for i, v := range s {
		switch v {
		case "(":
			z = append(z, " "+v)
		case ")", ",":
			z = append(z, v)
		default:
			switch {
			case rn.MatchString(v):
				z = append(z, "n")
			case i == 0:
				z = append(z, v)
			default:
				z = append(z, " "+v)
			}
		}
	}

	k := strings.ToLower(strings.Join(z, ""))
	if _, ok := mysqlDatatypes[k]; ok {
		return true
	}

	return false
}

func (d MySQLDialect) keyword(s string) (bool, bool) {

	/*
	   MySQL keywords

	   https://dev.mysql.com/doc/refman/5.7/en/keywords.html

	   Those keywords that had no indicator if they were reserved or not were set to false.

	*/

	// map[keyword]isReserved
	var mysqlKeywords = map[string]bool{
		"ACCESSIBLE":                    true,
		"ACCOUNT":                       false,
		"ACTION":                        false,
		"ADD":                           true,
		"AFTER":                         false,
		"AGAINST":                       false,
		"AGGREGATE":                     false,
		"ALGORITHM":                     false,
		"ALL":                           true,
		"ALTER":                         true,
		"ALWAYS":                        false,
		"ANALYZE":                       true,
		"AND":                           true,
		"ANY":                           false,
		"ASCII":                         false,
		"ASC":                           true,
		"ASENSITIVE":                    true,
		"AS":                            true,
		"AT":                            false,
		"AUTOEXTEND_SIZE":               false,
		"AUTO_INCREMENT":                false,
		"AVG":                           false,
		"AVG_ROW_LENGTH":                false,
		"BACKUP":                        false,
		"BEFORE":                        true,
		"BEGIN":                         false,
		"BETWEEN":                       true,
		"BIGINT":                        true,
		"BINARY":                        true,
		"BINLOG":                        false,
		"BIT":                           false,
		"BLOB":                          true,
		"BLOCK":                         false,
		"BOOLEAN":                       false,
		"BOOL":                          false,
		"BOTH":                          true,
		"BTREE":                         false,
		"BYTE":                          false,
		"BY":                            true,
		"CACHE":                         false,
		"CALL":                          true,
		"CASCADED":                      false,
		"CASCADE":                       true,
		"CASE":                          true,
		"CATALOG_NAME":                  false,
		"CHAIN":                         false,
		"CHANGED":                       false,
		"CHANGE":                        true,
		"CHANNEL":                       false,
		"CHARACTER":                     true,
		"CHARSET":                       false,
		"CHAR":                          true,
		"CHECKSUM":                      false,
		"CHECK":                         true,
		"CIPHER":                        false,
		"CLASS_ORIGIN":                  false,
		"CLIENT":                        false,
		"CLOSE":                         false,
		"COALESCE":                      false,
		"CODE":                          false,
		"COLLATE":                       true,
		"COLLATION":                     false,
		"COLUMN_FORMAT":                 false,
		"COLUMN_NAME":                   false,
		"COLUMNS":                       false,
		"COLUMN":                        true,
		"COMMENT":                       false,
		"COMMIT":                        false,
		"COMMITTED":                     false,
		"COMPACT":                       false,
		"COMPLETION":                    false,
		"COMPRESSED":                    false,
		"COMPRESSION":                   false,
		"CONCURRENT":                    false,
		"CONDITION":                     true,
		"CONNECTION":                    false,
		"CONSISTENT":                    false,
		"CONSTRAINT_CATALOG":            false,
		"CONSTRAINT_NAME":               false,
		"CONSTRAINT_SCHEMA":             false,
		"CONSTRAINT":                    true,
		"CONTAINS":                      false,
		"CONTEXT":                       false,
		"CONTINUE":                      true,
		"CONVERT":                       true,
		"CPU":                           false,
		"CREATE":                        true,
		"CROSS":                         true,
		"CUBE":                          false,
		"CURRENT_DATE":                  true,
		"CURRENT":                       false,
		"CURRENT_TIMESTAMP":             true,
		"CURRENT_TIME":                  true,
		"CURRENT_USER":                  true,
		"CURSOR_NAME":                   false,
		"CURSOR":                        true,
		"DATABASES":                     true,
		"DATABASE":                      true,
		"DATA":                          false,
		"DATAFILE":                      false,
		"DATE":                          false,
		"DATETIME":                      false,
		"DAY":                           false,
		"DAY_HOUR":                      true,
		"DAY_MICROSECOND":               true,
		"DAY_MINUTE":                    true,
		"DAY_SECOND":                    true,
		"DEALLOCATE":                    false,
		"DECIMAL":                       true,
		"DECLARE":                       true,
		"DEC":                           true,
		"DEFAULT_AUTH":                  false,
		"DEFAULT":                       true,
		"DEFINER":                       false,
		"DELAYED":                       true,
		"DELAY_KEY_WRITE":               false,
		"DELETE":                        true,
		"DESCRIBE":                      true,
		"DESC":                          true,
		"DES_KEY_FILE":                  false,
		"DETERMINISTIC":                 true,
		"DIAGNOSTICS":                   false,
		"DIRECTORY":                     false,
		"DISABLE":                       false,
		"DISCARD":                       false,
		"DISK":                          false,
		"DISTINCTROW":                   true,
		"DISTINCT":                      true,
		"DIV":                           true,
		"DO":                            false,
		"DOUBLE":                        true,
		"DROP":                          true,
		"DUAL":                          true,
		"DUMPFILE":                      false,
		"DUPLICATE":                     false,
		"DYNAMIC":                       false,
		"EACH":                          true,
		"ELSEIF":                        true,
		"ELSE":                          true,
		"ENABLE":                        false,
		"ENCLOSED":                      true,
		"ENCRYPTION":                    false,
		"END":                           false,
		"ENDS":                          false,
		"ENGINE":                        false,
		"ENGINES":                       false,
		"ENUM":                          false,
		"ERROR":                         false,
		"ERRORS":                        false,
		"ESCAPED":                       true,
		"ESCAPE":                        false,
		"EVENT":                         false,
		"EVENTS":                        false,
		"EVERY":                         false,
		"EXCHANGE":                      false,
		"EXECUTE":                       false,
		"EXISTS":                        true,
		"EXIT":                          true,
		"EXPANSION":                     false,
		"EXPIRE":                        false,
		"EXPLAIN":                       true,
		"EXPORT":                        false,
		"EXTENDED":                      false,
		"EXTENT_SIZE":                   false,
		"FALSE":                         true,
		"FAST":                          false,
		"FAULTS":                        false,
		"FETCH":                         true,
		"FIELDS":                        false,
		"FILE_BLOCK_SIZE":               false,
		"FILE":                          false,
		"FILTER":                        false,
		"FIRST":                         false,
		"FIXED":                         false,
		"FLOAT4":                        true,
		"FLOAT8":                        true,
		"FLOAT":                         true,
		"FLUSH":                         false,
		"FOLLOWS":                       false,
		"FORCE":                         true,
		"FOREIGN":                       true,
		"FORMAT":                        false,
		"FOR":                           true,
		"FOUND":                         false,
		"FROM":                          true,
		"FULL":                          false,
		"FULLTEXT":                      true,
		"FUNCTION":                      false,
		"GENERAL":                       false,
		"GENERATED":                     true,
		"GEOMETRYCOLLECTION":            false,
		"GEOMETRY":                      false,
		"GET_FORMAT":                    false,
		"GET":                           true,
		"GLOBAL":                        false,
		"GRANTS":                        false,
		"GRANT":                         true,
		"GROUP_REPLICATION":             false,
		"GROUP":                         true,
		"HANDLER":                       false,
		"HASH":                          false,
		"HAVING":                        true,
		"HELP":                          false,
		"HIGH_PRIORITY":                 true,
		"HOST":                          false,
		"HOSTS":                         false,
		"HOUR":                          false,
		"HOUR_MICROSECOND":              true,
		"HOUR_MINUTE":                   true,
		"HOUR_SECOND":                   true,
		"IDENTIFIED":                    false,
		"IF":                            true,
		"IGNORE_SERVER_IDS":             false,
		"IGNORE":                        true,
		"IMPORT":                        false,
		"INDEXES":                       false,
		"INDEX":                         true,
		"INFILE":                        true,
		"INITIAL_SIZE":                  false,
		"INNER":                         true,
		"INOUT":                         true,
		"INSENSITIVE":                   true,
		"INSERT_METHOD":                 false,
		"INSERT":                        true,
		"INSTALL":                       false,
		"INSTANCE":                      false,
		"INT1":                          true,
		"INT2":                          true,
		"INT3":                          true,
		"INT4":                          true,
		"INT8":                          true,
		"INTEGER":                       true,
		"INTERVAL":                      true,
		"INTO":                          true,
		"IN":                            true,
		"INT":                           true,
		"INVOKER":                       false,
		"IO_AFTER_GTIDS":                true,
		"IO_BEFORE_GTIDS":               true,
		"IO":                            false,
		"IO_THREAD":                     false,
		"IPC":                           false,
		"ISOLATION":                     false,
		"ISSUER":                        false,
		"IS":                            true,
		"ITERATE":                       true,
		"JOIN":                          true,
		"JSON":                          false,
		"KEY_BLOCK_SIZE":                false,
		"KEYS":                          true,
		"KEY":                           true,
		"KILL":                          true,
		"LANGUAGE":                      false,
		"LAST":                          false,
		"LEADING":                       true,
		"LEAVES":                        false,
		"LEAVE":                         true,
		"LEFT":                          true,
		"LESS":                          false,
		"LEVEL":                         false,
		"LIKE":                          true,
		"LIMIT":                         true,
		"LINEAR":                        true,
		"LINESTRING":                    false,
		"LINES":                         true,
		"LIST":                          false,
		"LOAD":                          true,
		"LOCAL":                         false,
		"LOCALTIMESTAMP":                true,
		"LOCALTIME":                     true,
		"LOCKS":                         false,
		"LOCK":                          true,
		"LOGFILE":                       false,
		"LOGS":                          false,
		"LONGBLOB":                      true,
		"LONGTEXT":                      true,
		"LONG":                          true,
		"LOOP":                          true,
		"LOW_PRIORITY":                  true,
		"MASTER_AUTO_POSITION":          false,
		"MASTER_BIND":                   true,
		"MASTER_CONNECT_RETRY":          false,
		"MASTER_DELAY":                  false,
		"MASTER":                        false,
		"MASTER_HEARTBEAT_PERIOD":       false,
		"MASTER_HOST":                   false,
		"MASTER_LOG_FILE":               false,
		"MASTER_LOG_POS":                false,
		"MASTER_PASSWORD":               false,
		"MASTER_PORT":                   false,
		"MASTER_RETRY_COUNT":            false,
		"MASTER_SERVER_ID":              false,
		"MASTER_SSL_CA":                 false,
		"MASTER_SSL_CAPATH":             false,
		"MASTER_SSL_CERT":               false,
		"MASTER_SSL_CIPHER":             false,
		"MASTER_SSL_CRL":                false,
		"MASTER_SSL_CRLPATH":            false,
		"MASTER_SSL":                    false,
		"MASTER_SSL_KEY":                false,
		"MASTER_SSL_VERIFY_SERVER_CERT": true,
		"MASTER_TLS_VERSION":            false,
		"MASTER_USER":                   false,
		"MATCH":                         true,
		"MAX_CONNECTIONS_PER_HOUR":      false,
		"MAX_QUERIES_PER_HOUR":          false,
		"MAX_ROWS":                      false,
		"MAX_SIZE":                      false,
		"MAX_STATEMENT_TIME":            false,
		"MAX_UPDATES_PER_HOUR":          false,
		"MAX_USER_CONNECTIONS":          false,
		"MAXVALUE":                      true,
		"MEDIUMBLOB":                    true,
		"MEDIUM":                        false,
		"MEDIUMINT":                     true,
		"MEDIUMTEXT":                    true,
		"MEMORY":                        false,
		"MERGE":                         false,
		"MESSAGE_TEXT":                  false,
		"MICROSECOND":                   false,
		"MIDDLEINT":                     true,
		"MIGRATE":                       false,
		"MIN_ROWS":                      false,
		"MINUTE":                        false,
		"MINUTE_MICROSECOND":            true,
		"MINUTE_SECOND":                 true,
		"MODE":                          false,
		"MODIFIES":                      true,
		"MODIFY":                        false,
		"MOD":                           true,
		"MONTH":                         false,
		"MULTILINESTRING":               false,
		"MULTIPOINT":                    false,
		"MULTIPOLYGON":                  false,
		"MUTEX":                         false,
		"MYSQL_ERRNO":                   false,
		"NAME":                          false,
		"NAMES":                         false,
		"NATIONAL":                      false,
		"NATURAL":                       true,
		"NCHAR":                         false,
		"NDBCLUSTER":                    false,
		"NDB":                           false,
		"NEVER":                         false,
		"NEW":                           false,
		"NEXT":                          false,
		"NODEGROUP":                     false,
		"NO":                            false,
		"NONBLOCKING":                   false,
		"NONE":                          false,
		"NOT":                           true,
		"NO_WAIT":                       false,
		"NO_WRITE_TO_BINLOG":            true,
		"NULL":                          true,
		"NUMBER":                        false,
		"NUMERIC":                       true,
		"NVARCHAR":                      false,
		"OFFSET":                        false,
		"OLD_PASSWORD":                  false,
		"ONE":                           false,
		"ONLY":                          false,
		"ON":                            true,
		"OPEN":                          false,
		"OPTIMIZER_COSTS":               true,
		"OPTIMIZE":                      true,
		"OPTIONALLY":                    true,
		"OPTIONS":                       false,
		"OPTION":                        true,
		"ORDER":                         true,
		"OR":                            true,
		"OUTER":                         true,
		"OUTFILE":                       true,
		"OUT":                           true,
		"OWNER":                         false,
		"PACK_KEYS":                     false,
		"PAGE":                          false,
		"PARSE_GCOL_EXPR":               true,
		"PARSER":                        false,
		"PARTIAL":                       false,
		"PARTITIONING":                  false,
		"PARTITIONS":                    false,
		"PARTITION":                     true,
		"PASSWORD":                      false,
		"PHASE":                         false,
		"PLUGIN_DIR":                    false,
		"PLUGIN":                        false,
		"PLUGINS":                       false,
		"POINT":                         false,
		"POLYGON":                       false,
		"PORT":                          false,
		"PRECEDES":                      false,
		"PRECISION":                     true,
		"PREPARE":                       false,
		"PRESERVE":                      false,
		"PREV":                          false,
		"PRIMARY":                       true,
		"PRIVILEGES":                    false,
		"PROCEDURE":                     true,
		"PROCESSLIST":                   false,
		"PROFILE":                       false,
		"PROFILES":                      false,
		"PROXY":                         false,
		"PURGE":                         true,
		"QUARTER":                       false,
		"QUERY":                         false,
		"QUICK":                         false,
		"RANGE":                         true,
		"READ_ONLY":                     false,
		"READS":                         true,
		"READ":                          true,
		"READ_WRITE":                    true,
		"REAL":                          true,
		"REBUILD":                       false,
		"RECOVER":                       false,
		"REDO_BUFFER_SIZE":              false,
		"REDOFILE":                      false,
		"REDUNDANT":                     false,
		"REFERENCES":                    true,
		"REGEXP":                        true,
		"RELAY":                         false,
		"RELAYLOG":                      false,
		"RELAY_LOG_FILE":                false,
		"RELAY_LOG_POS":                 false,
		"RELAY_THREAD":                  false,
		"RELEASE":                       true,
		"RELOAD":                        false,
		"REMOVE":                        false,
		"RENAME":                        true,
		"REORGANIZE":                    false,
		"REPAIR":                        false,
		"REPEATABLE":                    false,
		"REPEAT":                        true,
		"REPLACE":                       true,
		"REPLICATE_DO_DB":               false,
		"REPLICATE_DO_TABLE":            false,
		"REPLICATE_IGNORE_DB":           false,
		"REPLICATE_IGNORE_TABLE":        false,
		"REPLICATE_REWRITE_DB":          false,
		"REPLICATE_WILD_DO_TABLE":       false,
		"REPLICATE_WILD_IGNORE_TABLE":   false,
		"REPLICATION":                   false,
		"REQUIRE":                       true,
		"RESET":                         false,
		"RESIGNAL":                      true,
		"RESTORE":                       false,
		"RESTRICT":                      true,
		"RESUME":                        false,
		"RETURNED_SQLSTATE":             false,
		"RETURNS":                       false,
		"RETURN":                        true,
		"REVERSE":                       false,
		"REVOKE":                        true,
		"RIGHT":                         true,
		"RLIKE":                         true,
		"ROLLBACK":                      false,
		"ROLLUP":                        false,
		"ROTATE":                        false,
		"ROUTINE":                       false,
		"ROW_COUNT":                     false,
		"ROW":                           false,
		"ROW_FORMAT":                    false,
		"ROWS":                          false,
		"RTREE":                         false,
		"SAVEPOINT":                     false,
		"SCHEDULE":                      false,
		"SCHEMA_NAME":                   false,
		"SCHEMAS":                       true,
		"SCHEMA":                        true,
		"SECOND":                        false,
		"SECOND_MICROSECOND":            true,
		"SECURITY":                      false,
		"SELECT":                        true,
		"SENSITIVE":                     true,
		"SEPARATOR":                     true,
		"SERIAL":                        false,
		"SERIALIZABLE":                  false,
		"SERVER":                        false,
		"SESSION":                       false,
		"SET":                           true,
		"SHARE":                         false,
		"SHOW":                          true,
		"SHUTDOWN":                      false,
		"SIGNAL":                        true,
		"SIGNED":                        false,
		"SIMPLE":                        false,
		"SLAVE":                         false,
		"SLOW":                          false,
		"SMALLINT":                      true,
		"SNAPSHOT":                      false,
		"SOCKET":                        false,
		"SOME":                          false,
		"SONAME":                        false,
		"SOUNDS":                        false,
		"SOURCE":                        false,
		"SPATIAL":                       true,
		"SPECIFIC":                      true,
		"SQL_AFTER_GTIDS":               false,
		"SQL_AFTER_MTS_GAPS":            false,
		"SQL_BEFORE_GTIDS":              false,
		"SQL_BIG_RESULT":                true,
		"SQL_BUFFER_RESULT":             false,
		"SQL_CACHE":                     false,
		"SQL_CALC_FOUND_ROWS":           true,
		"SQLEXCEPTION":                  true,
		"SQL_NO_CACHE":                  false,
		"SQL_SMALL_RESULT":              true,
		"SQLSTATE":                      true,
		"SQL_THREAD":                    false,
		"SQL":                           true,
		"SQL_TSI_DAY":                   false,
		"SQL_TSI_HOUR":                  false,
		"SQL_TSI_MINUTE":                false,
		"SQL_TSI_MONTH":                 false,
		"SQL_TSI_QUARTER":               false,
		"SQL_TSI_SECOND":                false,
		"SQL_TSI_WEEK":                  false,
		"SQL_TSI_YEAR":                  false,
		"SQLWARNING":                    true,
		"SSL":                           true,
		"STACKED":                       false,
		"START":                         false,
		"STARTING":                      true,
		"STARTS":                        false,
		"STATS_AUTO_RECALC":             false,
		"STATS_PERSISTENT":              false,
		"STATS_SAMPLE_PAGES":            false,
		"STATUS":                        false,
		"STOP":                          false,
		"STORAGE":                       false,
		"STORED":                        true,
		"STRAIGHT_JOIN":                 true,
		"STRING":                        false,
		"SUBCLASS_ORIGIN":               false,
		"SUBJECT":                       false,
		"SUBPARTITION":                  false,
		"SUBPARTITIONS":                 false,
		"SUPER":                         false,
		"SUSPEND":                       false,
		"SWAPS":                         false,
		"SWITCHES":                      false,
		"TABLE_CHECKSUM":                false,
		"TABLE_NAME":                    false,
		"TABLES":                        false,
		"TABLESPACE":                    false,
		"TABLE":                         true,
		"TEMPORARY":                     false,
		"TEMPTABLE":                     false,
		"TERMINATED":                    true,
		"TEXT":                          false,
		"THAN":                          false,
		"THEN":                          true,
		"TIME":                          false,
		"TIMESTAMPADD":                  false,
		"TIMESTAMPDIFF":                 false,
		"TIMESTAMP":                     false,
		"TINYBLOB":                      true,
		"TINYINT":                       true,
		"TINYTEXT":                      true,
		"TO":                            true,
		"TRAILING":                      true,
		"TRANSACTION":                   false,
		"TRIGGERS":                      false,
		"TRIGGER":                       true,
		"TRUE":                          true,
		"TRUNCATE":                      false,
		"TYPE":                          false,
		"TYPES":                         false,
		"UNCOMMITTED":                   false,
		"UNDEFINED":                     false,
		"UNDO_BUFFER_SIZE":              false,
		"UNDOFILE":                      false,
		"UNDO":                          true,
		"UNICODE":                       false,
		"UNINSTALL":                     false,
		"UNION":                         true,
		"UNIQUE":                        true,
		"UNKNOWN":                       false,
		"UNLOCK":                        true,
		"UNSIGNED":                      true,
		"UNTIL":                         false,
		"UPDATE":                        true,
		"UPGRADE":                       false,
		"USAGE":                         true,
		"USE_FRM":                       false,
		"USER":                          false,
		"USER_RESOURCES":                false,
		"USE":                           true,
		"USING":                         true,
		"UTC_DATE":                      true,
		"UTC_TIMESTAMP":                 true,
		"UTC_TIME":                      true,
		"VALIDATION":                    false,
		"VALUE":                         false,
		"VALUES":                        true,
		"VARBINARY":                     true,
		"VARCHARACTER":                  true,
		"VARCHAR":                       true,
		"VARIABLES":                     false,
		"VARYING":                       true,
		"VIEW":                          false,
		"VIRTUAL":                       true,
		"WAIT":                          false,
		"WARNINGS":                      false,
		"WEEK":                          false,
		"WEIGHT_STRING":                 false,
		"WHEN":                          true,
		"WHERE":                         true,
		"WHILE":                         true,
		"WITHOUT":                       false,
		"WITH":                          true,
		"WORK":                          false,
		"WRAPPER":                       false,
		"WRITE":                         true,
		"X509":                          false,
		"XA":                            false,
		"XID":                           false,
		"XML":                           false,
		"XOR":                           true,
		"YEAR":                          false,
		"YEAR_MONTH":                    true,
		"ZEROFILL":                      true,
	}

	v, ok := mysqlKeywords[strings.ToUpper(s)]

	return ok, v
}

// IsKeyword returns a boolean indicating if the supplied string
// is considered to be a keyword in MySQL
func (d MySQLDialect) IsKeyword(s string) bool {

	isKey, _ := d.keyword(s)
	return isKey
}

// IsReservedKeyword returns a boolean indicating if the supplied
// string is considered to be a reserved keyword in MySQL
func (d MySQLDialect) IsReservedKeyword(s string) bool {

	isKey, isReserved := d.keyword(s)

	if isKey {
		return isReserved
	}
	return false
}

// IsOperator returns a boolean indicating if the supplied string
// is considered to be an operator in MySQL
func (d MySQLDialect) IsOperator(s string) bool {

	var mysqlOperators = map[string]bool{
		"^":   true,
		"~":   true,
		"<":   true,
		"<<":  true,
		"<=":  true,
		"<=>": true,
		"<>":  true,
		"=":   true,
		">":   true,
		">=":  true,
		">>":  true,
		"|":   true,
		"||":  true,
		"-":   true,
		"->":  true,
		"->>": true,
		":=":  true,
		"!":   true,
		"!=":  true,
		"/":   true,
		"*":   true,
		"&":   true,
		"&&":  true,
		"%":   true,
		"+":   true,
	}

	_, ok := mysqlOperators[s]
	return ok
}

// IsLabel returns a boolean indicating if the supplied string
// is considered to be a label in MySQL
func (d MySQLDialect) IsLabel(s string) bool {
	if len(s) < 2 {
		return false
	}
	if string(s[len(s)-1]) != ":" {
		return false
	}
	if !d.IsIdentifier(s[0 : len(s)-2]) {
		return false
	}
	return true
}

// IsIdentifier returns a boolean indicating if the supplied
// string is considered to be a non-quoted MySQL identifier.
func (d MySQLDialect) IsIdentifier(s string) bool {

	/*

	   From the documentation:

	   * Permitted characters in unquoted identifiers:
	       ASCII: [0-9,a-z,A-Z$_] (basic Latin letters, digits 0-9, dollar, underscore)
	       Extended: U+0080 .. U+FFFF
	   * ASCII NUL (U+0000) and supplementary characters (U+10000 and higher) are not permitted in quoted or unquoted identifiers.
	   * Identifiers may begin with a digit but unless quoted may not consist solely of digits.
	   * Database, table, and column names cannot end with space characters.

	*/

	const identChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_$"
	const digitChars = "0123456789"

	allDigits := true

	chr := strings.Split(s, "")
	for i := 0; i < len(chr); i++ {

		matches := strings.Contains(identChars, chr[i])
		if !matches && chr[i] != "." {
			return false
		}

		if !strings.Contains(digitChars, chr[i]) {
			allDigits = false
		}
	}

	if allDigits {
		return false
	}

	return true
}
