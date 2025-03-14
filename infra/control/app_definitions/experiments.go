package app_definitions

const (
	ExperimentDbxClientConditionerNarrow20  = "dbx_client_conditioner_narrow20"  // 429 error for 20% traffic
	ExperimentDbxClientConditionerNarrow40  = "dbx_client_conditioner_narrow40"  // 429 error for 40% traffic
	ExperimentDbxClientConditionerNarrow100 = "dbx_client_conditioner_narrow100" // 429 error for 100% traffic

	ExperimentDbxClientConditionerError20  = "dbx_client_conditioner_error20"  // 500 error for 20% traffic
	ExperimentDbxClientConditionerError40  = "dbx_client_conditioner_error40"  // 500 error for 40% traffic
	ExperimentDbxClientConditionerError100 = "dbx_client_conditioner_error100" // 500 error for 100% traffic

	ExperimentDbxAuthRedirect           = "dbx_auth_redirect"
	ExperimentDbxAuthCourseGrainedScope = "dbx_auth_course_grained_scope"
	ExperimentDbxDisableAutoPathRoot    = "dbx_disable_auto_path_root"

	// Execute batch sequentially in same batchId
	ExperimentBatchSequential = "batch_sequential"
	// Execute batch with random batchId order
	ExperimentBatchRandom = "batch_random"
	// Execute batch from the largest batch
	ExperimentBatchBalance = "batch_balance"
	// ExperimentBatchNonDurable use non durable queue
	ExperimentBatchNonDurable = "batch_non_durable"

	// Do not hard limit window size
	ExperimentCongestionWindowNoLimit = "congestion_window_no_limit"
	// Aggressive initial window size
	ExperimentCongestionWindowAggressive = "congestion_window_aggressive"
	// Reduce call create_folder on sync
	ExperimentFileSyncDisableReduceCreateFolder = "file_sync_disable_reduce_create_folder"
	// Use legacy local to dropbox sync connector
	ExperimentFileSyncLegacyLocalToDropboxConnector = "legacy_local_to_dbx_connector"
	// Use non-cache dropbox file system
	ExperimentFileSyncNoCacheDropboxFileSystem = "use_no_cache_dbxfs"

	// ExperimentReportAllColumns Do not hide columns
	ExperimentReportAllColumns = "report_all_columns"

	// ExperimentProfileMemory Enable memory profiler
	ExperimentProfileMemory = "profile_memory"

	// ExperimentProfileCpu Enable CPU profiler
	ExperimentProfileCpu = "profile_cpu"

	// ExperimentKvsBadger Use Badger as KVS engine
	ExperimentKvsBadger = "kvs_badger"

	// ExperimentKvsBadgerTurnstile Use Bitcask as KVS engine with turnstile
	ExperimentKvsBadgerTurnstile = "kvs_badger_turnstile"

	// ExperimentKvsBitcask Use Bitcask as KVS engine
	ExperimentKvsBitcask = "kvs_bitcask"

	// ExperimentKvsBitcaskTurnstile Use Bitcask as KVS engine with turnstile
	ExperimentKvsBitcaskTurnstile = "kvs_bitcask_turnstile"

	// ExperimentKvsSqlite Use SQLite as KVS engine
	ExperimentKvsSqlite = "kvs_sqlite"

	// ExperimentKvsSqliteTurnstile Use SQLite as KVS engine with turnstile
	ExperimentKvsSqliteTurnstile = "kvs_sqlite_turnstile"

	// ExperimentDbxDownloadBlock download by block
	ExperimentDbxDownloadBlock = "dbx_download_block"

	// ExperimentSuppressProgress suppress progress
	ExperimentSuppressProgress = "suppress_progress"

	// ExperimentValidateNetworkConnectionOnBootstrap validate network connection on bootstrap
	ExperimentValidateNetworkConnectionOnBootstrap = "validate_network_connection_on_bootstrap"
)

var (
	ExperimentalFeatures = []string{
		ExperimentBatchBalance,
		ExperimentBatchNonDurable,
		ExperimentBatchRandom,
		ExperimentBatchSequential,
		ExperimentCongestionWindowAggressive,
		ExperimentCongestionWindowNoLimit,
		ExperimentDbxAuthCourseGrainedScope,
		ExperimentDbxAuthRedirect,
		ExperimentDbxClientConditionerError100,
		ExperimentDbxClientConditionerError20,
		ExperimentDbxClientConditionerError40,
		ExperimentDbxClientConditionerNarrow100,
		ExperimentDbxClientConditionerNarrow20,
		ExperimentDbxClientConditionerNarrow40,
		ExperimentDbxDisableAutoPathRoot,
		ExperimentDbxDownloadBlock,
		ExperimentFileSyncDisableReduceCreateFolder,
		ExperimentFileSyncLegacyLocalToDropboxConnector,
		ExperimentFileSyncNoCacheDropboxFileSystem,
		ExperimentKvsBadger,
		ExperimentKvsBadgerTurnstile,
		ExperimentKvsBitcask,
		ExperimentKvsBitcaskTurnstile,
		ExperimentKvsSqlite,
		ExperimentKvsSqliteTurnstile,
		ExperimentProfileCpu,
		ExperimentProfileMemory,
		ExperimentReportAllColumns,
		ExperimentSuppressProgress,
		ExperimentValidateNetworkConnectionOnBootstrap,
	}
)
