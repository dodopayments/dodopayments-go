package cache

import (
	"context"
	"strings"
)

// InvalidationRule defines a cache invalidation rule.
type InvalidationRule struct {
	// Operation is the write operation that triggers invalidation.
	// Format: "Service.Method" (e.g., "Payment.New", "Customer.Update")
	Operation string

	// Invalidates specifies which cache entries should be invalidated.
	// Supports wildcards with "Service.*" or specific operations.
	// Format: "Service.Method" or "Service.*"
	Invalidates []string
}

// InvalidationRegistry manages cache invalidation rules.
type InvalidationRegistry struct {
	rules []InvalidationRule
}

// NewInvalidationRegistry creates a new invalidation registry with default rules.
func NewInvalidationRegistry() *InvalidationRegistry {
	registry := &InvalidationRegistry{
		rules: make([]InvalidationRule, 0),
	}
	
	// Add default invalidation rules
	registry.addDefaultRules()
	
	return registry
}

// addDefaultRules adds sensible default invalidation rules.
func (r *InvalidationRegistry) addDefaultRules() {
	// Payment operations
	r.AddRule(InvalidationRule{
		Operation:   "Payment.New",
		Invalidates: []string{"Payment.List", "Payment.ListAutoPaging"},
	})
	r.AddRule(InvalidationRule{
		Operation:   "Payment.Get", // If payment status changes
		Invalidates: []string{"Payment.*"},
	})
	
	// Subscription operations
	r.AddRule(InvalidationRule{
		Operation:   "Subscription.New",
		Invalidates: []string{"Subscription.List", "Subscription.ListAutoPaging"},
	})
	r.AddRule(InvalidationRule{
		Operation:   "Subscription.Update",
		Invalidates: []string{"Subscription.Get", "Subscription.List", "Subscription.ListAutoPaging"},
	})
	
	// Customer operations
	r.AddRule(InvalidationRule{
		Operation:   "Customer.New",
		Invalidates: []string{"Customer.List", "Customer.ListAutoPaging"},
	})
	r.AddRule(InvalidationRule{
		Operation:   "Customer.Update",
		Invalidates: []string{"Customer.Get", "Customer.List", "Customer.ListAutoPaging"},
	})
	
	// Product operations
	r.AddRule(InvalidationRule{
		Operation:   "Product.New",
		Invalidates: []string{"Product.List", "Product.ListAutoPaging"},
	})
	r.AddRule(InvalidationRule{
		Operation:   "Product.Update",
		Invalidates: []string{"Product.Get", "Product.List", "Product.ListAutoPaging"},
	})
	
	// Discount operations
	r.AddRule(InvalidationRule{
		Operation:   "Discount.New",
		Invalidates: []string{"Discount.List", "Discount.ListAutoPaging"},
	})
	r.AddRule(InvalidationRule{
		Operation:   "Discount.Update",
		Invalidates: []string{"Discount.Get", "Discount.List", "Discount.ListAutoPaging"},
	})
	
	// Addon operations
	r.AddRule(InvalidationRule{
		Operation:   "Addon.New",
		Invalidates: []string{"Addon.List", "Addon.ListAutoPaging"},
	})
	r.AddRule(InvalidationRule{
		Operation:   "Addon.Update",
		Invalidates: []string{"Addon.Get", "Addon.List", "Addon.ListAutoPaging"},
	})
}

// AddRule adds an invalidation rule to the registry.
func (r *InvalidationRegistry) AddRule(rule InvalidationRule) {
	r.rules = append(r.rules, rule)
}

// GetInvalidations returns the cache keys that should be invalidated for a given operation.
func (r *InvalidationRegistry) GetInvalidations(operation string) []string {
	var invalidations []string
	
	for _, rule := range r.rules {
		if rule.Operation == operation {
			invalidations = append(invalidations, rule.Invalidates...)
		}
	}
	
	return invalidations
}

// Invalidator handles cache invalidation based on operations.
type Invalidator struct {
	registry *InvalidationRegistry
	config   *CacheConfig
}

// NewInvalidator creates a new cache invalidator.
func NewInvalidator(config *CacheConfig) *Invalidator {
	return &Invalidator{
		registry: NewInvalidationRegistry(),
		config:   config,
	}
}

// Invalidate performs cache invalidation for a given operation.
func (i *Invalidator) Invalidate(ctx context.Context, service, method string) error {
	// Only perform invalidation if strategy is operation-based or hybrid
	if i.config.InvalidationStrategy == TimeBased {
		return nil
	}
	
	operation := service + "." + method
	invalidations := i.registry.GetInvalidations(operation)
	
	if len(invalidations) == 0 {
		return nil
	}
	
	// Invalidate each rule
	for _, invalidation := range invalidations {
		if err := i.invalidateRule(ctx, invalidation); err != nil {
			return err
		}
	}
	
	return nil
}

// invalidateRule invalidates a single invalidation rule.
func (i *Invalidator) invalidateRule(ctx context.Context, rule string) error {
	parts := strings.Split(rule, ".")
	if len(parts) != 2 {
		return nil
	}
	
	service := parts[0]
	method := parts[1]
	
	// Handle wildcards
	if method == "*" {
		// Delete all cache entries for this service
		prefix := "dodo:" + service + ":"
		return i.config.Cache.DeletePrefix(ctx, prefix)
	}
	
	// Delete specific operation
	// Note: This is a simplified version. In production, you might want to track
	// which cache keys exist for each operation to delete them more efficiently.
	prefix := "dodo:" + service + ":" + method + ":"
	return i.config.Cache.DeletePrefix(ctx, prefix)
}

// AddRule adds a custom invalidation rule.
func (i *Invalidator) AddRule(rule InvalidationRule) {
	i.registry.AddRule(rule)
}
