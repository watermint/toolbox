package app

const (
	ExperimentKvsStorageUseInMemory       = "kvs_use_inmemory"
	ExperimentKvsStorageCompressionZstd   = "kvs_compress_zstd"
	ExperimentKvsStorageCompressionSnappy = "kvs_compress_snappy"

	ExperimentDbxClientConditionerNarrow20  = "dbx_client_conditioner_narrow20"  // 429 error for 20% traffic
	ExperimentDbxClientConditionerNarrow40  = "dbx_client_conditioner_narrow40"  // 429 error for 40% traffic
	ExperimentDbxClientConditionerNarrow100 = "dbx_client_conditioner_narrow100" // 429 error for 100% traffic
)
