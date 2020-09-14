package app

const (
	ExperimentKvsStorageUseBadger               = "kvs_use_badger"
	ExperimentKvsStorageUseBitcask              = "kvs_use_bitcask"
	ExperimentKvsStorageBadgerUseInMemory       = "kvs_badger_use_inmemory"
	ExperimentKvsStorageBadgerCompressionZstd   = "kvs_badger_compress_zstd"
	ExperimentKvsStorageBadgerCompressionSnappy = "kvs_badger_compress_snappy"

	ExperimentDbxClientConditionerNarrow20  = "dbx_client_conditioner_narrow20"  // 429 error for 20% traffic
	ExperimentDbxClientConditionerNarrow40  = "dbx_client_conditioner_narrow40"  // 429 error for 40% traffic
	ExperimentDbxClientConditionerNarrow100 = "dbx_client_conditioner_narrow100" // 429 error for 100% traffic

	ExperimentDbxClientConditionerError20  = "dbx_client_conditioner_error20"  // 500 error for 20% traffic
	ExperimentDbxClientConditionerError40  = "dbx_client_conditioner_error40"  // 500 error for 40% traffic
	ExperimentDbxClientConditionerError100 = "dbx_client_conditioner_error100" // 500 error for 100% traffic

	// Execute batch sequentially in same batchId
	ExperimentBatchSequential = "batch_sequential"
	// Execute batch with random batchId order
	ExperimentBatchRandom = "batch_random"
)
