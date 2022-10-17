SOURCE := $(shell git rev-parse --show-toplevel)

include $(SOURCE)/scripts/make/dev.mk
include $(SOURCE)/scripts/make/build.mk
