package policy

import (
	"time"
	"windops/internal/fault"
)

func EvaluateQueryFilterAndCountMatch(ctx Context) (Result, error) {
	if err := requireTenant(ctx, "policy.query_consistency"); err != nil {
		return Result{}, err
	}
	if ctx.Used < 0 || ctx.Capacity < 0 {
		return Result{}, fault.New(fault.CodeInvalid, "policy.query_consistency", "query counts cannot be negative")
	}
	if ctx.Used > ctx.Capacity {
		return Result{}, fault.New(fault.CodePrecondition, "policy.query_consistency", "page item count cannot exceed total")
	}
	if !sameQueryFilters(ctx.Status, ctx.RelatedStatus, ctx.Region, ctx.Metadata["count_region"]) {
		return deny("query_filter_mismatch", "item query and count query use different filters"), nil
	}
	result := allow("item query and count query use matching filters")
	result.Quantity = ctx.Capacity
	return result, nil
}

var _ = time.Second
var _ = fault.CodeInvalid
