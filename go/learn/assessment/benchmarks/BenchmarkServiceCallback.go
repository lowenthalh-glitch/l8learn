/*
 * Copyright 2026 Sharon Aicler (saichler@gmail.com)
 *
 * Licensed under the Apache License, Version 2.0
 */
package benchmarks

import (
	"github.com/saichler/l8common/go/common"
	"github.com/saichler/l8learn/go/types/learn"
	"github.com/saichler/l8types/go/ifs"
)

func newBenchmarkServiceCallback(vnic ifs.IVNic) ifs.IServiceCallback {
	return common.NewValidation(&learn.Benchmark{}, vnic).
		Require(func(v interface{}) string { return v.(*learn.Benchmark).BenchmarkId }, "BenchmarkId").
		Require(func(v interface{}) string { return v.(*learn.Benchmark).Name }, "Name").
		Build()
}
