package config

// GetFullModelPricing returns the complete ModelPricing entry for the given model,
// including cache read/write rates. Returns (p, true) when at least one rate is
// non-zero or Currency is explicitly set (currency-only counts as configured so
// that free/self-hosted models produce cost=0 telemetry). Returns (zero, false)
// when the model has no pricing entry at all. Accepts both short names and
// provider model IDs.
func (c *RouterConfig) GetFullModelPricing(modelName string) (ModelPricing, bool) {
	if modelConfig, ok := c.resolveModelConfig(modelName); ok {
		p := modelConfig.Pricing
		if p.PromptPer1M != 0 || p.CompletionPer1M != 0 ||
			p.CacheReadPer1M != 0 || p.CacheWritePer1M != 0 || p.Currency != "" {
			if p.Currency == "" {
				p.Currency = "USD"
			}
			return p, true
		}
	}
	return ModelPricing{}, false
}
