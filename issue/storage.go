package issue

var latest map[string]Component

// RefreshComponents Merges with the latest data and returns changed components, which should be notified
func RefreshComponents(components []Component) []Component {
	var changed []Component
	if latest == nil {
		latest = make(map[string]Component)
		for _, component := range components {
			if !component.IsOperational() {
				changed = append(changed, component)
			}
			latest[component.Title] = component
		}
		return changed
	}

	for _, component := range components {
		different := false
		for s, c := range latest {
			if c.Title == component.Title {
				different = c.GetStatus().Name != component.GetStatus().Name
			}
			latest[s] = component
		}
		if different {
			changed = append(changed, component)
		}
	}
	return changed
}
