package app_msg_container_impl

func keyWithPrefix(prefix, key string) string {
	if prefix != "" {
		return prefix + "." + key
	} else {
		return key
	}
}
